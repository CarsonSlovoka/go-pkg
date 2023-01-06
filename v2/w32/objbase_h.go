package w32

// https://learn.microsoft.com/en-us/windows/win32/api/objbase/ne-objbase-coinit#constants
const (
	COINIT_APARTMENTTHREADED = 0x2
	COINIT_MULTITHREADED     = 0x0
	COINIT_DISABLE_OLE1DDE   = 0x4
	COINIT_SPEED_OVER_MEMORY = 0x8
)
