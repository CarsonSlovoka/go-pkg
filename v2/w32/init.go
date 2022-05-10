package w32

import "syscall"

type DllName string

const (
	DN_USER32   DllName = "User32.dll"
	DN_KERNEL32         = "Kernel32.dll"
)

type ProcName string

type dLL struct {
	name     DllName
	procMap  map[ProcName]*syscall.LazyProc
	mustProc func(name ProcName) *syscall.LazyProc
}

// defaultMustProc 當您從dll中呼叫任何的Proc程序，該程序名稱必須事前申明需要用到
func defaultMustProc(dll dLL, name ProcName) *syscall.LazyProc {
	proc, exists := dll.procMap[name]
	if !exists {
		panic("The proc is not exist in the dll.")
	}
	return proc
}

func newDll(name DllName, procList []ProcName) *dLL {
	dll := dLL{name: name}
	lazyDLL := syscall.NewLazyDLL(string(dll.name))
	dll.procMap = make(map[ProcName]*syscall.LazyProc)
	for _, procName := range procList {
		dll.procMap[procName] = lazyDLL.NewProc(string(procName))
	}
	dll.mustProc = func(name ProcName) *syscall.LazyProc {
		return defaultMustProc(dll, name)
	}
	return &dll
}
