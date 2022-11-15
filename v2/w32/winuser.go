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

func LOWORD(dwValue uint32) uint16 {
	return uint16(dwValue)
}

func HIWORD(dwValue uint32) uint16 {
	return uint16((dwValue >> 16) & 0xffff)
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmdiwindoww
const (
	CW_USEDEFAULT int32 = -2147483648
)

type WNDENUMPROC func(hWnd HWND, lParam LPARAM) BOOL

// WNDPROC https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
type WNDPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) LRESULT

// MENUITEMINFO https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32
	FType         uint32
	FState        uint32
	WID           uint32
	HSubMenu      HMENU
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    uintptr
	DwTypeData    *uint16 /* LPWSTR */
	Cch           uint32
	HbmpItem      HBITMAP
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
const (
	MIIM_BITMAP     uint32 = 0x00000080
	MIIM_CHECKMARKS        = 0x00000008
	MIIM_DATA              = 0x00000020
	MIIM_FTYPE             = 0x00000100
	MIIM_ID                = 0x00000002
	MIIM_STATE             = 0x00000001
	MIIM_STRING            = 0x00000040
	MIIM_SUBMENU           = 0x00000004
	MIIM_TYPE              = 0x00000010
)
