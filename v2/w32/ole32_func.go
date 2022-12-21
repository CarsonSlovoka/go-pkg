//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNCLSIDFromProgID  ProcName = "CLSIDFromProgID"
	PNCLSIDFromString  ProcName = "CLSIDFromString"
	PNCoCreateInstance ProcName = "CoCreateInstance"
	PNCoInitialize     ProcName = "CoInitialize"
	PNCoInitializeEx   ProcName = "CoInitializeEx"
	PNCoUnInitialize   ProcName = "CoUninitialize"
)

type Ole32DLL struct {
	*dLL
}

// NewOle32DLL You can import the interested process only instead of the whole process.
// We will load the entire process list if you do not give any parameters.
// We encourage you to fill in only the items you are using, as importing the all item will cause unnecessary waste.
func NewOle32DLL(procList ...ProcName) *Ole32DLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNCLSIDFromProgID,
			PNCLSIDFromString,
			PNCoCreateInstance,
			PNCoInitialize,
			PNCoInitializeEx,
			PNCoUnInitialize,
		}
	}
	dll := newDll(DNOle32, procList)
	return &Ole32DLL{dll}
}

// CLSIDFromProgID https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
// This function can return the following values.
// S_OK, CO_E_CLASSSTRING, REGDB_E_WRITEREGDB
func (dll *Ole32DLL) CLSIDFromProgID(progID string, guid *GUID /*[out]*/) syscall.Errno {
	proc := dll.mustProc(PNCLSIDFromProgID)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(progID),
		uintptr(unsafe.Pointer(guid)),
	)
	return syscall.Errno(r)
}

// CLSIDFromString https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromstring
// return values:
// NOERROR, CO_E_CLASSSTRING, REGDB_E_CLASSNOTREG, REGDB_E_READREGDB
func (dll *Ole32DLL) CLSIDFromString(str string, guid *GUID /*[out]*/) syscall.Errno {
	proc := dll.mustProc(PNCLSIDFromString)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(str),
		uintptr(unsafe.Pointer(guid)),
	)
	return syscall.Errno(hr)
}

// CoCreateInstance https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
// This function can return the following values.
// S_OK, REGDB_E_CLASSNOTREG, CLASS_E_NOAGGREGATION, E_NOINTERFACE, E_POINTER
func (dll *Ole32DLL) CoCreateInstance(clsID *GUID,
	dwClsContext uint32, // https://learn.microsoft.com/en-us/windows/win32/api/wtypesbase/ne-wtypesbase-clsctx
	riID *GUID,
	// unk *IUnknown, // [out] // 不可以放這邊，離開函數之後數值會不見
) (*IUnknown, syscall.Errno) {
	proc := dll.mustProc(PNCoCreateInstance)
	if riID == nil {
		riID = IID_IUnknown
	}
	var unk *IUnknown
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(clsID)),
		0,
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(riID)),
		uintptr(unsafe.Pointer(&unk)),
	)
	return unk, syscall.Errno(hr)
}

// CoInitialize https://learn.microsoft.com/en-us/windows/win32/api/objbase/nf-objbase-coinitialize
// This parameter is reserved and must be NULL.
func (dll *Ole32DLL) CoInitialize(pvReserved LPVOID) HRESULT {
	proc := dll.mustProc(PNCoInitialize)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(pvReserved),
	)
	return HRESULT(r1)
}

// CoInitializeEx https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
// dwCoInit: COINIT_MULTITHREADED, ...
func (dll *Ole32DLL) CoInitializeEx(
	pvReserved LPVOID, // This parameter is reserved and must be NULL.
	dwCoInit uint32,
) HRESULT {
	proc := dll.mustProc(PNCoInitializeEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(0),
		uintptr(dwCoInit),
	)
	return HRESULT(r1)
}

// CoUnInitialize https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func (dll *Ole32DLL) CoUnInitialize() {
	proc := dll.mustProc(PNCoUnInitialize)
	_, _, _ = syscall.SyscallN(proc.Addr())
}
