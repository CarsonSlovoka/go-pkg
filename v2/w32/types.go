//go:build windows

package w32

import "unsafe"

type (
	ATOM          uint16
	HANDLE        uintptr
	HGLOBAL       HANDLE
	HINSTANCE     HANDLE
	LCID          uint32
	LCTYPE        uint32
	LANGID        uint16
	HMODULE       uintptr
	HWINEVENTHOOK HANDLE
	HRSRC         uintptr
)

type (
	HACCEL    HANDLE
	HCURSOR   HANDLE
	HDWP      HANDLE
	HICON     HANDLE
	HMENU     HANDLE
	HMONITOR  HANDLE
	HRAWINPUT HANDLE
	HWND      HANDLE
)

type (
	LPCVOID unsafe.Pointer
)

type Handle interface {
	~uintptr // uintptr | HWND | HANDLE //不需要這樣寫，使用~，其表示只要「原型」是如此就算
}

type Hwnd interface {
	~uintptr
}
