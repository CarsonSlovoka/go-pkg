//go:build windows

package wgo

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
)

func GetProcessEntryByName(kernel *w32.Kernel32DLL, exeName string) (entrySlice []w32.PROCESSENTRY32W, errno syscall.Errno) {
	var handleSnapshot w32.HANDLE
	handleSnapshot, errno = kernel.CreateToolHelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if uintptr(handleSnapshot) == w32.INVALID_HANDLE_VALUE {
		return nil, errno
	}

	defer func() {
		if errno == 0 {
			_, errno = kernel.CloseHandle(handleSnapshot)
		} else {
			_, _ = kernel.CloseHandle(handleSnapshot) // 使用之前的錯誤代碼
		}
	}()

	var ok int32
	entry := w32.NewPROCESSENTRY32W()

	ok, errno = kernel.Process32First(handleSnapshot, entry)
	if ok == 0 {
		if errno == w32.ERROR_NO_MORE_FILES {
			return nil, 0
		}
		return nil, errno
	}

	var numProcess = 0
	for {
		ok, errno = kernel.Process32Next(handleSnapshot, entry)
		if ok == 0 {
			if errno == w32.ERROR_NO_MORE_FILES {
				return entrySlice, 0
			}
			return entrySlice, errno
		}
		numProcess++

		if entry.ExeFileName() == exeName {
			if entrySlice == nil {
				entrySlice = make([]w32.PROCESSENTRY32W, 0)
			}
			entrySlice = append(entrySlice, *entry)
		}
	}
}

// KillProcess Call GetProcessEntryByName to get entrySlice
func KillProcess(kernel *w32.Kernel32DLL, entrySlice []w32.PROCESSENTRY32W, callback func(deleteEntry *w32.PROCESSENTRY32W)) syscall.Errno {
	var (
		handle w32.HANDLE
		ok     bool
		errno  syscall.Errno
	)
	for _, entry := range entrySlice {
		handle, _ = kernel.OpenProcess(w32.PROCESS_TERMINATE, false, entry.Th32ProcessID)
		if handle == 0 {
			continue
		}
		if ok, errno = kernel.TerminateProcess(handle, 0); !ok {
			return errno
		}
		if ok, errno = kernel.CloseHandle(handle); !ok {
			return errno
		}
		callback(&entry)
	}
	return 0
}
