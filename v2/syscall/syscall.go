package syscall

import "syscall"

var DllKernel32 *syscall.LazyDLL // syscall.NewLazyDLL("kernel32.dll")
