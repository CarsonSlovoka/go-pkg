package w32

type SECURITY_ATTRIBUTES struct {
	Length               uint32
	LpSecurityDescriptor LPVOID
	BInheritHandle       BOOL
}

// https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject#return-value
const (
	WAIT_OBJECT_0  = 0
	WAIT_ABANDONED = 0x00000080
	WAIT_FAILED    = 0xffffffff
)

const (
	IGNORE   = 0          // Ignore signal
	INFINITE = 0xFFFFFFFF // Infinite timeout
)
