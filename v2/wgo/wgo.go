package wgo

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
)

type WGO struct {
	kernel *w32.Kernel32DLL
}

type dllLike interface {
	Name() string
	Call(name string, args ...uintptr) (r1, r2 uintptr, err syscall.Errno)
}

func NewWGO(dlls ...dllLike) *WGO {
	w := new(WGO)

	for _, dll := range dlls {
		switch dll.Name() {
		case string(w32.DNKernel32):
			w.kernel = dll.(*w32.Kernel32DLL)
		}
	}

	if w.kernel == nil {
		w.kernel = w32.NewKernel32DLL(
			w32.PNCreateToolHelp32Snapshot,
			w32.PNCloseHandle,
			w32.PNOpenProcess,
			w32.PNProcess32First,
			w32.PNProcess32Next,
		)
	}

	return w
}
