package w32

type Kernel32DLL struct {
	*dLL
}

func NewKernel32DLL(procList []ProcName) *Kernel32DLL {
	dll := newDll(DN_KERNEL32, procList)
	return &Kernel32DLL{dll}
}
