package w32

type SECURITY_ATTRIBUTES struct {
	Length               uint32
	LpSecurityDescriptor LPVOID
	BInheritHandle       BOOL
}
