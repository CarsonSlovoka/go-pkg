package syscall

// CloseHandle Closes an open object handle.
// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle?redirectedfrom=MSDN
func CloseHandle(handle uintptr) error {
	procCloseHandle := DllKernel32.NewProc("CloseHandle")
	val, _, err := procCloseHandle.Call(handle)
	if val == 0 {
		return err
	}
	return nil
}
