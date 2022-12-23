package wgo_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"syscall"
)

// 如果您使用COM來調用EXCEL，可能終止的時候還會有殘留的EXCEL，需要手動終止
func ExampleWGO_KillProcess() {
	if entrySlice, errno := wGo.GetProcessEntry(func(entry *w32.PROCESSENTRY32W) bool {
		return entry.ExeFileName() == "EXCEL.EXE" // && entry.Th32ProcessID != 1844
	}); errno == 0 {
		wGo.KillProcess(entrySlice, func(entry *w32.PROCESSENTRY32W, errno syscall.Errno) {
			if errno == 0 {
				log.Println("delete", entry.ExeFileName(), entry.Th32ProcessID)
			}
		})
	}
	// Output:
}
