//go:build windows

package w32

import (
	"unsafe"
)

func safeArrayFromByteSlice(slice []byte) *SafeArray {
	array, _ := OleAutDll.SafeArrayCreateVector(VT_UI1, 0, uint32(len(slice)))

	if array == nil {
		panic("Could not convert []byte to SAFEARRAY")
	}

	for i, v := range slice {
		_ = OleAutDll.SafeArrayPutElement(array, int32(i), uintptr(unsafe.Pointer(&v)))
	}
	return array
}

func safeArrayFromStringSlice(slice []string) *SafeArray {
	array, _ := OleAutDll.SafeArrayCreateVector(VT_BSTR, 0, uint32(len(slice)))

	if array == nil {
		panic("Could not convert []string to SafeArray")
	}
	// SysAllocStringLen(s)
	for i, v := range slice {
		_ = OleAutDll.SafeArrayPutElement(array, int32(i), uintptr(unsafe.Pointer(OleAutDll.SysAllocStringLen(v))))
	}
	return array
}
