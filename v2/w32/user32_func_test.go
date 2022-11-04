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
