package w32

import "unsafe"

// https://learn.microsoft.com/en-us/windows/win32/menurc/resource-types
const (
	RT_ACCELERATOR  uintptr = 9
	RT_ANICURSOR            = 21
	RT_ANIICON              = 22
	RT_BITMAP               = 2
	RT_CURSOR               = 1
	RT_DIALOG               = 5
	RT_DLGINCLUDE           = 17
	RT_FONT                 = 8
	RT_FONTDIR              = 7
	RT_GROUP_CURSOR         = RT_CURSOR + 11
	RT_GROUP_ICON           = RT_ICON + 11
	RT_ICON                 = 3
	RT_HTML                 = 23
	RT_MANIFEST             = 24
	RT_MENU                 = 4
	RT_MESSAGETABLE         = 11
	RT_PLUGPLAY             = 19
	RT_RCDATA               = 10
	RT_STRING               = 6
	RT_VERSION              = 16
	RT_VXD                  = 20
)

// MakeIntResource https://learn.microsoft.com/en-us/windows/win32/menurc/resource-types
func MakeIntResource(id uintptr) *uint16 {
	return (*uint16)(unsafe.Pointer(id))
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmdiwindoww
const (
	CW_USEDEFAULT int32 = -2147483648
)

type WNDENUMPROC func(hWnd HWND, lParam LPARAM) BOOL

// WNDPROC https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
type WNDPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) LRESULT
