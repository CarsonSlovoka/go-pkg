package syscall

import (
	"fmt"
	"syscall"
	"unsafe"
)

// CreateMutex You can use it to restrict to a single instance of executable
// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createmutexW#return-value
func CreateMutex(name string) (uintptr, error) {

	if DllKernel32 == nil {
		return 0, fmt.Errorf(`you need call 'syscall2.DllKernel32 = syscall.NewLazyDLL("kernel32.dll")' first'`)
	}

	procCreateMutex := DllKernel32.NewProc("CreateMutexW")
	lpName, _ := syscall.UTF16PtrFromString(name) // LPCWSTR
	handleID, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(lpName)),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return handleID, nil
	default:
		return handleID, err
	}
}
