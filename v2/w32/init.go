//go:build windows

package w32

import "syscall"

type DllName string

const (
	DNUser32   DllName = "User32.dll"
	DNKernel32 DllName = "Kernel32.dll"
	DNShell32  DllName = "Shell32.dll"
	DNGdi32    DllName = "Gdi32.dll"
	DNNtdll    DllName = "ntdll.dll"
	DNOle32    DllName = "Ole32.dll"
	DNOleAut32 DllName = "OleAut32.dll" // ole auto
	DNAdApi32  DllName = "Advapi32.dll"
	DNPsApi    DllName = "Psapi.dll"
)

type ProcName string

type dLL struct {
	name     DllName
	procMap  map[ProcName]*syscall.LazyProc
	mustProc func(name ProcName) *syscall.LazyProc
}

func (d *dLL) Name() string {
	return string(d.name)
}

func (d *dLL) Call(name string, args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	proc := d.mustProc(ProcName(name))
	return syscall.SyscallN(proc.Addr(), args...)
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
