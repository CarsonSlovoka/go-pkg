package w32

import (
	"syscall"
	"unsafe"
)

func StrToLPCWSTR(str string) uintptr {
	if str == "" {
		return 0
	}
	lpcwstr, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(lpcwstr))
}
