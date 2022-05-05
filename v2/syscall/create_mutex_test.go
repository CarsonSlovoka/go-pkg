package syscall_test

import (
	syscall2 "github.com/CarsonSlovoka/go-pkg/v2/syscall"
	"syscall"
	"testing"
)

func TestCreateMutex(t *testing.T) {
	syscall2.DllKernel32 = syscall.NewLazyDLL("kernel32.dll")
	handle, err := syscall2.CreateMutex("hello world")
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = syscall2.CreateMutex("hello world")
	if err == nil || err != syscall.ERROR_ALREADY_EXISTS {
		t.Error("should panic")
	}

	if err = syscall2.CloseHandle(handle); err != nil {
		t.Error(err)
	}
	// syscall2.CloseHandle(handle) // If the handle does not exist will panic! Don't do that.

	// We can create again since we have closed.
	handle, _ = syscall2.CreateMutex("hello world")
	if err = syscall2.CloseHandle(handle); err != nil {
		t.Error(err)
	}
}
