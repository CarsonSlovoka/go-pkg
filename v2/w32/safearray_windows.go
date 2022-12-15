package w32

import "unsafe"

// safeArrayGetElementString retrieves element at given index and converts to string.
func safeArrayGetElementString(safearray *SafeArray, index int32) (str string, err error) {
	var element *uint16
	err = OleAutDll.SafeArrayGetElement(safearray, index, uintptr(unsafe.Pointer(&element)))
	str = BstrToString(element)
	OleAutDll.SysFreeString(element)
	return
}
