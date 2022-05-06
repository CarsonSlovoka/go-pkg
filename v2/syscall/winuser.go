package syscall

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
	"unsafe"
)

// MessageBoxW
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func MessageBoxW[H w32.Handle](proc *syscall.LazyProc, handle H, caption, body string, btnFlag uintptr) uintptr {
	if proc.Name != "MessageBoxW" {
		panic(proc.Name != "MessageBoxW")
	}
	lpCaption, _ := syscall.UTF16PtrFromString(caption) // LPCWSTR
	lpText, _ := syscall.UTF16PtrFromString(body)       // LPCWSTR
	responseVal, _, _ := proc.Call(uintptr(handle),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		btnFlag)

	return responseVal
}

func MessageBox[H w32.Handle](proc *syscall.LazyProc, handle H, caption, body string, btnFlag uintptr) uintptr {
	return MessageBoxW(proc, handle, caption, body, btnFlag)
}

func MessageBoxOK[H w32.Handle](proc *syscall.LazyProc, handle H, caption, body string) uintptr {
	return MessageBoxW(proc, handle, caption, body, 0x00000000)
}

func MessageBoxYesNo[H w32.Handle](proc *syscall.LazyProc, handle H, caption, body string) uintptr {
	return MessageBoxW(proc, handle, caption, body, 0x00000004)
}

func MessageBoxYesNoCancel[H w32.Handle](proc *syscall.LazyProc, handle H, caption, body string) uintptr {
	return MessageBoxW(proc, handle, caption, body, 0x00000003)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
//func GetSystemMetrics(proc *syscall.LazyProc, nIndex int)
