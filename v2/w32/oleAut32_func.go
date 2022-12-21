//go:build windows

package w32

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	PNSafeArrayAccessData        ProcName = "SafeArrayAccessData"
	PNSafeArrayAllocData         ProcName = "SafeArrayAllocData"
	PNSafeArrayAllocDescriptor   ProcName = "SafeArrayAllocDescriptor"
	PNSafeArrayAllocDescriptorEx ProcName = "SafeArrayAllocDescriptorEx"

	PNSafeArrayCopy     ProcName = "SafeArrayCopy"
	PNSafeArrayCopyData ProcName = "SafeArrayCopyData"

	PNSafeArrayCreate         ProcName = "SafeArrayCreate"
	PNSafeArrayCreateEx       ProcName = "SafeArrayCreateEx"
	PNSafeArrayCreateVector   ProcName = "SafeArrayCreateVector"
	PNSafeArrayCreateVectorEx ProcName = "SafeArrayCreateVectorEx"

	PNSafeArrayDestroy           ProcName = "SafeArrayDestroy"
	PNSafeArrayDestroyData       ProcName = "SafeArrayDestroyData"
	PNSafeArrayDestroyDescriptor ProcName = "SafeArrayDestroyDescriptor"

	PNSafeArrayGetDim        ProcName = "SafeArrayGetDim"
	PNSafeArrayGetElement    ProcName = "SafeArrayGetElement"
	PNSafeArrayGetElemSize   ProcName = "SafeArrayGetElemsize"
	PNSafeArrayGetIID        ProcName = "SafeArrayGetIID"
	PNSafeArrayGetLBound     ProcName = "SafeArrayGetLBound"
	PNSafeArrayGetRecordInfo ProcName = "SafeArrayGetRecordInfo"
	PNSafeArrayGetUBound     ProcName = "SafeArrayGetUBound"
	PNSafeArrayGetVarType    ProcName = "SafeArrayGetVartype"

	PNSafeArrayPutElement ProcName = "SafeArrayPutElement"

	PNSafeArrayUnAccessData ProcName = "SafeArrayUnaccessData"
	PNSafeArrayUnlock       ProcName = "SafeArrayUnlock"

	PNSafeArraySetRecordInfo ProcName = "SafeArraySetRecordInfo"

	PNSysAllocStringLen ProcName = "SysAllocStringLen"
	PNSysFreeString     ProcName = "SysFreeString"
	PNSysStringLen      ProcName = "SysStringLen"

	PNVariantClear            ProcName = "VariantClear"
	PNVariantInit             ProcName = "VariantInit"
	PNVariantTimeToSystemTime ProcName = "VariantTimeToSystemTime"
)

// OleAut32DLL Automation
type OleAut32DLL struct {
	*dLL
}

// NewOleAut32DLL You can import the interested process only instead of the whole process.
// We will load the entire process list if you do not give any parameters.
// We encourage you to fill in only the items you are using, as importing the all item will cause unnecessary waste.
func NewOleAut32DLL(procList ...ProcName) *OleAut32DLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNSafeArrayAccessData,
			PNSafeArrayAllocData,
			PNSafeArrayAllocDescriptor,
			PNSafeArrayAllocDescriptorEx,

			PNSafeArrayCopy,
			PNSafeArrayCopyData,

			PNSafeArrayCreate,
			PNSafeArrayCreateEx,
			PNSafeArrayCreateVector,
			PNSafeArrayCreateVectorEx,

			PNSafeArrayDestroy,
			PNSafeArrayDestroyData,
			PNSafeArrayDestroyDescriptor,

			PNSafeArrayGetDim,
			PNSafeArrayGetElement,
			PNSafeArrayGetElemSize,
			PNSafeArrayGetIID,
			PNSafeArrayGetLBound,
			PNSafeArrayGetRecordInfo,
			PNSafeArrayGetUBound,
			PNSafeArrayGetVarType,

			PNSafeArrayPutElement,

			PNSafeArrayUnAccessData,
			PNSafeArrayUnlock,

			PNSafeArraySetRecordInfo,

			PNSysAllocStringLen,
			PNSysFreeString,
			PNSysStringLen,

			PNVariantClear,
			PNVariantInit,
			PNVariantTimeToSystemTime,
		}
	}
	dll := newDll(DNOleAut32, procList)
	return &OleAut32DLL{dll}
}

var OleAutDll *OleAut32DLL

func init() {
	OleAutDll = NewOleAut32DLL()
}

// SafeArrayAccessData https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayaccessdata
func (dll OleAut32DLL) SafeArrayAccessData(safeArray *SafeArray) (uintptr, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayAccessData)
	var element uintptr
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&element)),
	)
	return element, syscall.Errno(hr)
}

// SafeArrayAllocData https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayallocdata
func (dll OleAut32DLL) SafeArrayAllocData(safeArray *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayAllocData)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
	)
	return syscall.Errno(hr)
}

// SafeArrayAllocDescriptor https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayallocdescriptor
func (dll OleAut32DLL) SafeArrayAllocDescriptor(cDims uint32) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayAllocDescriptor)
	var safeArray *SafeArray
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(cDims),
		uintptr(unsafe.Pointer(&safeArray)),
	)
	return safeArray, syscall.Errno(hr)
}

// SafeArrayAllocDescriptorEx https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayallocdescriptorex
func (dll OleAut32DLL) SafeArrayAllocDescriptorEx(vt VarType, cDims uint32) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayAllocDescriptorEx)
	var safeArray *SafeArray
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(vt),
		uintptr(cDims),
		uintptr(unsafe.Pointer(&safeArray)),
	)
	return safeArray, syscall.Errno(hr)
}

// SafeArrayCopy https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycopy
func (dll OleAut32DLL) SafeArrayCopy(original *SafeArray) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayCopy)
	var out *SafeArray
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(original)),
		uintptr(unsafe.Pointer(&out)),
	)
	return out, syscall.Errno(hr)
}

// SafeArrayCopyData https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycopydata
func (dll OleAut32DLL) SafeArrayCopyData(original *SafeArray, duplicate *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayCopyData)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(original)),
		uintptr(unsafe.Pointer(duplicate)),
	)
	return syscall.Errno(hr)
}

// SafeArrayCreate https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycreate
// Return value: A safe array descriptor, or null if the array could not be created.
func (dll OleAut32DLL) SafeArrayCreate(vt VarType, cDims uint32, bounds *SafeArrayBound) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayCreate)
	r, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(vt),
		uintptr(cDims),
		uintptr(unsafe.Pointer(bounds)),
	)
	return (*SafeArray)(unsafe.Pointer(r)), errno
}

// SafeArrayCreateEx https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycreateex
func (dll OleAut32DLL) SafeArrayCreateEx(vt VarType, cDims uint32, bounds *SafeArrayBound, pvExtra uintptr) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayCreateEx)
	r, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(vt),
		uintptr(cDims),
		uintptr(unsafe.Pointer(bounds)),
		pvExtra,
	)
	return (*SafeArray)(unsafe.Pointer(r)), errno
}

// SafeArrayCreateVector https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycreatevector
func (dll OleAut32DLL) SafeArrayCreateVector(vt VT,
	lBound int32, // The lower bound for the array. This parameter can be negative.
	cElements uint32, // The number of elements in the array.
) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayCreateVector)
	r, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(vt),
		uintptr(lBound),
		uintptr(cElements),
	)
	return (*SafeArray)(unsafe.Pointer(r)), errno
}

// SafeArrayCreateVectorEx https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycreatevectorex
func (dll OleAut32DLL) SafeArrayCreateVectorEx(vt VarType,
	lBound int32, // The lower bound for the array. This parameter can be negative.
	cElements uint32, // The number of elements in the array.
	extra uintptr,
) (*SafeArray, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayCreateVectorEx)
	r, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(vt),
		uintptr(lBound),
		uintptr(cElements),
		extra,
	)
	return (*SafeArray)(unsafe.Pointer(r)), errno
}

// SafeArrayDestroy https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraydestroy
func (dll OleAut32DLL) SafeArrayDestroy(pSA *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayDestroy)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSA)),
	)
	return syscall.Errno(hr)
}

// SafeArrayDestroyData https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraydestroydata
func (dll OleAut32DLL) SafeArrayDestroyData(pSA *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayDestroyData)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSA)),
	)
	return syscall.Errno(hr)
}

// SafeArrayDestroyDescriptor https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraydestroydescriptor
func (dll OleAut32DLL) SafeArrayDestroyDescriptor(pSA *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayDestroyDescriptor)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSA)),
	)
	return syscall.Errno(hr)
}

// SafeArrayGetDim https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetdim
// The number of dimensions in the array.
func (dll OleAut32DLL) SafeArrayGetDim(pSA *SafeArray) uint32 {
	proc := dll.mustProc(PNSafeArrayGetDim)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSA)),
	)
	return uint32(r)
}

// SafeArrayGetElement https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetelement
func (dll OleAut32DLL) SafeArrayGetElement(safeArray *SafeArray, index int32,
	pvData uintptr, // [out]
) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayGetElement)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&index)),
		pvData,
	)
	return syscall.Errno(hr)
}

// SafeArrayGetElemSize https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetelemsize
// The size of an element in a safe array, in bytes.
func (dll OleAut32DLL) SafeArrayGetElemSize(safeArray *SafeArray) uint32 {
	proc := dll.mustProc(PNSafeArrayGetElemSize)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
	)
	return *(*uint32)(unsafe.Pointer(r))
}

// SafeArrayGetIID https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetiid
// The size of an element in a safe array, in bytes.
func (dll OleAut32DLL) SafeArrayGetIID(safeArray *SafeArray) (*GUID, syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayGetIID)
	var guid *GUID
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&guid)),
	)
	return guid, syscall.Errno(hr)
}

// SafeArrayGetLBound https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetlbound
// Gets the lower bound for any dimension of the specified safe array.
func (dll OleAut32DLL) SafeArrayGetLBound(safeArray *SafeArray, nDim uint32) (lowerBound int32, errno syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayGetLBound)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(nDim),
		uintptr(unsafe.Pointer(&lowerBound)),
	)
	return lowerBound, syscall.Errno(hr)
}

// SafeArrayGetRecordInfo https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetrecordinfo
func (dll OleAut32DLL) SafeArrayGetRecordInfo(safeArray *SafeArray, nDim uint32) (recordInfo any, errno syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayGetRecordInfo)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&recordInfo)),
	)
	return recordInfo, syscall.Errno(hr)
}

// SafeArrayGetUBound https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetubound
// Gets the upper bound for any dimension of the specified safe array.
func (dll OleAut32DLL) SafeArrayGetUBound(safeArray *SafeArray, nDim uint32) (upperBound int32, errno syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayGetUBound)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(nDim),
		uintptr(unsafe.Pointer(&upperBound)),
	)
	return upperBound, syscall.Errno(hr)
}

// SafeArrayGetVarType https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetvartype
// Gets the VarType stored in the specified safe array.
func (dll OleAut32DLL) SafeArrayGetVarType(safeArray *SafeArray) (vt VT, errno syscall.Errno) {
	proc := dll.mustProc(PNSafeArrayGetVarType)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&vt)),
	)
	return vt, syscall.Errno(hr)
}

// SafeArrayPutElement https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayputelement
// Gets the VarType stored in the specified safe array.
func (dll OleAut32DLL) SafeArrayPutElement(safeArray *SafeArray, index int32, element uintptr) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayPutElement)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(index),
		element,
	)
	return syscall.Errno(hr)
}

// SafeArrayUnAccessData https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayunaccessdata
func (dll OleAut32DLL) SafeArrayUnAccessData(safeArray *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayUnAccessData)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
	)
	return syscall.Errno(hr)
}

// SafeArrayUnlock https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayunaccessdata
// Decrements the lock count of an array so it can be freed or resized.
func (dll OleAut32DLL) SafeArrayUnlock(safeArray *SafeArray) syscall.Errno {
	proc := dll.mustProc(PNSafeArrayUnlock)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
	)
	return syscall.Errno(hr)
}

// SafeArraySetRecordInfo https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraysetrecordinfo
// Sets the record info in the specified safe array.
func (dll OleAut32DLL) SafeArraySetRecordInfo(safeArray *SafeArray,
	recordInfo any, // [out]
) syscall.Errno {
	proc := dll.mustProc(PNSafeArraySetRecordInfo)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&recordInfo)),
	)
	return syscall.Errno(hr)
}

// SysAllocStringLen  https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysallocstringlen
// Deallocates a string allocated previously by SysAllocString, SysAllocStringByteLen, SysReAllocString, SysAllocStringLen, or SysReAllocStringLen
func (dll OleAut32DLL) SysAllocStringLen(utf8Str string) *uint16 {
	proc := dll.mustProc(PNSysAllocStringLen)
	// UintptrFromStr(str) // 可以直接用這個取，但是長度會不知道
	utf16Str := utf16.Encode([]rune(utf8Str + "\x00"))
	r, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(&utf16Str[0])),
		uintptr(len(utf16Str)-1),
	)
	return (*uint16)(unsafe.Pointer(r))
}

// SysFreeString https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysfreestring
// Deallocates a string allocated previously by SysAllocString, SysAllocStringByteLen, SysReAllocString, SysAllocStringLen, or SysReAllocStringLen
func (dll OleAut32DLL) SysFreeString(bStrString *uint16) {
	proc := dll.mustProc(PNSysFreeString)
	_, _, _ = syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(bStrString)),
	)
}

// SysStringLen  https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysstringlen
func (dll OleAut32DLL) SysStringLen(bStrString *uint16) uint32 {
	proc := dll.mustProc(PNSysStringLen)
	numberOfChar, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(bStrString)),
	)
	return uint32(numberOfChar)
}

// VariantClear https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
func (dll OleAut32DLL) VariantClear(v *VARIANT) syscall.Errno {
	proc := dll.mustProc(PNVariantClear)
	hr, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(&v)),
	)
	return syscall.Errno(hr)
}

// VariantInit https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantinit
func (dll OleAut32DLL) VariantInit(v *VARIANT) {
	proc := dll.mustProc(PNVariantInit)
	_, _, _ = syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(&v)),
	)
}

// VariantTimeToSystemTime https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-varianttimetosystemtime
func (dll *OleAut32DLL) VariantTimeToSystemTime(vTime float64,
	st *syscall.Systemtime, // [OUT] 2.0表示1900/1/1, 2.0 1900/1/2, 2.5 1900/1/2 12:00:00
) bool {
	proc := dll.mustProc(PNVariantTimeToSystemTime)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(vTime),
		uintptr(unsafe.Pointer(&st)),
	)
	return r != 0
}
