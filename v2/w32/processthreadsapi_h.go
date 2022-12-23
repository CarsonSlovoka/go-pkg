package w32

type STARTUPINFO struct {
	Cb            uint32
	lpReserved    *uint16
	LpDesktop     *uint16
	LpTitle       *uint16
	X             uint32
	Y             uint32
	XSize         uint32
	YSize         uint32
	XCountChars   uint32
	YCountChars   uint32
	FillAttribute uint32
	Flags         uint32
	ShowWindow    uint16
	cbReserved2   uint16
	lpReserved2   *byte
	HStdInput     HANDLE
	HStdOutput    HANDLE
	HStdError     HANDLE
}

type PROCESS_INFORMATION struct {
	HProcess  HANDLE
	HThread   HANDLE
	ProcessId uint32
	ThreadId  uint32
}
