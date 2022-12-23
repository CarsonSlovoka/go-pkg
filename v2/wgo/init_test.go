package wgo_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"github.com/CarsonSlovoka/go-pkg/v2/wgo"
)

var (
	wGo       *wgo.WGO
	kernelDll *w32.Kernel32DLL
)

func init() {
	kernelDll = w32.NewKernel32DLL()
	wGo = wgo.NewWGO(kernelDll)
}
