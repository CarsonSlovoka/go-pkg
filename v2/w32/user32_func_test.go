package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"unsafe"
)

func ExampleUser32DLL_GetWindowText() {
	user32dll := w32.NewUser32DLL(
		w32.PNGetForegroundWindow,
		w32.PNGetClassName,
		w32.PNGetWindowText,
	)

	curHwnd := user32dll.GetForegroundWindow()
	log.Println("current window HWND:", curHwnd) // 當前窗口的識別號

	clsName, err := user32dll.GetClassName(curHwnd)
	if err != nil {
		panic(err)
	}
	log.Println("window class Name:", clsName)

	winText, err := user32dll.GetWindowText(curHwnd)
	if err != nil {
		panic(err)
	}
	log.Println("window text Name:", winText)

	// Output:
}

func ExampleUser32DLL_MessageBox() {
	user32dll := w32.NewUser32DLL(
		w32.PNMessageBox,
	)

	hwndTop := w32.HWND_TOP
	response, _ := user32dll.MessageBox(hwndTop, "title", "body message", w32.MB_OK)
	switch response {
	case w32.IDYES:
		fmt.Println("Yes")
	case w32.IDNO:
		fmt.Println("No")
	default: // 基本上沒辦法有這個出現，對話框只有Yes,No可以選，右上角的打X也不能按
		fmt.Println("Unknown")
	}

	messageBox := user32dll.MessageBox
	_, _ = messageBox(hwndTop, "Test", "OK", w32.MB_OK)
	_, _ = messageBox(hwndTop, "Test", "Yes No Cancel", w32.MB_YESNOCANCEL)
	_, _ = messageBox(hwndTop, "Test", "OK", w32.MB_OK)
	_, _ = messageBox(hwndTop, "Test", "Help button", w32.MB_HELP)
	_, _ = messageBox(hwndTop, "Test", "OK CANCEL", w32.MB_OKCANCEL)
	_, _ = messageBox(hwndTop, "Test", "ABORT RETRY IGNORE", w32.MB_ABORTRETRYIGNORE)
	_, _ = messageBox(hwndTop, "Test", "RETRY CANCEL", w32.MB_RETRYCANCEL)
	_, _ = messageBox(hwndTop, "Test", "CANCEL TRY CONTINUE", w32.MB_CANCELTRYCONTINUE)

	// newline
	_, _ = messageBox(hwndTop, "Test", "row1\nrow2\nrow3", w32.MB_OK)

	body := `r1
r2
...
rn`
	_, _ = messageBox(hwndTop, "Test", body, w32.MB_OK)

	// Icon
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONSTOP|
		w32.MB_RIGHT| // text right-justified
		// w32.MB_TOPMOST,
		w32.MB_SYSTEMMODAL, // 比使用MB_TOPMOST好
	)
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONQUESTION)
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONWARNING)
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONINFORMATION)
	// Output
}

// 抓取icon畫在notepad應用程式上(如果要執行，請確保您有運行nodepad.exe)
// https://learn.microsoft.com/en-us/windows/win32/menurc/using-icons#creating-an-icon
func ExampleUser32DLL_DrawIcon() {
	user32dll := w32.NewUser32DLL(
		w32.PNLoadIcon,
		w32.PNDrawIcon,
		w32.PNGetDC,
		w32.PNFindWindow,

		w32.PNSendMessage,
		w32.PNFindWindow,

		w32.PNReleaseDC,
	)

	// 獲取HICON{question, chrome}
	var hIconQuestion, hIconChrome w32.HICON
	{
		var err error
		// 取得系統圖標question
		hIconQuestion, err = user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION))
		if err != nil {
			log.Println("系統圖標: QUESTION 找不到")
			return
		}

		// 取得chrome的圖標
		hwndChrome := user32dll.FindWindow("Chrome_WidgetWin_1", "")
		if hwndChrome == 0 {
			log.Println("找不到chrome窗口")
			return
		}

		hIcon, _, _ := user32dll.SendMessage(uintptr(hwndChrome), w32.WM_GETICON, w32.ICON_SMALL, 0)
		hIconChrome = w32.HICON(hIcon)
		if hIconChrome == 0 {
			log.Println("chrome圖標獲取失敗")

			// 嘗試使用LoadIcon函數取得
			hIconChrome, _ = user32dll.LoadIcon(uintptr(hwndChrome), w32.MakeIntResource(w32.IDI_APPLICATION))
			if hIconChrome == 0 {
				// Alternative method. Use OS default icon
				hIconChrome, _ = user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_APPLICATION))
			}
		}
	}

	// 建立HDC
	var curHDC w32.HDC
	{
		// 獲取notepad的hdc對象
		hwndNotepad := user32dll.FindWindow("Notepad", "")
		if hwndNotepad == 0 {
			log.Println("找不到Notepad窗口")
			return
		}
		curHDC = user32dll.GetDC(hwndNotepad)

		defer func() {
			if curHDC != 0 {
				if user32dll.ReleaseDC(hwndNotepad, curHDC) == 0 {
					log.Fatal("ReleaseDC")
				}
			}
		}()
	}

	// 將圖標輸出
	{
		for _, d := range []struct {
			x     int // 要畫在哪一個位置
			y     int
			hIcon w32.HICON
		}{
			{50, 100, hIconQuestion},
			{50, 200, hIconQuestion},
			{50, 300, hIconChrome},
		} {
			if err := user32dll.DrawIcon(curHDC, d.x, d.y, d.hIcon); err != nil {
				panic(err)
			}
		}
	}

	// Output:
}

// 本範例流程
// 1. HICON: 取得hIcon
// 2. ICONINFO: 由該hIcon透過GetIconInfo取得到ICONINFO的結構資料
// 3. init BITMAP: 建立空的BITMAP對象(初始化參照ICONINFO.HbmColor)
// 4. copy to BITMAP: 透過CopyImage將ICONINFO.HbmColor複製到我們所建立的BITMAP之中
func ExampleUser32DLL_GetIconInfo() {
	user32dll := w32.NewUser32DLL(
		w32.PNLoadIcon,
		w32.PNGetIconInfo,
		w32.PNCopyImage,
	)

	gdi32dll := w32.NewGdi32DLL(
		w32.PNGetObject,
		w32.PNDeleteObject,
	)

	hIconQuestion, err := user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION))
	if err != nil {
		return
	}

	var iInfo w32.ICONINFO
	if !user32dll.GetIconInfo(hIconQuestion, &iInfo) {
		return
	}
	// 當獲取成功之後，如果不用時，要把所有HBITMAP的對象釋放掉
	defer func() {
		if !gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmColor)) {
			fmt.Println("error DeleteObject HbmColor")
		}
		if !gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmMask)) {
			fmt.Println("error DeleteObject HbmMask")
		}
	}()
	log.Printf("%+v\n", iInfo)
	fmt.Println("ok")

	// 以下為copyImage的測試，沿用上面取得到的icon
	{
		// 以ICONINFO的資料建立一個空的BITMAP
		bmp := w32.Bitmap{}
		if gdi32dll.GetObject(w32.HANDLE(iInfo.HbmColor), int32(unsafe.Sizeof(bmp)), uintptr(unsafe.Pointer(&bmp))) == 0 {
			return
		}

		w32.NewUser32DLL()
		var hBmp w32.HBITMAP
		hBmpHandle, errno := user32dll.CopyImage(w32.HANDLE(iInfo.HbmColor), w32.IMAGE_BITMAP, 0, 0, 0)
		if errno != 0 {
			return
		}
		hBmp = w32.HBITMAP(hBmpHandle)
		fmt.Println("copyImage OK")
		log.Println(hBmp)

		if !gdi32dll.DeleteObject(w32.HGDIOBJ(hBmp)) {
			fmt.Println("error")
		}
	}
	// Output:
	// ok
	// copyImage OK
}

func ExampleUser32DLL_PostMessage() {
	user32dll := w32.NewUser32DLL(
		w32.PNPostMessage,
	)
	if _, _, err := user32dll.PostMessage(uintptr(w32.HWND_BROADCAST), w32.WM_FONTCHANGE, 0, 0); err != nil {
		panic(err)
	}
}

func ExampleUser32DLL_FindWindow() {
	user32dll := w32.NewUser32DLL(
		w32.PNFindWindow,
	)

	// Chrome
	// "Chrome_WidgetWin_1" You can find this information from Spy++ tool
	hwnd := user32dll.FindWindow("Chrome_WidgetWin_1", "")
	log.Println(hwnd)
	// Output:
}

func ExampleUser32DLL_FindWindowEx() {
	user32dll := w32.NewUser32DLL(w32.PNFindWindowEx)
	hwnd := user32dll.FindWindowEx(0, 0, "Notepad", "")
	log.Println(hwnd)
	// Output:
}

func ExampleUser32DLL_SendMessage() {
	user32dll := w32.NewUser32DLL(
		w32.PNSendMessage,
		w32.PNFindWindow,
	)

	hIcon, _, _ := user32dll.SendMessage(0, w32.WM_GETICON, w32.ICON_SMALL, 0)
	log.Println(hIcon)
	// Output:
}
