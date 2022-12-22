package wgo_test

import "github.com/CarsonSlovoka/go-pkg/v2/w32"

var (
	kernelDll *w32.Kernel32DLL
)

func init() {
	kernelDll = w32.NewKernel32DLL()
}
