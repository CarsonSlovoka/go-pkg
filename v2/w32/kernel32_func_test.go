package w32_test

import (
	"fmt"
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
	handle1, err := kernel32dll.CreateMutex(nil, false, "hello world")
	if err != 0 {
		t.Fatal(err)
	}

	// 再嘗試創建一個相同的Mutex，它會報錯，但是截至目前為止已經有兩個"hello world"所創建的Mutex(錯誤的也算在內)
	// 如果沒把所有該名稱相符的Mutex都Close掉，即便已經關閉了一個，之後再次創建仍然會報ERROR_ALREADY_EXISTS的訊息
	handle2, err := kernel32dll.CreateMutex(nil, false, "hello world")
	if err != syscall.ERROR_ALREADY_EXISTS {
		t.Fatal("not as expected")
	}

	if ok, err := kernel32dll.CloseHandle(handle1); !ok {
		t.Fatal(err)
	}
	if ok, err := kernel32dll.CloseHandle(handle2); !ok { // 如果我們省略了這個closeHandle之後再嘗試創建一次，會報ERROR_ALREADY_EXISTS的錯誤
		t.Fatal(err)
	}

	// We can create again since we have closed (close ALL)
	handle3, err := kernel32dll.CreateMutex(nil, false, "hello world")
	if err != w32.NO_ERROR {
		t.Fatal(err)
	}
	if ok, err := kernel32dll.CloseHandle(handle3); !ok {
		t.Fatal(err)
	}
}

func ExampleKernel32DLL_CreateMutex() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNCreateMutex,
		w32.PNCloseHandle,
	)
	handle, err := kernel32dll.CreateMutex(nil, false, "hello world")
	if err != 0 {
		return
	}

	if ok, err := kernel32dll.CloseHandle(handle); !ok {
		fmt.Println(err)
		return
	}

	// We can create again since we have closed.
	handle, err = kernel32dll.CreateMutex(nil, false, "hello world")
	if err != w32.NO_ERROR {
		return
	}
	if ok, _ := kernel32dll.CloseHandle(handle); !ok {
		return
	}
	fmt.Println("ok")
	// Output:
	// ok

}

func ExampleKernel32DLL_GetNativeSystemInfo() {
	kernel32dll := w32.NewKernel32DLL(w32.PNGetNativeSystemInfo)
	info := kernel32dll.GetNativeSystemInfo()
	// ProcessorArchitecture: https://github.com/CarsonSlovoka/go-pkg/blob/8d251b6a295cc4177593e9dae7455955e769e88d/v2/w32/const.go#L7-L12
	log.Println(info.ProcessorArchitecture)
	// Output:
}

func ExampleKernel32DLL_CreateFile() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNCreateFile,
		w32.PNCloseHandle,
	)
	testFilePath := "testdata/temp.txt"
	hFile, errno := kernel32dll.CreateFile(testFilePath,
		w32.GENERIC_READ|w32.GENERIC_WRITE,
		0,
		0,
		w32.CREATE_ALWAYS,
		w32.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if errno != 0 { // w32.NO_ERROR
		return
	}
	if ok, _ := kernel32dll.CloseHandle(hFile); !ok {
		return
	}
	defer os.Remove(testFilePath)
	hFile, errno = kernel32dll.CreateFile(testFilePath,
		w32.GENERIC_READ|w32.GENERIC_WRITE,
		0,
		0,
		w32.CREATE_ALWAYS, // 如果檔案已經存在，還是可以創建成功，但是錯誤代碼會回傳:183(ERROR_ALREADY_EXISTS)
		w32.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if ok, _ := kernel32dll.CloseHandle(hFile); !ok {
		return
	}
	if errno != w32.ERROR_ALREADY_EXISTS {
		return
	}
	fmt.Printf("%s\n", errno)
	fmt.Println("ok")
	// Output:
	// Cannot create a file when that file already exists.
	// ok
}

// https://learn.microsoft.com/en-us/windows/win32/menurc/using-resources
func ExampleKernel32DLL_CopyFile() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNCreateFile,
		w32.PNCloseHandle,
		w32.PNGetLastError,

		w32.PNCopyFile,
	)

	targetPath, _ := filepath.Abs("testdata/temp.txt") // 避免反斜線問題
	hFile, _ := kernel32dll.CreateFile(targetPath,     // name of file
		w32.GENERIC_READ|w32.GENERIC_WRITE, // access mode
		0,                                  // share mode
		0,                                  // default security
		w32.CREATE_ALWAYS,                  // create flags
		w32.FILE_ATTRIBUTE_NORMAL,          // file attributes
		0,                                  // no template
	)

	if int(hFile) == w32.INVALID_HANDLE_VALUE {
		fmt.Println("Could not create file.") // 有可能是目錄路徑不存在
		return
	}
	if ok, _ := kernel32dll.CloseHandle(hFile); !ok {
		fmt.Println("error: CloseHandle")
		return
	}

	_ = os.WriteFile(targetPath, []byte("Hello World"), 0x666)

	// test copy file
	{
		copyFilepath := "testdata/temp-copy.txt"
		if !kernel32dll.CopyFile(targetPath, copyFilepath, false) {
			fmt.Println("error: copy file")
			return
		}
		bs, _ := os.ReadFile(copyFilepath)
		_ = os.Remove(copyFilepath)
		fmt.Println(string(bs))
	}

	defer func() {
		if err := os.Remove(targetPath); err != nil { // 刪除測試檔案
			fmt.Println("error: remove file")
		}
	}()

	// Output:
	// Hello World
}

func ExampleKernel32DLL_GetModuleHandle() {
	kernel32dll := w32.NewKernel32DLL(
		w32.PNGetModuleHandle,
		w32.PNLoadLibrary,
		w32.PNFreeLibrary,
	)
	hModule := kernel32dll.GetModuleHandle("") // nil表示應用程式自己本身
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

		defer func(hModule w32.HMODULE) {
			// 注意freeLibrary的對象不要使用GetModuleHandle出來的handle，有可能會出問題！ 要使用LoadLibrary的handle
			if ok := kernel32dll.FreeLibrary(hModule); !ok {
				log.Fatal("FreeLibrary")
			}
		}(hExe)

		hModule2 := kernel32dll.GetModuleHandle(exePath)
		if hModule2 != 0 {
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

	var (
		hExe      w32.HMODULE
		hRes      w32.HRSRC
		lpResLock uintptr
	)
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
		w32.PNFindResource,
		w32.PNLoadResource,
		w32.PNSizeofResource,
		w32.PNLockResource,
	)
	user32dll := w32.NewUser32DLL(
		w32.PNLookupIconIdFromDirectoryEx,
		w32.PNCreateIconFromResourceEx,

		w32.PNFindWindow,
		w32.PNReleaseDC,
		w32.PNDrawIcon,
		w32.PNGetDC,
	)
	hExe := kernel32dll.LoadLibrary(exePath)
	if hExe == 0 {
		// not found
		return
	}

	hResource := kernel32dll.FindResource(hExe,
		w32.MakeIntResource(w32.UintptrFromStr("IDI_BTH_UD_TASK")), // w32.MakeIntResource(150) // 該資源有哪些ID，可以安裝Resource Hacker去查看。以微軟的fontview.exe，它擁有Icon Group: 150: 1033這個資源
		w32.MakeIntResource(w32.RT_GROUP_ICON),                     // w32.MakeIntResource(w32.UintptrFromStr("xfont"))
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
	var hdc w32.HDC
	{
		hwndNotepad := user32dll.FindWindow("Notepad", "")
		if hwndNotepad == 0 {
			return
		}
		hdc = user32dll.GetDC(hwndNotepad)

		defer func() {
			if hdc != 0 {
				if user32dll.ReleaseDC(hwndNotepad, hdc) == 0 {
					log.Fatal("ReleaseDC")
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
		w32.PNFindResource,
		w32.PNLoadResource,
		w32.PNSizeofResource,
		w32.PNLockResource,
	)
	user32dll := w32.NewUser32DLL(
		w32.PNCreateIconFromResourceEx,

		w32.PNFindWindow,
		w32.PNReleaseDC,
		w32.PNDrawIcon,
		w32.PNGetDC,
	)
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
	var hdc w32.HDC
	{
		hwndNotepad := user32dll.FindWindow("Notepad", "")
		if hwndNotepad == 0 {
			return
		}
		hdc = user32dll.GetDC(hwndNotepad)

		defer func() {
			if hdc != 0 {
				if user32dll.ReleaseDC(hwndNotepad, hdc) == 0 {
					log.Fatal("ReleaseDC")
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
