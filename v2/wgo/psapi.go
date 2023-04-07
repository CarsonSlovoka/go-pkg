//go:build windows

package wgo

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"path/filepath"
	"syscall"
)

// FindModuleFileName return: map[name]processID
func (w *WGO) FindModuleFileName(names ...string) map[string]uint32 {
	procIDs, numProcesses, _ := w.psApi.EnumProcesses(1024)

	var (
		hProcess  w32.HANDLE
		eno       syscall.Errno
		numModule uint32
		exePath   string
		i, j      uint32
	)

	var result map[string]uint32

	for i = uint32(0); i < numProcesses; i++ {
		hProcess, eno = w.kernel.OpenProcess(w32.PROCESS_QUERY_INFORMATION|w32.PROCESS_VM_READ, false, procIDs[i])
		if eno != 0 {
			continue
		}
		_, numModule, eno = w.psApi.EnumProcessModules(hProcess, 256)
		if eno != 0 {
			continue
		}

		if numModule > 0 {
			exePath, _ = w.psApi.GetModuleFileNameExW(hProcess, 0, 1024)
			for j = 0; j < uint32(len(names)); j++ {
				if filepath.Base(exePath) == names[j] {
					if result == nil {
						result = make(map[string]uint32, len(names))
					}
					result[names[j]] = procIDs[i]
					names = append(names[:j], names[j+1:]...) // 更新尋找清單，已經找到的不需要再找
					break
				}
			}
		}
		_ = w.kernel.CloseHandle(hProcess)
	}
	return result
}
