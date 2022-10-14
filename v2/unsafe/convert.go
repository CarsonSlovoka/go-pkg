package unsafe

import "unsafe"

// StrToBytes It doesn't copy the data, so it's more quickly.
func StrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesToStr It doesn't copy the data, so it's more quickly.
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
