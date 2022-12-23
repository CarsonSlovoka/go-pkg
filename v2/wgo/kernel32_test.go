package wgo_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"syscall"
	"testing"
)

func TestWGO_KillProcess(t *testing.T) {
	var (
		si w32.STARTUPINFO
		pi w32.PROCESS_INFORMATION
	)

	if errno := kernelDll.CreateProcess("", // // No module name (use command line)
		"notepad.exe",
		nil, nil, false, 0, 0, "",
		&si,
		&pi,
	); errno != 0 {
		log.Println(errno)
		return
	}
	defer func() {
		_ = kernelDll.CloseHandle(pi.HProcess)
		_ = kernelDll.CloseHandle(pi.HThread)
	}()

	entrySlice, errno := wGo.GetProcessEntry(func(entry *w32.PROCESSENTRY32W) bool {
		return entry.Th32ProcessID == pi.ProcessId
	})
	if errno != 0 {
		return
	}

	wGo.KillProcess(entrySlice, func(entry *w32.PROCESSENTRY32W, errno syscall.Errno) {
		if errno == 0 {
			fmt.Println("Kill Process: ", entry.ExeFileName())
		}
	})
	wGo.KillProcess(entrySlice, nil) // 可以傳nil，只是忽略刪除之後想要表達的事情而已

	// 由於該entry清單已經被我們移除，所以在嘗試刪除會得到Access is denied.的錯誤
	wGo.KillProcess(entrySlice, func(entry *w32.PROCESSENTRY32W, errno syscall.Errno) {
		fmt.Println(errno)
	})

	// Output:
	// Kill Process: notepad.exe
	// Access is denied.
}

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
