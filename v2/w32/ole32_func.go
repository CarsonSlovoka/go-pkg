//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNCLSIDFromProgID  ProcName = "CLSIDFromProgID"
	PNCLSIDFromString  ProcName = "CLSIDFromString"
	PNCoCreateGuid     ProcName = "CoCreateGuid"
	PNCoCreateInstance ProcName = "CoCreateInstance"
	PNCoInitialize     ProcName = "CoInitialize"
	PNCoInitializeEx   ProcName = "CoInitializeEx"
	PNCoTaskMemFree    ProcName = "CoTaskMemFree"
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
			PNCoCreateGuid,
			PNCoCreateInstance,
			PNCoInitialize,
			PNCoInitializeEx,
			PNCoTaskMemFree,
			PNCoUnInitialize,
		}
	}
	dll := newDll(DNOle32, procList)
	return &Ole32DLL{dll}
}

// CLSIDFromProgID https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
// The CLSID format is {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}.
// ProgID 是一個字串它記錄在: HKEY_LOCAL_MACHINE\SOFTWARE\Classes\<progID>\CLSID 可以知道此progID會對應到哪一個CLSID，
// 它所對應的CLSID紀錄在 HKEY_LOCAL_MACHINE\SOFTWARE\Classes\CLSID // https://learn.microsoft.com/en-us/windows/win32/com/progid
// This function can return the following values.
// S_OK, CO_E_CLASSSTRING, REGDB_E_WRITEREGDB
func (dll *Ole32DLL) CLSIDFromProgID(progID string, clsID *GUID /*[out]*/) syscall.Errno {
	proc := dll.mustProc(PNCLSIDFromProgID)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(progID),
		uintptr(unsafe.Pointer(clsID)),
	)
	return syscall.Errno(r)
}

// CLSIDFromString https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromstring
// 例如輸入: "{50AC103F-D235-4598-BBEF-98FE4D1A3AD4}" 會將此字串轉為GUID
// 可以使用 NewGUID 來代替此函數
// return values:
// NOERROR, CO_E_CLASSSTRING
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

// CoCreateGuid https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateguid?redirectedfrom=MSDN
// Creates a GUID, a unique 128-bit integer used for CLSIDs and interface identifiers.
func (dll *Ole32DLL) CoCreateGuid(guid *GUID) HRESULT {
	proc := dll.mustProc(PNCoCreateGuid)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(guid)),
	)
	return HRESULT(r1)
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

// CoTaskMemFree https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
// Frees a block of task memory previously allocated through a call to the CoTaskMemAlloc or CoTaskMemRealloc function.
func (dll *Ole32DLL) CoTaskMemFree(
	address unsafe.Pointer, // A pointer to the memory block to be freed. If this parameter is NULL, the function has no effect.
) {
	proc := dll.mustProc(PNCoTaskMemFree)
	_, _, _ = syscall.SyscallN(proc.Addr(),
		uintptr(address),
	)
}

// CoUnInitialize https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func (dll *Ole32DLL) CoUnInitialize() {
	proc := dll.mustProc(PNCoUnInitialize)
	_, _, _ = syscall.SyscallN(proc.Addr())
}
