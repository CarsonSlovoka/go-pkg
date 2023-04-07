package w32_test

import (
	"encoding/json"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"path/filepath"
	"syscall"
	"testing"
)

func TestPsApiDLL_EnumProcesses(t *testing.T) {
	const maxProcessCount uint32 = 9000000
	_, numProcesses, err := psApiDll.EnumProcesses(maxProcessCount)
	if err != 0 {
		t.Fatal("Failed to enumerate processes:", err)
		return
	}
	if numProcesses >= maxProcessCount {
		t.Fatal("理論上不會有那麼多process，所以process的數量要小於", maxProcessCount)
	}

	_, numProcesses, _ = psApiDll.EnumProcesses(1)
	if numProcesses <= 1 {
		t.Fatal("當給定的size不足的時候，會自動增加，所以理應當大於1")
	}
}

func ExamplePsApiDLL_EnumProcesses() {
	var (
		processIDs   []uint32
		numProcesses uint32
		eno          syscall.Errno
	)
	processIDs, numProcesses, eno = psApiDll.EnumProcesses(1024) // 一般都在400以內，所以通常用1024已經很足夠
	if eno != 0 {
		fmt.Println("Failed to enumerate processes:", eno)
		return
	}
	log.Println("numProcesses:", numProcesses)
	log.Println(processIDs)
	// Output:
}

func TestPsApiDLL_GetModuleFileNameExW(t *testing.T) {
	procIDs, numProcesses, err := psApiDll.EnumProcesses(512)
	if err != 0 {
		fmt.Println("Failed to enumerate processes:", err)
		return
	}
	log.Println("numProcesses:", numProcesses)
	moreThan256Modules := 0 // 只是用來統計有多少Process他的module有超過256個

	searchExeName := "notepad++.exe"
	var printDetails = false

	for i := uint32(0); i < numProcesses; i++ {
		needLog := false

		curPID := procIDs[i]
		hProcess, eno := kernelDll.OpenProcess(w32.PROCESS_QUERY_INFORMATION|w32.PROCESS_VM_READ, // GetModuleFileNameExW會需要用到PROCESS_VM_READ
			false, curPID,
		)
		if eno != 0 && eno != w32.ERROR_ACCESS_DENIED {
			log.Printf("Failed to open process %d: %s\n", curPID, eno)
			continue
		}
		// Get the path of the process's main module
		hModules, numModule, eno := psApiDll.EnumProcessModules(hProcess, 256)
		if eno != 0 && eno != w32.ERROR_INVALID_HANDLE {
			log.Printf("Failed to enumerate process modules for process %d: %s\n", curPID, eno)
			_ = kernelDll.CloseHandle(hProcess)
			continue
		}

		processInfo := struct { // 只是為了方便打印結果
			PID       uint32
			NumModule uint32
			ModuleMap map[string][]w32.HMODULE // [name]
		}{curPID, numModule, make(map[string][]w32.HMODULE, numModule)}

		if numModule > 200 { // 避免打印太多結果，只記錄模組特別多的資料
			needLog = true
		}

		var exePath string
		if numModule > 0 {
			exePath, _ = psApiDll.GetModuleFileNameExW(hProcess, 0, 1024) // 當hModule為0(NULL)則表示取該執行檔的路徑，相當於GetModuleFileNameExW(hProcess, hModules[0], 1024)

			if filepath.Base(exePath) == searchExeName && searchExeName != "" {
				log.Println(searchExeName, "found")
			}
		}

		for j := uint32(0); j < numModule; j++ {
			if j == 0 && needLog {
				log.Println(exePath)
			}

			var moduleFileName string
			moduleFileName, eno = psApiDll.GetModuleFileNameExA(hProcess, hModules[j], 1024)

			if hModules[j] != 0 { // 不紀錄0的資料
				processInfo.ModuleMap[moduleFileName] = append(processInfo.ModuleMap[moduleFileName], hModules[j])
			}

			if eno != 0 {
				log.Println(eno)
				continue
			}
			// log.Printf("PID: %05d, moduleFilename: %s", curPID, moduleFilename)
		}

		if processInfo.NumModule > 256 {
			moreThan256Modules += 1
		}

		if needLog && printDetails {
			bs, _ := json.MarshalIndent(processInfo, "", "  ")
			log.Println(string(bs))
		}

		_ = kernelDll.CloseHandle(hProcess)
	}
	log.Printf("how many processes have more than 256 modules: %d\n", moreThan256Modules)
}

func ExamplePsApiDLL_EnumProcessModules() {
	procIDs, numProcesses, _ := psApiDll.EnumProcesses(1024)
	searchExeName := "notepad++.exe"

	for i := uint32(0); i < numProcesses; i++ {
		curPID := procIDs[i]
		hProcess, eno := kernelDll.OpenProcess(w32.PROCESS_QUERY_INFORMATION|w32.PROCESS_VM_READ, // GetModuleFileNameExW會需要用到PROCESS_VM_READ
			false, curPID,
		)
		if eno != 0 {
			continue
		}
		_, numModule, eno := psApiDll.EnumProcessModulesEx(hProcess, 256, w32.LIST_MODULES_ALL)
		if eno != 0 {
			continue
		}

		if numModule > 0 {
			exePath, _ := psApiDll.GetModuleFileNameExW(hProcess, 0, 1024) // 當hModule為0(NULL)則表示取該執行檔的路徑，相當於GetModuleFileNameExW(hProcess, hModules[0], 1024)
			if filepath.Base(exePath) == searchExeName {
				log.Println(searchExeName, "found")
			}
		}
		_ = kernelDll.CloseHandle(hProcess)
	}
	// Output:
}
