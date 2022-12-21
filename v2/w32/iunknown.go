package w32

import (
	"syscall"
	"unsafe"
)

// IUnknown https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	RawVTable *any // 這是一個可以指向任何東西的指標，通常他會指到一個結構去，此結構會記錄所有該物件能使用的方法之位址，再依據您想要使用哪一個方法，用該位址來當作procAddr
}

func NewIUnknownInstance[T *GUID | string](ole *Ole32DLL,
	classGuid T,
	classContext uint32, // CLSCTX_SERVER
) (unknown *IUnknown, errno syscall.Errno) {
	var guid *GUID

	switch any(classGuid).(type) {
	case *GUID:
		guid = any(classGuid).(*GUID)
	case string:
		guid = new(GUID)
		if errno = ole.CLSIDFromProgID(any(classGuid).(string), guid); errno != 0 {
			return nil, errno
		}
	}
	unknown, errno = ole.CoCreateInstance(
		guid,
		classContext,
		IID_IUnknown,
	)
	return unknown, errno
}

func (unk *IUnknown) VTable() *IUnknownVTable {
	return (*IUnknownVTable)(unsafe.Pointer(unk.RawVTable))
}

// IUnknownVTable https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown#methods
type IUnknownVTable struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type LpUnKnown *IUnknown

// QueryInterface https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (unk *IUnknown) QueryInterface(
	iid *GUID,
	// disp *IDispatch // [out]   // 放這邊會有問題
) (*IDispatch, syscall.Errno) {
	var disp *IDispatch
	hr, _, _ := syscall.SyscallN(unk.VTable().QueryInterface, uintptr(unsafe.Pointer(unk)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&disp)),
	)
	return disp, syscall.Errno(hr)
}

// AddRef https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (unk *IUnknown) AddRef() int32 {
	r, _, _ := syscall.SyscallN(unk.VTable().AddRef, uintptr(unsafe.Pointer(&unk)))
	return int32(r)
}

// Release https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
// return value: The method returns the new reference count. This value is intended to be used only for test purposes.
func (unk *IUnknown) Release() uint32 {
	r, _, _ := syscall.SyscallN(unk.VTable().AddRef, uintptr(unsafe.Pointer(&unk)))
	return uint32(r)
}
