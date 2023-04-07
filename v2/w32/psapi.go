package w32

// https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocessmodulesex#parameters
const (
	LIST_MODULES_DEFAULT uint32 = 0
	LIST_MODULES_32BIT   uint32 = 0x01
	LIST_MODULES_64BIT   uint32 = 0x02
	LIST_MODULES_ALL     uint32 = 0x03
)
