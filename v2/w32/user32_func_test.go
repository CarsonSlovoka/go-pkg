package w32_test

import (
	"encoding/binary"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"
	"unicode/utf16"
	"unsafe"
)

func ExampleUser32DLL_GetWindowLongPtr() {
	hwnd := getTestHwnd()
	style, _ := userDll.GetWindowLongPtr(hwnd, w32.GWL_STYLE)
	style &^= w32.WS_MAXIMIZEBOX | w32.WS_MINIMIZEBOX // Remove the max and min buttons.
	style ^= w32.WS_THICKFRAME                        // Toggle, resizeable // 如果沒有此屬性就加上，有的話就移除
	if _, eno := userDll.SetWindowLongPtr(hwnd, w32.GWL_STYLE, style); eno != 0 {
		log.Println(eno)
	}
	// Output:
}

func ExampleUser32DLL_GetWindowText() {
	user32dll := w32.NewUser32DLL(
		w32.PNGetForegroundWindow,
		w32.PNGetClassName,
		w32.PNGetWindowText,
	)

	curHwnd := user32dll.GetForegroundWindow()
	log.Println("current window HWND:", curHwnd) // 當前窗口的識別號
	if curHwnd == 0 {
		return
	}

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

// Automatically close the message box
func TestMessageBox(t *testing.T) {
	user32dll := w32.NewUser32DLL(
		w32.PNMessageBox,
		w32.PNFindWindow,
		w32.PNSendMessage,
		w32.PNPostMessage,
	)
	go func() {
		_, _ = user32dll.MessageBox(0, "...", "TestBox", w32.MB_OK)
	}()
	time.Sleep(200 * time.Millisecond)
	hwnd := user32dll.FindWindow("", "TestBox")
	if hwnd != 0 {
		// _, _ = user32dll.PostMessage(hwnd, w32.WM_CLOSE, 0, 0) // 如果想直接送了就不管，可以使用這個
		_, _, _ = user32dll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)
	}
}

func ExampleUser32DLL_MessageBox() {
	user32dll := w32.NewUser32DLL(
		w32.PNMessageBox,
	)

	hwndTop := w32.HWND_TOP
	response, _ := user32dll.MessageBox(hwndTop, "body message", "title", w32.MB_OK)
	switch response {
	case w32.IDYES:
		fmt.Println("Yes")
	case w32.IDNO:
		fmt.Println("No")
	default: // 基本上沒辦法有這個出現，對話框只有Yes,No可以選，右上角的打X也不能按
		fmt.Println("Unknown")
	}

	messageBox := user32dll.MessageBox
	_, _ = messageBox(hwndTop, "OK", "Test", w32.MB_OK)
	_, _ = messageBox(hwndTop, "Yes No Cancel", "Test", w32.MB_YESNOCANCEL)
	_, _ = messageBox(hwndTop, "OK", "Test", w32.MB_OK)
	_, _ = messageBox(hwndTop, "Help button", "Test", w32.MB_HELP)
	_, _ = messageBox(hwndTop, "OK CANCEL", "Test", w32.MB_OKCANCEL)
	_, _ = messageBox(hwndTop, "ABORT RETRY IGNORE", "Test", w32.MB_ABORTRETRYIGNORE)
	_, _ = messageBox(hwndTop, "RETRY CANCEL", "Test", w32.MB_RETRYCANCEL)
	_, _ = messageBox(hwndTop, "CANCEL TRY CONTINUE", "Test", w32.MB_CANCELTRYCONTINUE)

	// newline
	_, _ = messageBox(hwndTop, "row1\nrow2\nrow3", "Test", w32.MB_OK)

	body := `r1
r2
...
rn`
	_, _ = messageBox(hwndTop, body, "Test", w32.MB_OK)

	// Icon
	_, _ = messageBox(0, "OK", "Test", w32.MB_OK|w32.MB_ICONSTOP|
		w32.MB_RIGHT| // text right-justified
		// w32.MB_TOPMOST,
		w32.MB_SYSTEMMODAL, // 比使用MB_TOPMOST好
	)
	_, _ = messageBox(0, "OK", "Test", w32.MB_OK|w32.MB_ICONQUESTION)
	_, _ = messageBox(0, "OK", "Test", w32.MB_OK|w32.MB_ICONWARNING)
	_, _ = messageBox(0, "OK", "Test", w32.MB_OK|w32.MB_ICONINFORMATION)
}

// 抓取icon畫在notepad應用程式上(如果要執行，請確保您有運行nodepad.exe)
// https://learn.microsoft.com/en-us/windows/win32/menurc/using-icons#creating-an-icon
func ExampleUser32DLL_DrawIcon() {
	user32dll := w32.NewUser32DLL()

	// 獲取HICON{question, chrome}
	var hIconQuestion, hIconChrome w32.HICON
	{
		var errno syscall.Errno

		// 取得系統圖標question
		hIconQuestion, errno = user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION))
		if hIconQuestion == 0 {
			log.Printf("系統圖標: QUESTION 找不到 %s\n", errno)
			return
		}

		// 取得chrome的圖標
		hwndChrome := user32dll.FindWindow("Chrome_WidgetWin_1", "")
		if hwndChrome == 0 {
			log.Println("找不到chrome窗口")
			return
		}

		hIcon, _, _ := user32dll.SendMessage(hwndChrome, w32.WM_GETICON, w32.ICON_SMALL, 0)
		hIconChrome = w32.HICON(hIcon)
		if hIconChrome == 0 {
			log.Println("chrome圖標獲取失敗")

			// 嘗試使用LoadIcon函數取得
			hIconChrome, _ = user32dll.LoadIcon(w32.HINSTANCE(hwndChrome), w32.MakeIntResource(w32.IDI_APPLICATION))
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
			if errno := user32dll.DrawIcon(curHDC, d.x, d.y, d.hIcon); errno != 0 {
				log.Fatalf("%s", errno)
			}
		}
	}
	// Output:
}

func ExampleUser32DLL_DrawIconEx() {
	user32dll := w32.NewUser32DLL()
	hIcon, _ := user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION))

	// 準備一個作圖用的HDC, 我會建議畫在notepad上，可以方便查看
	var hdcScreen w32.HDC
	{
		hwndNotepad := user32dll.FindWindowEx(0, 0, "Notepad", "")
		if hwndNotepad == 0 { // 如果您當前的應用程式會刷新，那可能看不到畫的結果，因為馬上就會被更新掉
			hdcScreen = user32dll.GetDC(0)
			defer user32dll.ReleaseDC(0, hdcScreen)
		} else {
			hdcScreen = user32dll.GetDC(hwndNotepad)
			defer user32dll.ReleaseDC(hwndNotepad, hdcScreen)
		}
	}
	_ = user32dll.DrawIconEx(hdcScreen, 10, 20, hIcon, 0, 0, 0, 0, w32.DI_NORMAL)

	var xLeft int32 = 40
	for i, d := range []struct {
		width, height int32
		diFlag        uint32
	}{
		{0, 0, w32.DI_DEFAULTSIZE},                    // w, h用SM_CXICON, SM_CYICON取代
		{w32.SM_CXICON, w32.SM_CYICON, w32.DI_NORMAL}, // w=11, h=12
		{0, 0, w32.DI_NORMAL},                         // w=0, h=0, DI_DEFAULTSIZE沒有設定 => 原始資源大小
		{64, 128, w32.DI_NORMAL},                      // 自定義大小 w=64, h=128

		// 以下大小都是原尺寸
		{0, 0, w32.DI_IMAGE},  // 整張圖
		{0, 0, w32.DI_MASK},   // Mask的區塊
		{0, 0, w32.DI_NORMAL}, // 他會用IMAGE和MASK做運算，結果的圖形會只有Mask的部分會呈現出來
	} {
		var yTop = 100 * (int32(i) + 1)
		_ = user32dll.DrawIconEx(hdcScreen, xLeft, yTop, hIcon, d.width, d.height, 0, 0, d.diFlag)
	}
	// Output:
}

// 本範例流程
// 1. HICON: 取得hIcon
// 2. ICONINFO: 由該hIcon透過GetIconInfo取得到ICONINFO的結構資料
// 3. init BITMAP: 建立空的BITMAP對象(初始化參照ICONINFO.HbmColor)
// 4. copy to BITMAP: 透過CopyImage將ICONINFO.HbmColor複製到我們所建立的BITMAP之中
// 5. (可選) 把圖標存成檔案(範例執行完之後會刪除)
func ExampleUser32DLL_GetIconInfo() {
	user32dll := w32.NewUser32DLL()
	gdi32dll := w32.NewGdi32DLL()

	hIconQuestion, errno := user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION))
	if hIconQuestion == 0 {
		log.Fatalf("%s", errno)
	}

	var iInfo w32.ICONINFO
	if user32dll.GetIconInfo(hIconQuestion, &iInfo) != 0 {
		return
	}
	// Remember to release when you are not using the HBITMAP.
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

	bmp := w32.Bitmap{}
	{
		modifyImg := true

		if !modifyImg {
			// 可以直接透過以下的方式就可以得到圖像，但我想要測試CopyImage，所以強制跑else的選項
			// 以ICONINFO的資料建立BITMAP
			if gdi32dll.GetObject(w32.HANDLE(iInfo.HbmColor), int32(unsafe.Sizeof(bmp)), uintptr(unsafe.Pointer(&bmp))) == 0 {
				return
			}
		} else {
			hwndCopy, errno := user32dll.CopyImage(w32.HANDLE(iInfo.HbmColor), w32.IMAGE_BITMAP, 10, 20, w32.LR_DEFAULTCOLOR)
			if hwndCopy == 0 {
				log.Printf("%s\n", errno)
				return
			}
			defer func() {
				if !gdi32dll.DeleteObject(w32.HGDIOBJ(hwndCopy)) {
					fmt.Println("error")
				}
			}()

			// 以該HWND的資料建立BITMAP
			if gdi32dll.GetObject(hwndCopy, int32(unsafe.Sizeof(bmp)), uintptr(unsafe.Pointer(&bmp))) == 0 {
				return
			}
		}
	}

	// Save Bitmap to a file.
	var (
		bitmapFileHeader w32.BitmapFileHeader // https://en.wikipedia.org/wiki/BMP_file_format#Bitmap_file_header
		bitmapInfoHeader w32.BitmapInfoHeader // https://en.wikipedia.org/wiki/BMP_file_format#DIB_header_(bitmap_information_header)
	)
	{
		bitmapInfoHeader = w32.BitmapInfoHeader{
			Size:  uint32(unsafe.Sizeof(bitmapInfoHeader)), // 40
			Width: bmp.Width, Height: bmp.Height,
			Planes:      1,
			BitCount:    32,
			Compression: w32.BI_RGB,
		}
		bmpSize := ((bmp.Width*int32(bitmapInfoHeader.BitCount) + 31) / 32) * 4 /* uint32 */ * bmp.Height

		sizeofDIB := 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)) + uint32(bmpSize)
		bitmapFileHeader = w32.BitmapFileHeader{
			Type:       0x4D42,    // BM. // B: 42, M: 4D  //  All the integer values are stored in little-endian format
			Size:       sizeofDIB, // HEADER + INFO + DATA
			OffsetBits: 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)),
		}

		hdc := user32dll.GetDC(0)

		kernel32dll := w32.NewKernel32DLL()

		var lpBitmap w32.LPVOID
		hDIB, _ := kernel32dll.GlobalAlloc(w32.GHND, w32.SIZE_T(bmpSize))
		lpBitmap, _ = kernel32dll.GlobalLock(hDIB)
		defer kernel32dll.GlobalFree(hDIB)

		gdi32dll.GetDIBits(
			hdc, iInfo.HbmColor,
			0,
			uint32(bmp.Height),
			lpBitmap, // [out]
			&w32.BitmapInfo{Header: bitmapInfoHeader},
			w32.DIB_RGB_COLORS,
		)
		_, _ = kernel32dll.GlobalUnlock(hDIB)
		outputBmpPath := "testdata/info.bmp"
		// Write: FileHeader, DIBHeader, bitmapData
		{
			f, err := os.Create(outputBmpPath)
			if err != nil {
				log.Fatal(err)
			}

			defer func() {
				if err = os.Remove(outputBmpPath); err != nil {
					log.Printf("could not remove the test data: %s\n", outputBmpPath)
				}
			}()

			// FileHeader
			_ = binary.Write(f, binary.LittleEndian, bitmapFileHeader)

			// DIB Header
			_ = binary.Write(f, binary.LittleEndian, bitmapInfoHeader)

			// bitmapData
			bmpDatas := make([]byte, bmpSize)
			for offset := uint32(0); offset < uint32(bmpSize); offset += 1 {
				curByteAddr := unsafe.Pointer(uintptr(lpBitmap) + uintptr(offset))
				bmpDatas[offset] = *(*byte)(curByteAddr)
			}
			_ = binary.Write(f, binary.LittleEndian, bmpDatas)

			_ = f.Close()
		}
	}
	// Output:
	// ok
}

func ExampleUser32DLL_PostMessage() {
	user32dll := w32.NewUser32DLL(
		w32.PNPostMessage,
	)
	if errno := user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_FONTCHANGE, 0, 0); errno != 0 {
		panic(fmt.Sprintf("%s", errno))
	}
}

func ExampleUser32DLL_PostThreadMessage() {
	threadID := kernelDll.GetCurrentThreadId()

	// 發送自定義消息至當前的隊列中
	messageID := w32.WM_USER + 1

	eno := userDll.PostThreadMessage(threadID, uint32(messageID), 0, 0)
	if eno != 0 {
		log.Println(eno)
	}
	// Output:
}

func ExampleUser32DLL_FindWindow() {
	user32dll := w32.NewUser32DLL(
		w32.PNFindWindow,
	)

	for _, d := range []struct {
		className  string
		windowName string
	}{
		// "Chrome_WidgetWin_1" You can find this information from Spy++ tool
		{"Chrome_WidgetWin_1", ""}, // 當省略windowName會以找到的第一筆當作依據
		{"", "設定"},                 // className也可以省略, 有多筆時以最先找到的優先
		{"Progman", "Program Manager"},
	} {
		hwnd := user32dll.FindWindow(d.className, d.windowName)
		log.Println(hwnd)
	}

	// Output:
}

// 使用FindWindowEx找到視窗的HWND接著發送WM_SETTEXT來改變標題名稱
// 接著查找子窗口Edit獲得當前正在編輯的文本內容
func ExampleUser32DLL_FindWindowEx() {
	hwnd := userDll.FindWindowEx(0, 0, "Notepad", "")
	// userDll.FindWindow("Notepad", "") // 同上
	if hwnd == 0 {
		return // 窗口不存在
	}

	// 修改標題
	_, _, eno := userDll.SendMessage(hwnd, w32.WM_SETTEXT, 0, uintptr(unsafe.Pointer(&utf16.Encode([]rune("hello world!" + "\x00"))[0])))
	if eno != 0 {
		log.Println(eno)
	}

	// 獲取記事本正在編輯的文本內容
	{
		// 找到記事本的視窗句柄
		// userDll.FindWindow("Notepad", "") // 使用這個只能得到記事本的窗口標題
		hwnd = userDll.FindWindowEx(hwnd, 0, "Edit", "") // 記事本中還有一個Edit這個子窗口，他可以響應WM_GETTEXT與WM_GETTEXTLENGTH這兩個消息 https://stackoverflow.com/a/63494725/9935654

		textLength, _, _ := userDll.SendMessage(hwnd, w32.WM_GETTEXTLENGTH, 0, 0) // 這的每一個元素是uint16
		textLength += 1

		// 向記事本發送WM_GETTEXT消息，獲取內容
		buf := make([]uint16, textLength)
		if _, _, eno = userDll.SendMessage(hwnd, w32.WM_GETTEXT, textLength, uintptr(unsafe.Pointer(&buf[0]))); eno != 0 {
			log.Fatal(eno)
		}

		log.Println("Text:", syscall.UTF16ToString(buf))
	}
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

// 這個範例會做以下三件事情:
// 1. 關閉螢幕 (之後等待兩秒
// 2. 打開螢幕 (在等待兩秒)
// 3. 電池已進入耗電模式 // 這個也會把螢幕關起來
func ExampleUser32DLL_PostMessage_SC_MONITORPOWER() {
	user32dll := w32.NewUser32DLL(w32.PNPostMessage)

	// -1: ^uintptr(0)
	// -2: ^uintptr(1)
	// -3: ^uintptr(2)
	// ...

	log.Println("screenOff")
	_ = user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_SYSCOMMAND, w32.SC_MONITORPOWER, 2)
	time.Sleep(2 * time.Second)

	log.Println("screenOn")
	_ = user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_SYSCOMMAND, w32.SC_MONITORPOWER, ^uintptr(0)) // -1
	time.Sleep(2 * time.Second)

	log.Println("screenLowPower")
	_ = user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_SYSCOMMAND, w32.SC_MONITORPOWER, 1)
}

func ExampleUser32DLL_GetMenu() {
	hwnd := getTestHwnd()
	menu := userDll.GetMenu(hwnd)
	if menu == 0 {
		log.Println("menu not found.")
		return
	}

	_ = userDll.AppendMenu(menu, w32.MF_STRING, 1000, "Hello world")
	menuFile := userDll.GetSubMenu(menu, 0)
	var menuItemInfo w32.MENUITEMINFO
	_ = userDll.InsertMenuItem(menuFile, uint32(userDll.MustGetMenuItemCount(menuFile)-1), 1, &w32.MENUITEMINFO{
		CbSize:     uint32(unsafe.Sizeof(menuItemInfo)),
		FMask:      w32.MIIM_ID | w32.MIIM_STRING,
		WID:        500,
		DwTypeData: &(utf16.Encode([]rune("my Plugin" + "\x00")))[0],
	})

	// Add separator
	_ = userDll.InsertMenu(menuFile, uint32(userDll.MustGetMenuItemCount(menuFile)-1), w32.MF_BYPOSITION|w32.MF_SEPARATOR, nil, "")

	// TODO: 如何偵測點擊自訂清單的動作
	{

	}

	// Output:
}

// 異動標題欄位按下右鍵，所彈出來的選單{恢復、移動、最大化、最小化、關閉} => {Hello world}
func ExampleUser32DLL_GetSystemMenu() {
	hwnd := getTestHwnd()
	menu := userDll.GetSystemMenu(hwnd, false)
	if menu == 0 {
		return
	}
	_ = userDll.AppendMenu(menu, w32.MF_STRING, 1000, "Hello world")

	// Delete all the items except the last.
	/*
		for userDll.MustGetMenuItemCount(menu) > 1 {
			if eno := userDll.DeleteMenu(menu, 0, w32.MF_BYPOSITION); eno != 0 {
				log.Println(eno)
			}
		}
	*/
	if eno := userDll.DeleteMenu(menu, 1000, w32.MF_BYCOMMAND); eno != 0 {
		log.Println(eno)
	}

	// Optional 本範例就算不呼叫DrawMenuBar依然可以正常工作
	if eno := userDll.DrawMenuBar(hwnd); eno != 0 {
		log.Println(eno)
	}

	// Output:
}

func ExampleUser32DLL_LoadImage() {
	user32dll := w32.NewUser32DLL(w32.PNLoadImage)
	hicon, errno := user32dll.LoadImage( // returns a HANDLE so we have to cast to HICON
		0,                         // hInstance must be NULL when loading from a file
		"testdata/img/golang.ico", // the icon file name
		w32.IMAGE_ICON,            // specifies that the file is an icon
		0,                         // width of the image (we'll specify default later on)
		0,                         // height of the image

		w32.LR_LOADFROMFILE| // we want to load a file (as opposed to a resource)
			w32.LR_DEFAULTSIZE| // default metrics uses the SM_CXICON or SM_CXCURSOR, SM_CYICON or SM_CYCURSOR
			w32.LR_SHARED, // let the system release the handle when it's no longer used
	)
	if hicon == 0 {
		fmt.Printf("%s", errno)
	}

	hicon, errno = user32dll.LoadImage( // returns a HANDLE so we have to cast to HICON
		0,                         // hInstance must be NULL when loading from a file
		"testdata/img/golang.ico", // the icon file name
		w32.IMAGE_ICON,            // specifies that the file is an icon

		5,  // 如果尺寸超過原始圖片大小，會無效
		10, // 如果尺寸超過原始圖片大小，會無效

		w32.LR_LOADFROMFILE| // we want to load a file (as opposed to a resource)
			w32.LR_SHARED, // let the system release the handle when it's no longer used
	)

	if hicon == 0 {
		fmt.Printf("%s", errno)
	}

	// Output:
}

func ExampleUser32DLL_ClientToScreen() {
	hwnd := userDll.FindWindow("Notepad", "")
	if hwnd == 0 {
		return
	}
	var pt w32.POINT
	pt.X = 100
	pt.Y = 300
	if userDll.ClientToScreen(hwnd, &pt) {
		log.Printf("Screen coordinates: (%d, %d)\n", pt.X, pt.Y)
	} else {
		log.Println("ClientToScreen failed.")
	}
	// Output:
}

// https://stackoverflow.com/a/68845977/9935654
// https://learn.microsoft.com/en-us/windows/win32/menurc/using-menus
func ExampleUser32DLL_CreatePopupMenu() {
	user32dll := w32.NewUser32DLL()
	gdi32dll := w32.NewGdi32DLL(w32.PNDeleteObject)
	kernel32dll := w32.NewKernel32DLL(w32.PNGetModuleHandle)

	var (
		hIcon w32.HICON
		iInfo w32.ICONINFO
	)
	{
		hIcon = w32.HICON(user32dll.MustLoadImage(
			0,
			"testdata/img/golang.ico",
			w32.IMAGE_ICON,
			0, 0,
			w32.LR_LOADFROMFILE|w32.LR_DEFAULTSIZE|w32.LR_SHARED,
		))

		if user32dll.GetIconInfo(hIcon, &iInfo) != 0 {
			return
		}

		// Remember to release when you are not using the HBITMAP.
		defer func() {
			_ = gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmColor))
			_ = gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmMask))
		}()
	}

	// Create Menu (PopupMenu)
	var hMenu w32.HMENU
	{
		hMenu = user32dll.CreatePopupMenu()
		_ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1023, "Open")
		_ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1024, "TODO")

		// 設定含有icon的Menu
		// _, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1024, "Hello") // 可以先指定string，再用SetMenuItemInfo添加icon或者直接在SetMenuItemInfo添加string或icon都可以
		var menuItemInfo w32.MENUITEMINFO
		menuItemInfo = w32.MENUITEMINFO{
			CbSize: uint32(unsafe.Sizeof(menuItemInfo)),

			// FMask是一個開關，當有設定這些mask，某些欄位設定數值才會有意義
			FMask: w32.MIIM_BITMAP | // sets the hbmpItem member.
				w32.MIIM_ID | // sets the wID member.
				w32.MIIM_STRING, // sets the dwTypeData member.
			// w32.MIIM_STATE, // set the FState member.
			// FState:     w32.MFS_CHECKED, // 圖片可能會檔住勾選的圖示，所以在使用圖示之後，不建議在使用此項目
			WID:        1024,
			DwTypeData: &(utf16.Encode([]rune("Hello" + "\x00")))[0], // 修改原有的名稱
			HbmpItem:   iInfo.HbmColor,
		}
		_ = user32dll.SetMenuItemInfo(hMenu, 1024, false, &menuItemInfo) // 將menuItemInfo放到此ID項目去
		// _ = user32dll.SetMenuItemInfo(hMenu, 1 /*放在第二個項目*/, true, &menuItemInfo) // 同上

		_ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1025, "Exit program")
		_ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1026, "Advanced")

		// SetMenuItemInfo如果該pos已經有其他項目存在，會直接覆蓋
		// _ = user32dll.SetMenuItemInfo(hMenu, 2, true, &menuItemInfo) // 會覆蓋掉Exit program

		// SEPARATOR
		{
			// 在後面一個項目插入分隔線
			_ = user32dll.AppendMenu(hMenu, w32.MF_STRING,
				0,  // 當建造的是separate line，此時的uID沒有意義(隨便給都沒差，最後會被當成0)，也不能被SetMenuItemInfo所更改
				"", // 傳空字串會建立分隔線 separate line
			)

			// 將第二個項目插入分隔線(也就是原本的第二個項目會變成第三個)
			_ = userDll.InsertMenu(hMenu, 1 /*下標值從0開始*/, w32.MF_BYPOSITION|w32.MF_SEPARATOR, nil, "")

			// 在最後一個項目插入分隔線
			countMenuItems, _ := userDll.GetMenuItemCount(hMenu)
			_ = userDll.InsertMenu(hMenu, uint32(countMenuItems), w32.MF_BYPOSITION|w32.MF_SEPARATOR, nil, "")
		}

		// 用AppendMenu與SetMenuItemInfo插配來產生新增項目的效果
		{
			menuItemInfo2 := menuItemInfo // copy
			title, _ := syscall.UTF16PtrFromString("insrtMenuTest")
			menuItemInfo2.DwTypeData = title
			menuItemInfo2.WID = 1027
			_ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1027, "TODO") // 因為SetMenuItemInfo不能創造新的item，所以要在建立一個
			_ = user32dll.SetMenuItemInfo(hMenu, 1027, false, &menuItemInfo2)

			// 直接用InsertMenuItem來插入新的項目會比較快
			countMenuItems, _ := userDll.GetMenuItemCount(hMenu)
			_ = userDll.InsertMenuItem(hMenu, uint32(countMenuItems), 1, &w32.MENUITEMINFO{
				CbSize: uint32(unsafe.Sizeof(menuItemInfo)),
				FMask: w32.MIIM_ID | w32.MIIM_STRING |
					w32.MIIM_STATE,
				FState:     w32.MFS_CHECKED | w32.MFS_GRAYED,
				WID:        1028,
				DwTypeData: &(utf16.Encode([]rune("test State" + "\x00")))[0],
			})
		}

		// Submenu
		{
			subMenu := userDll.CreateMenu()
			defer user32dll.DestroyMenu(subMenu)
			_ = user32dll.AppendMenu(subMenu, w32.MF_STRING, 10001, "submenu item1")
			_ = user32dll.AppendMenu(subMenu, w32.MF_STRING, 10002, "submenu item2")

			pMsg, _ := syscall.UTF16PtrFromString("sumMenu")
			_ = userDll.InsertMenuItem(hMenu, 1028, 0, &w32.MENUITEMINFO{
				CbSize:     uint32(unsafe.Sizeof(menuItemInfo)),
				FMask:      w32.MIIM_ID | w32.MIIM_STRING | w32.MIIM_SUBMENU,
				WID:        10000,
				HSubMenu:   subMenu,
				DwTypeData: pMsg,
			})
			fmt.Println("GetMenuItemCount:", userDll.MustGetMenuItemCount(hMenu)) // 10 // subMenu的指項目不會納入計算

			// DeleteMenu
			{
				_ = user32dll.AppendMenu(subMenu, w32.MF_STRING, 10003, "submenu item3")
				_ = user32dll.AppendMenu(subMenu, w32.MF_STRING, 10004, "submenu item4")

				if eno := userDll.DeleteMenu(subMenu, 2, w32.MF_BYPOSITION); eno != 0 {
					fmt.Println(eno)
				}

				if eno := userDll.DeleteMenu(subMenu, 10004, w32.MF_BYCOMMAND); eno != 0 {
					fmt.Println(eno)
				}
			}

			// GetMenuItemID Test
			{
				fmt.Println("Menu ID 0:", userDll.GetMenuItemID(hMenu, 0)) // 1023
				fmt.Println("Menu ID 1:", userDll.GetMenuItemID(hMenu, 1)) // 分隔線的項目回傳的id都是0

				fmt.Println(userDll.GetMenuItemID(hMenu, userDll.MustGetMenuItemCount(hMenu)) == 0xffffffff) // true，因為是zero-based，所以這個項目一定不存在

				fmt.Println("Menu ID second last == 0xffffffff:", userDll.GetMenuItemID(hMenu, userDll.MustGetMenuItemCount(hMenu)-2) == 0xffffffff) // 如果該項目是subMenu，那麼會無法取得到該項目

				fmt.Println("Menu ID last:", userDll.GetMenuItemID(hMenu, userDll.MustGetMenuItemCount(hMenu)-1)) // 1028 // 我們的subMenu是插入在1028的位子，即取代原本1028的位子，原1028會往後移動，所以最後一個還是1028

				mySubMenu := userDll.GetSubMenu(hMenu, userDll.MustGetMenuItemCount(hMenu)-2)
				fmt.Println("submenu ID 0:", userDll.GetMenuItemID(mySubMenu, 0)) // 10001
			}
		}

		// 預設選項設定
		{
			// _ = user32dll.SetMenuDefaultItem(hMenu, 1026, false) // 表示ID:1026為預設選項
			// _ = user32dll.SetMenuDefaultItem(hMenu, 2, true) // 使用第3個選項來當作預設選項(起始index從0開始)
			// _ = user32dll.SetMenuDefaultItem(hMenu, 0xffffffff, true) // no default item
			_ = user32dll.SetMenuDefaultItem(hMenu, 0, true) // 設定預設選項為第一個(會加粗顯示)
		}

		defer func() {
			if errno := user32dll.DestroyMenu(hMenu); errno != 0 {
				log.Printf("%s\n", errno)
			}
		}()
	}

	ch := make(chan w32.HWND)
	go func(chanWin chan<- w32.HWND) {
		// define WndProc
		wndProcFuncPtr := syscall.NewCallback(w32.WndProc(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			switch uMsg {
			case w32.WM_DESTROY:
				log.Println("WM_DESTROY")
				user32dll.PostQuitMessage(0)
				return 0
			case w32.WM_RBUTTONDOWN:
				log.Println("WM_RBUTTONDOWN")
				// Show the menu
				// hwnd = user32dll.GetForegroundWindow() // 如果當前所在的窗口非自己所建，在TrackPopupMenu可能會遇到The parameter is incorrect.的問題
				user32dll.SetForegroundWindow(hwnd)
				var pos w32.POINT
				if errno := user32dll.GetCursorPos(&pos); errno != 0 {
					fmt.Printf("GetCursorPos %s", errno)
				}
				if wParam != 123 {
					if result, errno := user32dll.TrackPopupMenu(hMenu, w32.TPM_LEFTALIGN, pos.X, pos.Y, 0, hwnd, nil); result == 0 {
						// 如果出現The parameter is incorrect. 問題可能在於hwnd本身，如果該hwnd是您所建立的視窗就一定沒有問題，但若不是就可能會導致該問題發生，建議GetForegroundWindow要慎用
						log.Printf("Error TrackPopupMenu %s\n", errno)
					}
				} else {
					cmd, _ := user32dll.TrackPopupMenu(hMenu, w32.TPM_LEFTALIGN|w32.TPM_RETURNCMD, pos.X, pos.Y, 0, hwnd, nil)
					if cmd > 0 {
						log.Println("TrackPopupMenu with TPM_RETURNCMD")
						_, _, _ = user32dll.SendMessage(hwnd, w32.WM_COMMAND, uintptr(cmd), 0)
					}
				}
			case w32.WM_COMMAND:
				id := w32.LOWORD(wParam)
				switch id {
				case 500:
					log.Println("sys open")
				case 1023:
					log.Println("open")
				case 1024:
					log.Println("hello")
				case 1025:
					log.Println("1025")
					_ = user32dll.PostMessage(hwnd, w32.WM_DESTROY, 0, 0)
				case 10001:
					log.Println("submenu-item1")
				case 10002:
					log.Println("submenu item2")
				}
			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam)
		}))

		const (
			wndClassName  = "classCreatePopupMenu"
			wndWindowName = "windowCreatePopupMenu"
		)

		// Register
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(wndClassName)
		hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))

		if atom, errno := user32dll.RegisterClass(&w32.WNDCLASS{
			Style:         w32.CS_HREDRAW | w32.CS_HREDRAW,
			HbrBackground: w32.COLOR_WINDOW,
			WndProc:       wndProcFuncPtr,
			HInstance:     hInstance,
			HIcon:         hIcon,
			ClassName:     pUTF16ClassName,
		}); atom == 0 {
			fmt.Printf("%s", errno)
			chanWin <- 0
			return
		}

		// Create window
		hwnd, errno := user32dll.CreateWindowEx(0,
			wndClassName,
			wndWindowName,
			w32.WS_OVERLAPPEDWINDOW,

			// Size and position
			w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,

			0, // Parent window
			0, // Menu
			hInstance,
			0, // Additional application data
		)

		if hwnd == 0 {
			fmt.Printf("%s\n", errno)
			if errno2 := user32dll.UnregisterClass(wndClassName, hInstance); errno2 != 0 {
				fmt.Printf("Error UnregisterClass: %s", errno2)
			}
			chanWin <- hwnd
			return
		}

		// Make sure it can be unregistered when exiting.
		defer func() {
			if errno2 := user32dll.UnregisterClass(wndClassName, hInstance); errno2 != 0 {
				log.Printf("Error UnregisterClass: %s", errno2)
			} else {
				log.Println("OK UnregisterClass")
			}

			close(chanWin)
		}()

		chanWin <- hwnd

		var msg w32.MSG
		for {
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
				break
			}
			user32dll.TranslateMessage(&msg)
			user32dll.DispatchMessage(&msg)
		}
	}(ch)

	hwnd := <-ch
	user32dll.ShowWindow(hwnd, w32.SW_SHOW) // 如果沒有顯示，對於不使用TPM_RETURNCMD的選單，不會觸發WM_COMMAND，也就是雖然選單會出來，但選中的項目沒有任何意義，但對有設計TPM_RETURNCMD則不影響，選中的行為仍有效

	_ = user32dll.PostMessage(hwnd, w32.WM_RBUTTONDOWN, 0, 0)   // 選單測試
	_ = user32dll.PostMessage(hwnd, w32.WM_RBUTTONDOWN, 123, 0) // with TPM_RETURNCMD

	// 🕹️ 如果您要手動嘗試，請把以下的SendMessage.WM_CLOSE註解掉，避免自動關閉
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)

	// wait user close the window
	<-ch

	// Output:
	// GetMenuItemCount: 10
	// Menu ID 0: 1023
	// Menu ID 1: 0
	// true
	// Menu ID second last == 0xffffffff: true
	// Menu ID last: 1028
	// submenu ID 0: 10001
}

// https://learn.microsoft.com/en-us/windows/win32/learnwin32/creating-a-window
// https://learn.microsoft.com/en-us/windows/win32/winmsg/using-messages-and-message-queues#creating-a-message-loop
func ExampleUser32DLL_CreateWindowEx() {
	user32dll := w32.NewUser32DLL()
	kernel32dll := w32.NewKernel32DLL()

	hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
	// hwndParent := user32dll.FindWindow("Notepad", "") // 可以把窗口附加到其他的應用程式之下
	hwndParent := w32.HWND(0)

	// 創建測試用結構，非必要

	// 此結構主要用來測試SetWindowLongPtr, GetWindowLongPtr所設計，可以自由設計
	type Response struct {
		// Msg    string // 不建議對SetWindowLongPtr的資料用string，最好固定大小，不能可能會遇到memory的錯誤
		Msg    [256]byte
		MsgLen uint16
		Status uint32
	}

	// 用來測試CREATESTRUCT
	type AppData struct {
		title string // 這個沒有像GetWindowLongPtr遇到memory溢位的錯誤，不過如果要透過unsafe去轉換，最好都限定長度，不然轉換雖然可以過，但是在調用時，長度可能會抓的很大而導致出錯。
		id    uint32
	}

	// 通知外層主程式用
	ch := make(chan w32.HWND)
	wg := sync.WaitGroup{}
	// 新建一個執行緒來專門處理視窗{建立、消息循環}
	go func(className, windowName string, channel chan<- w32.HWND, wg *sync.WaitGroup) {
		wg.Add(1)

		// define ProcFunc // https://learn.microsoft.com/en-us/windows/win32/learnwin32/writing-the-window-procedure
		wndProcFuncPtr := syscall.NewCallback(w32.WndProc(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			// log.Printf("uMsg:%d\n", uMsg)
			switch uMsg {
			case w32.WM_GETMINMAXINFO: // 首次使用CreateWindowEx會先觸發此msg // https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-getminmaxinfo
				// wParam not used
				log.Println("WM_GETMINMAXINFO")
				minmaxInfo := *(*w32.MINMAXINFO)(unsafe.Pointer(lParam))
				log.Printf("%#v", minmaxInfo)
			case w32.WM_NCCREATE: // 首次建立視窗會觸發此MSG
				log.Println("WM_NCCREATE")
				// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-nccreate
				fmt.Println("WM_NCCREATE")
				// 對於WM_NCCREATE的回傳值: 「true會繼續創建; False(0)將會導致CreateWindowEx得到的hwnd為0」。可以倚靠DefWindowProc來自動幫我們計算回傳值
			case w32.WM_CREATE: // 觸發完WM_NCCREATE會再跑WM_CREATE
				log.Println("WM_CREATE")
				pCreate := *((*w32.CREATESTRUCT)(unsafe.Pointer(lParam))) // 注意您呼叫的函數如果是用W就對應CREATESTRUCTW 用A則對應CREATESTRUCTA
				apData := *((*AppData)(unsafe.Pointer(pCreate.LpCreateParams)))
				fmt.Println(apData.title)
				fmt.Println(apData.id)
				msg := "Msg from WM_CREATE"
				response := &Response{Status: 200, MsgLen: uint16(uintptr(len(msg)))}
				copy(response.Msg[:], msg)
				// 不建議用SetWindowLongPtr，有時候放進去的內容會不如預期，推測可能與go回收機制有關
				_, _ = user32dll.SetWindowLongPtr(hwnd, w32.GWLP_USERDATA, uintptr(unsafe.Pointer(response)))
			case w32.WM_CLOSE: // Pressed Close Button (X) / Alt+F4 / "Close" in context menu // 在這之後它會調用WM_DESTROY
				log.Println("WM_CLOSE")
				// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-close
				if errno := user32dll.DestroyWindow(hwnd); errno != 0 {
					fmt.Printf("[DestroyWindow] %s\n", errno)
				}
			case w32.WM_DESTROY:
				log.Println("WM_DESTROY")
				user32dll.PostQuitMessage(0)
				return 0 // 要有返回不能再靠DefWindowProc，不然GetMessage不會結束
			case w32.WM_NCDESTROY: // WM_QUIT會觸發此MSG
				log.Println("WM_NCDESTROY")
				return 0
			case w32.WM_SHOWWINDOW:
				log.Println("WM_SHOWWINDOW")
			case w32.WM_MOVE:
			case w32.WM_ACTIVATE:
				log.Println("WM_NCDESTROY")
				// https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-activate
			case w32.WM_SIZE:
				// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-size
			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam) // default window proc
		}))

		pUTF16ClassName, _ := syscall.UTF16PtrFromString(className)
		wc := w32.WNDCLASS{
			// Style:       0, // 可以不給，或者w32.CS_NOCLOSE禁用右上角的關閉按鈕) // CS_指的是class的style
			WndProc: wndProcFuncPtr, // 每次有消息，就會送通知到此函數
			// CbClsExtra:    0,
			// CbWndExtra:    0,
			HInstance: hInstance,
			HIcon:     user32dll.MustLoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION /* w32.IDI_WINLOGO */)), // 可以不給, 用預設0 會是: IDI_WINLOGO
			HCursor:   user32dll.MustLoadCursor(0, w32.MakeIntResource(w32.IDC_CROSS /* w32.IDC_ARROW */)),    // 可以不給, 用預設0 會是: IDC_ARROW
			// HbrBackground: 0,
			// LpszMenuName:  nil,
			ClassName: pUTF16ClassName,
		}

		// 確保沒有殘留的資料
		if errno := user32dll.UnregisterClass(className, hInstance); errno != 0 {
			log.Println("clear previous RegisterClass")
		} else {
			log.Printf("%s", errno) // Class does not exist.
		}

		if atom, errno := user32dll.RegisterClass(&wc); atom == 0 {
			log.Printf("%s", errno)
			close(channel)
			wg.Done()
			return
		}

		var (
			hwnd  w32.HWND
			errno syscall.Errno
		)

		defer func() {
			if errno = user32dll.UnregisterClass(className, hInstance); errno != 0 {
				log.Printf("[UnregisterClass] %s\n", errno)
			}
			close(channel)
		}()

		fmt.Println("CreateWindowEx")
		if hwnd, errno = user32dll.CreateWindowEx(0, /* | w32.WS_EX_TOOLWINDOW 如果不想要讓應用程式出現在任務欄可以考慮加上此屬性 */
			className,
			windowName,              // 視窗左上角的標題名稱
			w32.WS_OVERLAPPEDWINDOW, // 這項包含了: WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX // 如果不想要最小和最大化按鈕要在這邊調整，而關閉按鈕則是需要透過class註冊的時候就設定要還是不要

			// Size and position
			w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,

			hwndParent, // 0, // Parent window
			0,          // Menu
			hInstance,
			uintptr(unsafe.Pointer(&AppData{"Demo-CreateWindowEx", 6})), // Additional application data // 可以不給(設定為0). 如果有給，這個資料會在WM_CREATE的時候傳入給lParam
		); hwnd == 0 {
			fmt.Printf("%s", errno)
			return
		} else {
			// test FindWindow
			{
				hwnd2 := user32dll.FindWindow(className, windowName)
				log.Println("hwnd == hwnd2 ", hwnd == hwnd2) // true
			}

			channel <- hwnd
		}

		// 消息循環
		go func() {
			var msg w32.MSG
			for {
				if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
					log.Println("===quit the window===")
					wg.Done()
					break
				}
				user32dll.TranslateMessage(&msg)
				user32dll.DispatchMessage(&msg)
			}
		}()
		wg.Wait()
	}("class_ExampleUser32DLL_CreateWindowEx", "ExampleUser32DLL_CreateWindowEx", ch, &wg)

	// 如果視窗成功被建立會傳送該hwnd
	hwnd, isOpen := <-ch

	// 如果視窗建立失敗，會得到空hwnd，不做其他處理直接返回
	if !isOpen || hwnd == 0 {
		return
	}

	// 不確定什麼原因，在整個專案進行go test的時候，會卡死無法結束，所以在用一個routine包起來
	go func() {
		// 以下為模擬外層程式，向視窗發送訊息
		{
			fmt.Println("ShowWindow")
			if false {
				_, _ = user32dll.SetWindowLongPtr(hwnd, w32.GWL_EXSTYLE, uintptr(user32dll.GetWindowLong(hwnd, w32.GWL_EXSTYLE)|w32.WS_EX_TOOLWINDOW)) // 可以之後透過此方法來修改WS_EX屬性
				_, _ = user32dll.SetWindowLongPtr(hwnd, w32.GWL_STYLE, uintptr(user32dll.GetWindowLong(hwnd, w32.GWL_STYLE)&^w32.WS_THICKFRAME))       // 移除WS某屬性
			}

			user32dll.ShowWindow(hwnd, w32.SW_MAXIMIZE)

			fmt.Println("CloseWindow")                            // 僅是縮小視窗
			if errno := user32dll.CloseWindow(hwnd); errno != 0 { // close只是把它縮小並沒有真正關閉
				fmt.Printf("[CloseWindow] %s\n", errno)
			}
		}

		// 測試來自於視窗所寫入的使用者資料
		// 這種用法是取記憶體中的資訊，所以不管哪一個視窗還是程式，只要知道確切的hwnd還有類型(GWLP_USERDATA, ...)，就可以強制轉換來取得資料(前提是該記憶體位置已經有被寫入該資料，也就是一定要有人用SetWindowLongPtr先放資料進去)
		if userDataPtr, _ := user32dll.GetWindowLongPtr(hwnd, w32.GWLP_USERDATA); userDataPtr != 0 {
			res := *((*Response)(unsafe.Pointer(userDataPtr)))
			if uintptr(res.MsgLen) <= unsafe.Sizeof(res.Msg) { // set資料的時候可能會發生問題，導致此長度已經不正確，對於不正確的結果就不顯示
				log.Printf("%s\n", string(res.Msg[:res.MsgLen])) // Msg from WM_CREATE
			}
			log.Println(res.Status) // 200
		}

		fmt.Println("DestroyWindow")
		// user32dll.DestroyWindow(hwnd) // 注意！ DestroyWindow不要在外面呼叫，需要在callback之中運行, 不然可能會得到錯誤: Access is denied.
		_, _, _ = user32dll.SendMessage(hwnd, w32.WM_DESTROY, 0, 0) // 如果您想要在視窗上進行操作，可以把這列註解，運行的時候再去手動關閉視窗即可結束
	}()

	maxTry := 2
	for {
		select {
		case _, isOpen = <-ch:
			if !isOpen {
				log.Print("bye")
				return
			}
		case <-time.After(2 * time.Second):
			log.Println("timeout")
			wg.Done()
			maxTry--
			if maxTry == 0 {
				log.Println("reach maxTry")
				return
			}
		}
	}

	// Output:
	// CreateWindowEx
	// WM_NCCREATE
	// Demo-CreateWindowEx
	// 6
	// ShowWindow
	// CloseWindow
	// DestroyWindow
}

func ExampleUser32DLL_GetWindowThreadProcessId() {
	user32dll := w32.NewUser32DLL(w32.PNGetForegroundWindow, w32.PNGetWindowThreadProcessId)
	hwndTarget := user32dll.GetForegroundWindow()

	var (
		processID uint32
		threadID  uint32
	)
	threadID = user32dll.GetWindowThreadProcessId(hwndTarget, &processID)
	log.Printf("threadID: %d\n", threadID)
	log.Printf("processID: %d\n", processID)
	// Output:
}

func ExampleUser32DLL_SetWindowPos() {
	var r w32.RECT
	r.Right = 1000
	r.Bottom = 600
	hwnd := userDll.FindWindow("Notepad", "")
	if hwnd == 0 {
		hwnd = userDll.GetDesktopWindow()
	}
	if eno := userDll.AdjustWindowRect(&r, w32.WS_OVERLAPPEDWINDOW, false); eno != 0 {
		fmt.Println(eno)
	}
	if eno := userDll.SetWindowPos(hwnd, 0,
		r.Left, r.Top, r.Width(), r.Height(),
		w32.SWP_NOZORDER|
			w32.SWP_NOMOVE, // 忽略x, y即窗口的位置不變，指改變寬度與高度
	); eno != 0 {
		fmt.Println(eno)
	}

	// Output:
}

// https://learn.microsoft.com/en-us/windows/win32/winmsg/using-hooks
// https://social.msdn.microsoft.com/Forums/en-US/d5e1ec20-9ff2-4cd4-918a-02560d473845/winapi-hook-procedure-and-reading-message-details?forum=windowsgeneraldevelopmentissues
func ExampleUser32DLL_SetWindowsHookEx() {
	user32dll := w32.NewUser32DLL()
	kernel32dll := w32.NewKernel32DLL(w32.PNGetModuleHandle, w32.PNGetCurrentThreadId)

	// https://learn.microsoft.com/en-us/windows/win32/winmsg/about-hooks#wh_journalrecord
	var (
		// enableCallWnd   = true // 因為要和其他goroutine溝通，用純數值沒辦法統一，必須採用指標才能到處都能溝通
		enableCallWnd    = new(bool)
		hCallWndHook     w32.HHOOK
		hCallWndHookProc w32.HOOKPROC

		enableCBT        = new(bool)
		hCBTProcHook     w32.HHOOK
		hCBTProcHookProc w32.HOOKPROC

		enableLLMouseHook = new(bool) // LL: Low Level
		hLLMouseHook      w32.HHOOK
		hLLMouseHookProc  w32.HOOKPROC

		enableLLKeyboardHook = new(bool)
		hLLKeyboardHook      w32.HHOOK
		hLLKeyboardHookProc  w32.HOOKPROC
	)
	*enableCallWnd = true
	*enableCBT = true
	*enableLLMouseHook = true
	*enableLLKeyboardHook = true

	// define HookProc
	{
		hCallWndHookProc = func(nCode int32, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			// https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644975(v=vs.85)
			if nCode < 0 || !*enableCallWnd {
				return user32dll.CallNextHookEx(hCallWndHook, nCode, wParam, lParam)
			}
			switch nCode {
			case w32.HC_ACTION:
				//  wParam: If the message was sent by the current thread, it is nonzero; otherwise, it is zero.
				isCurrentThreadSent := wParam != 0
				cwpStruct := *(*w32.CWPSTRUCT)(unsafe.Pointer(lParam))
				log.Printf("CallWndProc - isCurrentThreadSend: %t - cwpStruct: %+v\n", isCurrentThreadSent, cwpStruct)
				// 當isCurrentThreadSent為true，則cwpStruct.Hwnd表示我們所新建的視窗hwnd
			default:
				break
			}
			return user32dll.CallNextHookEx(hCallWndHook, nCode, wParam, lParam)
		}

		hCBTProcHookProc = func(nCode int32, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			// https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644977(v=vs.85)
			if nCode < 0 || !*enableCBT {
				return user32dll.CallNextHookEx(hCBTProcHook, nCode, wParam, lParam)
			}
			switch nCode {
			case w32.HCBT_ACTIVATE:
				log.Println("HCBT_ACTIVATE") // The system is about to activate a window.
			case w32.HCBT_CLICKSKIPPED:
				log.Println("HCBT_CLICKSKIPPED")
			case w32.HCBT_CREATEWND:
				log.Println("HCBT_CREATEWND")
			case w32.HCBT_DESTROYWND:
				log.Println("HCBT_DESTROYWND") // A window is about to be destroyed.
			case w32.HCBT_KEYSKIPPED:
				log.Println("HCBT_KEYSKIPPED")
			case w32.HCBT_MINMAX:
				log.Println("HCBT_MINMAX") // A window is about to be minimized or maximized.
			case w32.HCBT_MOVESIZE:
				log.Println("HCBT_MOVESIZE") // A window is about to be moved or sized.
			case w32.HCBT_QS:
				log.Println("HCBT_QS")
			case w32.HCBT_SETFOCUS:
				log.Println("HCBT_SETFOCUS")
			case w32.HCBT_SYSCOMMAND:
				log.Println("HCBT_SYSCOMMAND")
			default:
			}
			return user32dll.CallNextHookEx(hCBTProcHook, nCode, wParam, lParam)
		}

		hLLMouseHookProc = func(nCode int32, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			// LowLevelMouseProc https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644986(v=vs.85)
			// MouseProc: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644988(v=vs.85)
			if nCode < 0 || !*enableLLMouseHook {
				return user32dll.CallNextHookEx(hLLMouseHook, nCode, wParam, lParam)
			}

			if nCode == w32.HC_ACTION {
				mouseMsgID := wParam // The identifier of the mouse message.
				msLLHookStruct := *(*w32.MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
				var (
					isXButton1Done bool
					isXButton2Done bool
				)

				if mouseMsgID != w32.WM_MOUSEWHEEL {
					isXButton1Done = (uint16((msLLHookStruct.MouseData>>16)&0xffff) & w32.MK_XBUTTON1) == w32.MK_XBUTTON1
					isXButton2Done = (uint16((msLLHookStruct.MouseData>>16)&0xffff) & w32.MK_XBUTTON2) == w32.MK_XBUTTON2
				}

				type MouseMsg struct {
					X              int32
					Y              int32
					isXButton1Done bool // XButton一定都會是false，因為這個訊息只在WM_X或者WM_NCX開頭的訊息才會有紀錄，而這些訊息都沒有涵蓋在LowLevelMouseProc所能接收到的訊息之中，所以必定為false
					isXButton2Done bool
					isInjected     bool
					isILInjected   bool
					time           uint32
				}
				mouseMsg := MouseMsg{
					msLLHookStruct.Pt.X,
					msLLHookStruct.Pt.Y,
					isXButton1Done,
					isXButton2Done,
					(msLLHookStruct.Flags & w32.LLMHF_INJECTED) == w32.LLMHF_INJECTED,
					(msLLHookStruct.Flags & w32.LLMHF_LOWER_IL_INJECTED) == w32.LLMHF_LOWER_IL_INJECTED,
					msLLHookStruct.Time}
				switch mouseMsgID {
				// double click都沒有成功，可以是被單下的click所佔據
				case w32.WM_MBUTTONDBLCLK:
					log.Printf("WM_MBUTTONDBLCLK %+v", mouseMsg)
				case w32.WM_LBUTTONDBLCLK:
					log.Printf("WM_LBUTTONDBLCLK %+v", mouseMsg)
				case w32.WM_LBUTTONDOWN:
					log.Printf("WM_LBUTTONDOWN %+v", mouseMsg)
				case w32.WM_LBUTTONUP:
					log.Printf("WM_LBUTTONUP %+v", mouseMsg)
				case w32.WM_RBUTTONDOWN:
					log.Printf("WM_RBUTTONDOWN %+v", mouseMsg)
				case w32.WM_RBUTTONUP:
					log.Printf("WM_RBUTTONUP %+v", mouseMsg)
				case w32.WM_MOUSEWHEEL:
					// WHEEL_DELTA
					wheelDelta := int16((msLLHookStruct.MouseData >> 16) & 0xffff)
					isForwardScroll := wheelDelta > 0 // 是否是往前滾(遠離自己)、反之就往後(面向自己)
					log.Printf("WM_MOUSEWHEEL: wheelDelta: %d isForwardScroll: %t\n time: %d", wheelDelta, isForwardScroll, msLLHookStruct.Time)
				default:
					// 這幾類訊息也可以收到
					// WM_MOUSEMOVE 512
					// WM_MBUTTONDOWN 519
					// WM_MBUTTONUP 520
					log.Printf("mouseMsgID:%d | %+v", mouseMsgID, mouseMsg)
				}
			}
			return user32dll.CallNextHookEx(hLLMouseHook, nCode, wParam, lParam)
		}

		hLLKeyboardHookProc = func(nCode int32, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			// LowLevelKeyboardProc: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644985(v=vs.85)
			// KeyboardProc: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms644984(v=vs.85)
			// 注意以上兩個類似，但是他們的wParam和lParam是有差異的，所以要注意，不要看錯文檔
			if nCode < 0 || !*enableLLKeyboardHook {
				return user32dll.CallNextHookEx(hLLKeyboardHook, nCode, wParam, lParam)
			}

			kbDllHookStruct := *(*w32.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
			isExtendedKey := (kbDllHookStruct.Flags | w32.LLKHF_EXTENDED) == w32.LLKHF_EXTENDED
			lowerILInjected := (kbDllHookStruct.Flags | w32.LLKHF_LOWER_IL_INJECTED) == w32.LLKHF_LOWER_IL_INJECTED // 這個如果是1，那麼injected(LLKHF_INJECTED)也會是1
			injected := (kbDllHookStruct.Flags | w32.LLKHF_INJECTED) == w32.LLKHF_INJECTED
			isAltDone := (kbDllHookStruct.Flags | w32.LLKHF_ALTDOWN) == w32.LLKHF_ALTDOWN
			isKeyReleased := (kbDllHookStruct.Flags | w32.LLKHF_UP) == w32.LLKHF_UP
			durationMS := kbDllHookStruct.Time // 這個是指事件創立完成自本次事件經過了多少毫秒，所以這個數值在只會越來越大 // https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime

			logMsg := func(title string) {
				log.Printf(title+" vkCode: %d isExtendedKey: %t lowerILInjected: %t injected: %t isAltDone: %t isKeyReleased: %t duration: %d ms\n",
					kbDllHookStruct.VkCode,
					isExtendedKey,
					lowerILInjected, injected,
					isAltDone,
					isKeyReleased,
					durationMS,
				)
			}

			if nCode == w32.HC_ACTION {
				// wParam就這四種而已{WM_KEYDOWN, WM_KEYUP, WM_SYSKEYDOWN, WM_SYSKEYUP}
				switch wParam {
				case w32.WM_SYSKEYUP:
					logMsg("[WM_SYSKEYUP]")
				case w32.WM_SYSKEYDOWN:
					logMsg("[WM_SYSKEYDOWN]")
				case w32.WM_KEYDOWN:
					logMsg("[WM_KEYDOWN]")
				case w32.WM_KEYUP:
					logMsg("[WM_KEYUP]")
					switch kbDllHookStruct.VkCode { // 有區分左右鍵
					case w32.VK_F1:
						log.Println("LLKeyboardHookProc VK_F1")
					case w32.VK_RCONTROL:
						log.Println("LLKeyboardHookProc VK_RCONTROL")
					case w32.VK_LCONTROL:
						log.Println("LLKeyboardHookProc VK_LCONTROL")
					case w32.VK_RSHIFT:
						log.Println("LLKeyboardHookProc VK_RSHIFT")
					case w32.VK_LSHIFT:
						log.Println("LLKeyboardHookProc VK_LSHIFT")
					case w32.VK_RMENU:
						log.Println("LLKeyboardHookProc VK_RMENU (ALT)")
					case w32.VK_LMENU:
						log.Println("LLKeyboardHookProc VK_LMENU (ALT)")
					}
				}
			}
			return user32dll.CallNextHookEx(hLLKeyboardHook, nCode, wParam, lParam)
		}
	}

	ch := make(chan w32.HWND)
	go func(chanWin chan<- w32.HWND) {
		// 定義處理視窗訊息的函數
		wndProcFuncPtr := syscall.NewCallback(w32.WndProc(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			switch uMsg {
			case w32.WM_DESTROY:
				log.Println("WM_DESTROY")
				for _, d := range []struct {
					w32.HHOOK
					Msg string
				}{
					{hCallWndHook, "CallWnd"},
					{hCBTProcHook, "CBT"},
					{hLLMouseHook, "Mouse"},
					{hLLKeyboardHook, "Keyboard"},
				} {
					if d.HHOOK == 0 {
						continue
					}
					if en := user32dll.UnhookWindowsHookEx(d.HHOOK); en != 0 {
						fmt.Printf("%s", en)
						continue
					}
					log.Printf("Unhook: %s\n", d.Msg)
				}
				user32dll.PostQuitMessage(0)
				return 0

			case w32.WM_CREATE:
				var errno syscall.Errno
				threadID := kernel32dll.GetCurrentThreadId()

				// init HHOOK
				{
					// Local HOOK (第三個參數0, 給第四個參數: threadID)
					if hCallWndHook, errno = user32dll.SetWindowsHookEx(w32.WH_CALLWNDPROC, hCallWndHookProc, 0, threadID); hCallWndHook == 0 {
						log.Printf("Error [WH_CALLWNDPROC] %s", errno)
					}

					if hCBTProcHook, errno = user32dll.SetWindowsHookEx(w32.WH_CBT, hCBTProcHookProc, 0, threadID); hCBTProcHook == 0 {
						log.Printf("Error [WH_CBT] %s", errno)
					}

					// Global HOOK (給第三個參數, 第四個參數為0)
					// LL 可以進行全局Hook，否則全局Hook無法建立成功，會遇到錯誤: Cannot set nonlocal hook without a module handle.
					hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
					if hLLMouseHook, errno = user32dll.SetWindowsHookEx(w32.WH_MOUSE_LL, hLLMouseHookProc, hInstance, 0); hLLMouseHook == 0 {
						log.Printf("Error [WH_MOUSE_LL] %s", errno)
					}
					if hLLKeyboardHook, errno = user32dll.SetWindowsHookEx(w32.WH_KEYBOARD_LL, hLLKeyboardHookProc, hInstance, 0); hLLKeyboardHook == 0 {
						log.Printf("Error [WH_KEYBOARD_LL] %s", errno)
					}

				}

			// 以下與hook無關，純粹是WM_MOUSEWHEEL用法的範例
			case w32.WM_MOUSEWHEEL:
				// lParam
				xPos := w32.GET_X_LPARAM(lParam)
				yPos := w32.GET_Y_LPARAM(lParam)

				// wparam
				keyState := w32.GET_KEYSTATE_WPARAM(wParam)      // LOWORD
				wheelDelta := w32.GET_WHEEL_DELTA_WPARAM(wParam) // 往下滾(朝自己)為負, 往上滾為正 // HIWORD
				isCtrlDone := (keyState & w32.MK_CONTROL) == w32.MK_CONTROL
				isLButtonDone := (keyState & w32.MK_LBUTTON) == w32.MK_LBUTTON
				isMouseBtnDone := (keyState & w32.MK_MBUTTON) == w32.MK_MBUTTON
				isRButtonDone := (keyState & w32.MK_RBUTTON) == w32.MK_RBUTTON
				isSHIFTDone := (keyState & w32.MK_SHIFT) == w32.MK_SHIFT
				isXButton1Done := (keyState & w32.MK_XBUTTON1) == w32.MK_XBUTTON1
				isXButton2Done := (keyState & w32.MK_XBUTTON2) == w32.MK_XBUTTON2
				log.Printf(`
xPos:%d yPos:%d
isCtrlDone:%t
isLButtonDone:%t
isMouseBtnDone:%t wheelDelta: %d
isRButtonDone:%t
isSHIFTDone:%t
isXButton1Done:%t
isXButton2Done:%t
`,
					xPos, yPos,
					isCtrlDone, isLButtonDone,
					isMouseBtnDone, wheelDelta,
					isRButtonDone, isSHIFTDone, isXButton1Done, isXButton2Done)

			// 以下是對在視窗中按下的按鍵做判斷(注意hook可能會搶訊息，所以該訊息可能被hook擷取導致這邊的訊息不觸發)
			// https://learn.microsoft.com/en-us/windows/win32/inputdev/about-keyboard-input?redirectedfrom=MSDN#keystroke-messages
			case w32.WM_KEYDOWN:
				fallthrough
			case w32.WM_KEYUP:
				fallthrough
			case w32.WM_SYSKEYDOWN:
				fallthrough
			case w32.WM_SYSKEYUP:
				virtualKeyCode := w32.LOWORD(wParam) // https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes?redirectedfrom=MSDN
				keyFlag := w32.HIWORD(lParam)
				var scanCode uint16
				scanCode = uint16(w32.LOBYTE(uintptr(keyFlag)))
				isExtendedKey := (keyFlag & w32.KF_EXTENDED) == w32.KF_EXTENDED // https://learn.microsoft.com/en-us/windows/win32/inputdev/about-keyboard-input?redirectedfrom=MSDN#extended-key-flag
				if isExtendedKey {
					// extended-key flag, 1 if scancode has 0xE0 prefix
					scanCode = w32.MAKEWORD(uintptr(scanCode), 0xE0)
				}
				wasKeyDown := (keyFlag & w32.KF_REPEAT) == w32.KF_REPEAT // 按住不放掉會是true
				repeatCount := w32.LOWORD(lParam)
				isKeyReleased := (keyFlag & w32.KF_UP) == w32.KF_UP // 按住不放掉就會是false

				// 由於每按下一個按鍵，其實會有包含兩個行為: {Done, Up}, 所以我們指判定up的時候才動作
				if isKeyReleased {
					// 我們想設定開關，來開啟或者暫停HOOK
					switch virtualKeyCode {
					case w32.VK_F1:
						*enableCallWnd = !*enableCallWnd
						log.Printf("VK_F1: enableCallWnd:%t\n", *enableCallWnd)
					case w32.VK_F2:
						*enableCBT = !*enableCBT
						log.Printf("VK_F2: enableCBT:%t\n", *enableCBT)
					case w32.VK_F3:
						*enableLLMouseHook = !*enableLLMouseHook
						log.Printf("VK_F3: enableMouseHook:%t\n", *enableLLMouseHook)
					case w32.VK_F4:
						*enableLLKeyboardHook = !*enableLLKeyboardHook
						log.Printf("VK_F4: enableLLKeyboardHook:%t\n", *enableLLKeyboardHook)
					case w32.VK_F11:
						*enableCallWnd = false
						*enableCBT = false
						*enableLLMouseHook = false
						*enableLLKeyboardHook = false
						log.Println("disable all hook")
					case w32.VK_F12:
						*enableCallWnd = true
						*enableCBT = true
						*enableLLMouseHook = true
						*enableLLKeyboardHook = true
						log.Println("enable all hook")
					}
				}

				switch virtualKeyCode {
				case w32.VK_LBUTTON:
					log.Println("VK_LBUTTON")

				// 一般的特殊按鍵{alt, shift, ctrl}是沒辦法被區分是按左邊還是右邊，需要特殊處理才能得知
				case w32.VK_SHIFT: // converts to VK_LSHIFT or VK_RSHIFT
					fallthrough
				case w32.VK_CONTROL: // converts to VK_LCONTROL or VK_RCONTROL
					fallthrough
				case w32.VK_MENU: // ALT  // converts to VK_LMENU or VK_RMENU
					virtualKeyCode = w32.LOWORD(uintptr(user32dll.MapVirtualKey(uint32(scanCode), w32.MAPVK_VSC_TO_VK_EX)))
					switch virtualKeyCode {
					case w32.VK_LSHIFT:
						log.Println("VK_LSHIFT")
					case w32.VK_RSHIFT:
						log.Println("VK_RSHIFT")
					case w32.VK_LCONTROL:
						log.Println("VK_LCONTROL")
					case w32.VK_RCONTROL:
						log.Println("VK_RCONTROL")
					case w32.VK_LMENU:
						log.Println("VK_LMENU")
					case w32.VK_RMENU:
						log.Println("VK_RMENU")
					}
				}
				log.Printf("vk_keyCode:%d wasKeyDown:%t repeatCount:%d isKeyReleased:%t", virtualKeyCode, wasKeyDown, repeatCount, isKeyReleased)

			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam)
		}))

		const (
			wndClassName  = "classSetWindowsHookEx"
			wndWindowName = "windowSetWindowsHookEx"
		)

		// 類別名稱註冊
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(wndClassName)
		hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))

		if atom, errno := user32dll.RegisterClass(&w32.WNDCLASS{
			Style:         w32.CS_HREDRAW | w32.CS_HREDRAW,
			HbrBackground: w32.COLOR_WINDOW,
			WndProc:       wndProcFuncPtr,
			HInstance:     hInstance,
			ClassName:     pUTF16ClassName,
		}); atom == 0 {
			fmt.Printf("%s", errno)
			chanWin <- 0
			return
		}

		// 確保程式結束之後能解除註冊名稱
		defer func() {
			if en2 := user32dll.UnregisterClass(wndClassName, hInstance); en2 != 0 {
				log.Printf("Error UnregisterClass: %s", en2)
			} else {
				log.Println("OK UnregisterClass")
			}

			// 通知外部程式用
			close(chanWin)
		}()

		// Create window
		hwnd, errno := user32dll.CreateWindowEx(0,
			wndClassName,
			wndWindowName,
			w32.WS_OVERLAPPEDWINDOW,

			// Size and position
			w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,

			0, // Parent window
			0, // Menu
			hInstance,
			0, // Additional application data
		)

		if hwnd == 0 {
			fmt.Printf("%s\n", errno)
			return
		}

		chanWin <- hwnd

		var msg w32.MSG
		for {
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 /* || msg.Message == w32.WM_QUIT */ {
				break
			}
			user32dll.TranslateMessage(&msg)
			user32dll.DispatchMessage(&msg)
		}
	}(ch)
	hwnd, isOpen := <-ch
	if !isOpen {
		return
	}
	log.Println(hwnd)

	user32dll.ShowWindow(hwnd, w32.SW_SHOW)

	// 🕹️ 如果您要手動嘗試，請把以下的SendMessage.WM_CLOSE註解掉，避免自動關閉
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)

	<-ch
	// Output:
}

func ExampleUser32DLL_RegisterHotKey() {
	user32dll := w32.NewUser32DLL()
	kernel32dll := w32.NewKernel32DLL()

	const (
		HokeyIDHello = 123
		HokeyIDBye   = 124
	)

	ch := make(chan w32.HWND)
	go func(className, windowName string, channel chan<- w32.HWND) {
		wndProcFuncPtr := syscall.NewCallback(w32.WndProc(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			switch uMsg {
			case w32.WM_DESTROY:
				log.Println("WM_DESTROY")
				for _, hotkeyID := range []int32{HokeyIDHello, HokeyIDBye} {
					if en := user32dll.UnregisterHotKey(hwnd, hotkeyID); en != 0 {
						log.Printf("Error [UnregisterHotKey] %s", en)
					}
				}
				user32dll.PostQuitMessage(0)
				return 0
			case w32.WM_CREATE:
				if en := user32dll.RegisterHotKey(hwnd, HokeyIDHello, w32.MOD_CONTROL|w32.MOD_ALT, w32.VK_F1); en != 0 {
					log.Println(en)
				}
				if en := user32dll.RegisterHotKey(hwnd, HokeyIDBye, w32.MOD_ALT, 65 /* A */); en != 0 {
					log.Println(en)
				}

			case w32.WM_HOTKEY:
				switch wParam {
				case HokeyIDHello:
					log.Println("hello")
					if lParam == 5 {
						fmt.Println("hello")
					}
				case HokeyIDBye:
					log.Println("bye~")
					_ = user32dll.PostMessage(hwnd, w32.WM_CLOSE, 0, 0)
				}
			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam) // default window proc
		}))

		hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(className)
		wc := w32.WNDCLASS{
			WndProc:   wndProcFuncPtr,
			HInstance: hInstance,
			ClassName: pUTF16ClassName,
		}

		if atom, errno := user32dll.RegisterClass(&wc); atom == 0 {
			fmt.Printf("%s", errno)
			return
		}

		defer func() {
			if en := user32dll.UnregisterClass(className, hInstance); en != 0 {
				fmt.Printf("[UnregisterClass] %s\n", en)
			}
			close(channel)
		}()

		if hwnd, errno := user32dll.CreateWindowEx(0,
			className, windowName,
			w32.WS_OVERLAPPEDWINDOW,
			w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,
			0, 0,
			hInstance, 0,
		); hwnd == 0 {
			fmt.Printf("%s", errno)
			return
		} else {
			channel <- hwnd
		}

		var msg w32.MSG
		for {
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
				break
			}
			user32dll.TranslateMessage(&msg)
			user32dll.DispatchMessage(&msg)
		}
	}("classRegisterHotKey", "windowRegisterHotKey", ch)

	hwnd, isOpen := <-ch
	if !isOpen {
		return
	}

	// user32dll.ShowWindow(hwnd, w32.SW_SHOW) // 不需要顯示視窗一樣可以觸發hotkey
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_HOTKEY, HokeyIDHello, 5)
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_HOTKEY, HokeyIDHello, 0)
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)
	<-ch

	// Output:
	// hello
}

// 按下Ctrl+1複製，並保存在變數buf1之中，按下Alt+1可以將buf1變數的內容寫入到剪貼簿並且貼上
// 熱鍵Ctrl+2複製選取內容到buf2變數; 熱鍵Alt+2將buf2的變數內容複製到剪貼簿並且貼上
func ExampleUser32DLL_RegisterHotKey_clipboard() {
	const (
		HokeyIDCtrl1 = w32.WM_APP + 1
		HokeyIDAlt1  = w32.WM_APP + 2

		HokeyIDCtrl2 = w32.WM_APP + 3
		HokeyIDAlt2  = w32.WM_APP + 4

		WMUpdateWindow = w32.WM_APP + 1024
	)

	var (
		buf1 = ""
		buf2 = ""
	)

	var (
		inputCtrlC [4]w32.INPUT
		inputCtrlV [4]w32.INPUT
	)

	// init hotkey
	{
		// C
		// 按下
		inputCtrlC[0].Type = w32.INPUT_KEYBOARD
		inputCtrlC[0].Ki().Vk = w32.VK_CONTROL
		inputCtrlC[1].Type = w32.INPUT_KEYBOARD
		inputCtrlC[1].Ki().Vk = w32.VK_KEY_C

		// 彈起
		inputCtrlC[2].Type = w32.INPUT_KEYBOARD
		inputCtrlC[2].Ki().Vk = w32.VK_KEY_C
		inputCtrlC[2].Ki().Flags = w32.KEYEVENTF_KEYUP

		inputCtrlC[3].Type = w32.INPUT_KEYBOARD
		inputCtrlC[3].Ki().Vk = w32.VK_CONTROL
		inputCtrlC[3].Ki().Flags = w32.KEYEVENTF_KEYUP

		// V
		inputCtrlV[0].Type = w32.INPUT_KEYBOARD
		inputCtrlV[0].Ki().Vk = w32.VK_CONTROL
		inputCtrlV[1].Type = w32.INPUT_KEYBOARD
		inputCtrlV[1].Ki().Vk = w32.VK_KEY_V

		inputCtrlV[2].Type = w32.INPUT_KEYBOARD
		inputCtrlV[2].Ki().Vk = w32.VK_KEY_V
		inputCtrlV[2].Ki().Flags = w32.KEYEVENTF_KEYUP

		inputCtrlV[3].Type = w32.INPUT_KEYBOARD
		inputCtrlV[3].Ki().Vk = w32.VK_CONTROL
		inputCtrlV[3].Ki().Flags = w32.KEYEVENTF_KEYUP
	}

	copy2clipboard := func() string {
		if _, eno := userDll.SendInput(4, &inputCtrlC[0], int32(unsafe.Sizeof(inputCtrlC[0]))); eno != 0 {
			log.Printf("SendInput error: %s\n", eno)
			return ""
		}
		time.Sleep(200 * time.Millisecond) // 等待剪貼簿複製完成

		// 從剪貼簿獲取資料
		if eno := userDll.OpenClipboard(0); eno != 0 {
			log.Printf("open clipboard error %s\n", eno)
			return ""
		}
		defer func() {
			if eno := userDll.CloseClipboard(); eno != 0 {
				log.Printf("CloseClipboard error %s\n", eno)
			}
		}()

		var clipboardText string
		if hClipboardData, eno := userDll.GetClipboardData(w32.CF_UNICODETEXT); eno == 0 {
			lpMemData, _ := kernelDll.GlobalLock(w32.HGLOBAL(hClipboardData))
			clipboardText = syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(lpMemData))[:])
			if _, eno = kernelDll.GlobalUnlock(w32.HGLOBAL(hClipboardData)); eno != 0 {
				log.Printf("GlobalUnlock error. %s", eno)
			}
		} else {
			log.Printf("GetClipboardData error %s\n", eno)
			return ""
		}
		return clipboardText
	}

	paste := func(text string) {
		// 將資料餵入剪貼簿
		eno := userDll.OpenClipboard(0)
		if eno != 0 {
			log.Printf("open clipboard error %s\n", eno)
			return
		}
		if eno = userDll.EmptyClipboard(); eno != 0 {
			log.Printf("EmptyClipboard error %s\n", eno)
			if eno = userDll.CloseClipboard(); eno != 0 {
				log.Printf("CloseClipboard error %s\n", eno)
			}
			return
		}
		data := utf16.Encode([]rune(text + "\x00"))
		size := len(data) * int(unsafe.Sizeof(data[0]))
		hMem, eno := kernelDll.GlobalAlloc(w32.GMEM_MOVEABLE, w32.SIZE_T(size))
		if eno == 0 {
			defer kernelDll.GlobalFree(hMem)
			lpMemData, _ := kernelDll.GlobalLock(hMem)
			kernelDll.StrCpyW(uintptr(lpMemData), &data[0])
			if _, eno = kernelDll.GlobalUnlock(hMem); eno == 0 {
				if _, eno = userDll.SetClipboardData(w32.CF_UNICODETEXT, w32.HANDLE(hMem)); eno != 0 {
					log.Printf("SetClipboardData error: %s\n", eno)
				}
			}
		} else {
			log.Printf("GlobalAlloc error %s\n", eno)
		}
		if eno = userDll.CloseClipboard(); eno != 0 {
			log.Printf("CloseClipboard error %s\n", eno)
		}

		// 貼上
		time.Sleep(250 * time.Millisecond) // 剪貼簿關閉之後要等待一段時間，等待的時間也不行太短100，否則接下來的SendInput都可能會異常，貼不到資料
		if _, eno = userDll.SendInput(4, &inputCtrlV[0], int32(unsafe.Sizeof(inputCtrlV[0]))); eno != 0 {
			log.Printf("SendInput error: %s\n", eno)
		}
	}

	const fontSize = 36
	hFontSystem := gdiDll.CreateFont(
		int32(fontSize), 0, 0, 0,
		w32.FW_DONTCARE,
		0, 1, 0, w32.DEFAULT_CHARSET,
		w32.OUT_OUTLINE_PRECIS,
		w32.CLIP_DEFAULT_PRECIS,
		w32.CLEARTYPE_QUALITY,
		w32.DEFAULT_PITCH,
		"System",
	)
	if hFontSystem == 0 {
		return
	}
	defer func() {
		if gdiDll.DeleteObject(w32.HGDIOBJ(hFontSystem)) {
			log.Println("DeleteObject HFONT")
		}
	}()

	opt := &w32.WindowOptions{Width: 800, Height: 600,
		ClassName: "ExampleUser32DLL_RegisterHotKey_clipboard", // 如果class名稱已經存在會遇到Class already exists的錯誤
	}
	opt.WndProc = func(hwnd w32.HWND, uMsg uint32, wParam w32.WPARAM, lParam w32.LPARAM) uintptr {
		switch uMsg {
		case w32.WM_CREATE:
			userDll.ShowWindow(hwnd, w32.SW_SHOW)
			// hwnd Null 可以註冊全局熱鍵, 不過這種情況下，Msg的接收要寫在消息循環之中，因為他不綁hwnd
			if en := userDll.RegisterHotKey(0, HokeyIDCtrl1, w32.MOD_CONTROL, w32.VK_KEY_1); en != 0 {
				log.Println(en)
			}
			if en := userDll.RegisterHotKey(0, HokeyIDAlt1, w32.MOD_ALT, w32.VK_KEY_1); en != 0 {
				log.Println(en)
			}
			if en := userDll.RegisterHotKey(0, HokeyIDCtrl2, w32.MOD_CONTROL, w32.VK_KEY_2); en != 0 {
				log.Println(en)
			}
			if en := userDll.RegisterHotKey(0, HokeyIDAlt2, w32.MOD_ALT, w32.VK_KEY_2); en != 0 {
				log.Println(en)
			}
			// userDll.RegisterHotKey(0, HokeyIDTest, 0, w32.VK_F1) // 可以只單個按鍵: 例如F1 (輔助鍵不需要可以給NULL)，
			go func() {
				<-time.After(2 * time.Second)
				log.Println("2秒已到，自動關閉程式. (如果要測試，可以自行延長秒數)")
				_, _, _ = userDll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)
			}()

		case w32.WM_DESTROY:
			for _, hotkeyID := range []int32{HokeyIDCtrl1, HokeyIDCtrl2, HokeyIDAlt1, HokeyIDAlt2} {
				if en := userDll.UnregisterHotKey(0, hotkeyID); en != 0 {
					log.Printf("Error [UnregisterHotKey] %s", en)
				}
			}
			userDll.PostQuitMessage(0)
			return 0
		case WMUpdateWindow:
			hdc := userDll.GetDC(hwnd)
			defer func() {
				userDll.ReleaseDC(hwnd, hdc)
			}()
			// 清空
			var rect w32.RECT
			_ = userDll.GetClientRect(hwnd, &rect)
			hRgnBackground := gdiDll.CreateRectRgnIndirect(&rect)
			defer gdiDll.DeleteObject(w32.HGDIOBJ(hRgnBackground))
			hBrush := gdiDll.CreateSolidBrush(w32.RGB(255, 255, 255))
			defer gdiDll.DeleteObject(w32.HGDIOBJ(hBrush))
			gdiDll.FillRgn(hdc, hRgnBackground, hBrush)

			windowWidth := rect.Right - rect.Left
			// windowHeight := rect.Bottom - rect.Top
			gdiDll.SelectObject(hdc, w32.HGDIOBJ(hFontSystem))
			for h, text := range []string{buf1, buf2} {
				userDll.SetRect(&rect, 10, int32((h)*fontSize), windowWidth, int32((h+1)*fontSize))
				userDll.DrawText(hdc, fmt.Sprintf("[%d] %s", h, text), -1, &rect, w32.DT_NOCLIP)
			}
			return 0
		}
		return uintptr(userDll.DefWindowProc(hwnd, w32.UINT(uMsg), wParam, lParam))
	}

	wnd, err := createWindow("multi clipboard", opt)
	if err != nil {
		log.Fatal(err)
	}
	wnd.Run(func(msg *w32.MSG) bool {
		if msg.Message != w32.WM_HOTKEY {
			return true
		}
		log.Println("WM_HOTKEY received")
		var curBuf *string
		switch msg.WParam {
		// copy
		case HokeyIDCtrl2:
			curBuf = &buf2
			fallthrough
		case HokeyIDCtrl1:
			if curBuf == nil {
				curBuf = &buf1
			}
			*curBuf = copy2clipboard()                              // 複製到剪貼簿，也複製到該變數
			_ = userDll.PostMessage(wnd.hwnd, WMUpdateWindow, 0, 0) // 在畫面上顯示剪貼簿的內容
		// paste
		case HokeyIDAlt2:
			curBuf = &buf2
			fallthrough
		case HokeyIDAlt1:
			if curBuf == nil {
				curBuf = &buf1
			}
			paste(*curBuf)
		}
		return false
	})
	// Output:
}

func ExampleUser32DLL_BeginPaint() {
	user32dll := w32.NewUser32DLL()
	kernel32dll := w32.NewKernel32DLL()
	gdi32dll := w32.NewGdi32DLL()

	ch := make(chan w32.HWND)
	go func(className, windowName string, channel chan<- w32.HWND) {
		wndProcFuncPtr := syscall.NewCallback(w32.WndProc(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			switch uMsg {
			case w32.WM_DESTROY:
				user32dll.PostQuitMessage(0)
				return 0
			case w32.WM_CREATE:
				user32dll.ShowWindow(hwnd, w32.SW_SHOW)
			case w32.WM_PAINT:
				log.Println("WM_PAINT")
				var paintStruct w32.PAINTSTRUCT
				user32dll.BeginPaint(hwnd, &paintStruct)

				hdc := user32dll.GetDC(hwnd)
				defer func() {
					user32dll.ReleaseDC(hwnd, hdc)
				}()

				var rect w32.RECT
				_ = user32dll.GetClientRect(hwnd, &rect)

				// Background Color
				hRgnBackground := gdi32dll.CreateRectRgnIndirect(&rect)
				defer gdi32dll.DeleteObject(w32.HGDIOBJ(hRgnBackground))
				hBrush := gdi32dll.CreateSolidBrush(w32.RGB(128, 128, 128))
				defer gdi32dll.DeleteObject(w32.HGDIOBJ(hBrush))

				gdi32dll.FillRgn(hdc, hRgnBackground, hBrush)

				gdi32dll.SetTextColor(hdc, w32.RGB(255, 128, 0))

				user32dll.DrawText(hdc, "Hello World 您好 世界", -1, &rect, w32.DT_NOCLIP)
				gdi32dll.TextOut(hdc, 100, 200, "Hi, 您好")
				user32dll.EndPaint(hwnd, &paintStruct)
			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam) // default window proc
		}))

		hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(className)
		wc := w32.WNDCLASS{
			WndProc:   wndProcFuncPtr,
			HInstance: hInstance,
			ClassName: pUTF16ClassName,
		}

		if atom, errno := user32dll.RegisterClass(&wc); atom == 0 {
			fmt.Printf("%s", errno)
			return
		}

		defer func() {
			if en := user32dll.UnregisterClass(className, hInstance); en != 0 {
				fmt.Printf("[UnregisterClass] %s\n", en)
			}
			close(channel)
		}()

		if hwnd, errno := user32dll.CreateWindowEx(0,
			className, windowName,
			w32.WS_OVERLAPPEDWINDOW,
			w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,
			0, 0,
			hInstance, 0,
		); hwnd == 0 {
			fmt.Printf("%s", errno)
			return
		} else {
			channel <- hwnd
		}

		var msg w32.MSG
		for {
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
				break
			}
			user32dll.TranslateMessage(&msg)
			user32dll.DispatchMessage(&msg)
		}
	}("classExampleBeginPaint", "windowExampleBeginPaint", ch)

	hwnd, isOpen := <-ch
	if !isOpen {
		return
	}

	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_PAINT, 0, 0)
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)

	<-ch
	// Output:
}

func TestINPUT_Hi(t *testing.T) {
	var input w32.INPUT
	input.Hi().Msg = 0
	input.Hi().LParamL = 0
	input.Hi().LParamH = 0
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendinput
func ExampleUser32DLL_SendInput() {
	user32dll := w32.NewUser32DLL()

	log.Println("Example: [INPUT_KEYBOARD] Sending 'Win-D' to Display and hide the desktop.")
	{
		var input [4]w32.INPUT
		input[0].Type = w32.INPUT_KEYBOARD
		input[0].Ki().Vk = w32.VK_LWIN
		input[1].Type = w32.INPUT_KEYBOARD
		input[1].Ki().Vk = w32.VK_KEY_D

		input[2].Type = w32.INPUT_KEYBOARD
		input[2].Ki().Vk = w32.VK_KEY_D
		input[2].Ki().Flags = w32.KEYEVENTF_KEYUP
		input[3].Type = w32.INPUT_KEYBOARD
		input[3].Ki().Vk = w32.VK_LWIN
		input[3].Ki().Flags = w32.KEYEVENTF_KEYUP

		// Display the desktop at the first one.
		if n, errno := user32dll.SendInput(4, &input[0], int32(unsafe.Sizeof(input[0]))); n == 0 {
			fmt.Printf("%s", errno)
		}
		// And try again will restore
		time.Sleep(200 * time.Millisecond) // If running too fast may not work.
		_, _ = user32dll.SendInput(4, &input[0], int32(unsafe.Sizeof(input[0])))
	}

	log.Println("Example: [INPUT_MOUSE]")
	{
		var input w32.INPUT
		// input.Type = w32.INPUT_MOUSE // not necessary, since INPUT_MOUSE is zero
		input.Mi().Dx = 10000
		input.Mi().Dy = 30000
		input.Mi().Flags = w32.MOUSEEVENTF_MOVE | w32.MOUSEEVENTF_ABSOLUTE
		if n, errno := user32dll.SendInput(1, &input, int32(unsafe.Sizeof(input))); n == 0 {
			log.Printf("%s", errno)
		}
	}

	// Output:
}

func ExampleUser32DLL_SendInput_keyboard() {
	user32dll := w32.NewUser32DLL()
	myStr := "Hello World 您好世界 !"
	var input []w32.INPUT
	uint16Array, _ := syscall.UTF16FromString(myStr)
	input = make([]w32.INPUT, len(uint16Array)*2) // *2: down and up

	for i, uint16Val := range uint16Array {
		input[2*i].Type = w32.INPUT_KEYBOARD
		input[2*i].Ki().Flags = w32.KEYEVENTF_UNICODE
		input[2*i].Ki().Scan = uint16Val

		input[2*i+1].Type = w32.INPUT_KEYBOARD
		input[2*i+1].Ki().Flags = w32.KEYEVENTF_UNICODE | w32.KEYEVENTF_KEYUP
		input[2*i+1].Ki().Scan = uint16Val
	}

	hwnd := user32dll.FindWindow("Notepad", "")
	hdc := user32dll.GetDC(hwnd)
	defer user32dll.ReleaseDC(hwnd, hdc)

	if hwnd != 0 {
		user32dll.ShowWindow(hwnd, w32.SW_NORMAL) // 在視窗最小化時SetForegroundWindow或者SetActiveWindow都沒有用
		user32dll.SetForegroundWindow(hwnd)
	}

	if n, errno := user32dll.SendInput(uint32(len(uint16Array)*2), &input[0], int32(unsafe.Sizeof(input[0]))); n == 0 {
		fmt.Printf("%s", errno)
	}

	// Output:
}

func TestUser32DLL_EnumWindows(t *testing.T) {
	user32dll := w32.NewUser32DLL()
	kernel32dll := w32.NewKernel32DLL()

	enumFuncOK := w32.WndEnumProc(func(hwnd w32.HWND, lParam w32.LPARAM) w32.BOOL {

		if user32dll.IsWindowVisible(hwnd) {
			var err error
			var className, windowName string
			if windowName, err = user32dll.GetWindowText(hwnd); err != nil {
				return 1
			}
			if className, err = user32dll.GetClassName(hwnd); err != nil {
				return 1
			}
			// log.Printf() // 很奇怪會相衝的樣子，導致顯示不完整，整個Example沒跑完
			_, _ = fmt.Fprintf(os.Stderr, "%d _ %s _ %s\n", hwnd, className, windowName)
		}
		return 1
	})

	type MyData struct {
		id  uint32
		tag [4]byte
	}

	enumFuncOK2 := w32.WndEnumProc(func(hwnd w32.HWND, lParam w32.LPARAM) w32.BOOL {
		d := *(*MyData)(unsafe.Pointer(lParam))
		if d.id == 666 {
			log.Println(string(d.tag[:])) // 123, wall
		}
		return 888 // 不建議這樣用，因為只要非0就還會繼續，最終的回傳值取決於最後一個回傳的結果
	})

	enumFuncErr := w32.WndEnumProc(func(hwnd w32.HWND, lParam w32.LPARAM) w32.BOOL {
		if user32dll.IsWindowVisible(hwnd) {
			if windowName, err := user32dll.GetWindowText(hwnd); err != nil {
				kernel32dll.SetLastError(w32.ERROR_INVALID_DATA) // 自定義錯誤訊息
				// kernel32dll.SetLastError(uint32(err.(syscall.Errno)))
				return 0 // 只要回傳0，整個enum就會終止，不會再繼續下去
			} else {
				log.Println(windowName)
			}
		}
		return 1
	})

	// 在github.action測試，有可能會發生列舉不完導致超時的錯誤，所以設定此機制來避免
	ch := make(chan string)
	go func() {
		// Example 1 常規用法. 忽略錯誤且不計較回傳值
		log.Println("test no error")
		_, _ = user32dll.EnumWindows(enumFuncOK, 0)

		// Example 2 傳遞參數給WndEnumProc, 並且接收回傳值
		data := MyData{123, [4]byte{'w', 'a', 'l', 'l'}}
		rtnVal, _ := user32dll.EnumWindows(enumFuncOK2, w32.LPARAM(unsafe.Pointer(&data)))
		fmt.Println(rtnVal)

		// Example 3 模擬錯誤的情況. 注意enumWindows當傳遞的函數傳回0之後就會直接終止，若不為0則會繼續直到窮舉完畢
		log.Println("test error")
		if r, errno := user32dll.EnumWindows(enumFuncErr, 0); r == 0 {
			log.Printf("%s\n", errno)
		}
		ch <- "finish"
	}()

	select {
	case <-ch:
		return
	case <-time.After(2 * time.Second):
		log.Println("timeout")
	}

	// Output:
	// 888
}

// Github actions執行會有問題，所以不實際跑
func ExampleUser32DLL_EnumDesktopWindows() {
	var curDesktopHandList []w32.HWND
	r, eno := userDll.EnumDesktopWindows(0, func(hwnd w32.HWND, lParam w32.LPARAM) w32.BOOL {
		curDesktopHandList = append(curDesktopHandList, hwnd)
		var className, windowName string
		var err error
		if windowName, err = userDll.GetWindowText(hwnd); err != nil {
			return 1 // 繼續列舉，回傳0就會列舉結束
		}
		if className, err = userDll.GetClassName(hwnd); err != nil {
			return 1
		}
		_, _ = fmt.Fprintf(os.Stderr, "%d _ %s _ %s\n", hwnd, className, windowName)
		return 1
	}, 0)
	log.Println(r, eno) // 1, 0
	log.Println(curDesktopHandList)
}

// 打印出桌面名稱
// https://learn.microsoft.com/en-us/windows/win32/winstation/window-station-and-desktop-creation
// Github actions執行會有問題，所以不實際跑
func ExampleUser32DLL_EnumDesktops() {
	log.Println(
		userDll.EnumDesktops(0, func(name string, lParam w32.LPARAM) w32.BOOL {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", name) // WinSta0, Service-0x0-3e7$
			return 1
		}, 0),
	)
}

func ExampleUser32DLL_PrintWindow() {
	user32dll := w32.NewUser32DLL()
	hwndDst := user32dll.FindWindow("ApplicationFrameWindow", "小算盤")
	hdcDst := user32dll.GetDC(hwndDst)
	defer user32dll.ReleaseDC(hwndDst, hdcDst)
	hwndSrc := user32dll.FindWindow("Notepad", "")
	if !user32dll.PrintWindow(hwndSrc, hdcDst, 0) {
		log.Println("not ok")
	}

	// Output:
}

func ExampleUser32DLL_IsIconic() {
	user32dll := w32.NewUser32DLL()
	hwnd := user32dll.GetDesktopWindow()
	if user32dll.IsIconic(hwnd) {
		log.Println("is minimized")
	}
	// Output:
}

func ExampleUser32DLL_GetWindowLong() {
	user32dll := w32.NewUser32DLL()
	hwnd := user32dll.GetDesktopWindow()
	if user32dll.GetWindowLong(hwnd, w32.GWL_STYLE)&w32.WS_MINIMIZE == w32.WS_MINIMIZE {
		log.Println("is minimized")
	}
	// Output:
}

type HookData struct {
	Type int32 // WH_KEYBOARD_LL, WH_MOUSE_LL,...
	*w32.HHOOK
	*w32.HOOKPROC
}

func StartLowLevelHook(user32dll *w32.User32DLL, chQuit chan<- bool, wg *sync.WaitGroup, hwndScreen w32.HWND, hookDatas ...HookData) {
	if len(hookDatas) == 0 {
		close(chQuit)
		return
	}

	var errno syscall.Errno
	kernel32dll := w32.NewKernel32DLL(w32.PNGetModuleHandle)
	hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))

	for _, h := range hookDatas {
		if *h.HHOOK, errno = user32dll.SetWindowsHookEx(h.Type, *h.HOOKPROC, hInstance, 0); *h.HHOOK == 0 {
			log.Printf("Error [WH_KEYBOARD_LL] %s", errno)
			wg.Done()
			return
		}
	}

	defer func() {
		for _, h := range hookDatas {
			if user32dll.UnhookWindowsHookEx(*h.HHOOK) != 0 {
				log.Printf("UnhookWindowsHookEx OK")
			}
		}
		close(chQuit)
	}()

	go func() {
		var msg w32.MSG
		if status, _ := user32dll.GetMessage(&msg, hwndScreen, 0, 0); status <= 0 {
			return
		}
	}()
	wg.Wait()
}

// Press F9 can capture the specific device screen and show it on the notepad.
func Test_captureWindowByHotkey(t *testing.T) {
	user32dll := w32.NewUser32DLL()
	gdi32dll := w32.NewGdi32DLL()

	// 🕹️ change this such that you have time to play more.
	const timeout = 2 * time.Second

	hwndScreen := user32dll.GetDesktopWindow()
	hdcScreen := user32dll.GetDC(hwndScreen)
	defer user32dll.ReleaseDC(hwndScreen, hdcScreen)

	// Calculator
	hwndC := user32dll.FindWindow("ApplicationFrameWindow", "小算盤")
	if hwndC == 0 {
		log.Println("Calculator not found")
		hwndC = hwndScreen // In order to allow the test to continue, use this instead of it.
	}

	hwndDst := user32dll.FindWindow("Notepad", "")
	if hwndDst == 0 {
		log.Println("Notepad not found")
	}
	hdcDst := user32dll.GetWindowDC(hwndDst)
	defer user32dll.ReleaseDC(hwndDst, hdcDst)

	ch := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	var (
		hLLKeyboardHook     w32.HHOOK
		hLLKeyboardHookProc w32.HOOKPROC
	)

	// init hookProc
	hLLKeyboardHookProc = func(nCode int32, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
		if nCode < 0 {
			return user32dll.CallNextHookEx(hLLKeyboardHook, nCode, wParam, lParam)
		}
		kbDllHookStruct := *(*w32.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		var rectTarget w32.RECT
		if nCode == w32.HC_ACTION {
			switch wParam {
			case w32.WM_KEYUP:
				switch kbDllHookStruct.VkCode {
				case w32.VK_F6:
					wg.Done() // Exit program
				case w32.VK_F9:
					// user32dll.GetWindowLong(hwndC, w32.GWL_STYLE)&w32.WS_MINIMIZE == w32.WS_MINIMIZE // below are better
					if user32dll.IsIconic(hwndC) {
						user32dll.ShowWindow(hwndC, w32.SW_NORMAL)
					}

					if errno := user32dll.GetWindowRect(hwndC, &rectTarget); errno != 0 {
						log.Println(errno)
						return user32dll.CallNextHookEx(hLLKeyboardHook, nCode, wParam, lParam)
					}

					if errno := gdi32dll.BitBlt(
						hdcDst, 0, 0, rectTarget.Width(), rectTarget.Height(),
						hdcScreen, rectTarget.Left, rectTarget.Top, w32.SRCCOPY,
					); errno != 0 {
						log.Println(errno)
					}
				}
			}
		}
		return user32dll.CallNextHookEx(hLLKeyboardHook, nCode, wParam, lParam)
	}
	go StartLowLevelHook(user32dll, ch, &wg, hwndScreen, HookData{w32.WH_KEYBOARD_LL, &hLLKeyboardHook, &hLLKeyboardHookProc})

	for {
		select {
		case _, isOpen := <-ch:
			if !isOpen {
				return
			}
		case <-time.After(timeout):
			log.Println("timeout")
			wg.Done()
		}
	}
}

// Capture the Calculator's screen every 0.1 seconds and then draw it on the notepad.
func Test_captureWindow(t *testing.T) {
	user32dll := w32.NewUser32DLL()
	gdi32dll := w32.NewGdi32DLL()

	// 🕹️ change this such that you have time to play more.
	const timeout = 2 * time.Second

	hwndScreen := user32dll.GetDesktopWindow()
	hdcScreen := user32dll.GetDC(hwndScreen)
	defer user32dll.ReleaseDC(hwndScreen, hdcScreen)

	// Calculator
	hwndC := user32dll.FindWindow("ApplicationFrameWindow", "小算盤")
	if hwndC == 0 {
		log.Println("Calculator not found")
		hwndC = hwndScreen // In order to allow the test to continue, use this instead of it.
	}

	hwndDst := user32dll.FindWindow("Notepad", "")
	if hwndDst == 0 {
		log.Println("Notepad not found")
		return
	}
	hdcDst := user32dll.GetWindowDC(hwndDst)
	defer user32dll.ReleaseDC(hwndDst, hdcDst)

	ch := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		var rectNotepad w32.RECT
		var rectC w32.RECT
		gdi32dll.SetStretchBltMode(hdcDst, w32.HALFTONE)
		for {
			select {
			case _, isOpen := <-ch:
				if !isOpen {
					wg.Done()
					return
				}
			default:
				time.Sleep(100 * time.Millisecond)
			}
			if user32dll.GetWindowRect(hwndDst, &rectNotepad) != 0 {
				continue
			}
			if user32dll.GetWindowRect(hwndC, &rectC) != 0 {
				continue
			}

			gdi32dll.StretchBlt(
				hdcDst, rectNotepad.Width()/4, rectNotepad.Height()/4, rectNotepad.Width()/2, rectNotepad.Height()/2, // draw screen on the center
				hdcScreen, rectC.Left, rectC.Top, rectC.Width(), rectC.Height(), w32.SRCCOPY,
			)
		}
	}()

	for {
		select {
		case _, isOpen := <-ch:
			if !isOpen {
				wg.Wait()
				return
			}
		case <-time.After(timeout):
			log.Println("timeout")
			close(ch)
		}
	}
}

func ExampleUser32DLL_SetClipboardData() {
	_ = userDll.OpenClipboard(0)
	_ = userDll.EmptyClipboard()

	text := `
Hello World 您好 世界
123
`
	// data := syscall.StringToUTF16(text)
	data := utf16.Encode([]rune(text + "\x00"))
	size := len(data) * int(unsafe.Sizeof(data[0]))
	hMem, eno := kernelDll.GlobalAlloc(w32.GMEM_MOVEABLE, w32.SIZE_T(size))
	if eno == 0 {
		defer kernelDll.GlobalFree(hMem)

		// 鎖定內存
		lpMemData, _ := kernelDll.GlobalLock(hMem)

		// 寫入資料
		kernelDll.StrCpyW(uintptr(lpMemData), &data[0])

		// 解鎖後才能被調用
		if _, eno = kernelDll.GlobalUnlock(hMem); eno != 0 {
			log.Println(eno)
		}
	}

	if _, eno = userDll.SetClipboardData(w32.CF_UNICODETEXT, w32.HANDLE(hMem)); eno != 0 {
		log.Println(eno)
	}
	_ = userDll.CloseClipboard()

	// 驗證
	_ = userDll.OpenClipboard(0)

	hClipboardData, eno := userDll.GetClipboardData(w32.CF_UNICODETEXT)
	if hClipboardData != 0 {
		lpMemData, _ := kernelDll.GlobalLock(w32.HGLOBAL(hClipboardData))
		clipboardText := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(lpMemData))[:])
		_, _ = kernelDll.GlobalUnlock(w32.HGLOBAL(hClipboardData))
		log.Println("Clipboard text:", clipboardText)
		fmt.Println(text == clipboardText)
	} else {
		log.Println(eno)
	}
	_ = userDll.CloseClipboard()

	// Output:
	// true
}

// 複製圖片(點集資料)到剪貼簿去
func TestUser32DLL_SetClipboardData(t *testing.T) {
	// 核心:
	// SetClipboardData(w32.CF_DIB, w32.HANDLE(hDIB))
	// 其中hDIB要有{DIB HEADER, 點集資料}

	var x, y, width, height int32
	x = 100
	y = 200
	width = 300
	height = 400

	// screenshot
	var bitmapInfoHeader w32.BitmapInfoHeader
	size := width * height * 4
	// 用截圖的方式來模擬點集資料生成
	pixels := make([]byte, size)
	{
		hdc := userDll.GetDC(0)
		defer userDll.ReleaseDC(0, hdc)

		hMemDC := gdiDll.CreateCompatibleDC(hdc)
		defer gdiDll.DeleteDC(hMemDC)

		hBitmap := gdiDll.CreateCompatibleBitmap(hdc, width, height)
		defer gdiDll.DeleteObject(w32.HGDIOBJ(hBitmap))
		hBitmapOld := gdiDll.SelectObject(hMemDC, w32.HGDIOBJ(hBitmap))
		defer gdiDll.SelectObject(hMemDC, hBitmapOld)

		_ = gdiDll.BitBlt(hMemDC,
			0, 0,
			width, height,
			hdc,
			x, y,
			w32.SRCCOPY)

		var bitmap w32.Bitmap
		gdiDll.GetObject(w32.HANDLE(hBitmap), int32(unsafe.Sizeof(bitmap)), uintptr(unsafe.Pointer(&bitmap)))

		bitmapInfoHeader = w32.BitmapInfoHeader{
			Size:        uint32(unsafe.Sizeof(bitmapInfoHeader)), // 也可以直接寫40
			Width:       bitmap.Width,
			Height:      bitmap.Height,
			Planes:      1,
			BitCount:    32,
			Compression: w32.BI_RGB,
		}

		bmpSize := ((bitmap.Width*int32(bitmapInfoHeader.BitCount) + 31) / 32) * 4 * bitmap.Height
		hDIB, _ := kernelDll.GlobalAlloc(w32.GHND, w32.SIZE_T(bmpSize))
		defer kernelDll.GlobalFree(hDIB)

		var lpBitmap w32.LPVOID
		lpBitmap, _ = kernelDll.GlobalLock(hDIB)

		gdiDll.GetDIBits(
			hMemDC, hBitmap, 0,
			uint32(bitmap.Height),
			lpBitmap, // [out]
			&w32.BitmapInfo{Header: bitmapInfoHeader},
			w32.DIB_RGB_COLORS,
		)
		copy(pixels, (*[1 << 25]byte)(unsafe.Pointer(lpBitmap))[:size])

		if _, eno := kernelDll.GlobalUnlock(hDIB); eno != 0 {
			t.Fatal(eno)
		}
	}

	// SetClipboardData(w32.CF_DIB, w32.HANDLE(hDIB))
	{
		hDIB2, _ := kernelDll.GlobalAlloc(w32.GMEM_MOVEABLE,
			w32.SIZE_T(uint32(unsafe.Sizeof(bitmapInfoHeader))+uint32(height*width)*4), // 如果是rgba就需要*4
		)
		if hDIB2 == 0 {
			t.Fatal("GlobalAlloc error")
			return
		}
		defer kernelDll.GlobalFree(hDIB2)
		var lpDIB w32.LPVOID
		lpDIB, _ = kernelDll.GlobalLock(hDIB2)

		// 寫入DIB HEADER
		*(*w32.BitmapInfoHeader)(unsafe.Pointer(lpDIB)) = bitmapInfoHeader

		// 寫入點集資料
		// 1GB  1 << 30,
		// 32MB 1 << 25
		// 將pixels內容複製到(lpDIB+40)指向的記憶體位置，其可容納內存空間為(32MB)，以確保內容不會被截斷或超出範圍
		copy((*[1 << 25]byte)(unsafe.Pointer(uintptr(lpDIB) + unsafe.Sizeof(bitmapInfoHeader)))[:len(pixels)],
			pixels,
		)

		// 解鎖後才能被調用
		if _, eno := kernelDll.GlobalUnlock(hDIB2); eno != 0 {
			t.Fatal(eno)
		}

		_ = userDll.OpenClipboard(0)
		defer func() {
			_ = userDll.CloseClipboard()
		}()
		_ = userDll.EmptyClipboard()

		if _, eno := userDll.SetClipboardData(w32.CF_DIB, w32.HANDLE(hDIB2)); eno != 0 {
			t.Fatal(eno)
		}
	}
}

// 模擬滑鼠移動到指定的位置之後雙擊左鍵
func ExampleUser32DLL_SetCursorPos() {
	eno := userDll.SetCursorPos(100, 200)
	fmt.Println(eno != 0)

	input := w32.INPUT{
		Type: w32.INPUT_MOUSE,
	}
	input.Mi().Flags = w32.MOUSEEVENTF_LEFTDOWN | w32.MOUSEEVENTF_LEFTUP
	_, _ = userDll.SendInput(1, &input, int32(unsafe.Sizeof(input)))
	time.Sleep(50 * time.Millisecond)
	_, _ = userDll.SendInput(1, &input, int32(unsafe.Sizeof(input)))

	// Output:
	// true
}
