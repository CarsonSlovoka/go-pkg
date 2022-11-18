package w32_test

import (
	"encoding/binary"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"syscall"
	"testing"
	"time"
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
			if ok, errno := user32dll.DrawIcon(curHDC, d.x, d.y, d.hIcon); !ok {
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
	_, _ = user32dll.DrawIconEx(hdcScreen, 10, 20, hIcon, 0, 0, 0, 0, w32.DI_NORMAL)

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
		_, _ = user32dll.DrawIconEx(hdcScreen, xLeft, yTop, hIcon, d.width, d.height, 0, 0, d.diFlag)
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
	if !user32dll.GetIconInfo(hIconQuestion, &iInfo) {
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
			Type:       0x4D42,    // BM. // B: 42, M: 4D  //  All of the integer values are stored in little-endian format
			Size:       sizeofDIB, // HEADER + INFO + DATA
			OffsetBits: 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)),
		}

		hdc := user32dll.GetDC(0)

		kernel32dll := w32.NewKernel32DLL()

		var lpBitmap w32.LPVOID
		hDIB, _ := kernel32dll.GlobalAlloc(w32.GHND, w32.SIZE_T(bmpSize))
		lpBitmap, _ = kernel32dll.GlobalLock(hDIB)
		defer func() {
			kernel32dll.GlobalUnlock(hDIB)
			kernel32dll.GlobalFree(hDIB)
		}()
		gdi32dll.GetDIBits(
			hdc, iInfo.HbmColor,
			0,
			w32.UINT(bmp.Height),
			lpBitmap, // [out]
			&w32.BitmapInfo{Header: bitmapInfoHeader},
			w32.DIB_RGB_COLORS,
		)
		outputBmpPath := "testdata/info.bmp"
		// Write: FileHeader, DIPHeader, bitmapData
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

			// DIP Header
			_ = binary.Write(f, binary.LittleEndian, bitmapInfoHeader)

			// bitmapData
			bmpDatas := make([]byte, sizeofDIB)
			var offset uint32 = 0
			for offset = 0; offset < sizeofDIB; offset += 1 {
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
	if ok, errno := user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_FONTCHANGE, 0, 0); !ok {
		panic(fmt.Sprintf("%s", errno))
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
	_, _ = user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_SYSCOMMAND, w32.SC_MONITORPOWER, 2)
	time.Sleep(2 * time.Second)

	log.Println("screenOn")
	_, _ = user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_SYSCOMMAND, w32.SC_MONITORPOWER, ^uintptr(0)) // -1
	time.Sleep(2 * time.Second)

	log.Println("screenLowPower")
	_, _ = user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_SYSCOMMAND, w32.SC_MONITORPOWER, 1)
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

// https://stackoverflow.com/a/68845977/9935654
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

		if !user32dll.GetIconInfo(hIcon, &iInfo) {
			return
		}

		// Remember to release when you are not using the HBITMAP.
		defer func() {
			_ = gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmColor))
			_ = gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmMask))
		}()
	}

	// Create Menu
	var hMenu w32.HMENU
	{
		hMenu = user32dll.CreatePopupMenu()
		_, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1023, "Open")

		// 設定含有icon的Menu
		// _, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1024, "Hello") // 可以先指定string，再用SetMenuItemInfo添加icon或者直接在SetMenuItemInfo添加string或icon都可以
		var menuItemInfo w32.MENUITEMINFO
		pMsg, _ := syscall.UTF16PtrFromString("Hello")
		menuItemInfo = w32.MENUITEMINFO{
			CbSize: uint32(unsafe.Sizeof(menuItemInfo)),

			// FMask是一個開關，當有設定這些mask，某些欄位設定數值才會有意義
			FMask: w32.MIIM_BITMAP | // sets the hbmpItem member.
				w32.MIIM_ID | // sets the wID member.
				w32.MIIM_STRING, // sets the dwTypeData member.

			WID:        1024,
			DwTypeData: pMsg,
			HbmpItem:   iInfo.HbmColor,
		}
		_, _ = user32dll.SetMenuItemInfo(hMenu, 1024, false, &menuItemInfo)

		_, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1025, "Exit program")

		defer func() {
			if ok, errno := user32dll.DestroyMenu(hMenu); !ok {
				log.Printf("%s\n", errno)
			}
		}()
	}

	ch := make(chan w32.HWND)
	go func(chanWin chan<- w32.HWND) {
		// define WNDPROC
		wndProcFuncPtr := syscall.NewCallback(w32.WNDPROC(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
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
				if ok, errno := user32dll.GetCursorPos(&pos); !ok {
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
				case 1023:
					log.Println("open")
				case 1024:
					log.Println("hello")
				case 1025:
					log.Println("1025")
					_, _ = user32dll.PostMessage(hwnd, w32.WM_DESTROY, 0, 0)
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
			LpfnWndProc:   wndProcFuncPtr,
			HInstance:     hInstance,
			HIcon:         hIcon,
			LpszClassName: pUTF16ClassName,
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
			if ok, errno2 := user32dll.UnregisterClass(wndClassName, hInstance); !ok {
				fmt.Printf("Error UnregisterClass: %s", errno2)
			}
			chanWin <- hwnd
			return
		}

		// Make sure it can be unregistered when exiting.
		defer func() {
			if ok, errno2 := user32dll.UnregisterClass(wndClassName, hInstance); !ok {
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

	_, _ = user32dll.PostMessage(hwnd, w32.WM_RBUTTONDOWN, 0, 0)   // 選單測試
	_, _ = user32dll.PostMessage(hwnd, w32.WM_RBUTTONDOWN, 123, 0) // with TPM_RETURNCMD

	// 🕹️ 如果您要手動嘗試，請把以下的SendMessage.WM_CLOSE註解掉，避免自動關閉
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)

	// wait user close the window
	<-ch

	// Output:
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

	// 新建一個執行緒來專門處理視窗{建立、消息循環}
	go func(channel chan<- w32.HWND) {
		// define ProcFunc // https://learn.microsoft.com/en-us/windows/win32/learnwin32/writing-the-window-procedure
		wndProcFuncPtr := syscall.NewCallback(w32.WNDPROC(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
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
				if ok, errno := user32dll.DestroyWindow(hwnd); !ok {
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

		const className = "myClassName"
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(className)
		wc := w32.WNDCLASS{
			// Style:       0, // 可以不給，或者w32.CS_NOCLOSE禁用右上角的關閉按鈕) // CS_指的是class的style
			LpfnWndProc: wndProcFuncPtr, // 每次有消息，就會送通知到此函數
			// CbClsExtra:    0,
			// CbWndExtra:    0,
			HInstance: hInstance,
			HIcon:     user32dll.MustLoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION /* w32.IDI_WINLOGO */)), // 可以不給, 用預設0 會是: IDI_WINLOGO
			HCursor:   user32dll.MustLoadCursor(0, w32.MakeIntResource(w32.IDC_CROSS /* w32.IDC_ARROW */)),    // 可以不給, 用預設0 會是: IDC_ARROW
			// HbrBackground: 0,
			// LpszMenuName:  nil,
			LpszClassName: pUTF16ClassName,
		}

		if atom, errno := user32dll.RegisterClass(&wc); atom == 0 {
			fmt.Printf("%s", errno)
			return
		}

		var (
			hwnd  w32.HWND
			errno syscall.Errno
			ok    bool
		)

		defer func() {
			if ok, errno = user32dll.UnregisterClass(className, hInstance); !ok {
				fmt.Printf("[UnregisterClass] %s", errno)
			}
		}()

		fmt.Println("CreateWindowEx")
		const windowName = "myWindowName" // 視窗左上角的標題名稱
		if hwnd, errno = user32dll.CreateWindowEx(0,
			className,
			windowName,
			w32.WS_OVERLAPPEDWINDOW, // 這項包含了: WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX // 如果不想要最小和最大化按鈕要在這邊調整，而關閉按鈕則是需要透過class註冊的時候就設定要還是不要

			// Size and position
			w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,

			hwndParent, // 0, // Parent window
			0,          // Menu
			hInstance,
			uintptr(unsafe.Pointer(&AppData{"Demo-CreateWindowEx", 6})), // Additional application data // 可以不給(設定為0). 如果有給，這個資料會在WM_CREATE的時候傳入給lParam
		); hwnd == 0 {
			fmt.Printf("%s", errno)
			channel <- 0
			return
		} else {
			// test FindWindow
			{
				hwnd2 := user32dll.FindWindow("myClassName", windowName)
				log.Println(hwnd == hwnd2) // true
			}

			channel <- hwnd
		}

		// 消息循環
		var msg w32.MSG
		for {
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
				fmt.Println("===quit the window===")
				channel <- 0
				break
			}
			user32dll.TranslateMessage(&msg)
			user32dll.DispatchMessage(&msg)
		}
	}(ch)

	// 如果視窗成功被建立會傳送該hwnd
	hwnd := <-ch

	// 如果視窗建立失敗，會得到空hwnd，不做其他處理直接返回
	if hwnd == 0 {
		return
	}

	// 以下為模擬外層程式，向視窗發送訊息
	{
		fmt.Println("ShowWindow")
		user32dll.ShowWindow(hwnd, w32.SW_MAXIMIZE)

		fmt.Println("CloseWindow")                         // 僅是縮小視窗
		if ok, errno := user32dll.CloseWindow(hwnd); !ok { // close只是把它縮小並沒有真正關閉
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

	// time.Sleep(time.Second * 5) // 可以暫停一段時間，之後再終止，當您設定CS_NOCLOSE，需要自己去關閉視窗
	_, _, _ = user32dll.SendMessage(hwnd, w32.WM_DESTROY, 0, 0) // 如果您想要在視窗上進行操作，可以把這列註解，運行的時候再去手動關閉視窗即可結束

	<-ch // wait window close

	fmt.Print("bye")

	// Output:
	// CreateWindowEx
	// WM_NCCREATE
	// Demo-CreateWindowEx
	// 6
	// ShowWindow
	// CloseWindow
	// DestroyWindow
	// ===quit the window===
	// bye
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
		wndProcFuncPtr := syscall.NewCallback(w32.WNDPROC(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
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
					if ok, errno := user32dll.UnhookWindowsHookEx(d.HHOOK); !ok {
						fmt.Printf("%s", errno)
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
			LpfnWndProc:   wndProcFuncPtr,
			HInstance:     hInstance,
			LpszClassName: pUTF16ClassName,
		}); atom == 0 {
			fmt.Printf("%s", errno)
			chanWin <- 0
			return
		}

		// 確保程式結束之後能解除註冊名稱
		defer func() {
			if ok, errno2 := user32dll.UnregisterClass(wndClassName, hInstance); !ok {
				log.Printf("Error UnregisterClass: %s", errno2)
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
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
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
		wndProcFuncPtr := syscall.NewCallback(w32.WNDPROC(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			switch uMsg {
			case w32.WM_DESTROY:
				log.Println("WM_DESTROY")
				for _, hotkeyID := range []int32{HokeyIDHello, HokeyIDBye} {
					if ok, errno := user32dll.UnregisterHotKey(hwnd, hotkeyID); !ok {
						log.Printf("Error [UnregisterHotKey] %s", errno)
					}
				}
				log.Println("haha")
				user32dll.PostQuitMessage(0)
				return 0
			case w32.WM_CREATE:
				if ok, errno := user32dll.RegisterHotKey(hwnd, HokeyIDHello, w32.MOD_CONTROL|w32.MOD_ALT, w32.VK_F1); !ok {
					log.Printf("%s\n", errno)
				}
				if ok, errno := user32dll.RegisterHotKey(hwnd, HokeyIDBye, w32.MOD_ALT, 65 /* A */); !ok {
					log.Printf("%s\n", errno)
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
					_, _ = user32dll.PostMessage(hwnd, w32.WM_CLOSE, 0, 0)
				}
			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam) // default window proc
		}))

		hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(className)
		wc := w32.WNDCLASS{
			LpfnWndProc:   wndProcFuncPtr,
			HInstance:     hInstance,
			LpszClassName: pUTF16ClassName,
		}

		if atom, errno := user32dll.RegisterClass(&wc); atom == 0 {
			fmt.Printf("%s", errno)
			return
		}

		defer func() {
			if ok, errno := user32dll.UnregisterClass(className, hInstance); !ok {
				fmt.Printf("[UnregisterClass] %s", errno)
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
