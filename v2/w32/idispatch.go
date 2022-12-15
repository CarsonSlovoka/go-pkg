package w32

import (
	"syscall"
	"unsafe"
)

// IDispatch
// Methods: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch#methods
type IDispatch struct {
	IUnknown
}

type IDispatchVTable struct { // 🧙 注意！ 順序不可以亂放，否則disp.VTable().xxx取到的函數會是錯的
	IUnknownVTable   // 所有的其他類型都一定會含有IUnknown所提供的方法，記得要補上
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

func (d *IDispatch) VTable() *IDispatchVTable {
	if d == nil {
		panic("IDispatch is nil")
	}
	return (*IDispatchVTable)(unsafe.Pointer(d.RawVTable))
}

// GetIDsOfNames https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-getidsofnames
// LCID kernel32dll.GetUserDefaultLCID()
func (d *IDispatch) GetIDsOfNames(
	iid *GUID, // Reserved for future use. Must be IID_NULL.
	names []string,
	nNames uint32,
	lcID LCID,
) ([]DISPID, syscall.Errno) {
	if iid == nil {
		iid = IID_NULL
	}

	if nNames == 0 {
		nNames = uint32(len(names))
	}

	dispIDs := make([]DISPID, len(names))
	wNames := make([]*uint16, len(names)) // winapi用的是utf16，go字串是utf8，要轉換
	for i := 0; i < len(names); i++ {
		pWStr, err := syscall.UTF16PtrFromString(names[i])
		if err != nil {
			panic(err)
		}
		wNames[i] = pWStr
	}

	// log.Println(d.VTable().GetIDsOfNames) // 如果您知道位址可以直接寫入也是一樣，例如: uintptr(140715441321184), 如果函數調用失敗不妨看看成功的例子他的數值是多少，來判斷是否函數的順序寫錯 // 這個是固定的數值，表示某一函數，如果d.VTable()也就是IDispatchVTable，依照該函數順序去取得相關函數
	hr, _, _ := syscall.SyscallN(d.VTable().GetIDsOfNames, uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(iid)), // Reserved for future use. Must be IID_NULL.
		uintptr(unsafe.Pointer(&wNames[0])),
		uintptr(nNames),
		uintptr(lcID),
		uintptr(unsafe.Pointer(&dispIDs[0])),
	)
	return dispIDs, syscall.Errno(hr)
}

// GetTypeInfo https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (d *IDispatch) GetTypeInfo(lcid LCID) (*ITypeInfo, syscall.Errno) {
	var tinfo *ITypeInfo
	hr, _, _ := syscall.SyscallN(d.VTable().GetTypeInfo,
		uintptr(unsafe.Pointer(d)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(&tinfo)),
	)
	return tinfo, syscall.Errno(hr)
}

// GetTypeInfoCount https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
// Retrieves the number of type information interfaces that an object provides (either 0 or 1).
func (d *IDispatch) GetTypeInfoCount() (c uint32, errno syscall.Errno) {
	hr, _, _ := syscall.SyscallN(d.VTable().GetTypeInfoCount, uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(&c)),
	)
	errno = syscall.Errno(hr)
	return
}

// Invoke https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
func (d *IDispatch) Invoke(
	disID DISPID, // Use GetIDsOfNames or the object's documentation to obtain the dispatch identifier.
	rIID *GUID,   // Reserved for future use. Must be IID_NULL.
	lcID LCID,
	dispatchFlag uint16, // DISPATCH_METHOD, DISPATCH_PROPERTYGET, DISPATCH_PROPERTYPUT, DISPATCH_PROPERTYPUTREF

	dispParams *DispParams, // [in, out]
	result *VARIANT,        // [out]
	exceptInfo *EXCEPINFO,  // [out]
	puArgErr *uint32,       // [out]
) syscall.Errno {

	if rIID == nil {
		rIID = IID_NULL
	}

	var argErr uintptr
	if puArgErr != nil {
		argErr = uintptr(unsafe.Pointer(&puArgErr))
	}

	hr, _, _ := syscall.SyscallN(d.VTable().Invoke, uintptr(unsafe.Pointer(d)),
		uintptr(disID),
		uintptr(unsafe.Pointer(rIID)), // Reserved for future use. Must be IID_NULL.
		uintptr(lcID),
		uintptr(dispatchFlag),

		uintptr(unsafe.Pointer(dispParams)), // 不能給0
		uintptr(unsafe.Pointer(result)),
		uintptr(unsafe.Pointer(exceptInfo)),
		argErr,
	)

	return syscall.Errno(hr)
}
