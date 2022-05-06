package syscall_test

import (
	syscall2 "github.com/CarsonSlovoka/go-pkg/v2/syscall"
	"syscall"
	"testing"
)

func TestCreateMutex(t *testing.T) {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	procCreteMutex := kernel32DLL.NewProc("CreateMutexW")
	procCloseHandle := kernel32DLL.NewProc("CloseHandle")

	handle, err := syscall2.CreateMutexW(procCreteMutex, "hello world")
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = syscall2.CreateMutexW(procCreteMutex, "hello world")
	if err == nil || err != syscall.ERROR_ALREADY_EXISTS {
		t.Error("should panic")
	}

	if err = syscall2.CloseHandle(procCloseHandle, handle); err != nil {
		t.Error(err)
	}
	// err = syscall2.CloseHandle(procCloseHandle, handle) // If you are debugging it will panic!
	// fmt.Printf("%+v\n%d", err, err.(syscall.Errno))     // The Handle is invalid.  6

	// We can create again since we have closed.
	handle, _ = syscall2.CreateMutexW(procCreteMutex, "hello world")
	if err = syscall2.CloseHandle(procCloseHandle, handle); err != nil {
		t.Error(err)
	}
}
