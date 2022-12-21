package w32

import "unsafe"

// DispParams https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-dispparams
type DispParams struct {
	VArgs      uintptr // *VARIANT
	NamedArgs  uintptr // *DISPID
	CArgs      uint32
	CNamedArgs uint32
}

func (d DispParams) Variants() []*VARIANT {
	var i uint32
	vs := make([]*VARIANT, d.CArgs)
	for i = 0; i < d.CArgs; i++ {
		vs[i] = (*VARIANT)(unsafe.Pointer(d.VArgs + (uintptr(i) * unsafe.Sizeof(vs[0]))))
	}
	return vs
}

func (d DispParams) DispID() []*DISPID {
	var i uint32
	vs := make([]*DISPID, d.CNamedArgs)
	for i = 0; i < d.CNamedArgs; i++ {
		vs[i] = (*DISPID)(unsafe.Pointer(d.NamedArgs + (uintptr(i) * unsafe.Sizeof(vs[0]))))
	}
	return vs
}

const (
	DISPID_UNKNOWN     = -1
	DISPID_VALUE       = 0
	DISPID_PROPERTYPUT = -3
	DISPID_NEWENUM     = -4
	DISPID_EVALUATE    = -5
	DISPID_CONSTRUCTOR = -6
	DISPID_DESTRUCTOR  = -7
	DISPID_COLLECT     = -8
)

// EXCEPINFO https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
type EXCEPINFO struct {
	WCode             uint16
	WReserved         uint16
	BstrSource        uint16
	BstrDescription   uint16
	BstrHelpFile      uint16
	DwHelpContext     uint32
	PvReserved        uintptr
	PfnDeferredFillIn uintptr
	Scode             int32
}
