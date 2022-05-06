package syscall

import (
	"syscall"
	"unsafe"
)

// CreateMutexW You can use it to restrict to a single instance of executable
// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createmutexW#return-value
func CreateMutexW(proc *syscall.LazyProc, name string) (uintptr, error) {
	if proc.Name != "CreateMutexW" {
		panic("proc.Name != CreateMutexW")
	}
	lpName, _ := syscall.UTF16PtrFromString(name) // LPCWSTR
	if handleID, _, err := proc.Call(
		0,
		0,
		uintptr(unsafe.Pointer(lpName)),
	); err.(syscall.Errno) == 0 {
		return handleID, nil
	} else {
		return handleID, err
	}
}
