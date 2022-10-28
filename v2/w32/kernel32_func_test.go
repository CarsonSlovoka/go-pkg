package w32_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestCreateMutex(t *testing.T) {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNCreateMutex,
		w32.PNCloseHandle,
	)
	handle, err := kernel32dll.CreateMutex("hello world")
	if err != nil {
		t.Fatal(err)
	}

	_, err = kernel32dll.CreateMutex("hello world")
	if err == nil || err != syscall.ERROR_ALREADY_EXISTS {
		t.Error("not as expected")
	}

	if err = kernel32dll.CloseHandle(handle); err != nil {
		t.Error(err)
	}
	// err = kernel32dll.CloseHandle(handle) // If you are debugging it will panic!
	// fmt.Printf("%+v\n%d", err, err.(syscall.Errno))     // The Handle is invalid.  6

	// We can create again since we have closed.
	handle, _ = kernel32dll.CreateMutex("hello world")
	if err = kernel32dll.CloseHandle(handle); err != nil {
		t.Error(err)
	}
}

// TODO 有待驗證
// https://learn.microsoft.com/en-us/windows/win32/menurc/using-resources
func ExampleKernel32DLL_CreateFile() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNCreateFile,
	)
	hFile := kernel32dll.CreateFile("resinfo.txt", // name of file
		w32.GENERIC_READ|w32.GENERIC_WRITE, // access mode
		0,                                  // share mode
		0,                                  // default security
		w32.CREATE_ALWAYS,                  // create flags
		w32.FILE_ATTRIBUTE_NORMAL,          // file attributes
		0,                                  // no template
	)
	if int(hFile) == w32.INVALID_HANDLE_VALUE {
		log.Fatal("Could not open file.")
	}
	log.Println(hFile)
}

func ExampleKernel32DLL_GetModuleHandle() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNGetModuleHandle,
		w32.PNLoadLibrary,
		w32.PNFreeLibrary,
	)
	hModule := kernel32dll.GetModuleHandle(nil) // nil表示應用程式自己本身
	if hModule == 0 {
		log.Fatal("GetModuleHandle")
	}
	log.Println(hModule)

	// 範例二，載入「其它」檔案
	// GetModuleHandle的項目，要取得到資料的前提是該資料並須先被載入，所以我們要用LoadLibrary把該資料載入
	{
		exePath := filepath.Join(os.Getenv("windir"), "System32/fontview.exe")
		if _, err := os.Stat(exePath); os.IsNotExist(err) {
			return // 無法載入。 // 不確定github.action的虛擬環境是否有此檔案，所以忽略
		}
		hExe := kernel32dll.LoadLibrary(exePath)
		if hExe == 0 {
			log.Fatal("LoadLibrary")
		}

		defer func(hmodule uintptr) {
			// 注意freeLibrary的對象不要使用GetModuleHandle出來的handle，有可能會出問題！ 要使用LoadLibrary的handle
			if ok := kernel32dll.FreeLibrary(hmodule); !ok {
				log.Fatal("FreeLibrary")
			}
		}(hExe)

		uint16prtModulePath, err := syscall.UTF16PtrFromString(exePath)
		if err != nil {
			return
		}
		hModule2 := kernel32dll.GetModuleHandle(uint16prtModulePath)
		if hModule != 0 {
			log.Println(hModule2)
		}
	}
	// Output:
}

// https://learn.microsoft.com/en-us/windows/win32/menurc/using-resources#updating-resources
// 從A應用程式抓取其資源，放入到B應用程式之中
func ExampleKernel32DLL_UpdateResource() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNLoadLibrary,
		w32.PNFindResource,
		w32.PNLoadResource,
		w32.PNLockResource,
		w32.PNBeginUpdateResource,
		w32.PNSizeofResource,
		w32.PNUpdateResource,
		w32.PNEndUpdateResource,
		w32.PNFreeLibrary,
	)

	// 最好透過filepath，避免反斜線錯邊的問題(backslash or forward slash)
	sourcePath := filepath.Join("testdata/exe", "writeWithFont.exe") // 資源的來源
	targetPath := filepath.Join("testdata/exe", "write.exe")         // 我們不想讓原始應用程式被修改，所以準備copy一份
	targetBytes, _ := os.ReadFile(targetPath)
	targetPath = filepath.Join("testdata/exe", "write-copy.exe") // 把資源放到這個應用程式之中

	if f, err := os.Create(targetPath); err != nil {
		panic(err)
	} else {
		_, _ = f.Write(targetBytes)
		_ = f.Close()

		defer func() {
			_ = os.Remove(targetPath) // 結束之後把測試檔案刪掉，如果您要查看範例結果，請在這邊下中斷點，避免看不到輸出的結果檔案
		}()
	}

	/* writeFile檔案已經存在時，會報錯
	if err := os.WriteFile(targetPath, targetBytes, 0x666); err != nil {
		panic(err)
	}
	*/

	var lpResLock, hExe uintptr
	var hRes w32.HRSRC
	{
		hExe = kernel32dll.LoadLibrary(sourcePath)
		if hExe == 0 {
			log.Fatal("Could not load exe.")
		}

		defer func() {
			e := recover()
			if !kernel32dll.FreeLibrary(hExe) {
				log.Fatal("Could not free executable.")
			}
			if e != nil {
				panic("should panic")
			}
		}()

		hRes = kernel32dll.FindResource(hExe,
			w32.MakeIntResource(666), // 這是該應用程式所對應的resourceID
			w32.MakeIntResource(w32.RT_FONT),
		)
		if hRes == 0 {
			log.Fatal("Could not locate font.")
		}

		hResLoad := kernel32dll.LoadResource(hExe, hRes)
		if hResLoad == 0 {
			log.Fatal("Could not locate font resource.")
		}

		lpResLock = kernel32dll.LockResource(hResLoad)
		if lpResLock == 0 {
			log.Fatal("Could not lock dialog box.")
		}

	}

	hUpdateRes := kernel32dll.BeginUpdateResource(targetPath, false)
	if hUpdateRes == 0 {
		log.Fatal("Could not open file for writing.")
		return
	}

	if ok := kernel32dll.UpdateResource(hUpdateRes,
		w32.RT_FONT,
		w32.MakeIntResource(666),
		w32.MakeLangID(w32.LANG_ENGLISH, w32.SUBLANG_ENGLISH_US),
		lpResLock,
		kernel32dll.SizeofResource(hExe, hRes),
	); !ok {
		log.Fatal("Could not add resource.")
	}

	if !kernel32dll.EndUpdateResource(hUpdateRes, false) {
		log.Fatal("Could not write changes to file.")
	}

	// Output:
}

// 本範例抓取指定應用程式中groupIcon.ID=xxx的資料，再從中挑選出比較合適的ICON來畫在notepad上
// https://learn.microsoft.com/en-us/windows/win32/menurc/using-icons#creating-an-icon
// https://blog.csdn.net/qq_41490873/article/details/112250964?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-112250964-blog-121427409.pc_relevant_layerdownloadsortv1&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-112250964-blog-121427409.pc_relevant_layerdownloadsortv1&utm_relevant_index=1
func ExampleKernel32DLL_FindResource() {
	exePath := filepath.Join(os.Getenv("windir"), "System32/bthudtask.exe")
	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		return
	}
	kernel32dll := w32.NewKernel32DLL(
		w32.PNLoadLibrary,
		w32.PNCreateFile,
		w32.PNFindResource,
		w32.PNLoadResource,
		w32.PNSizeofResource,
		w32.PNLockResource,
	)
	user32dll := w32.NewUser32DLL([]w32.ProcName{
		w32.PNLookupIconIdFromDirectoryEx,
		w32.PNCreateIconFromResourceEx,

		w32.PNFindWindow,
		w32.PNReleaseDC,
		w32.PNDrawIcon,
		w32.PNGetDC,
	})
	hExe := kernel32dll.LoadLibrary(exePath)
	if hExe == 0 {
		// not found
		return
	}

	hResource := kernel32dll.FindResource(hExe,
		w32.MakeIntResource(w32.StrToLPCWSTR("IDI_BTH_UD_TASK")), // w32.MakeIntResource(150) // 該資源有哪些ID，可以安裝Resource Hacker去查看。以微軟的fontview.exe，它擁有Icon Group: 150: 1033這個資源
		w32.MakeIntResource(w32.RT_GROUP_ICON),                   // w32.MakeIntResource(w32.StrToLPCWSTR("xfont"))
	)

	hMem := kernel32dll.LoadResource(hExe, hResource)

	lpResource := kernel32dll.LockResource(hMem)

	// Get the identifier of the icon that is most appropriate
	// for the video display
	nID := user32dll.LookupIconIdFromDirectoryEx(lpResource,
		true,
		w32.SM_CXICON,
		w32.SM_CYICON,
		w32.LR_DEFAULTCOLOR,
	)

	// Find the bits for the nID icon.
	hResource = kernel32dll.FindResource(hExe,
		w32.MakeIntResource(uintptr(nID)),
		w32.MakeIntResource(w32.RT_ICON),
	)

	// Load and lock the icon.
	hMem = kernel32dll.LoadResource(hExe, hResource)
	lpResource = kernel32dll.LockResource(hMem)

	hIcon1 := user32dll.CreateIconFromResourceEx(lpResource,
		kernel32dll.SizeofResource(hExe, hResource), true, 0x00030000,
		w32.SM_CXICON, w32.SM_CYICON, w32.LR_DEFAULTCOLOR)

	// init HDC
	var hdc uintptr
	{
		hwndNotepad, err := user32dll.FindWindow("Notepad", "")
		if err != nil {
			return
		}
		hdc = user32dll.GetDC(hwndNotepad)

		defer func() {
			if hdc != 0 {
				if err = user32dll.ReleaseDC(hwndNotepad, hdc); err != nil {
					log.Fatal(err)
				}
			}
		}()
	}

	// Draw the icon in the client area.
	if err := user32dll.DrawIcon(hdc, 10, 20, hIcon1); err != nil {
		log.Fatal("DrawIcon")
	}
	// Output:
}

// 類似ExampleKernel32DLL_FindResource，不過本範例直接抓取ICON不再從Icon Group去找尋
// 建議您安裝[Resource Hacker](http://www.angusj.com/resourcehacker/)去查看微軟的fontView.exe會對本範例更了解
func ExampleKernel32DLL_FindResource_icon() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNLoadLibrary,
		w32.PNCreateFile,
		w32.PNFindResource,
		w32.PNLoadResource,
		w32.PNSizeofResource,
		w32.PNLockResource,
	)
	user32dll := w32.NewUser32DLL([]w32.ProcName{
		w32.PNCreateIconFromResourceEx,

		w32.PNFindWindow,
		w32.PNReleaseDC,
		w32.PNDrawIcon,
		w32.PNGetDC,
	})
	exePath := filepath.Join(os.Getenv("windir"), "System32/fontview.exe")
	hExe := kernel32dll.LoadLibrary(exePath)
	if hExe == 0 {
		// not found
		return
	}

	hResource := kernel32dll.FindResource(hExe,
		w32.MakeIntResource(1), // 抓取ICON中ID為1的資源
		w32.MakeIntResource(w32.RT_ICON),
	)

	// 載入資源兩個動作:
	// 1. LoadResource
	// 2. LockResource
	hMem := kernel32dll.LoadResource(hExe, hResource)
	lpResource := kernel32dll.LockResource(hMem)

	hIcon := user32dll.CreateIconFromResourceEx(lpResource,
		kernel32dll.SizeofResource(hExe, hResource),
		true, 0x00030000,
		w32.SM_CXICON, w32.SM_CYICON, w32.LR_DEFAULTCOLOR)

	// init HDC
	var hdc uintptr
	{
		hwndNotepad, err := user32dll.FindWindow("Notepad", "")
		if err != nil {
			return
		}
		hdc = user32dll.GetDC(hwndNotepad)

		defer func() {
			if hdc != 0 {
				if err = user32dll.ReleaseDC(hwndNotepad, hdc); err != nil {
					log.Fatal(err)
				}
			}
		}()
	}

	// Draw the icon in the client area.
	if err := user32dll.DrawIcon(hdc, 10, 20, hIcon); err != nil {
		log.Fatal("DrawIcon")
	}
	// Output:
}
