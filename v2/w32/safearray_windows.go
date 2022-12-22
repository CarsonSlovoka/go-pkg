package w32

/*
// safeArrayGetElementString retrieves element at given index and converts to string.
func safeArrayGetElementString(safeArray *SafeArray, index int32) (str string, err error) {
	var element *uint16
	err = OleAutDll.SafeArrayGetElement(safeArray, index, uintptr(unsafe.Pointer(&element)))
	str = BstrToString(*(**uint16)(unsafe.Pointer(element)))
	OleAutDll.SysFreeString(element)
	return
}
*/
