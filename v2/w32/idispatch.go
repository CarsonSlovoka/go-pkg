package w32

import (
	"fmt"
	"math/big"
	"syscall"
	"time"
	"unsafe"
)

var preferLCID LCID = LOCALE_USER_DEFAULT

func SetPreferLCID(id LCID) {
	preferLCID = id
}

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
	rIID *GUID, // Reserved for future use. Must be IID_NULL.
	lcID LCID,
	dispatchFlag DispatchFlag,

	dispParams *DispParams, // [in, out]

	// result *VARIANT,     // 🧙 不能放裡面

	exceptInfo *EXCEPINFO, // [out]
	puArgErr *uint32, // [out]
) (result *VARIANT, errno syscall.Errno) {

	if rIID == nil {
		rIID = IID_NULL
	}

	var argErr uintptr
	if puArgErr != nil {
		argErr = uintptr(unsafe.Pointer(&puArgErr))
	}

	result = new(VARIANT)
	OleAutDll.VariantInit(result)

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
	return result, syscall.Errno(hr)
}

func (d *IDispatch) PropertyPut(name string, params ...any) (*VARIANT, syscall.Errno) {
	return invoke(d, name, DISPATCH_PROPERTYPUT, params...)
}

func (d *IDispatch) MustPropertyPut(name string, params ...any) *VARIANT {
	result, errno := invoke(d, name, DISPATCH_PROPERTYPUT, params...)
	if errno != 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return result
}

func (d *IDispatch) PropertyPutRef(name string, params ...any) (*VARIANT, syscall.Errno) {
	return invoke(d, name, DISPATCH_PROPERTYPUTREF, params...)
}

func (d *IDispatch) MustPropertyPutRef(name string, params ...any) *VARIANT {
	result, errno := invoke(d, name, DISPATCH_PROPERTYPUTREF, params...)
	if errno != 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return result
}

func (d *IDispatch) PropertyGet(name string, params ...any) (*VARIANT, syscall.Errno) {
	return invoke(d, name, DISPATCH_PROPERTYGET, params...)
}

func (d *IDispatch) MustPropertyGet(name string, params ...any) *VARIANT {
	result, errno := invoke(d, name, DISPATCH_PROPERTYGET, params...)
	if errno != 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return result
}

func (d *IDispatch) Method(name string, params ...any) (*VARIANT, syscall.Errno) {
	return invoke(d, name, DISPATCH_METHOD, params...)
}

func (d *IDispatch) MustMethod(name string, params ...any) (result *VARIANT) {
	var errno syscall.Errno
	result, errno = invoke(d, name, DISPATCH_METHOD, params...)
	if errno != 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return result
}

func invoke(d *IDispatch, name string, dispatchFlag DispatchFlag, params ...any) (outVariant *VARIANT, errno syscall.Errno) {
	var (
		propertyNameID DISPID
		names          []DISPID
	)
	names, errno = d.GetIDsOfNames(IID_NULL, []string{name}, 1, preferLCID)
	if errno != 0 {
		return nil, errno
	}
	propertyNameID = names[0]

	dispParams := new(DispParams) // 都沒有數值需要填還是需要指向一個空內容，不能是nil，否則會錯

	// Set NameArgs if and only if the property is the PUT.
	if dispatchFlag&DISPATCH_PROPERTYPUT != 0 {
		nameArgs := [1]DISPID{DISPID_PROPERTYPUT}
		dispParams.NamedArgs = uintptr(unsafe.Pointer(&nameArgs[0]))
		dispParams.CNamedArgs = 1
	} else if dispatchFlag&DISPATCH_PROPERTYPUTREF != 0 {
		nameArgs := [1]DISPID{DISPID_PROPERTYPUT}
		dispParams.NamedArgs = uintptr(unsafe.Pointer(&nameArgs[0]))
		dispParams.CNamedArgs = 1
	}

	// Set Params
	var vargs []VARIANT // 寫在外面是應付當為VT_BSTR或VT_BYREF，最後需要再釋放
	if len(params) > 0 {
		vargs = make([]VARIANT, len(params))
		for i, v := range params {
			n := len(params) - i - 1 // https://go.dev/play/p/cbE_62Q1-96 因為defer有後進先出(LIFO)的特性，所以排序也要在顛倒，才會是原本的順序
			OleAutDll.VariantInit(&vargs[n])
			switch vv := v.(type) {
			case bool:
				if vv {
					vargs[n] = NewVariant(VT_BOOL, 0xffff)
				} else {
					vargs[n] = NewVariant(VT_BOOL, 0)
				}
			case *bool:
				vargs[n] = NewVariant(VT_BOOL|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*bool)))))
			case uint8:
				vargs[n] = NewVariant(VT_I1, int64(v.(uint8)))
			case *uint8:
				vargs[n] = NewVariant(VT_I1|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*uint8)))))
			case int8:
				vargs[n] = NewVariant(VT_I1, int64(v.(int8)))
			case *int8:
				vargs[n] = NewVariant(VT_I1|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*int8)))))
			case int16:
				vargs[n] = NewVariant(VT_I2, int64(v.(int16)))
			case *int16:
				vargs[n] = NewVariant(VT_I2|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*int16)))))
			case uint16:
				vargs[n] = NewVariant(VT_UI2, int64(v.(uint16)))
			case *uint16:
				vargs[n] = NewVariant(VT_UI2|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*uint16)))))
			case int32:
				vargs[n] = NewVariant(VT_I4, int64(v.(int32)))
			case *int32:
				vargs[n] = NewVariant(VT_I4|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*int32)))))
			case uint32:
				vargs[n] = NewVariant(VT_UI4, int64(v.(uint32)))
			case *uint32:
				vargs[n] = NewVariant(VT_UI4|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*uint32)))))
			case int64:
				vargs[n] = NewVariant(VT_I8, int64(v.(int64)))
			case *int64:
				vargs[n] = NewVariant(VT_I8|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*int64)))))
			case uint64:
				vargs[n] = NewVariant(VT_UI8, int64(uintptr(v.(uint64))))
			case *uint64:
				vargs[n] = NewVariant(VT_UI8|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*uint64)))))
			case int:
				vargs[n] = NewVariant(VT_I4, int64(v.(int)))
			case *int:
				vargs[n] = NewVariant(VT_I4|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*int)))))
			case uint:
				vargs[n] = NewVariant(VT_UI4, int64(v.(uint)))
			case *uint:
				vargs[n] = NewVariant(VT_UI4|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*uint)))))
			case float32:
				vargs[n] = NewVariant(VT_R4, *(*int64)(unsafe.Pointer(&vv)))
			case *float32:
				vargs[n] = NewVariant(VT_R4|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*float32)))))
			case float64:
				vargs[n] = NewVariant(VT_R8, *(*int64)(unsafe.Pointer(&vv)))
			case *float64:
				vargs[n] = NewVariant(VT_R8|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*float64)))))
			case *big.Int:
				vargs[n] = NewVariant(VT_DECIMAL, v.(*big.Int).Int64())
			case string:
				vargs[n] = NewVariant(VT_BSTR, int64(uintptr(unsafe.Pointer(OleAutDll.SysAllocStringLen(v.(string))))))
			case *string:
				vargs[n] = NewVariant(VT_BSTR|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*string)))))
			case time.Time:
				s := vv.Format("2006-01-02 15:04:05")
				vargs[n] = NewVariant(VT_BSTR, int64(uintptr(unsafe.Pointer(OleAutDll.SysAllocStringLen(s)))))
			case *time.Time:
				s := vv.Format("2006-01-02 15:04:05")
				vargs[n] = NewVariant(VT_BSTR|VT_BYREF, int64(uintptr(unsafe.Pointer(&s))))
			case *IDispatch:
				vargs[n] = NewVariant(VT_DISPATCH, int64(uintptr(unsafe.Pointer(v.(*IDispatch)))))
			case **IDispatch:
				vargs[n] = NewVariant(VT_DISPATCH|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(**IDispatch)))))
			case nil:
				vargs[n] = NewVariant(VT_NULL, 0)
			case *VARIANT:
				vargs[n] = NewVariant(VT_VARIANT|VT_BYREF, int64(uintptr(unsafe.Pointer(v.(*VARIANT)))))
			case []byte:
				safeByteArray := safeArrayFromByteSlice(v.([]byte))
				vargs[n] = NewVariant(VT_ARRAY|VT_UI1, int64(uintptr(unsafe.Pointer(safeByteArray))))
				defer OleAutDll.VariantClear(&vargs[n])
			case []string:
				safeByteArray := safeArrayFromStringSlice(v.([]string))
				vargs[n] = NewVariant(VT_ARRAY|VT_BSTR, int64(uintptr(unsafe.Pointer(safeByteArray))))
				defer OleAutDll.VariantClear(&vargs[n])
			default:
				panic("unknown type")
			}
		}
		dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))
		dispParams.CArgs = uint32(len(params))
	}

	var exceptInfo EXCEPINFO
	outVariant, errno = d.Invoke(propertyNameID, IID_NULL, preferLCID, dispatchFlag, dispParams, &exceptInfo, nil)

	if vargs != nil {
		for i, varg := range vargs {
			n := len(params) - i - 1
			if varg.VT == VT_BSTR && varg.Val != 0 {
				OleAutDll.SysFreeString((*uint16)(unsafe.Pointer(uintptr(varg.Val))))
			}
			if varg.VT == (VT_BSTR|VT_BYREF) && varg.Val != 0 {
				*(params[n].(*string)) = LpOleStrToString(*(**uint16)(unsafe.Pointer(uintptr(varg.Val))))
			}
		}
	}

	return outVariant, errno
}
