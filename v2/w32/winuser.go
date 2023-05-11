package w32

import (
	"unsafe"
)

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

func LOWORD[T uintptr | WPARAM | LPARAM](v T) uint16 {
	return uint16(v & 0xffff)
}

func HIWORD[T uintptr | WPARAM | LPARAM](v T) uint16 {
	return uint16((v >> 16) & 0xffff)
}

func LOBYTE[T uintptr | WPARAM | LPARAM](v T) uint8 {
	return uint8(v & 0xff)
}

/* 有用到時再新增，避免測試覆蓋率降低
func HIBYTE[T uintptr | WPARAM | LPARAM](v T) uint8 {
	return uint8((v >> 8) & 0xff)
}
*/

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createmdiwindoww
const (
	CW_USEDEFAULT int32 = -2147483648
)

type WndEnumProc func(hWnd HWND, lParam LPARAM) BOOL

// WndProc https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
type WndProc func(hwnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT

// Constants for MENUITEMINFO.fState
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
const (
	MFS_GRAYED    = 0x00000003
	MFS_DISABLED  = MFS_GRAYED
	MFS_CHECKED   = MF_CHECKED
	MFS_HILITE    = MF_HILITE
	MFS_ENABLED   = MF_ENABLED
	MFS_UNCHECKED = MF_UNCHECKED
	MFS_UNHILITE  = MF_UNHILITE
	MFS_DEFAULT   = MF_DEFAULT
)

// MENUITEMINFO https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32 // MIIM_BITMAP, MIIM_ID, MIIM_STRING, MIIM_SUBMENU, MIIM_STATE...
	FType         uint32
	FState        uint32 // MFS_CHECKED, MFS_DEFAULT, ...
	WID           uint32 // 當您SetMenuItemInfo是使用Pos的方式時，此時就需要靠WID來辨識真正的ID
	HSubMenu      HMENU
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    uintptr
	DwTypeData    *uint16 // LPWSTR // &(utf16.Encode([]rune(MyString + "\x00")))[0] 或者 syscall.UTF16PtrFromString("MyString")
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

// HOOK
// https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644984(v=vs.85)
const (
	HC_ACTION   = 0
	HC_NOREMOVE = 3
)

// MOUSEHOOKSTRUCT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mousehookstruct?redirectedfrom=MSDN
type MOUSEHOOKSTRUCT struct {
	Pt           POINT
	Hwnd         HWND
	wHitTestCode UINT
	dwExtraInfo  ULONG_PTR
}

// CWPSTRUCT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-cwpstruct?redirectedfrom=MSDN
type CWPSTRUCT struct {
	LParam  LPARAM
	WParam  WPARAM
	Message UINT
	Hwnd    HWND
}

const (
	KF_EXTENDED uint16 = 0x0100 // https://learn.microsoft.com/en-us/windows/win32/inputdev/about-keyboard-input?redirectedfrom=MSDN#extended-key-flag INSERT, DEL, HOME, END, PAGE UP, PAGE DOWN, NUM LOCK, BREAK, 0-9 ...
	KF_DLGMODE         = 0x0800
	KF_MENUMODE        = 0x1000
	KF_ALTDOWN         = 0x2000
	KF_REPEAT          = 0x4000 // wasKeyDown
	KF_UP              = 0x8000
)

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-kbdllhookstruct?redirectedfrom=MSDN#members
const (
	LLKHF_EXTENDED                 = uint32(KF_EXTENDED) >> 8 // 1
	LLKHF_LOWER_IL_INJECTED uint32 = 0x00000002
	LLKHF_INJECTED          uint32 = 0x00000010
	LLKHF_ALTDOWN                  = uint32(KF_ALTDOWN) >> 8 // 32
	LLKHF_UP                       = uint32(KF_UP) >> 8      // 1000 0000 = 128
)

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msllhookstruct?redirectedfrom=MSDN
const (
	LLMHF_INJECTED          = 0x00000001
	LLMHF_LOWER_IL_INJECTED = 0x00000002
)

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeyexa
const (
	MAPVK_VK_TO_VSC    = 0 // 不區分
	MAPVK_VSC_TO_VK    = 1 // 若翻譯成功(不為0)，返回一個不區分左右建的虛擬鍵
	MAPVK_VK_TO_CHAR   = 2
	MAPVK_VSC_TO_VK_EX = 3 // 區分左右
	MAPVK_VK_TO_VSC_EX = 4 // 區分
)

// GET_KEYSTATE_WPARAM https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-get_keystate_wparam
func GET_KEYSTATE_WPARAM[T uintptr | WPARAM](wParam T) uint16 {
	return LOWORD(wParam)
}

// GET_WHEEL_DELTA_WPARAM https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-get_wheel_delta_wparam
func GET_WHEEL_DELTA_WPARAM[T uintptr | WPARAM](wparam T) int16 {
	return int16(HIWORD(wparam))
}

// KBDLLHOOKSTRUCT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-kbdllhookstruct?redirectedfrom=MSDN
type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DxExtraInfo ULONG_PTR
}

// MSLLHOOKSTRUCT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msllhookstruct?redirectedfrom=MSDN
type MSLLHOOKSTRUCT struct {
	Pt          POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DxExtraInfo ULONG_PTR
}

// WHEEL_DELTA
// positive: that the wheel was rotated forward, away from the user;
// negative: that the wheel was rotated backward, toward the user.
const WHEEL_DELTA = 120 // Default value for rolling one notch

// CBT Hook Codes
const (
	HCBT_MOVESIZE     = 0 // A window is about to be moved or sized.
	HCBT_MINMAX       = 1 // A window is about to be minimized or maximized.
	HCBT_QS           = 2
	HCBT_CREATEWND    = 3 // A window is about to be created.
	HCBT_DESTROYWND   = 4 // A window is about to be destroyed.
	HCBT_ACTIVATE     = 5 // The system is about to activate a window.
	HCBT_CLICKSKIPPED = 6
	HCBT_KEYSKIPPED   = 7
	HCBT_SYSCOMMAND   = 8
	HCBT_SETFOCUS     = 9 // A window is about to receive the keyboard focus.
)

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerhotkey#parameters
const (
	MOD_ALT      = 0x0001
	MOD_CONTROL  = 0x0002
	MOD_SHIFT    = 0x0004
	MOD_WIN      = 0x0008
	MOD_NOREPEAT = 0x4000
)

// INPUT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-input
type INPUT struct {
	Type     uint32 // INPUT_MOUSE, INPUT_KEYBOARD, INPUT_HARDWARE // To know which struct should be used.
	padding1 [pad4for64_0for32]byte
	data     [8 * pad4for64_3for32]byte
}

func (input *INPUT) Mi() *MOUSEINPUT {
	return (*MOUSEINPUT)(unsafe.Pointer(&input.data[0]))
}
func (input *INPUT) Ki() *KEYBDINPUT {
	return (*KEYBDINPUT)(unsafe.Pointer(&input.data[0]))
}
func (input *INPUT) Hi() *HARDWAREINPUT {
	return (*HARDWAREINPUT)(unsafe.Pointer(&input.data[0]))
}

// KEYBDINPUT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-keybdinput
type KEYBDINPUT struct {
	Vk        uint16
	Scan      uint16 // If the Flags specifies KEYEVENTF_UNICODE, wScan specifies a Unicode character
	Flags     uint32
	Time      uint32
	ExtraInfo ULONG_PTR
}

// MOUSEINPUT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mouseinput
type MOUSEINPUT struct {
	Dx        int32
	Dy        int32
	MouseData uint32 // for Flags:{MOUSEEVENTF_WHEEL(120 WHEEL_DELTA,-120), MOUSEEVENTF_XDOWN, MOUSEEVENTF_XUP} used. Otherwise should be zero.
	Flags     uint32
	Time      uint32
	ExtraInfo ULONG_PTR
}

// HARDWAREINPUT https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-hardwareinput
// https://www.computerhope.com/jargon/i/inputdev.htm
type HARDWAREINPUT struct {
	Msg     uint32
	LParamL uint16 // The low-order word of the lParam parameter for uMsg.
	LParamH uint16 // The high-order word of the lParam parameter for uMsg.
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-printwindow#parameters
const (
	PW_CLIENTONLY        = 0x00000001
	PW_RENDERFULLCONTENT = 0x00000002
)
