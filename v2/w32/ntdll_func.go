//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNRtlGetVersion ProcName = "RtlGetVersion"
)

type NtDLL struct {
	*dLL
}

func NewNtDLL(procList ...ProcName) *NtDLL {
	dll := newDll(DNNtdll, procList)
	return &NtDLL{dll}
}

// RtlGetVersion
// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/wdm/nf-wdm-rtlgetversion
func (dll *NtDLL) RtlGetVersion() *RTL_OSVERSIONINFOEXW {
	var info RTL_OSVERSIONINFOEXW
	info.OSVersionInfoSize = 284 // 5*4 + 128*2 + 3*2 + 2*1 // 文檔中已經註明，這個欄位的值要再呼叫函數前先告知, 如果用unsafe.Sizeof(info)要在轉型，麻煩
	proc := dll.mustProc(PNRtlGetVersion)
	_, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(&info)))
	return &info
}
