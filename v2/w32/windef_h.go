package w32

func MAKEWORD[T WPARAM | LPARAM | uintptr](low, high T) uint16 {
	return uint16(uint8(low)) | (uint16(uint8(high)))<<8
}

/*
func MakePoints[T WPARAM | LPARAM | uintptr](lParam T) POINTS {
	return POINTS{
		int16(lParam),
		int16(lParam >> 16),
	}
}
*/

// MakePoint 參數為8byte，但是在這種狀況，實際上只用了32bit
// lParam
// 0000 0000 0000 0000 0000 0000 0000 0000 | 0000 0000 0000 0000 0000 0000 0000 0000
// |------------>未使用<------------------- | |----> y <-------|  |------> x <------|
func MakePoint[T WPARAM | LPARAM | uintptr](lParam T) POINT {
	return POINT{
		int32(int16(lParam)),
		int32(int16(lParam >> 16)),
	}
}
