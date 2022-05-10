package w32

import "syscall"

type DllName string

const (
	DN_USER32 DllName = "User32.dll"
)

type ProcName string

type DllType interface {
	User32DLL
}

func defaultMustProc[D DllType](dll D, name ProcName) *syscall.LazyProc {
	proc, exists := dll.procMap[name]
	if !exists {
		panic("The proc is not exist in the dll.")
	}
	return proc
}

// NewDll 指派您要使用哪一個dll，以及使用該dll中的那些proc
func NewDll[D DllType](name DllName, procList []ProcName) *D {
	dll := D{}
	user32DLL := syscall.NewLazyDLL(string(name))
	for _, name := range procList {
		dll.procMap[name] = user32DLL.NewProc(string(name))
	}
	dll.mustProc = func(name ProcName) *syscall.LazyProc {
		return defaultMustProc(dll, name)
	}
	return &dll
}
