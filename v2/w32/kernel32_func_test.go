package w32_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
	"testing"
)

func TestCreateMutex(t *testing.T) {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNCreateMutex,
		w32.PNCloseHandle,
	)
	handle, err := kernel32dll.CreateMutex("hello world")
	if err != nil {
		t.Fatal(err)
	}

	_, err = kernel32dll.CreateMutex("hello world")
	if err == nil || err != syscall.ERROR_ALREADY_EXISTS {
		t.Error("not as expected")
	}

	if err = kernel32dll.CloseHandle(handle); err != nil {
		t.Error(err)
	}
	// err = kernel32dll.CloseHandle(handle) // If you are debugging it will panic!
	// fmt.Printf("%+v\n%d", err, err.(syscall.Errno))     // The Handle is invalid.  6

	// We can create again since we have closed.
	handle, _ = kernel32dll.CreateMutex("hello world")
	if err = kernel32dll.CloseHandle(handle); err != nil {
		t.Error(err)
	}
}
