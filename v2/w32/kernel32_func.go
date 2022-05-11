//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNCreateMutex         ProcName = "CreateMutexW"
	PNCloseHandle         ProcName = "CloseHandle"
	PNGetNativeSystemInfo ProcName = "GetNativeSystemInfo"
)

type Kernel32DLL struct {
	*dLL
}

func NewKernel32DLL(procList []ProcName) *Kernel32DLL {
	dll := newDll(DN_KERNEL32, procList)
	return &Kernel32DLL{dll}
}

// CloseHandle Closes an open object handle.
// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle?redirectedfrom=MSDN
func (dll *Kernel32DLL) CloseHandle(handle uintptr) error {
	proc := dll.mustProc(PNCloseHandle)
	// r0, _, err := proc.Call(handle) // 其為syscall.SyscallN的封裝(多了檢查的動作)，如果已經確定，可以直接用syscall.SyscallN會更有效率
	r0, _, errno := syscall.SyscallN(proc.Addr(), handle)
	if r0 == 0 {
		return errno
	}
	return nil
}

// CreateMutex You can use it to restrict to a single instance of executable
// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createmutexW#return-value
func (dll *Kernel32DLL) CreateMutex(name string) (handle uintptr, err error) {
	proc := dll.mustProc(PNCreateMutex)
	lpName, _ := syscall.UTF16PtrFromString(name) // LPCWSTR
	handle, _, errno := syscall.SyscallN(proc.Addr(), 0, 0, uintptr(unsafe.Pointer(lpName)))
	/*
		handle, _, err = proc.Call(
			0,
			0,
			uintptr(unsafe.Pointer(lpName)),
		)
		if err.(syscall.Errno) == 0 {
			return handle, nil
		}
	*/
	if errno == 0 {
		return handle, nil
	}
	return handle, errno
}

// GetNativeSystemInfo
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getnativesysteminfo
func (dll *Kernel32DLL) GetNativeSystemInfo() (info SYSTEM_INFO) {
	proc := dll.mustProc(PNGetNativeSystemInfo)
	_, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(&info)))
	return
}
