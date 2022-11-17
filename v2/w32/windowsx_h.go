package w32

// GET_X_LPARAM https://learn.microsoft.com/en-us/windows/win32/api/windowsx/nf-windowsx-get_x_lparam
func GET_X_LPARAM[T uintptr | WPARAM | LPARAM](lp T) int32 {
	// ((int)(short)LOWORD(lp))
	return int32(int16(LOWORD(lp)))
}

func GET_Y_LPARAM[T uintptr | WPARAM | LPARAM](lp T) int32 {
	return int32(int16(HIWORD(lp)))
}
