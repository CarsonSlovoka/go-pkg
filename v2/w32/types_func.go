package w32

import (
	"syscall"
	"unsafe"
)

// StrToCharPtr
// 可以用在A的函數 procA.Call(uintptr(unsafe.Pointer(StrToCharPtr(name))))
func StrToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

func UintptrFromBool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// UintptrFromStr 用在W的方法中
// 先將str轉成utf16ptr再轉成uintptr
func UintptrFromStr(str string) uintptr {
	// wchars := utf16.Encode([]rune(str + "\x00"))
	// return &wchars[0] // *uint16

	if str == "" {
		return 0
	}
	lpcwstr, err := syscall.UTF16PtrFromString(str) // *uint16, err
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(lpcwstr))
}
