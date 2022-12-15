//go:build windows

package w32

import "unsafe"

const hexTable = "0123456789ABCDEF"
const emptyGUID = "{00000000-0000-0000-0000-000000000000}"

// GUID is Windows API specific GUID type.
// ---d1---- d2 - d3 - d4a- d4b
// 12345678-1234-1234-1234-123456789ABC
//
// This exists to match Windows GUID type for direct passing for COM.
// Format is in xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxxxxxx.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (g *GUID) String() string {
	if g == nil {
		return emptyGUID
	}
	c := make([]byte, 38)
	c[0] = '{'
	putCharFromUint32(c[1:9], g.Data1)
	c[9] = '-'
	putCharFromUint16(c[10:14], g.Data2)
	c[14] = '-'
	putCharFromUint16(c[15:19], g.Data3)
	c[19] = '-'
	putCharFromBytes(c[20:24], g.Data4[0:2])
	c[24] = '-'
	putCharFromBytes(c[25:37], g.Data4[2:8])
	c[37] = '}'
	return *(*string)(unsafe.Pointer(&c))
}

func (g *GUID) Equal(g2 *GUID) bool {
	return *g == *g2 // https://stackoverflow.com/q/24534072/9935654
}

func NewGUID(guid string) *GUID {
	d := []byte(guid)
	var d1, d2, d3, d4a, d4b []byte

	switch len(d) {
	case 38:
		if d[0] != '{' || d[37] != '}' {
			return nil
		}
		d = d[1:37]
		fallthrough
	case 36:
		if d[8] != '-' || d[13] != '-' || d[18] != '-' || d[23] != '-' {
			return nil
		}
		d1 = d[0:8]
		d2 = d[9:13]
		d3 = d[14:18]
		d4a = d[19:23]
		d4b = d[24:36]
	case 32:
		d1 = d[0:8]
		d2 = d[8:12]
		d3 = d[12:16]
		d4a = d[16:20]
		d4b = d[20:32]
	default:
		return nil
	}

	var g GUID
	var ok1, ok2, ok3, ok4 bool
	g.Data1, ok1 = uint32FromHexStr(d1)
	g.Data2, ok2 = uint16FromHexStr(d2)
	g.Data3, ok3 = uint16FromHexStr(d3)
	g.Data4, ok4 = byte64FromHexStr(d4a, d4b)
	if ok1 && ok2 && ok3 && ok4 {
		return &g
	}
	return nil
}

func uint16FromHexStr(s []byte) (uint16, bool) {
	if len(s) != 4 {
		return 0, false
	}
	n1, ok1 := byteFromNibble(s[0], s[1])
	n2, ok2 := byteFromNibble(s[2], s[3])
	return uint16(n1)<<8 | uint16(n2),
		ok1 && ok2
}

func uint32FromHexStr(s []byte) (uint32, bool) {
	if len(s) != 8 {
		return 0, false
	}
	n1, ok1 := byteFromNibble(s[0], s[1])
	n2, ok2 := byteFromNibble(s[2], s[3])
	n3, ok3 := byteFromNibble(s[4], s[5])
	n4, ok4 := byteFromNibble(s[6], s[7])
	return uint32(n1)<<24 | uint32(n2)<<16 | uint32(n3)<<8 | uint32(n4),
		ok1 && ok2 && ok3 && ok4
}

func byte64FromHexStr(s1 []byte, s2 []byte) (b [8]byte, ok bool) {
	if len(s1) != 4 && len(s2) != 12 {
		return b, false
	}

	var i = 0
	for i = 0; i < 2; i++ {
		b[i], ok = byteFromNibble(s1[2*i], s1[2*i+1])
		if !ok {
			return
		}
	}
	for i = 2; i < 8; i++ {
		b[i], ok = byteFromNibble(s2[2*(i-2)], s2[2*(i-2)+1])
		if !ok {
			return
		}
	}
	return
}

func byteFromNibble(c1, c2 byte) (byte, bool) {
	n1, ok1 := decFromChar(c1)
	n2, ok2 := decFromChar(c2)
	return (n1 << 4) | n2,
		ok1 && ok2
}

func decFromChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}
	return 0, false
}

func putCharFromUint32(b []byte /* [out] */, v uint32) {
	b[0] = hexTable[v>>28]
	b[1] = hexTable[v>>24&0xf]
	b[2] = hexTable[v>>20&0xf]
	b[3] = hexTable[v>>16&0xf]
	b[4] = hexTable[v>>12&0xf]
	b[5] = hexTable[v>>8&0xf]
	b[6] = hexTable[v>>4&0xf]
	b[7] = hexTable[v&0xf]
}

func putCharFromUint16(b []byte /* [out] */, v uint16) {
	b[0] = hexTable[v>>12]
	b[1] = hexTable[v>>8&0xf]
	b[2] = hexTable[v>>4&0xf]
	b[3] = hexTable[v&0xf]
}

// 這裡的一個字元，都表示16進位表的某一個元素
// 故該字元可以由1個nibble組成 or 4bit or half of a byte
// 所以我們每一個byte裡面又要對半砍
func putCharFromBytes(dst []byte /* [out] */, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[2*i] = hexTable[src[i]>>4]
		dst[2*i+1] = hexTable[src[i]&0xf]
	}
}
