package syscall

import (
	"syscall"
	"unsafe"
)

func MessageBox(caption, body string, btnFlag uintptr) uintptr {
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
	var user32DLL = syscall.NewLazyDLL("user32.dll")
	var procMessageBox = user32DLL.NewProc("MessageBoxW") // Return value: Type int
	hwnd := uintptr(0x00)
	lpCaption, _ := syscall.UTF16PtrFromString(caption) // LPCWSTR
	lpText, _ := syscall.UTF16PtrFromString(body)       // LPCWSTR
	responseVal, _, _ := procMessageBox.Call(hwnd,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		btnFlag)
	return responseVal
}

func MessageBoxOK(caption, body string) uintptr {
	return MessageBox(caption, body, 0x00000000)
}

func MessageBoxYesNo(caption, body string) uintptr {
	return MessageBox(caption, body, 0x00000004)
}

func MessageBoxYesNoCancel(caption, body string) uintptr {
	return MessageBox(caption, body, 0x00000003)
}
