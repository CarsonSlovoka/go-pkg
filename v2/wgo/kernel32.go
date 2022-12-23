//go:build windows

package wgo

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
)

func (w *WGO) GetProcessEntry(filter func(entry *w32.PROCESSENTRY32W) bool) (entrySlice []w32.PROCESSENTRY32W, errno syscall.Errno) {
	var handleSnapshot w32.HANDLE
	handleSnapshot, errno = w.kernel.CreateToolHelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if uintptr(handleSnapshot) == w32.INVALID_HANDLE_VALUE {
		return nil, errno
	}

	defer func() {
		if errno == 0 {
			errno = w.kernel.CloseHandle(handleSnapshot)
		} else {
			_ = w.kernel.CloseHandle(handleSnapshot) // 使用之前的錯誤代碼
		}
	}()

	var ok int32
	entry := w32.NewPROCESSENTRY32W()

	ok, errno = w.kernel.Process32First(handleSnapshot, entry)
	if ok == 0 {
		if errno == w32.ERROR_NO_MORE_FILES {
			return nil, 0
		}
		return nil, errno
	}

	var numProcess = 0
	for {
		ok, errno = w.kernel.Process32Next(handleSnapshot, entry)
		if ok == 0 {
			if errno == w32.ERROR_NO_MORE_FILES {
				return entrySlice, 0
			}
			return entrySlice, errno
		}
		numProcess++

		if filter != nil && !filter(entry) {
			continue
		}
		if entrySlice == nil {
			entrySlice = make([]w32.PROCESSENTRY32W, 0)
		}
		entrySlice = append(entrySlice, *entry)
	}
}

// KillProcess Call GetProcessEntry to get entrySlice
func (w *WGO) KillProcess(
	entrySlice []w32.PROCESSENTRY32W,
	callback func(entry *w32.PROCESSENTRY32W, errno syscall.Errno),
) {
	var (
		handle         w32.HANDLE
		errno1, errno2 syscall.Errno
	)
	for _, entry := range entrySlice {
		handle, _ = w.kernel.OpenProcess(w32.PROCESS_TERMINATE, false, entry.Th32ProcessID)
		if handle == 0 {
			continue
		}
		errno1 = w.kernel.TerminateProcess(handle, 0)
		errno2 = w.kernel.CloseHandle(handle)
		if callback != nil {
			if errno1 == 0 && errno2 == 0 {
				callback(&entry, 0)
			} else if errno1 != 0 {
				callback(&entry, errno1)
			} else {
				callback(&entry, errno2)
			}
		}
	}
}
