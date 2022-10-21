//go:build windows

package w32

// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
const (
	PROCESSOR_ARCHITECTURE_AMD64   = 9
	PROCESSOR_ARCHITECTURE_ARM     = 5
	PROCESSOR_ARCHITECTURE_ARM64   = 12
	PROCESSOR_ARCHITECTURE_IA64    = 6
	PROCESSOR_ARCHITECTURE_INTEL   = 0
	PROCESSOR_ARCHITECTURE_UNKNOWN = 0xFFFF
)

// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-geticon#parameters
const (
	ICON_SMALL  = 0
	ICON_BIG    = 1
	ICON_SMALL2 = 2
)
