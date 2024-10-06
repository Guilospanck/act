//go:build (!windows && !plan9 && !openbsd) || (!windows && !plan9 && !mips64)

package container

import (
	"fmt"
	"math"
	"os"
	"syscall"

	"github.com/creack/pty"
)

func getSysProcAttr(_ string, tty bool) *syscall.SysProcAttr {
	if tty {
		return &syscall.SysProcAttr{
			Setsid:  true,
			Setctty: true,
		}
	}
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func openPty() (*os.File, *os.File, error) {
	return pty.Open()
}

func safeConversionFromInt64IntoUint32(value int64) (uint32, error) {
	if value < 0 || value > math.MaxUint32 {
		return 0, fmt.Errorf("value outside of uint32 range")
	}

	return uint32(value), nil
}
