// Helper for converting SafeArray to array of objects.

package w32

import (
	"syscall"
	"unsafe"
)

type SafeArrayConversion struct {
	Array *SafeArray
}

// TODO
/*
func (sac *SafeArrayConversion) ToStringArray() (strings []string) {
	totalElements, _ := sac.TotalElements(0)
	strings = make([]string, totalElements)

	for i := int32(0); i < totalElements; i++ {
		strings[i], _ = safeArrayGetElementString(sac.Array, i)
	}

	return
}
*/

func (sac *SafeArrayConversion) ToByteArray() (bytes []byte) {
	totalElements, _ := sac.TotalElements(0)
	bytes = make([]byte, totalElements)

	for i := int32(0); i < totalElements; i++ {
		_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&bytes[int32(i)])))
	}

	return
}

func (sac *SafeArrayConversion) ToValueArray() (values []interface{}) {
	totalElements, _ := sac.TotalElements(0)
	values = make([]interface{}, totalElements)
	vt, _ := OleAutDll.SafeArrayGetVarType(sac.Array)

	for i := int32(0); i < totalElements; i++ {
		switch vt {
		case VT_BOOL:
			var v bool
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_I1:
			var v int8
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_I2:
			var v int16
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_I4:
			var v int32
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_I8:
			var v int64
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_UI1:
			var v uint8
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_UI2:
			var v uint16
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_UI4:
			var v uint32
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_UI8:
			var v uint64
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_R4:
			var v float32
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v
		case VT_R8:
			var v float64
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v

			/* TODO
			case VT_BSTR:
				v, _ := safeArrayGetElementString(sac.Array, i)
				values[i] = v
			*/
		case VT_VARIANT:
			var v VARIANT
			_ = OleAutDll.SafeArrayGetElement(sac.Array, i, uintptr(unsafe.Pointer(&v)))
			values[i] = v.Value()
			v.Clear()
		default:
			// TODO
		}
	}

	return
}

func (sac *SafeArrayConversion) GetType() (varType VT, err error) {
	return OleAutDll.SafeArrayGetVarType(sac.Array)
}

func (sac *SafeArrayConversion) GetDimensions() uint32 {
	return OleAutDll.SafeArrayGetDim(sac.Array)
}

func (sac *SafeArrayConversion) GetSize() uint32 {
	return OleAutDll.SafeArrayGetElemSize(sac.Array)
}

func (sac *SafeArrayConversion) TotalElements(index uint32) (totalElements int32, errno syscall.Errno) {
	if index < 1 {
		index = 1
	}

	// Get array bounds
	var LowerBounds int32
	var UpperBounds int32

	LowerBounds, errno = OleAutDll.SafeArrayGetLBound(sac.Array, index)
	if errno != 0 {
		return
	}

	UpperBounds, errno = OleAutDll.SafeArrayGetUBound(sac.Array, index)
	if errno != 0 {
		return
	}

	totalElements = UpperBounds - LowerBounds + 1
	return
}

// Release Safe Array memory
func (sac *SafeArrayConversion) Release() {
	OleAutDll.SafeArrayDestroy(sac.Array)
}
