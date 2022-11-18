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
