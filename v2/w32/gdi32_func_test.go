package w32_test

import (
	"encoding/binary"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

// 添加字型，不需要安裝。重開機或執行RemoveFontResource將會被移除
func ExampleGdi32DLL_AddFontResource() {
	ttfPath := "./testdata/fonts/teamviewer15.otf"
	gdi32dll := w32.NewGdi32DLL(w32.PNAddFontResource, w32.PNRemoveFontResource)
	user32dll := w32.NewUser32DLL(w32.PNPostMessage)
	numFont := gdi32dll.AddFontResource(ttfPath)
	if numFont == 0 {
		return
	}

	defer func() {
		if ok := gdi32dll.RemoveFontResource(ttfPath); ok == 0 {
			log.Fatal("error RemoveFontResource")
		}
	}()

	_, _, err := user32dll.PostMessage(uintptr(w32.HWND_BROADCAST), w32.WM_FONTCHANGE, 0, 0)
	if err != syscall.Errno(0x0) {
		log.Fatal(err)
	}

	log.Println(numFont)
	// Output:
}

// https://www.codeguru.com/multimedia/how-to-use-a-font-without-installing-it/
func ExampleGdi32DLL_AddFontResourceEx() {
	ttfPath := "./testdata/fonts/teamviewer15.otf"
	gdi32dll := w32.NewGdi32DLL(w32.PNAddFontResourceEx, w32.PNRemoveFontResourceEx)
	user32dll := w32.NewUser32DLL(w32.PNPostMessage)
	numFont := gdi32dll.AddFontResourceEx(ttfPath,
		w32.FR_NOT_ENUM, // 若使用FR_PRIVATE程式結束會自動刪除，同時FR_PRIVATE沒有辦法讓其他應用程式訪問到該字型，即其他應用程式沒辦法選到該字型；但是FR_NOT_ENUM可以讓其他應用程式選到該字型
		0)
	if numFont == 0 {
		return
	}

	defer func() {
		if err := gdi32dll.RemoveFontResourceEx(ttfPath,
			w32.FR_NOT_ENUM, // flag 要與AddFontResourceEx所使用的flag一致
			0,
		); err == 0 {
			log.Fatal(err)
		}
	}()

	_, _, err := user32dll.PostMessage(uintptr(w32.HWND_BROADCAST), w32.WM_FONTCHANGE, 0, 0)
	if err != syscall.Errno(0x0) {
		log.Fatal(err)
	}

	log.Println(numFont)
	// Output:
}

func ExampleNewFontMemResource() {
	kernel32dll := w32.NewKernel32DLL(w32.PNLoadLibrary)
	hExe := kernel32dll.LoadLibrary("./testdata/exe/writeWithFont.exe")
	fontMemResource, err := w32.NewFontMemResource(hExe, w32.MakeIntResource(666)) // 該應用程式的RT_FONT資源下存在一個ID為666的字型檔案。實際上的ID代碼會依應用程式而定，非定值
	if err != nil {
		panic(err)
	}
	defer fontMemResource.Remove()
	fmt.Println("ok")
	// Output:
	// ok
}

// CaptureAnImage https://learn.microsoft.com/en-us/windows/win32/gdi/capturing-an-image
// https://zh.wikipedia.org/zh-tw/BMP
// 本範例簡述: 抓取當前的視窗，畫在notepad上，之後再保存在檔案之中，完成後檔案(testdata/captureNotepad.bmp)會刪除
func ExampleGdi32DLL_CreateCompatibleBitmap() {
	user32dll := w32.NewUser32DLL(
		w32.PNFindWindow,
		w32.PNGetDC,
		w32.PNGetClientRect,
		w32.PNGetSystemMetrics,
		w32.PNReleaseDC,
	)
	gdi32dll := w32.NewGdi32DLL(
		w32.PNCreateCompatibleDC,
		w32.PNSetStretchBltMode,
		w32.PNStretchBlt,
		w32.PNSelectObject,
		w32.PNCreateCompatibleBitmap,
		w32.PNBitBlt,
		w32.PNGetDIBits,
		w32.PNGetObject,
		w32.PNDeleteObject,
	)
	kernel32dll := w32.NewKernel32DLL(
		w32.PNGlobalAlloc,
		w32.PNGlobalLock,
		w32.PNCreateFile,
		w32.PNWriteFile,
		w32.PNGlobalUnlock,
		w32.PNGlobalFree,
		w32.PNCloseHandle,
	)

	var hwndNotepad w32.HWND
	hwndNotepad = user32dll.FindWindow("Notepad", "")
	if hwndNotepad == 0 {
		if err := exec.Command("notepad.exe").Start(); err != nil {
			log.Println(err)
			fmt.Println("ok")
			return
		}
		time.Sleep(200 * time.Millisecond) // waiting for the notepad.exe to open
		hwndNotepad = user32dll.FindWindow("Notepad", "")
		if hwndNotepad == 0 {
			log.Println("notepad.exe not found.")
			fmt.Println("ok")
			return
		}
	}
	hdcScreen := user32dll.GetDC(0)
	defer user32dll.ReleaseDC(0, hdcScreen)
	hdcNotepad := user32dll.GetDC(hwndNotepad)
	defer user32dll.ReleaseDC(hwndNotepad, hdcNotepad)

	// Create a compatible DC, which is used in a BitBlt from the window DC
	// memoryDC
	hdcMemNotepad := gdi32dll.CreateCompatibleDC(hdcNotepad)
	if hdcMemNotepad == 0 {
		log.Fatal("CreateCompatibleDC has failed")
	}
	defer gdi32dll.DeleteObject(w32.HGDIOBJ(hdcMemNotepad))

	// Get the client area for size calculation.
	var rcClient w32.RECT
	if ok, errno := user32dll.GetClientRect(hwndNotepad, &rcClient); !ok {
		log.Fatalf("GetClientRect errno:%d\n", errno)
	}

	// 把screen的圖畫在notepad上
	// 完成hdcNotepad的內容
	{
		// This is the best stretch mode.
		// 對HDC設定BltMode, 這個設定最接近原圖
		gdi32dll.SetStretchBltMode(hdcNotepad, w32.HALFTONE)

		// 將src的資源傳輸到dst中去
		if !gdi32dll.StretchBlt(hdcNotepad,
			0, 0, rcClient.Right, rcClient.Bottom,
			hdcScreen,
			0, 0,
			user32dll.GetSystemMetrics(w32.SM_CXSCREEN), user32dll.GetSystemMetrics(w32.SM_CYSCREEN),
			w32.SRCCOPY,
		) {
			log.Fatal("StretchBlt has failed")
		}
	}

	// 建立HBITMAP (由HDC來幫忙建立)
	var hbmNotepad w32.HBITMAP
	{
		// Create a compatible bitmap from the Window DC.
		// 我們可以再對該hdc做篩選(挑選出想要的地方)
		hbmNotepad = gdi32dll.CreateCompatibleBitmap(hdcNotepad,
			rcClient.Right-rcClient.Left,
			rcClient.Bottom-rcClient.Top,
		)
		if hbmNotepad == 0 {
			log.Fatal("CreateCompatibleBitmap Failed")
		}
		defer gdi32dll.DeleteObject(w32.HGDIOBJ(hbmNotepad))
	}

	// 完成HBITMAP的內容
	// 1. hdcMem選擇HBITMAP
	// 2. 將hdc的內容傳送到hdcMem上
	{
		// Select the compatible bitmap into the compatible memory DC.
		gdi32dll.SelectObject(hdcMemNotepad, w32.HGDIOBJ(hbmNotepad))

		// Bit block transfer into our compatible memory DC.
		if ok, errno := gdi32dll.BitBlt(hdcMemNotepad,
			0, 0,
			rcClient.Right-rcClient.Left, rcClient.Bottom-rcClient.Top,
			hdcNotepad,
			0, 0,
			w32.SRCCOPY); !ok {
			log.Fatalf("Bit-block has failed. errno: %s\n", errno) // errno有個好處，它有處理字串，所以不需要用%d，不然還要再去查數字的意思反而麻煩。
		}
	}

	// Get the BITMAP from the HBITMAP.
	var bmpNotepad w32.Bitmap
	// 透過HBITMAP來建立BITMAP
	gdi32dll.GetObject(w32.HANDLE(hbmNotepad), int32(unsafe.Sizeof(bmpNotepad)), uintptr(unsafe.Pointer(&bmpNotepad)))

	var bitmapInfoHeader w32.BitmapInfoHeader
	bitmapInfoHeader = w32.BitmapInfoHeader{
		Size:  uint32(unsafe.Sizeof(bitmapInfoHeader)), // 也可以直接寫40
		Width: bmpNotepad.Width, Height: bmpNotepad.Height,
		Planes:      1,
		BitCount:    32,
		Compression: w32.BI_RGB,
		// XPelsPerMeter: 2400, // 可以用0就好
		// YPelsPerMeter: 2400, // 可以用0就好
	}

	// 透過公式計算大小: https://en.wikipedia.org/wiki/BMP_file_format#Pixel_storage
	bmpSize := ((bmpNotepad.Width*int32(bitmapInfoHeader.BitCount) + 31) / 32) * 4 /* uint32 */ * bmpNotepad.Height // size 2682368 bytes => 2619KB

	hDIB, _ := kernel32dll.GlobalAlloc(w32.GHND, w32.SIZE_T(bmpSize))
	// Unlock and Free the DIB from the heap.
	defer func() {
		kernel32dll.GlobalUnlock(hDIB)
		kernel32dll.GlobalFree(hDIB)
	}()

	// 找到bitmap的資料起始位置lpBitmap
	var lpBitmap w32.LPVOID
	lpBitmap, _ = kernel32dll.GlobalLock(hDIB)
	// Gets the "bits" from the bitmap, and copies them into a buffer
	// that's pointed to by lpbitmap.
	gdi32dll.GetDIBits(
		hdcNotepad, hbmNotepad, 0,
		w32.UINT(bmpNotepad.Height),
		lpBitmap, // [out]
		&w32.BitmapInfo{Header: bitmapInfoHeader},
		w32.DIB_RGB_COLORS,
	)

	// Add the size of the headers to the size of the bitmap to get the total file size.
	var bitmapFileHeader w32.BitmapFileHeader
	// 注意uint32(unsafe.Sizeof(bitmapFileHeader))算出來的會是16，正確的應該是14
	sizeofDIB := 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)) + uint32(bmpSize)
	bitmapFileHeader = w32.BitmapFileHeader{
		Type:       0x4D42,    // BM. // B: 42, M: 4D  // 因為BitmapFile所有的描述都要用"little-endian"讀取，所以要反過來寫4D42
		Size:       sizeofDIB, // HEADER + INFO + DATA
		OffsetBits: 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)),
	}

	outputBmpPath := "testdata/captureNotepad.bmp"
	// 寫檔，寫入FileHeader, DIPHeader, bitmapData
	{
		f, err := os.Create(outputBmpPath)
		if err != nil {
			log.Fatal(err)
		}
		// FileHeader
		_ = binary.Write(f, binary.LittleEndian, bitmapFileHeader)

		// DIP Header
		_ = binary.Write(f, binary.LittleEndian, bitmapInfoHeader)

		// 其實可以直接透過以下這段把數值也順便寫入，即可完成。但我們因為要展示kernel32dll.CreateFile，所以寫入data的部分還是交由它去完成
		if false {
			// bitmapData
			bmpDatas := make([]byte, sizeofDIB)
			var offset uint32
			for offset = 0; offset < sizeofDIB; offset++ {
				curByteAddr := unsafe.Pointer(uintptr(lpBitmap) + uintptr(offset)) // 計算當前要寫入的byte位址在哪 // 我們是一個byte一個byte寫入，所以大小都是1
				bmpDatas[offset] = *(*byte)(curByteAddr)
			}
			_ = binary.Write(f, binary.LittleEndian, bmpDatas)
		}

		_ = f.Close()
	}

	// 上述故意少寫了bitmapData，以下透過kernel32dll.CreateFile來寫入資料
	// A file is created, this is where we will save the screen capture.
	hFile, errno := kernel32dll.CreateFile(outputBmpPath,
		w32.FILE_APPEND_DATA, // w32.GENERIC_WRITE <-- 用這個會新建，會把舊的資料刪除
		0,
		0,
		w32.OPEN_ALWAYS, // w32.CREATE_ALWAYS 用這個也會把舊的資料刪除
		w32.FILE_ATTRIBUTE_NORMAL,
		0,
	)

	if int(hFile) == w32.INVALID_HANDLE_VALUE {
		log.Fatalf("%s\n", errno)
	}

	defer func() {
		if err := os.Remove(outputBmpPath); err != nil {
			log.Printf("delete test file. %q", outputBmpPath)
		}
	}()

	// 以下這種寫法和kernel32dll.WriteFile寫出來的內容是一樣的
	// // (*(*[size]byte)(unsafe.Pointer(&structValue)))[:]
	// _ = binary.Write(f, binary.LittleEndian, (*(*[14]byte)(unsafe.Pointer(&bitmapFileHeader)))[:])
	// _ = binary.Write(f, binary.LittleEndian, (*(*[uint32(unsafe.Sizeof(bitmapInfoHeader))]byte)(unsafe.Pointer(&bitmapInfoHeader)))[:])

	var dwBytesWritten uint32
	// FILE HEADER 不行用以下的方法寫，會有endian的問題
	// _, _ = kernel32dll.WriteFile(hFile, uintptr(unsafe.Pointer(&bitmapFileHeader)), 14, &dwBytesWritten, nil)
	// DIP HEADER 不行用以下的方法寫，會有endian的問題
	// _, _ = kernel32dll.WriteFile(hFile, uintptr(unsafe.Pointer(&bitmapInfoHeader)), uint32(unsafe.Sizeof(bitmapInfoHeader)), &dwBytesWritten, nil)
	// DATA
	_, _ = kernel32dll.WriteFile(hFile, uintptr(lpBitmap), sizeofDIB, &dwBytesWritten, nil)
	_, _ = kernel32dll.CloseHandle(hFile)

	fmt.Println("ok")

	// Output:
	// ok
}
