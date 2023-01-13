package w32

const (
	// NOTIFYICON_VERSION set NOTIFYICONDATA.uVersion with 0, 3 or 4
	NOTIFYICON_VERSION uint32 = 3

	// NOTIFYICON_VERSION_4 NTDDI_VERSION >= NTDDI_VISTA
	NOTIFYICON_VERSION_4 = 4 // https://github.com/pilight/windows-cli/blob/master/include/shellapi.h
)

// NOTIFYICONDATA
// https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw // WCHAR
// https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataa // CHAR
// 本結構採用的是W的版本，使用WCHAR
// SIZE: 968
type NOTIFYICONDATA struct {
	CbSize           uint32 // 慣用手法，透過結構的大小來區分到底是使用哪種型別, 版本, ... 而這裡主要是判別szTip的訊息長度是採用128還是64的版本
	HWnd             HWND
	UID              uint32 // 如果已經指定了GuidItem就可以省略UID
	UFlags           uint32 // 可以多種組合, NIF_ICON | NIF_GUID | NIF_SHOWTIP
	UCallbackMessage uint32
	HIcon            HICON
	// if ... // szTip有兩種選擇，一種長度採用128, 另一種長度使用64
	SzTip [128]uint16 // message when the mouse hovering // A為CHAR byte W為WCHAR: uint16
	// else ...
	// SzTip            [64]uint16
	// endif
	DwState     uint32
	DwStateMask uint32
	SzInfo      [256]uint16 // body message // A CHAR byte , W WCHAR: uint16

	/*
			union {
		    UINT uTimeout;
		    UINT uVersion;
		  }
	*/
	union1 uint32 // 可以表示:uTimeout或者uVersion，不過uTimeout只在windows XP有在用而已 // https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyicona#remarks

	SzInfoTitle  [64]uint16 // Title // A CHAR byte , W WCHAR: uint16
	DwInfoFlags  uint32     // NIIF_NONE, NIIF_INFO,
	GuidItem     GUID
	HBalloonIcon HICON
}

func (n *NOTIFYICONDATA) SetVersion(v uint32) {
	n.union1 = v
}

// SetTimeout uTimeout is valid only in Windows 2000 and Windows XP.
func (n *NOTIFYICONDATA) SetTimeout(v uint32) {
	n.union1 = v
}

type PNOTIFYICONDATA *NOTIFYICONDATA

// NotifyIcon constants
// https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
const (
	NIM_ADD        = 0x00000000
	NIM_MODIFY     = 0x00000001
	NIM_DELETE     = 0x00000002
	NIM_SETFOCUS   = 0x00000003
	NIM_SETVERSION = 0x00000004

	// uFlags
	NIF_MESSAGE  = 0x00000001
	NIF_ICON     = 0x00000002 // 程式主要圖標
	NIF_TIP      = 0x00000004
	NIF_STATE    = 0x00000008
	NIF_INFO     = 0x00000010
	NIF_GUID     = 0x00000020
	NIF_REALTIME = 0x00000040
	NIF_SHOWTIP  = 0x00000080

	NIS_HIDDEN     = 0x00000001
	NIS_SHAREDICON = 0x00000002

	// dwINfoFlags
	NIIF_NONE               = 0x00000000 // No icon.
	NIIF_INFO               = 0x00000001 // An information icon.
	NIIF_WARNING            = 0x00000002
	NIIF_ERROR              = 0x00000003
	NIIF_USER               = 0x00000004 // 自定義圖標, 在XP系統圖標看的是hIcon，新版的看的是hBalloonIcon, 建議兩個都設定 // Windows XP: hIcon as the notification, Windows Vista and later: hBalloonIcon as the notification balloon's title icon.
	NIIF_NOSOUND            = 0x00000010 // 不要音效
	NIIF_LARGE_ICON         = 0x00000020 // 圖標的尺寸為: SM_CXICON x SM_CYICON 可以使用[stock icons](https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ne-shellapi-shstockiconid) 也建議設定hicon與hBalloonIcon
	NIIF_RESPECT_QUIET_TIME = 0x00000080 // 適用在windows7之後的版本, 在登錄帳戶的第一個小時, 則會忽略掉氣球通知，不要打擾客戶(尊重他們)
	NIIF_ICON_MASK          = 0x0000000F // Reserved
)

type SHSTOCKICONID int32

const (
	SIID_DOCNOASSOC        SHSTOCKICONID = 0
	SIID_DOCASSOC                        = 1
	SIID_APPLICATION                     = 2
	SIID_FOLDER                          = 3
	SIID_FOLDEROPEN                      = 4
	SIID_DRIVE525                        = 5
	SIID_DRIVE35                         = 6
	SIID_DRIVEREMOVE                     = 7
	SIID_DRIVEFIXED                      = 8
	SIID_DRIVENET                        = 9
	SIID_DRIVENETDISABLED                = 10
	SIID_DRIVECD                         = 11
	SIID_DRIVERAM                        = 12
	SIID_WORLD                           = 13
	SIID_SERVER                          = 15
	SIID_PRINTER                         = 16
	SIID_MYNETWORK                       = 17
	SIID_FIND                            = 22
	SIID_HELP                            = 23
	SIID_SHARE                           = 28
	SIID_LINK                            = 29
	SIID_SLOWFILE                        = 30
	SIID_RECYCLER                        = 31
	SIID_RECYCLERFULL                    = 32
	SIID_MEDIACDAUDIO                    = 40
	SIID_LOCK                            = 47
	SIID_AUTOLIST                        = 49
	SIID_PRINTERNET                      = 50
	SIID_SERVERSHARE                     = 51
	SIID_PRINTERFAX                      = 52
	SIID_PRINTERFAXNET                   = 53
	SIID_PRINTERFILE                     = 54
	SIID_STACK                           = 55
	SIID_MEDIASVCD                       = 56
	SIID_STUFFEDFOLDER                   = 57
	SIID_DRIVEUNKNOWN                    = 58
	SIID_DRIVEDVD                        = 59
	SIID_MEDIADVD                        = 60
	SIID_MEDIADVDRAM                     = 61
	SIID_MEDIADVDRW                      = 62
	SIID_MEDIADVDR                       = 63
	SIID_MEDIADVDROM                     = 64
	SIID_MEDIACDAUDIOPLUS                = 65
	SIID_MEDIACDRW                       = 66
	SIID_MEDIACDR                        = 67
	SIID_MEDIACDBURN                     = 68
	SIID_MEDIABLANKCD                    = 69
	SIID_MEDIACDROM                      = 70
	SIID_AUDIOFILES                      = 71
	SIID_IMAGEFILES                      = 72
	SIID_VIDEOFILES                      = 73
	SIID_MIXEDFILES                      = 74
	SIID_FOLDERBACK                      = 75
	SIID_FOLDERFRONT                     = 76
	SIID_SHIELD                          = 77
	SIID_WARNING                         = 78
	SIID_INFO                            = 79
	SIID_ERROR                           = 80
	SIID_KEY                             = 81
	SIID_SOFTWARE                        = 82
	SIID_RENAME                          = 83
	SIID_DELETE                          = 84
	SIID_MEDIAAUDIODVD                   = 85
	SIID_MEDIAMOVIEDVD                   = 86
	SIID_MEDIAENHANCEDCD                 = 87
	SIID_MEDIAENHANCEDDVD                = 88
	SIID_MEDIAHDDVD                      = 89
	SIID_MEDIABLURAY                     = 90
	SIID_MEDIAVCD                        = 91
	SIID_MEDIADVDPLUSR                   = 92
	SIID_MEDIADVDPLUSRW                  = 93
	SIID_DESKTOPPC                       = 94
	SIID_MOBILEPC                        = 95
	SIID_USERS                           = 96
	SIID_MEDIASMARTMEDIA                 = 97
	SIID_MEDIACOMPACTFLASH               = 98
	SIID_DEVICECELLPHONE                 = 99
	SIID_DEVICECAMERA                    = 100
	SIID_DEVICEVIDEOCAMERA               = 101
	SIID_DEVICEAUDIOPLAYER               = 102
	SIID_NETWORKCONNECT                  = 103
	SIID_INTERNET                        = 104
	SIID_ZIPFILE                         = 105
	SIID_SETTINGS                        = 106
	SIID_DRIVEHDDVD                      = 132
	SIID_DRIVEBD                         = 133
	SIID_MEDIAHDDVDROM                   = 134
	SIID_MEDIAHDDVDR                     = 135
	SIID_MEDIAHDDVDRAM                   = 136
	SIID_MEDIABDROM                      = 137
	SIID_MEDIABDR                        = 138
	SIID_MEDIABDRE                       = 139
	SIID_CLUSTEREDDRIVE                  = 140
	SIID_MAX_ICONS                       = 175
)
