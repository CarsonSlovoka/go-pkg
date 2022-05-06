package syscall

import "syscall"

// CloseHandle Closes an open object handle.
// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle?redirectedfrom=MSDN
func CloseHandle(proc *syscall.LazyProc, handle uintptr) error {
	if proc.Name != "CloseHandle" {
		panic("proc.Name != CloseHandle")
	}
	val, _, err := proc.Call(handle)
	if val == 0 {
		return err
	}
	return nil
}
