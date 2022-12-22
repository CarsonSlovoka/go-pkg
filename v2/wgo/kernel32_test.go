package wgo_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/wgo"
	"log"
)

// 如果您使用COM來調用EXCEL，可能終止的時候還會有殘留的EXCEL，需要手動終止
func ExampleKillProcess() {
	if entrySlice, errno := wgo.GetProcessEntryByName(kernelDll, "EXCEL.EXE"); errno == 0 {
		if errno = wgo.KillProcess(kernelDll, entrySlice, nil); errno != 0 {
			log.Printf("%s", errno)
		}
	}
	// Output:
}
