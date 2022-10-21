package w32

import "unsafe"

// MakeIntResource https://learn.microsoft.com/en-us/windows/win32/menurc/resource-types
func MakeIntResource(id uintptr) *uint16 {
	return (*uint16)(unsafe.Pointer(id))
}
