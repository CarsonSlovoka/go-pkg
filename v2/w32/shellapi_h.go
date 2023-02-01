package w32

// https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw#remarks
const (
	NIN_BALLOONSHOW      = WM_USER + 2 // 0x402 // 1026
	NIN_BALLOONHIDE      = WM_USER + 3
	NIN_BALLOONTIMEOUT   = WM_USER + 4
	NIN_BALLOONUSERCLICK = WM_USER + 5
	NIN_POPUPOPEN        = WM_USER + 6
	NIN_POPUPCLOSE       = WM_USER + 7
)

//	ShellExecute() and ShellExecuteEx() error codes
//
// https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew#return-value
const (
	SE_ERR_FNF          = 2 // file not found
	SE_ERR_PNF          = 3 // path not found
	SE_ERR_ACCESSDENIED = 5 // access denied
	SE_ERR_OOM          = 8 // out of memory
	SE_ERR_DLLNOTFOUND  = 32

	SE_ERR_SHARE           = 26
	SE_ERR_ASSOCINCOMPLETE = 27
	SE_ERR_DDETIMEOUT      = 28
	SE_ERR_DDEFAIL         = 29
	SE_ERR_DDEBUSY         = 30
	SE_ERR_NOASSOC         = 31
)

// ShellExecuteEX SEE https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shellexecuteinfow#members
const (
	SEE_MASK_DEFAULT        = 0x00000000
	SEE_MASK_CLASSNAME      = 0x00000001       // ShellExeCuteInfo.lpClass is valid
	SEE_MASK_CLASSKEY       = 0x00000003       // ShellExeCuteInfo.hkeyClass is valid
	SEE_MASK_IDLIST         = 0x00000004       // ShellExeCuteInfo.lpIDList is valid
	SEE_MASK_INVOKEIDLIST   = 0x0000000c       // enable IContextMenu based verbs
	SEE_MASK_ICON           = 0x00000010       // not used
	SEE_MASK_HOTKEY         = 0x00000020       // ShellExeCuteInfo.dwHotKey is valid
	SEE_MASK_NOCLOSEPROCESS = 0x00000040       // ShellExeCuteInfo.hProcess
	SEE_MASK_CONNECTNETDRV  = 0x00000080       // enables re-connecting disconnected network drives
	SEE_MASK_NOASYNC        = 0x00000100       // block on the call until the invoke has completed, use for callers that exit after calling ShellExecuteEx()
	SEE_MASK_FLAG_DDEWAIT   = SEE_MASK_NOASYNC // Use SEE_MASK_NOASYNC instead of SEE_MASK_FLAG_DDEWAIT as it more accuratly describes the behavior
	SEE_MASK_DOENVSUBST     = 0x00000200       // indicates that ShellExeCuteInfo.lpFile contains env vars that should be expanded
	SEE_MASK_FLAG_NO_UI     = 0x00000400       // disable UI including error messages
	SEE_MASK_UNICODE        = 0x00004000
	SEE_MASK_NO_CONSOLE     = 0x00008000
	SEE_MASK_ASYNCOK        = 0x00100000

	SEE_MASK_HMONITOR     = 0x00200000 // ShellExeCuteInfo.hMonitor
	SEE_MASK_NOZONECHECKS = 0x00800000

	SEE_MASK_NOQUERYCLASSSTORE = 0x01000000
	SEE_MASK_WAITFORINPUTIDLE  = 0x02000000

	SEE_MASK_FLAG_LOG_USAGE = 0x04000000
)

// ShellExeCuteInfo https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shellexecuteinfow
type ShellExeCuteInfo struct {
	Size           uint32
	Mask           uint32
	Hwnd           HWND
	LpVerb         *uint16 // edit, explore, find, open (The file can be an executable file), print, runas, properties (Displays the file or folder's properties)
	LpFile         *uint16
	LpParams       *uint16
	LpDirectory    *uint16
	NShow          int32
	HInstApp       HINSTANCE
	LpIDList       uintptr
	LpClass        *uint16
	HKeyClass      HKEY
	DwHotKey       uint32
	HIconOrMonitor HANDLE
	HProcess       HANDLE
}
