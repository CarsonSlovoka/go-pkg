package w32

// MessageBox constants
const (
	MB_OK                   = 0x00000000
	MB_OKCANCEL             = 0x00000001
	MB_ABORTRETRYIGNORE     = 0x00000002
	MB_YESNOCANCEL          = 0x00000003
	MB_YESNO                = 0x00000004
	MB_RETRYCANCEL          = 0x00000005
	MB_CANCELTRYCONTINUE    = 0x00000006
	MB_ICONHAND             = 0x00000010
	MB_ICONQUESTION         = 0x00000020
	MB_ICONEXCLAMATION      = 0x00000030
	MB_ICONASTERISK         = 0x00000040
	MB_USERICON             = 0x00000080
	MB_ICONWARNING          = MB_ICONEXCLAMATION
	MB_ICONERROR            = MB_ICONHAND
	MB_ICONINFORMATION      = MB_ICONASTERISK
	MB_ICONSTOP             = MB_ICONHAND
	MB_DEFBUTTON1           = 0x00000000
	MB_DEFBUTTON2           = 0x00000100
	MB_DEFBUTTON3           = 0x00000200
	MB_DEFBUTTON4           = 0x00000300
	MB_APPLMODAL            = 0x00000000
	MB_SYSTEMMODAL          = 0x00001000
	MB_TASKMODAL            = 0x00002000
	MB_HELP                 = 0x00004000
	MB_SETFOREGROUND        = 0x00010000
	MB_DEFAULT_DESKTOP_ONLY = 0x00020000
	MB_TOPMOST              = 0x00040000
	MB_RIGHT                = 0x00080000
	MB_RTLREADING           = 0x00100000
	MB_SERVICE_NOTIFICATION = 0x00200000
)

// Dialog box command ids
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw#return-value
const (
	IDOK       = 1
	IDCANCEL   = 2
	IDABORT    = 3
	IDRETRY    = 4
	IDIGNORE   = 5
	IDYES      = 6
	IDNO       = 7
	IDCLOSE    = 8
	IDHELP     = 9
	IDTRYAGAIN = 10
	IDCONTINUE = 11
	IDTIMEOUT  = 32000
)

// System commands
const (
	SC_SIZE         = 0xF000
	SC_MOVE         = 0xF010
	SC_MINIMIZE     = 0xF020
	SC_MAXIMIZE     = 0xF030
	SC_NEXTWINDOW   = 0xF040
	SC_PREVWINDOW   = 0xF050
	SC_CLOSE        = 0xF060
	SC_VSCROLL      = 0xF070
	SC_HSCROLL      = 0xF080
	SC_MOUSEMENU    = 0xF090
	SC_KEYMENU      = 0xF100
	SC_ARRANGE      = 0xF110
	SC_RESTORE      = 0xF120
	SC_TASKLIST     = 0xF130
	SC_SCREENSAVE   = 0xF140
	SC_HOTKEY       = 0xF150
	SC_DEFAULT      = 0xF160
	SC_MONITORPOWER = 0xF170
	SC_CONTEXTHELP  = 0xF180
	SC_SEPARATOR    = 0xF00F
)

// Static control styles
const (
	SS_BITMAP          = 14
	SS_BLACKFRAME      = 7
	SS_BLACKRECT       = 4
	SS_CENTER          = 1
	SS_CENTERIMAGE     = 512
	SS_EDITCONTROL     = 0x2000
	SS_ENHMETAFILE     = 15
	SS_ETCHEDFRAME     = 18
	SS_ETCHEDHORZ      = 16
	SS_ETCHEDVERT      = 17
	SS_GRAYFRAME       = 8
	SS_GRAYRECT        = 5
	SS_ICON            = 3
	SS_LEFT            = 0
	SS_LEFTNOWORDWRAP  = 0xc
	SS_NOPREFIX        = 128
	SS_NOTIFY          = 256
	SS_OWNERDRAW       = 0xd
	SS_REALSIZECONTROL = 0x040
	SS_REALSIZEIMAGE   = 0x800
	SS_RIGHT           = 2
	SS_RIGHTJUST       = 0x400
	SS_SIMPLE          = 11
	SS_SUNKEN          = 4096
	SS_WHITEFRAME      = 9
	SS_WHITERECT       = 6
	SS_USERITEM        = 10
	SS_TYPEMASK        = 0x0000001F
	SS_ENDELLIPSIS     = 0x00004000
	SS_PATHELLIPSIS    = 0x00008000
	SS_WORDELLIPSIS    = 0x0000C000
	SS_ELLIPSISMASK    = 0x0000C000
)

// Button message constants
const (
	BM_CLICK    = 245
	BM_GETCHECK = 240
	BM_GETIMAGE = 246
	BM_GETSTATE = 242
	BM_SETCHECK = 241
	BM_SETIMAGE = 247
	BM_SETSTATE = 243
	BM_SETSTYLE = 244
)

// Button notifications
const (
	BCN_DROPDOWN     = 0xfffffb20
	BN_CLICKED       = 0
	BN_PAINT         = 1
	BN_HILITE        = 2
	BN_PUSHED        = BN_HILITE
	BN_UNHILITE      = 3
	BN_UNPUSHED      = BN_UNHILITE
	BN_DISABLE       = 4
	BN_DOUBLECLICKED = 5
	BN_DBLCLK        = BN_DOUBLECLICKED
	BN_SETFOCUS      = 6
	BN_KILLFOCUS     = 7
)

const (
	IMAGE_BITMAP      = 0
	IMAGE_ICON        = 1
	IMAGE_CURSOR      = 2
	IMAGE_ENHMETAFILE = 3
)

const (
	LR_DEFAULTCOLOR     = 0
	LR_MONOCHROME       = 1
	LR_COLOR            = 2
	LR_COPYRETURNORG    = 4
	LR_COPYDELETEORG    = 8
	LR_LOADFROMFILE     = 16
	LR_LOADTRANSPARENT  = 32
	LR_LOADREALSIZE     = 128
	LR_DEFAULTSIZE      = 0x0040
	LR_VGACOLOR         = 0x0080
	LR_LOADMAP3DCOLORS  = 4096
	LR_CREATEDIBSECTION = 8192
	LR_COPYFROMRESOURCE = 0x4000
	LR_SHARED           = 32768
)

// Button style constants
const (
	BS_3STATE          = 5
	BS_AUTO3STATE      = 6
	BS_AUTOCHECKBOX    = 3
	BS_AUTORADIOBUTTON = 9
	BS_BITMAP          = 128
	BS_BOTTOM          = 0x800
	BS_CENTER          = 0x300
	BS_CHECKBOX        = 2
	BS_DEFPUSHBUTTON   = 1
	BS_GROUPBOX        = 7
	BS_ICON            = 64
	BS_LEFT            = 256
	BS_LEFTTEXT        = 32
	BS_MULTILINE       = 0x2000
	BS_NOTIFY          = 0x4000
	BS_OWNERDRAW       = 0xB
	BS_PUSHBUTTON      = 0
	BS_PUSHLIKE        = 4096
	BS_RADIOBUTTON     = 4
	BS_RIGHT           = 512
	BS_RIGHTBUTTON     = 32
	BS_SPLITBUTTON     = 0x0000000c
	BS_TEXT            = 0
	BS_TOP             = 0x400
	BS_USERBUTTON      = 8
	BS_VCENTER         = 0xC00
	BS_FLAT            = 0x8000
)

const (
	PM_NOREMOVE = 0x0000
	PM_REMOVE   = 0x0001
	PM_NOYIELD  = 0x0002
)

// Button state constants
const (
	BST_CHECKED       = 1
	BST_INDETERMINATE = 2
	BST_UNCHECKED     = 0
	BST_FOCUS         = 8
	BST_PUSHED        = 4
)

// Predefined brushes constants
const (
	COLOR_3DDKSHADOW              = 21
	COLOR_3DFACE                  = 15
	COLOR_3DHILIGHT               = 20
	COLOR_3DHIGHLIGHT             = 20
	COLOR_3DLIGHT                 = 22
	COLOR_BTNHILIGHT              = 20
	COLOR_3DSHADOW                = 16
	COLOR_ACTIVEBORDER            = 10
	COLOR_ACTIVECAPTION           = 2
	COLOR_APPWORKSPACE            = 12
	COLOR_BACKGROUND              = 1
	COLOR_DESKTOP                 = 1
	COLOR_BTNFACE                 = 15
	COLOR_BTNHIGHLIGHT            = 20
	COLOR_BTNSHADOW               = 16
	COLOR_BTNTEXT                 = 18
	COLOR_CAPTIONTEXT             = 9
	COLOR_GRAYTEXT                = 17
	COLOR_HIGHLIGHT               = 13
	COLOR_HIGHLIGHTTEXT           = 14
	COLOR_INACTIVEBORDER          = 11
	COLOR_INACTIVECAPTION         = 3
	COLOR_INACTIVECAPTIONTEXT     = 19
	COLOR_INFOBK                  = 24
	COLOR_INFOTEXT                = 23
	COLOR_MENU                    = 4
	COLOR_MENUTEXT                = 7
	COLOR_SCROLLBAR               = 0
	COLOR_WINDOW                  = 5
	COLOR_WINDOWFRAME             = 6
	COLOR_WINDOWTEXT              = 8
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
)

// GetAncestor flags
const (
	GA_PARENT    = 1
	GA_ROOT      = 2
	GA_ROOTOWNER = 3
)

// GetWindowLong and GetWindowLongPtr constants
const (
	GWL_EXSTYLE     = -20
	GWL_STYLE       = -16
	GWL_WNDPROC     = -4
	GWLP_WNDPROC    = -4
	GWL_HINSTANCE   = -6
	GWLP_HINSTANCE  = -6
	GWL_HWNDPARENT  = -8
	GWLP_HWNDPARENT = -8
	GWL_ID          = -12
	GWLP_ID         = -12
	GWL_USERDATA    = -21
	GWLP_USERDATA   = -21
)

// Predefined window handles
const (
	HWND_BROADCAST = HWND(0xFFFF)
	HWND_BOTTOM    = HWND(1)
	HWND_NOTOPMOST = ^HWND(1) // -2
	HWND_TOP       = HWND(0)
	HWND_TOPMOST   = ^HWND(0) // -1
	HWND_DESKTOP   = HWND(0)
	HWND_MESSAGE   = ^HWND(2) // -3
)

// Predefined icon constants
const (
	IDI_APPLICATION = 32512
	IDI_HAND        = 32513
	IDI_QUESTION    = 32514
	IDI_EXCLAMATION = 32515
	IDI_ASTERISK    = 32516
	IDI_WINLOGO     = 32517
	IDI_SHIELD      = 32518
	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)

// Predefined cursor constants
const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_ICON        = 32641
	IDC_SIZE        = 32640
)

// GetSystemMetrics constants
const (
	SM_CXSCREEN             = 0
	SM_CYSCREEN             = 1
	SM_CXVSCROLL            = 2
	SM_CYHSCROLL            = 3
	SM_CYCAPTION            = 4
	SM_CXBORDER             = 5
	SM_CYBORDER             = 6
	SM_CXDLGFRAME           = 7
	SM_CYDLGFRAME           = 8
	SM_CYVTHUMB             = 9
	SM_CXHTHUMB             = 10
	SM_CXICON               = 11
	SM_CYICON               = 12
	SM_CXCURSOR             = 13
	SM_CYCURSOR             = 14
	SM_CYMENU               = 15
	SM_CXFULLSCREEN         = 16
	SM_CYFULLSCREEN         = 17
	SM_CYKANJIWINDOW        = 18
	SM_MOUSEPRESENT         = 19
	SM_CYVSCROLL            = 20
	SM_CXHSCROLL            = 21
	SM_DEBUG                = 22
	SM_SWAPBUTTON           = 23
	SM_RESERVED1            = 24
	SM_RESERVED2            = 25
	SM_RESERVED3            = 26
	SM_RESERVED4            = 27
	SM_CXMIN                = 28
	SM_CYMIN                = 29
	SM_CXSIZE               = 30
	SM_CYSIZE               = 31
	SM_CXFRAME              = 32
	SM_CYFRAME              = 33
	SM_CXMINTRACK           = 34
	SM_CYMINTRACK           = 35
	SM_CXDOUBLECLK          = 36
	SM_CYDOUBLECLK          = 37
	SM_CXICONSPACING        = 38
	SM_CYICONSPACING        = 39
	SM_MENUDROPALIGNMENT    = 40
	SM_PENWINDOWS           = 41
	SM_DBCSENABLED          = 42
	SM_CMOUSEBUTTONS        = 43
	SM_CXFIXEDFRAME         = SM_CXDLGFRAME
	SM_CYFIXEDFRAME         = SM_CYDLGFRAME
	SM_CXSIZEFRAME          = SM_CXFRAME
	SM_CYSIZEFRAME          = SM_CYFRAME
	SM_SECURE               = 44
	SM_CXEDGE               = 45
	SM_CYEDGE               = 46
	SM_CXMINSPACING         = 47
	SM_CYMINSPACING         = 48
	SM_CXSMICON             = 49
	SM_CYSMICON             = 50
	SM_CYSMCAPTION          = 51
	SM_CXSMSIZE             = 52
	SM_CYSMSIZE             = 53
	SM_CXMENUSIZE           = 54
	SM_CYMENUSIZE           = 55
	SM_ARRANGE              = 56
	SM_CXMINIMIZED          = 57
	SM_CYMINIMIZED          = 58
	SM_CXMAXTRACK           = 59
	SM_CYMAXTRACK           = 60
	SM_CXMAXIMIZED          = 61
	SM_CYMAXIMIZED          = 62
	SM_NETWORK              = 63
	SM_CLEANBOOT            = 67
	SM_CXDRAG               = 68
	SM_CYDRAG               = 69
	SM_SHOWSOUNDS           = 70
	SM_CXMENUCHECK          = 71
	SM_CYMENUCHECK          = 72
	SM_SLOWMACHINE          = 73
	SM_MIDEASTENABLED       = 74
	SM_MOUSEWHEELPRESENT    = 75
	SM_XVIRTUALSCREEN       = 76
	SM_YVIRTUALSCREEN       = 77
	SM_CXVIRTUALSCREEN      = 78
	SM_CYVIRTUALSCREEN      = 79
	SM_CMONITORS            = 80
	SM_SAMEDISPLAYFORMAT    = 81
	SM_IMMENABLED           = 82
	SM_CXFOCUSBORDER        = 83
	SM_CYFOCUSBORDER        = 84
	SM_TABLETPC             = 86
	SM_MEDIACENTER          = 87
	SM_STARTER              = 88
	SM_SERVERR2             = 89
	SM_CMETRICS             = 91
	SM_REMOTESESSION        = 0x1000
	SM_SHUTTINGDOWN         = 0x2000
	SM_REMOTECONTROL        = 0x2001
	SM_CARETBLINKINGENABLED = 0x2002
)

// ShowWindow constants
const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

// SetWindowPos flags
const (
	SWP_DRAWFRAME      = 0x0020
	SWP_FRAMECHANGED   = 0x0020
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOACTIVATE     = 0x0010
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOMOVE         = 0x0002
	SWP_NOSIZE         = 0x0001
	SWP_NOREDRAW       = 0x0008
	SWP_NOZORDER       = 0x0004
	SWP_SHOWWINDOW     = 0x0040
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOREPOSITION   = SWP_NOOWNERZORDER
	SWP_NOSENDCHANGING = 0x0400
	SWP_DEFERERASE     = 0x2000
	SWP_ASYNCWINDOWPOS = 0x4000
)

// UI state constants
const (
	UIS_SET        = 1
	UIS_CLEAR      = 2
	UIS_INITIALIZE = 3
)

// UI state constants
const (
	UISF_HIDEFOCUS = 0x1
	UISF_HIDEACCEL = 0x2
	UISF_ACTIVE    = 0x4
)

// Virtual key codes
const (
	VK_LBUTTON             = 1
	VK_RBUTTON             = 2
	VK_CANCEL              = 3
	VK_MBUTTON             = 4
	VK_XBUTTON1            = 5
	VK_XBUTTON2            = 6
	VK_BACK                = 8
	VK_TAB                 = 9
	VK_CLEAR               = 12
	VK_RETURN              = 13
	VK_SHIFT               = 16
	VK_CONTROL             = 17
	VK_MENU                = 18
	VK_PAUSE               = 19
	VK_CAPITAL             = 20
	VK_KANA                = 0x15
	VK_HANGEUL             = 0x15
	VK_HANGUL              = 0x15
	VK_JUNJA               = 0x17
	VK_FINAL               = 0x18
	VK_HANJA               = 0x19
	VK_KANJI               = 0x19
	VK_ESCAPE              = 0x1B
	VK_CONVERT             = 0x1C
	VK_NONCONVERT          = 0x1D
	VK_ACCEPT              = 0x1E
	VK_MODECHANGE          = 0x1F
	VK_SPACE               = 32
	VK_PRIOR               = 33
	VK_NEXT                = 34
	VK_END                 = 35
	VK_HOME                = 36
	VK_LEFT                = 37
	VK_UP                  = 38
	VK_RIGHT               = 39
	VK_DOWN                = 40
	VK_SELECT              = 41
	VK_PRINT               = 42
	VK_EXECUTE             = 43
	VK_SNAPSHOT            = 44
	VK_INSERT              = 45
	VK_DELETE              = 46
	VK_HELP                = 47
	VK_LWIN                = 0x5B
	VK_RWIN                = 0x5C
	VK_APPS                = 0x5D
	VK_SLEEP               = 0x5F
	VK_NUMPAD0             = 0x60
	VK_NUMPAD1             = 0x61
	VK_NUMPAD2             = 0x62
	VK_NUMPAD3             = 0x63
	VK_NUMPAD4             = 0x64
	VK_NUMPAD5             = 0x65
	VK_NUMPAD6             = 0x66
	VK_NUMPAD7             = 0x67
	VK_NUMPAD8             = 0x68
	VK_NUMPAD9             = 0x69
	VK_MULTIPLY            = 0x6A
	VK_ADD                 = 0x6B
	VK_SEPARATOR           = 0x6C
	VK_SUBTRACT            = 0x6D
	VK_DECIMAL             = 0x6E
	VK_DIVIDE              = 0x6F
	VK_F1                  = 0x70
	VK_F2                  = 0x71
	VK_F3                  = 0x72
	VK_F4                  = 0x73
	VK_F5                  = 0x74
	VK_F6                  = 0x75
	VK_F7                  = 0x76
	VK_F8                  = 0x77
	VK_F9                  = 0x78
	VK_F10                 = 0x79
	VK_F11                 = 0x7A
	VK_F12                 = 0x7B
	VK_F13                 = 0x7C
	VK_F14                 = 0x7D
	VK_F15                 = 0x7E
	VK_F16                 = 0x7F
	VK_F17                 = 0x80
	VK_F18                 = 0x81
	VK_F19                 = 0x82
	VK_F20                 = 0x83
	VK_F21                 = 0x84
	VK_F22                 = 0x85
	VK_F23                 = 0x86
	VK_F24                 = 0x87
	VK_NUMLOCK             = 0x90
	VK_SCROLL              = 0x91
	VK_LSHIFT              = 0xA0
	VK_RSHIFT              = 0xA1
	VK_LCONTROL            = 0xA2
	VK_RCONTROL            = 0xA3
	VK_LMENU               = 0xA4
	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	VK_OEM_1               = 0xBA
	VK_OEM_PLUS            = 0xBB
	VK_OEM_COMMA           = 0xBC
	VK_OEM_MINUS           = 0xBD
	VK_OEM_PERIOD          = 0xBE
	VK_OEM_2               = 0xBF
	VK_OEM_3               = 0xC0
	VK_OEM_4               = 0xDB
	VK_OEM_5               = 0xDC
	VK_OEM_6               = 0xDD
	VK_OEM_7               = 0xDE
	VK_OEM_8               = 0xDF
	VK_OEM_102             = 0xE2
	VK_PROCESSKEY          = 0xE5
	VK_PACKET              = 0xE7
	VK_ATTN                = 0xF6
	VK_CRSEL               = 0xF7
	VK_EXSEL               = 0xF8
	VK_EREOF               = 0xF9
	VK_PLAY                = 0xFA
	VK_ZOOM                = 0xFB
	VK_NONAME              = 0xFC
	VK_PA1                 = 0xFD
	VK_OEM_CLEAR           = 0xFE
)

// Window style constants
const (
	WS_OVERLAPPED       = 0x00000000
	WS_POPUP            = 0x80000000
	WS_CHILD            = 0x40000000
	WS_MINIMIZE         = 0x20000000
	WS_VISIBLE          = 0x10000000
	WS_DISABLED         = 0x08000000
	WS_CLIPSIBLINGS     = 0x04000000
	WS_CLIPCHILDREN     = 0x02000000
	WS_MAXIMIZE         = 0x01000000
	WS_CAPTION          = 0x00C00000
	WS_BORDER           = 0x00800000
	WS_DLGFRAME         = 0x00400000
	WS_VSCROLL          = 0x00200000
	WS_HSCROLL          = 0x00100000
	WS_SYSMENU          = 0x00080000
	WS_THICKFRAME       = 0x00040000
	WS_GROUP            = 0x00020000
	WS_TABSTOP          = 0x00010000
	WS_MINIMIZEBOX      = 0x00020000
	WS_MAXIMIZEBOX      = 0x00010000
	WS_TILED            = 0x00000000
	WS_ICONIC           = 0x20000000
	WS_SIZEBOX          = 0x00040000
	WS_OVERLAPPEDWINDOW = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	WS_POPUPWINDOW      = 0x80000000 | 0x00800000 | 0x00080000
	WS_CHILDWINDOW      = 0x40000000
)

// Extended window style constants
const (
	WS_EX_DLGMODALFRAME    = 0x00000001
	WS_EX_NOPARENTNOTIFY   = 0x00000004
	WS_EX_TOPMOST          = 0x00000008
	WS_EX_ACCEPTFILES      = 0x00000010
	WS_EX_TRANSPARENT      = 0x00000020
	WS_EX_MDICHILD         = 0x00000040
	WS_EX_TOOLWINDOW       = 0x00000080
	WS_EX_WINDOWEDGE       = 0x00000100
	WS_EX_CLIENTEDGE       = 0x00000200
	WS_EX_CONTEXTHELP      = 0x00000400
	WS_EX_RIGHT            = 0x00001000
	WS_EX_LEFT             = 0x00000000
	WS_EX_RTLREADING       = 0x00002000
	WS_EX_LTRREADING       = 0x00000000
	WS_EX_LEFTSCROLLBAR    = 0x00004000
	WS_EX_RIGHTSCROLLBAR   = 0x00000000
	WS_EX_CONTROLPARENT    = 0x00010000
	WS_EX_STATICEDGE       = 0x00020000
	WS_EX_APPWINDOW        = 0x00040000
	WS_EX_OVERLAPPEDWINDOW = 0x00000100 | 0x00000200
	WS_EX_PALETTEWINDOW    = 0x00000100 | 0x00000080 | 0x00000008
	WS_EX_LAYERED          = 0x00080000
	WS_EX_NOINHERITLAYOUT  = 0x00100000
	WS_EX_LAYOUTRTL        = 0x00400000
	WS_EX_COMPOSITED       = 0x02000000
	WS_EX_NOACTIVATE       = 0x08000000
)

// Window message constants
const (
	WM_APP                    = 32768
	WM_ACTIVATE               = 6
	WM_ACTIVATEAPP            = 28
	WM_AFXFIRST               = 864
	WM_AFXLAST                = 895
	WM_ASKCBFORMATNAME        = 780
	WM_CANCELJOURNAL          = 75
	WM_CANCELMODE             = 31
	WM_CAPTURECHANGED         = 533
	WM_CHANGECBCHAIN          = 781
	WM_CHAR                   = 258
	WM_CHARTOITEM             = 47
	WM_CHILDACTIVATE          = 34
	WM_CLEAR                  = 771
	WM_CLOSE                  = 16
	WM_COMMAND                = 273
	WM_COMMNOTIFY             = 68 /* OBSOLETE */
	WM_COMPACTING             = 65
	WM_COMPAREITEM            = 57
	WM_CONTEXTMENU            = 123
	WM_COPY                   = 769
	WM_COPYDATA               = 74
	WM_CREATE                 = 1
	WM_CTLCOLORBTN            = 309
	WM_CTLCOLORDLG            = 310
	WM_CTLCOLOREDIT           = 307
	WM_CTLCOLORLISTBOX        = 308
	WM_CTLCOLORMSGBOX         = 306
	WM_CTLCOLORSCROLLBAR      = 311
	WM_CTLCOLORSTATIC         = 312
	WM_CUT                    = 768
	WM_DEADCHAR               = 259
	WM_DELETEITEM             = 45
	WM_DESTROY                = 2
	WM_DESTROYCLIPBOARD       = 775
	WM_DEVICECHANGE           = 537
	WM_DEVMODECHANGE          = 27
	WM_DISPLAYCHANGE          = 126
	WM_DPICHANGED             = 0x02E0
	WM_DRAWCLIPBOARD          = 776
	WM_DRAWITEM               = 43
	WM_DROPFILES              = 563
	WM_ENABLE                 = 10
	WM_ENDSESSION             = 22
	WM_ENTERIDLE              = 289
	WM_ENTERMENULOOP          = 529
	WM_ENTERSIZEMOVE          = 561
	WM_ERASEBKGND             = 20
	WM_EXITMENULOOP           = 530
	WM_EXITSIZEMOVE           = 562
	WM_FONTCHANGE             = 29
	WM_GETDLGCODE             = 135
	WM_GETFONT                = 49
	WM_GETHOTKEY              = 51
	WM_GETICON                = 127
	WM_GETMINMAXINFO          = 36
	WM_GETTEXT                = 13
	WM_GETTEXTLENGTH          = 14
	WM_HANDHELDFIRST          = 856
	WM_HANDHELDLAST           = 863
	WM_HELP                   = 83
	WM_HOTKEY                 = 786
	WM_HSCROLL                = 276
	WM_HSCROLLCLIPBOARD       = 782
	WM_ICONERASEBKGND         = 39
	WM_INITDIALOG             = 272
	WM_INITMENU               = 278
	WM_INITMENUPOPUP          = 279
	WM_INPUT                  = 0x00FF
	WM_INPUTLANGCHANGE        = 81
	WM_INPUTLANGCHANGEREQUEST = 80
	WM_KEYDOWN                = 256
	WM_KEYUP                  = 257
	WM_KILLFOCUS              = 8
	WM_MDIACTIVATE            = 546
	WM_MDICASCADE             = 551
	WM_MDICREATE              = 544
	WM_MDIDESTROY             = 545
	WM_MDIGETACTIVE           = 553
	WM_MDIICONARRANGE         = 552
	WM_MDIMAXIMIZE            = 549
	WM_MDINEXT                = 548
	WM_MDIREFRESHMENU         = 564
	WM_MDIRESTORE             = 547
	WM_MDISETMENU             = 560
	WM_MDITILE                = 550
	WM_MEASUREITEM            = 44
	WM_GETOBJECT              = 0x003D
	WM_CHANGEUISTATE          = 0x0127
	WM_UPDATEUISTATE          = 0x0128
	WM_QUERYUISTATE           = 0x0129
	WM_UNINITMENUPOPUP        = 0x0125
	WM_MENURBUTTONUP          = 290
	WM_MENUCOMMAND            = 0x0126
	WM_MENUGETOBJECT          = 0x0124
	WM_MENUDRAG               = 0x0123
	WM_APPCOMMAND             = 0x0319
	WM_MENUCHAR               = 288
	WM_MENUSELECT             = 287
	WM_MOVE                   = 3
	WM_MOVING                 = 534
	WM_NCACTIVATE             = 134
	WM_NCCALCSIZE             = 131
	WM_NCCREATE               = 129
	WM_NCDESTROY              = 130
	WM_NCHITTEST              = 132
	WM_NCLBUTTONDBLCLK        = 163
	WM_NCLBUTTONDOWN          = 161
	WM_NCLBUTTONUP            = 162
	WM_NCMBUTTONDBLCLK        = 169
	WM_NCMBUTTONDOWN          = 167
	WM_NCMBUTTONUP            = 168
	WM_NCXBUTTONDOWN          = 171
	WM_NCXBUTTONUP            = 172
	WM_NCXBUTTONDBLCLK        = 173
	WM_NCMOUSEHOVER           = 0x02A0
	WM_NCMOUSELEAVE           = 0x02A2
	WM_NCMOUSEMOVE            = 160
	WM_NCPAINT                = 133
	WM_NCRBUTTONDBLCLK        = 166
	WM_NCRBUTTONDOWN          = 164
	WM_NCRBUTTONUP            = 165
	WM_NEXTDLGCTL             = 40
	WM_NEXTMENU               = 531
	WM_NOTIFY                 = 78
	WM_NOTIFYFORMAT           = 85
	WM_NULL                   = 0
	WM_PAINT                  = 15
	WM_PAINTCLIPBOARD         = 777
	WM_PAINTICON              = 38
	WM_PALETTECHANGED         = 785
	WM_PALETTEISCHANGING      = 784
	WM_PARENTNOTIFY           = 528
	WM_PASTE                  = 770
	WM_PENWINFIRST            = 896
	WM_PENWINLAST             = 911
	WM_POWER                  = 72
	WM_POWERBROADCAST         = 536
	WM_PRINT                  = 791
	WM_PRINTCLIENT            = 792
	WM_QUERYDRAGICON          = 55
	WM_QUERYENDSESSION        = 17
	WM_QUERYNEWPALETTE        = 783
	WM_QUERYOPEN              = 19
	WM_QUEUESYNC              = 35
	WM_QUIT                   = 18
	WM_RENDERALLFORMATS       = 774
	WM_RENDERFORMAT           = 773
	WM_SETCURSOR              = 32
	WM_SETFOCUS               = 7
	WM_SETFONT                = 48
	WM_SETHOTKEY              = 50
	WM_SETICON                = 128
	WM_SETREDRAW              = 11
	WM_SETTEXT                = 12
	WM_SETTINGCHANGE          = 26
	WM_SHOWWINDOW             = 24
	WM_SIZE                   = 5
	WM_SIZECLIPBOARD          = 779
	WM_SIZING                 = 532
	WM_SPOOLERSTATUS          = 42
	WM_STYLECHANGED           = 125
	WM_STYLECHANGING          = 124
	WM_SYSCHAR                = 262
	WM_SYSCOLORCHANGE         = 21
	WM_SYSCOMMAND             = 274
	WM_SYSDEADCHAR            = 263
	WM_SYSKEYDOWN             = 260
	WM_SYSKEYUP               = 261
	WM_TCARD                  = 82
	WM_THEMECHANGED           = 794
	WM_TIMECHANGE             = 30
	WM_TIMER                  = 275
	WM_UNDO                   = 772
	WM_USER                   = 1024
	WM_USERCHANGED            = 84
	WM_VKEYTOITEM             = 46
	WM_VSCROLL                = 277
	WM_VSCROLLCLIPBOARD       = 778
	WM_WINDOWPOSCHANGED       = 71
	WM_WINDOWPOSCHANGING      = 70
	WM_WININICHANGE           = 26
	WM_KEYFIRST               = 256
	WM_KEYLAST                = 264
	WM_SYNCPAINT              = 136
	WM_MOUSEACTIVATE          = 33
	WM_MOUSEMOVE              = 512
	WM_LBUTTONDOWN            = 513
	WM_LBUTTONUP              = 514
	WM_LBUTTONDBLCLK          = 515
	WM_RBUTTONDOWN            = 516
	WM_RBUTTONUP              = 517
	WM_RBUTTONDBLCLK          = 518
	WM_MBUTTONDOWN            = 519
	WM_MBUTTONUP              = 520
	WM_MBUTTONDBLCLK          = 521
	WM_MOUSEWHEEL             = 522
	WM_MOUSEFIRST             = 512
	WM_XBUTTONDOWN            = 523
	WM_XBUTTONUP              = 524
	WM_XBUTTONDBLCLK          = 525
	WM_MOUSELAST              = 525
	WM_MOUSEHOVER             = 0x2A1
	WM_MOUSELEAVE             = 0x2A3
	WM_CLIPBOARDUPDATE        = 0x031D
	WM_UNICHAR                = 0x0109
)

const (
	CHILDID_SELF      = 0
	INDEXID_OBJECT    = 0
	INDEXID_CONTAINER = 0

	OBJID_WINDOW            = int32(0x00000000)
	OBJID_SYSMENU           = int32(-((0xFFFFFFFF ^ 0xFFFFFFFF) + 1))
	OBJID_TITLEBAR          = int32(-((0xFFFFFFFE ^ 0xFFFFFFFF) + 1))
	OBJID_MENU              = int32(-((0xFFFFFFFD ^ 0xFFFFFFFF) + 1))
	OBJID_CLIENT            = int32(-((0xFFFFFFFC ^ 0xFFFFFFFF) + 1))
	OBJID_VSCROLL           = int32(-((0xFFFFFFFB ^ 0xFFFFFFFF) + 1))
	OBJID_HSCROLL           = int32(-((0xFFFFFFFA ^ 0xFFFFFFFF) + 1))
	OBJID_SIZEGRIP          = int32(-((0xFFFFFFF9 ^ 0xFFFFFFFF) + 1))
	OBJID_CARET             = int32(-((0xFFFFFFF8 ^ 0xFFFFFFFF) + 1))
	OBJID_CURSOR            = int32(-((0xFFFFFFF7 ^ 0xFFFFFFFF) + 1))
	OBJID_ALERT             = int32(-((0xFFFFFFF6 ^ 0xFFFFFFFF) + 1))
	OBJID_SOUND             = int32(-((0xFFFFFFF5 ^ 0xFFFFFFFF) + 1))
	OBJID_QUERYCLASSNAMEIDX = int32(-((0xFFFFFFF4 ^ 0xFFFFFFFF) + 1))
	OBJID_NATIVEOM          = int32(-((0xFFFFFFF0 ^ 0xFFFFFFFF) + 1))
)

// event constants
const (
	EVENT_MIN = 0x00000001
	EVENT_MAX = 0x7FFFFFFF

	EVENT_SYSTEM_SOUND                  = 0x0001
	EVENT_SYSTEM_ALERT                  = 0x0002
	EVENT_SYSTEM_FOREGROUND             = 0x0003
	EVENT_SYSTEM_MENUSTART              = 0x0004
	EVENT_SYSTEM_MENUEND                = 0x0005
	EVENT_SYSTEM_MENUPOPUPSTART         = 0x0006
	EVENT_SYSTEM_MENUPOPUPEND           = 0x0007
	EVENT_SYSTEM_CAPTURESTART           = 0x0008
	EVENT_SYSTEM_CAPTUREEND             = 0x0009
	EVENT_SYSTEM_MOVESIZESTART          = 0x000A
	EVENT_SYSTEM_MOVESIZEEND            = 0x000B
	EVENT_SYSTEM_CONTEXTHELPSTART       = 0x000C
	EVENT_SYSTEM_CONTEXTHELPEND         = 0x000D
	EVENT_SYSTEM_DRAGDROPSTART          = 0x000E
	EVENT_SYSTEM_DRAGDROPEND            = 0x000F
	EVENT_SYSTEM_DIALOGSTART            = 0x0010
	EVENT_SYSTEM_DIALOGEND              = 0x0011
	EVENT_SYSTEM_SCROLLINGSTART         = 0x0012
	EVENT_SYSTEM_SCROLLINGEND           = 0x0013
	EVENT_SYSTEM_SWITCHSTART            = 0x0014
	EVENT_SYSTEM_SWITCHEND              = 0x0015
	EVENT_SYSTEM_MINIMIZESTART          = 0x0016
	EVENT_SYSTEM_MINIMIZEEND            = 0x0017
	EVENT_SYSTEM_DESKTOPSWITCH          = 0x0020
	EVENT_SYSTEM_SWITCHER_APPGRABBED    = 0x0024
	EVENT_SYSTEM_SWITCHER_APPOVERTARGET = 0x0025
	EVENT_SYSTEM_SWITCHER_APPDROPPED    = 0x0026
	EVENT_SYSTEM_SWITCHER_CANCELLED     = 0x0027
	EVENT_SYSTEM_IME_KEY_NOTIFICATION   = 0x0029
	EVENT_SYSTEM_END                    = 0x00FF

	EVENT_OEM_DEFINED_START = 0x0101
	EVENT_OEM_DEFINED_END   = 0x01FF

	EVENT_CONSOLE_CARET             = 0x4001
	EVENT_CONSOLE_UPDATE_REGION     = 0x4002
	EVENT_CONSOLE_UPDATE_SIMPLE     = 0x4003
	EVENT_CONSOLE_UPDATE_SCROLL     = 0x4004
	EVENT_CONSOLE_LAYOUT            = 0x4005
	EVENT_CONSOLE_START_APPLICATION = 0x4006
	EVENT_CONSOLE_END_APPLICATION   = 0x4007
	EVENT_CONSOLE_END               = 0x40FF

	EVENT_UIA_EVENTID_START = 0x4E00
	EVENT_UIA_EVENTID_END   = 0x4EFF

	EVENT_UIA_PROPID_START = 0x7500
	EVENT_UIA_PROPID_END   = 0x75FF

	EVENT_OBJECT_CREATE                           = 0x8000
	EVENT_OBJECT_DESTROY                          = 0x8001
	EVENT_OBJECT_SHOW                             = 0x8002
	EVENT_OBJECT_HIDE                             = 0x8003
	EVENT_OBJECT_REORDER                          = 0x8004
	EVENT_OBJECT_FOCUS                            = 0x8005
	EVENT_OBJECT_SELECTION                        = 0x8006
	EVENT_OBJECT_SELECTIONADD                     = 0x8007
	EVENT_OBJECT_SELECTIONREMOVE                  = 0x8008
	EVENT_OBJECT_SELECTIONWITHIN                  = 0x8009
	EVENT_OBJECT_STATECHANGE                      = 0x800A
	EVENT_OBJECT_LOCATIONCHANGE                   = 0x800B
	EVENT_OBJECT_NAMECHANGE                       = 0x800C
	EVENT_OBJECT_DESCRIPTIONCHANGE                = 0x800D
	EVENT_OBJECT_VALUECHANGE                      = 0x800E
	EVENT_OBJECT_PARENTCHANGE                     = 0x800F
	EVENT_OBJECT_HELPCHANGE                       = 0x8010
	EVENT_OBJECT_DEFACTIONCHANGE                  = 0x8011
	EVENT_OBJECT_ACCELERATORCHANGE                = 0x8012
	EVENT_OBJECT_INVOKED                          = 0x8013
	EVENT_OBJECT_TEXTSELECTIONCHANGED             = 0x8014
	EVENT_OBJECT_CONTENTSCROLLED                  = 0x8015
	EVENT_SYSTEM_ARRANGMENTPREVIEW                = 0x8016
	EVENT_OBJECT_CLOAKED                          = 0x8017
	EVENT_OBJECT_UNCLOAKED                        = 0x8018
	EVENT_OBJECT_LIVEREGIONCHANGED                = 0x8019
	EVENT_OBJECT_HOSTEDOBJECTSINVALIDATED         = 0x8020
	EVENT_OBJECT_DRAGSTART                        = 0x8021
	EVENT_OBJECT_DRAGCANCEL                       = 0x8022
	EVENT_OBJECT_DRAGCOMPLETE                     = 0x8023
	EVENT_OBJECT_DRAGENTER                        = 0x8024
	EVENT_OBJECT_DRAGLEAVE                        = 0x8025
	EVENT_OBJECT_DRAGDROPPED                      = 0x8026
	EVENT_OBJECT_IME_SHOW                         = 0x8027
	EVENT_OBJECT_IME_HIDE                         = 0x8028
	EVENT_OBJECT_IME_CHANGE                       = 0x8029
	EVENT_OBJECT_TEXTEDIT_CONVERSIONTARGETCHANGED = 0x8030
	EVENT_OBJECT_END                              = 0x80FF

	EVENT_AIA_START = 0xa000
	EVENT_AIA_END   = 0xafff

	WINEVENT_OUTOFCONTEXT   = 0x0000
	WINEVENT_SKIPOWNTHREAD  = 0x0001
	WINEVENT_SKIPOWNPROCESS = 0x0002
	WINEVENT_INCONTEXT      = 0x0004
)

// mouse button constants
const (
	MK_CONTROL  = 0x0008
	MK_LBUTTON  = 0x0001
	MK_MBUTTON  = 0x0010
	MK_RBUTTON  = 0x0002
	MK_SHIFT    = 0x0004
	MK_XBUTTON1 = 0x0020
	MK_XBUTTON2 = 0x0040
)

// TrackPopupMenu[Ex] flags
const (
	TPM_CENTERALIGN     = 0x0004
	TPM_LEFTALIGN       = 0x0000
	TPM_RIGHTALIGN      = 0x0008
	TPM_BOTTOMALIGN     = 0x0020
	TPM_TOPALIGN        = 0x0000
	TPM_VCENTERALIGN    = 0x0010
	TPM_NONOTIFY        = 0x0080
	TPM_RETURNCMD       = 0x0100
	TPM_LEFTBUTTON      = 0x0000
	TPM_RIGHTBUTTON     = 0x0002
	TPM_HORNEGANIMATION = 0x0800
	TPM_HORPOSANIMATION = 0x0400
	TPM_NOANIMATION     = 0x4000
	TPM_VERNEGANIMATION = 0x2000
	TPM_VERPOSANIMATION = 0x1000
	TPM_HORIZONTAL      = 0x0000
	TPM_VERTICAL        = 0x0040
)

// WINDOWPLACEMENT flags
const (
	WPF_ASYNCWINDOWPLACEMENT = 0x0004
	WPF_RESTORETOMAXIMIZED   = 0x0002
	WPF_SETMINPOSITION       = 0x0001
)

// DrawText[Ex] format flags
const (
	DT_TOP                  = 0x00000000
	DT_LEFT                 = 0x00000000
	DT_CENTER               = 0x00000001
	DT_RIGHT                = 0x00000002
	DT_VCENTER              = 0x00000004
	DT_BOTTOM               = 0x00000008
	DT_WORDBREAK            = 0x00000010
	DT_SINGLELINE           = 0x00000020
	DT_EXPANDTABS           = 0x00000040
	DT_TABSTOP              = 0x00000080
	DT_NOCLIP               = 0x00000100
	DT_EXTERNALLEADING      = 0x00000200
	DT_CALCRECT             = 0x00000400
	DT_NOPREFIX             = 0x00000800
	DT_INTERNAL             = 0x00001000
	DT_EDITCONTROL          = 0x00002000
	DT_PATH_ELLIPSIS        = 0x00004000
	DT_END_ELLIPSIS         = 0x00008000
	DT_MODIFYSTRING         = 0x00010000
	DT_RTLREADING           = 0x00020000
	DT_WORD_ELLIPSIS        = 0x00040000
	DT_NOFULLWIDTHCHARBREAK = 0x00080000
	DT_HIDEPREFIX           = 0x00100000
	DT_PREFIXONLY           = 0x00200000
)

// Window class styles
const (
	CS_VREDRAW         = 0x00000001
	CS_HREDRAW         = 0x00000002
	CS_KEYCVTWINDOW    = 0x00000004
	CS_DBLCLKS         = 0x00000008
	CS_OWNDC           = 0x00000020
	CS_CLASSDC         = 0x00000040
	CS_PARENTDC        = 0x00000080
	CS_NOKEYCVT        = 0x00000100
	CS_NOCLOSE         = 0x00000200
	CS_SAVEBITS        = 0x00000800
	CS_BYTEALIGNCLIENT = 0x00001000
	CS_BYTEALIGNWINDOW = 0x00002000
	CS_GLOBALCLASS     = 0x00004000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)

// SystemParametersInfo actions
const (
	SPI_GETNONCLIENTMETRICS = 0x0029
	SPI_GETHIGHCONTRAST     = 0x0042
)

// Dialog styles
const (
	DS_ABSALIGN      = 0x0001
	DS_SYSMODAL      = 0x0002
	DS_3DLOOK        = 0x0004
	DS_FIXEDSYS      = 0x0008
	DS_NOFAILCREATE  = 0x0010
	DS_LOCALEDIT     = 0x0020
	DS_SETFONT       = 0x0040
	DS_MODALFRAME    = 0x0080
	DS_NOIDLEMSG     = 0x0100
	DS_SETFOREGROUND = 0x0200
	DS_CONTROL       = 0x0400
	DS_CENTER        = 0x0800
	DS_CENTERMOUSE   = 0x1000
	DS_CONTEXTHELP   = 0x2000
	DS_USEPIXELS     = 0x8000
	DS_SHELLFONT     = (DS_SETFONT | DS_FIXEDSYS)
)

// WM_GETDLGCODE return values
const (
	DLGC_BUTTON          = 0x2000
	DLGC_DEFPUSHBUTTON   = 0x0010
	DLGC_HASSETSEL       = 0x0008
	DLGC_RADIOBUTTON     = 0x0040
	DLGC_STATIC          = 0x0100
	DLGC_UNDEFPUSHBUTTON = 0x0020
	DLGC_WANTALLKEYS     = 0x0004
	DLGC_WANTARROWS      = 0x0001
	DLGC_WANTCHARS       = 0x0080
	DLGC_WANTMESSAGE     = 0x0004
	DLGC_WANTTAB         = 0x0002
)

// WM_ACTIVATE codes
const (
	WA_ACTIVE      = 1
	WA_CLICKACTIVE = 2
	WA_INACTIVE    = 0
)

// Owner drawing actions
const (
	ODA_DRAWENTIRE = 0x0001
	ODA_FOCUS      = 0x0002
	ODA_SELECT     = 0x0004
)

// Owner drawing states
const (
	ODS_CHECKED      = 0x0001
	ODS_COMBOBOXEDIT = 0x0002
	ODS_DEFAULT      = 0x0004
	ODS_DISABLED     = 0x0008
	ODS_FOCUS        = 0x0010
	ODS_GRAYED       = 0x0020
	ODS_SELECTED     = 0x0040
)

// Raw input device flags
const (
	RIDEV_APPKEYS      = 0x00000400
	RIDEV_CAPTUREMOUSE = 0x00000200
	RIDEV_DEVNOTIFY    = 0x00002000
	RIDEV_EXCLUDE      = 0x00000010
	RIDEV_EXINPUTSINK  = 0x00001000
	RIDEV_INPUTSINK    = 0x00000100
	RIDEV_NOHOTKEYS    = 0x00000200
	RIDEV_NOLEGACY     = 0x00000030
	RIDEV_PAGEONLY     = 0x00000020
	RIDEV_REMOVE       = 0x00000001
)

// Raw input device command flags
const (
	RID_HEADER = 0x10000005
	RID_INPUT  = 0x10000003
)

// Raw input type
const (
	RIM_TYPEHID      = 2
	RIM_TYPEKEYBOARD = 1
	RIM_TYPEMOUSE    = 0
)

// Raw input scan code information
const (
	RI_KEY_MAKE  = 0
	RI_KEY_BREAK = 1
	RI_KEY_E0    = 2
	RI_KEY_E1    = 4
)

// Raw input mouse state
const (
	MOUSE_MOVE_RELATIVE      = 0x00
	MOUSE_MOVE_ABSOLUTE      = 0x01
	MOUSE_VIRTUAL_DESKTOP    = 0x02
	MOUSE_ATTRIBUTES_CHANGED = 0x04
)

// Raw input transistion state of mouse buttons
const (
	RI_MOUSE_LEFT_BUTTON_DOWN   = 0x0001
	RI_MOUSE_LEFT_BUTTON_UP     = 0x0002
	RI_MOUSE_MIDDLE_BUTTON_DOWN = 0x0010
	RI_MOUSE_MIDDLE_BUTTON_UP   = 0x0020
	RI_MOUSE_RIGHT_BUTTON_DOWN  = 0x0004
	RI_MOUSE_RIGHT_BUTTON_UP    = 0x0008
	RI_MOUSE_BUTTON_1_DOWN      = 0x0001
	RI_MOUSE_BUTTON_1_UP        = 0x0002
	RI_MOUSE_BUTTON_2_DOWN      = 0x0004
	RI_MOUSE_BUTTON_2_UP        = 0x0008
	RI_MOUSE_BUTTON_3_DOWN      = 0x0010
	RI_MOUSE_BUTTON_3_UP        = 0x0020
	RI_MOUSE_BUTTON_4_DOWN      = 0x0040
	RI_MOUSE_BUTTON_4_UP        = 0x0080
	RI_MOUSE_BUTTON_5_DOWN      = 0x0100
	RI_MOUSE_BUTTON_5_UP        = 0x0200
	RI_MOUSE_WHEEL              = 0x0400
)

// Multi monitor constants
const (
	MONITOR_DEFAULTTONULL    = 0x0
	MONITOR_DEFAULTTOPRIMARY = 0x1
	MONITOR_DEFAULTTONEAREST = 0x2
)

// MONITORINFO flags
const (
	MONITORINFOF_PRIMARY = 0x1
)

// INPUT Type
const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

// MOUSEINPUT MouseData
const (
	XBUTTON1 = 0x0001
	XBUTTON2 = 0x0002
)

// MOUSEINPUT DwFlags
const (
	MOUSEEVENTF_ABSOLUTE        = 0x8000
	MOUSEEVENTF_HWHEEL          = 0x1000
	MOUSEEVENTF_MOVE            = 0x0001
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000
	MOUSEEVENTF_LEFTDOWN        = 0x0002
	MOUSEEVENTF_LEFTUP          = 0x0004
	MOUSEEVENTF_RIGHTDOWN       = 0x0008
	MOUSEEVENTF_RIGHTUP         = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      = 0x0020
	MOUSEEVENTF_MIDDLEUP        = 0x0040
	MOUSEEVENTF_VIRTUALDESK     = 0x4000
	MOUSEEVENTF_WHEEL           = 0x0800
	MOUSEEVENTF_XDOWN           = 0x0080
	MOUSEEVENTF_XUP             = 0x0100
)

// KEYBDINPUT DwFlags
const (
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002
	KEYEVENTF_SCANCODE    = 0x0008
	KEYEVENTF_UNICODE     = 0x0004
)

// GetWindow uCmd constants
const (
	GW_CHILD        = 5
	GW_ENABLEDPOPUP = 6
	GW_HWNDFIRST    = 0
	GW_HWNDLAST     = 1
	GW_HWNDNEXT     = 2
	GW_HWNDPREV     = 3
	GW_OWNER        = 4
)

// Standard clipboard formats
const (
	CF_BITMAP          = 2
	CF_DIB             = 8
	CF_DIBV5           = 17
	CF_DIF             = 5
	CF_DSPBITMAP       = 0x0082
	CF_DSPENHMETAFILE  = 0x008E
	CF_DSPMETAFILEPICT = 0x0083
	CF_DSPTEXT         = 0x0081
	CF_ENHMETAFILE     = 14
	CF_GDIOBJFIRST     = 0x0300
	CF_GDIOBJLAST      = 0x03FF
	CF_HDROP           = 15
	CF_LOCALE          = 16
	CF_METAFILEPICT    = 3
	CF_OEMTEXT         = 7
	CF_OWNERDISPLAY    = 0x0080
	CF_PALETTE         = 9
	CF_PENDATA         = 10
	CF_PRIVATEFIRST    = 0x0200
	CF_PRIVATELAST     = 0x02FF
	CF_RIFF            = 11
	CF_SYLK            = 4
	CF_TEXT            = 1
	CF_TIFF            = 6
	CF_UNICODETEXT     = 13
	CF_WAVE            = 12
)

// ScrollBar constants
const (
	SB_HORZ = 0
	SB_VERT = 1
	SB_CTL  = 2
	SB_BOTH = 3
)

// ScrollBar commands
const (
	SB_LINEUP        = 0
	SB_LINELEFT      = 0
	SB_LINEDOWN      = 1
	SB_LINERIGHT     = 1
	SB_PAGEUP        = 2
	SB_PAGELEFT      = 2
	SB_PAGEDOWN      = 3
	SB_PAGERIGHT     = 3
	SB_THUMBPOSITION = 4
	SB_THUMBTRACK    = 5
	SB_TOP           = 6
	SB_LEFT          = 6
	SB_BOTTOM        = 7
	SB_RIGHT         = 7
	SB_ENDSCROLL     = 8
)

// [Get|Set]ScrollInfo mask constants
const (
	SIF_RANGE           = 1
	SIF_PAGE            = 2
	SIF_POS             = 4
	SIF_DISABLENOSCROLL = 8
	SIF_TRACKPOS        = 16
	SIF_ALL             = SIF_RANGE + SIF_PAGE + SIF_POS + SIF_TRACKPOS
)

// DrawIconEx flags
const (
	DI_COMPAT      = 0x0004
	DI_DEFAULTSIZE = 0x0008
	DI_IMAGE       = 0x0002
	DI_MASK        = 0x0001
	DI_NOMIRROR    = 0x0010
	DI_NORMAL      = DI_IMAGE | DI_MASK
)

// WM_NCHITTEST constants
const (
	HTBORDER      = 18
	HTBOTTOM      = 15
	HTBOTTOMLEFT  = 16
	HTBOTTOMRIGHT = 17
	HTCAPTION     = 2
	HTCLIENT      = 1
	HTCLOSE       = 20
	HTERROR       = -2
	HTGROWBOX     = 4
	HTHELP        = 21
	HTHSCROLL     = 6
	HTLEFT        = 10
	HTMENU        = 5
	HTMAXBUTTON   = 9
	HTMINBUTTON   = 8
	HTNOWHERE     = 0
	HTREDUCE      = 8
	HTRIGHT       = 11
	HTSIZE        = 4
	HTSYSMENU     = 3
	HTTOP         = 12
	HTTOPLEFT     = 13
	HTTOPRIGHT    = 14
	HTTRANSPARENT = -1
	HTVSCROLL     = 7
	HTZOOM        = 9
)

// AnimateWindow flags
const (
	AW_ACTIVATE     = 0x00020000
	AW_BLEND        = 0x00080000
	AW_CENTER       = 0x00000010
	AW_HIDE         = 0x00010000
	AW_HOR_POSITIVE = 0x00000001
	AW_HOR_NEGATIVE = 0x00000002
	AW_SLIDE        = 0x00040000
	AW_VER_POSITIVE = 0x00000004
	AW_VER_NEGATIVE = 0x00000008
)

// Session ending constants
const (
	ENDSESSION_CLOSEAPP = 0x00000001
	ENDSESSION_CRITICAL = 0x40000000
	ENDSESSION_LOGOFF   = 0x80000000
)

// ChangeWindowMessageFilterEx constants
const (
	MSGFLT_RESET    = 0
	MSGFLT_ALLOW    = 1
	MSGFLT_DISALLOW = 2

	MSGFLTINFO_NONE                     = 0
	MSGFLTINFO_ALREADYALLOWED_FORWND    = 1
	MSGFLTINFO_ALREADYDISALLOWED_FORWND = 2
	MSGFLTINFO_ALLOWED_HIGHER           = 3
)

// TRACKMOUSEEVENT flags
const (
	TME_CANCEL    = 0x80000000
	TME_HOVER     = 0x00000001
	TME_LEAVE     = 0x00000002
	TME_NONCLIENT = 0x00000010
	TME_QUERY     = 0x40000000
)

// HIGHCONTRAST flags
const (
	HCF_HIGHCONTRASTON  = 0x00000001
	HCF_AVAILABLE       = 0x00000002
	HCF_HOTKEYACTIVE    = 0x00000004
	HCF_CONFIRMHOTKEY   = 0x00000008
	HCF_HOTKEYSOUND     = 0x00000010
	HCF_INDICATOR       = 0x00000020
	HCF_HOTKEYAVAILABLE = 0x00000040
)

// EDITWORDBREAKPROC codes
const (
	WB_LEFT        = 0
	WB_RIGHT       = 1
	WB_ISDELIMITER = 2
)

type NMBCDROPDOWN struct {
	Hdr      NMHDR
	RcButton RECT
}

type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}

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

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type RAWINPUTDEVICE struct {
	UsUsagePage uint16
	UsUsage     uint16
	DwFlags     uint32
	HwndTarget  HWND
}

type RAWINPUTHEADER struct {
	DwType  uint32
	DwSize  uint32
	HDevice HANDLE
	WParam  uintptr
}

type RAWINPUTMOUSE struct {
	Header RAWINPUTHEADER
	Data   RAWMOUSE
}

type RAWINPUTKEYBOARD struct {
	Header RAWINPUTHEADER
	Data   RAWKEYBOARD
}

type RAWINPUTHID struct {
	Header RAWINPUTHEADER
	Data   RAWHID
}

type RAWMOUSE struct {
	UsFlags            uint16
	UsButtonFlags      uint16
	UsButtonData       uint16
	Pad_cgo_0          [2]byte
	UlRawButtons       uint32
	LLastX             int32
	LLastY             int32
	UlExtraInformation uint32
}

type RAWKEYBOARD struct {
	MakeCode         uint16
	Flags            uint16
	Reserved         int16
	VKey             uint16
	Message          uint32
	ExtraInformation uint32
}

type RAWHID struct {
	DwSizeHid uint32
	DwCount   uint32
	BRawData  [1]byte
}

type NMHDR struct {
	HwndFrom HWND
	IdFrom   uintptr
	Code     uint32
}

type CREATESTRUCT struct {
	CreateParams    uintptr
	Instance        HINSTANCE
	Menu            HMENU
	Parent          HWND
	Cy              int32
	Cx              int32
	Y               int32
	X               int32
	Style           int32
	Name, ClassName uintptr
	ExStyle         uint32
}

type CHANGEFILTERSTRUCT struct {
	size      uint32
	extStatus uint32
}

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

type TPMPARAMS struct {
	CbSize    uint32
	RcExclude RECT
}

type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    POINT
	PtMaxPosition    POINT
	RcNormalPosition RECT
}

type DRAWTEXTPARAMS struct {
	CbSize        uint32
	ITabLength    int32
	ILeftMargin   int32
	IRightMargin  int32
	UiLengthDrawn uint32
}

type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

type MINMAXINFO struct {
	PtReserved     POINT
	PtMaxSize      POINT
	PtMaxPosition  POINT
	PtMinTrackSize POINT
	PtMaxTrackSize POINT
}

type NONCLIENTMETRICS struct {
	CbSize           uint32
	IBorderWidth     int32
	IScrollWidth     int32
	IScrollHeight    int32
	ICaptionWidth    int32
	ICaptionHeight   int32
	LfCaptionFont    LOGFONT
	ISmCaptionWidth  int32
	ISmCaptionHeight int32
	LfSmCaptionFont  LOGFONT
	IMenuWidth       int32
	IMenuHeight      int32
	LfMenuFont       LOGFONT
	LfStatusFont     LOGFONT
	LfMessageFont    LOGFONT
}

type MEASUREITEMSTRUCT struct {
	CtlType    uint32
	CtlID      uint32
	ItemID     int32
	ItemWidth  uint32
	ItemHeight uint32
	ItemData   uintptr
}

type DRAWITEMSTRUCT struct {
	CtlType    uint32
	CtlID      uint32
	ItemID     int32
	ItemAction uint32
	ItemState  uint32
	HwndItem   HWND
	HDC        HDC
	RcItem     RECT
	ItemData   uintptr
}

type ICONINFO struct {
	FIcon    BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

type MOUSE_INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KEYBD_INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

type HARDWARE_INPUT struct {
	Type uint32
	Hi   HARDWAREINPUT
}

type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
	Unused  [16]byte
}

type SCROLLINFO struct {
	CbSize    uint32
	FMask     uint32
	NMin      int32
	NMax      int32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}

type WINDOWPOS struct {
	Hwnd            HWND
	HwndInsertAfter HWND
	X               int32
	Y               int32
	Cx              int32
	Cy              int32
	Flags           uint32
}

type TRACKMOUSEEVENT struct {
	CbSize      uint32
	DwFlags     uint32
	HwndTrack   HWND
	DwHoverTime uint32
}

type HIGHCONTRAST struct {
	CbSize            uint32
	DwFlags           uint32
	LpszDefaultScheme *uint16
}
