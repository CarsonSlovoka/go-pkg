package wgo

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
)

type WGO struct {
	kernel *w32.Kernel32DLL
	psApi  *w32.PsApiDLL
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
		case string(w32.DNPsApi):
			w.psApi = dll.(*w32.PsApiDLL)
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

	if w.psApi == nil {
		w.psApi = w32.NewPsApiDLL(
			w32.PNGetModuleFileNameExW,

			w32.PNEnumProcesses,
			w32.PNEnumProcessModules,
		)
	}

	return w
}
