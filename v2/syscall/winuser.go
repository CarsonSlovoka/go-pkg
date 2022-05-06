package syscall

import (
	"syscall"
	"unsafe"
)

// MessageBoxW
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func MessageBoxW(proc *syscall.LazyProc, caption, body string, btnFlag uintptr) uintptr {
	if proc.Name != "MessageBoxW" {
		panic(proc.Name != "MessageBoxW")
	}
	hwnd := uintptr(0x00)
	lpCaption, _ := syscall.UTF16PtrFromString(caption) // LPCWSTR
	lpText, _ := syscall.UTF16PtrFromString(body)       // LPCWSTR
	responseVal, _, _ := proc.Call(hwnd,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		btnFlag)
	return responseVal
}

func MessageBoxOK(proc *syscall.LazyProc, caption, body string) uintptr {
	return MessageBoxW(proc, caption, body, 0x00000000)
}

func MessageBoxYesNo(proc *syscall.LazyProc, caption, body string) uintptr {
	return MessageBoxW(proc, caption, body, 0x00000004)
}

func MessageBoxYesNoCancel(proc *syscall.LazyProc, caption, body string) uintptr {
	return MessageBoxW(proc, caption, body, 0x00000003)
}
