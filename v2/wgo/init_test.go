package wgo_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"github.com/CarsonSlovoka/go-pkg/v2/wgo"
)

var (
	wGo *wgo.WGO
)

func init() {
	kernel32dll := w32.NewKernel32DLL()
	wGo = wgo.NewWGO(kernel32dll)
}
