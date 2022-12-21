//go:build amd64

package w32

// VARIANT https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-variant?source=recommendations
type VARIANT struct {
	VT         VT      //  2
	wReserved1 uint16  //  4
	wReserved2 uint16  //  6
	wReserved3 uint16  //  8
	Val        int64   // 16
	_          [8]byte // 24
}
