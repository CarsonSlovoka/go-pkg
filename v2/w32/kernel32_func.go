package w32

import "syscall"

type Kernel32DLL struct {
	procMap  map[ProcName]*syscall.LazyProc
	mustProc func(name ProcName) *syscall.LazyProc
}
