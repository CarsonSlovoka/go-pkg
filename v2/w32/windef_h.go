package w32

func MAKEWORD[T WPARAM | LPARAM | uintptr](low, high T) uint16 {
	return uint16(uint8(low)) | (uint16(uint8(high)))<<8
}
