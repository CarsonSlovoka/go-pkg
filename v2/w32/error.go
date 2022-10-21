package w32

import (
	"fmt"
)

func lastError(win32FuncName string) error {
	kernel32dll := NewKernel32DLL(
		PNGetLastError,
	)
	errno := kernel32dll.GetLastError()
	if errno != ERROR_SUCCESS {
		return fmt.Errorf("%s: Error %d\n", win32FuncName, errno)
	}

	return fmt.Errorf("%s\n", win32FuncName)
}
