package w32_test

import (
	"encoding/binary"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"os/exec"
	"time"
	"unsafe"
)

func ExampleRGB() {
	fmt.Println(0xff80ff == w32.RGB(255, 128, 255))
	// Output:
	// true
}

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

	ok, errno := user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_FONTCHANGE, 0, 0)
	if !ok {
		log.Fatal(fmt.Sprintf("%s", errno))
	}

	log.Println(numFont)
	// Output:
}

// https://www.codeguru.com/multimedia/how-to-use-a-font-without-installing-it/
func ExampleGdi32DLL_AddFontResourceEx() {
	ttfPath := "./testdata/fonts/teamviewer15.otf"
	gdi32dll := w32.NewGdi32DLL(w32.PNAddFontResourceEx, w32.PNRemoveFontResourceEx)
	user32dll := w32.NewUser32DLL(w32.PNPostMessage)

	// FR_PRIVATE 就只有自己本身程式(呼叫者)能用
	// 至於FR_NOT_ENUM，您可以先把flag設定為0，開啟notepad選該字體並再Remove之前先選中該字體，接著再移除
	// 完成後不要更換字體(此時選單已經選不到該字體)
	// 再執行本程式一次，把flag改為FR_NOT_ENUM
	// 同樣的也在remove之前下斷點
	// 會發現字型有改變，但選單然仍選不到該字體
	// 也就是FR_NOT_ENUM可以讓選單不出現字體，但如果字體已經有被加載過，在還沒有重開機(或登出)前再用FR_NOT_ENUM，還是能讓應用程式顯示到該字體(前提是您不能更換字體)
	var flag uint32 = w32.FR_NOT_ENUM // w32.FR_PRIVATE // RemoveFontResourceEx要與AddFontResourceEx所使用的flag一致
	for {
		// 刪除舊有的資料
		if !gdi32dll.RemoveFontResourceEx(ttfPath,
			flag,
			0,
		) {
			break
		}
	}

	numFont := gdi32dll.AddFontResourceEx(ttfPath,
		flag,
		0)
	if numFont == 0 {
		return
	}

	defer func() {
		// 注意，應該要用迴圈不斷執行，直到刪不到東西為止(假設您不曉得到底成功加入了多少字體)
		for i := 1; ; i++ {
			if !gdi32dll.RemoveFontResourceEx(ttfPath, flag, 0) {
				break
			}
			log.Println("RemoveFontResourceEx:", i)
		}
	}()

	ok, errno := user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_FONTCHANGE, 0, 0)
	if !ok {
		log.Fatal(fmt.Sprintf("%s", errno))
	}

	log.Println(numFont)
	// Output:
}

func ExampleNewFontMemResource() {
	kernel32dll := w32.NewKernel32DLL(w32.PNLoadLibrary)
	hExe := kernel32dll.LoadLibrary("./testdata/exe/writeWithFont.exe")
	fontMemResource, errno := w32.NewFontMemResource(hExe, w32.MakeIntResource(666)) // 該應用程式的RT_FONT資源下存在一個ID為666的字型檔案。實際上的ID代碼會依應用程式而定，非定值
	if fontMemResource == nil {
		log.Fatal(errno)
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
	user32dll := w32.NewUser32DLL()
	gdi32dll := w32.NewGdi32DLL()
	kernel32dll := w32.NewKernel32DLL()

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

	if uintptr(hFile) == w32.INVALID_HANDLE_VALUE {
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

// Example_saveFileIconAsBitmap
// Step:
// get HICON
// get ICONINFO from HICON
// get Bitmap from ICONINFO
// save Bitmap to a File
// Note: To make the example look simple, I ignore all possible errors handling.
func Example_saveFileIconAsBitmap() {
	user32dll := w32.NewUser32DLL()
	kernel32dll := w32.NewKernel32DLL()
	gdi32dll := w32.NewGdi32DLL()

	var hIcon w32.HICON
	{
		// Because I already know the iconID(myResourceID) I want, I can load it directly.
		// If you want to find the appropriate iconID after searching through RT_GROUP_ICON you can refer to this example:
		// https://github.com/CarsonSlovoka/go-pkg/blob/34e5d2c1fc97bf149bf626acaaf8773fe1509d64/v2/w32/kernel32_func_test.go#L331-L353

		hmExe := kernel32dll.LoadLibrary("./testdata/exe/writeWithFont.exe") // writeWithFont.exe is in here: https://github.com/CarsonSlovoka/go-pkg/tree/983a2c1/v2/w32/testdata/exe
		myResourceID := uintptr(1)                                           // You can use resourceHacker.exe to help you find the ID.
		hRes, _ := kernel32dll.FindResource(hmExe,
			w32.MakeIntResource(myResourceID),
			w32.MakeIntResource(w32.RT_ICON),
		)
		hMem, _ := kernel32dll.LoadResource(hmExe, hRes)
		lpResource := kernel32dll.LockResource(hMem)

		hIcon = user32dll.CreateIconFromResourceEx(lpResource,
			kernel32dll.MustSizeofResource(hmExe, hRes), true, 0x00030000,
			32, 32, // size X, Y
			w32.LR_DEFAULTCOLOR,
		)
	}

	var iInfo w32.ICONINFO
	{
		if !user32dll.GetIconInfo(hIcon, &iInfo) {
			return
		}
		// Remember to release when you are not using the HBITMAP.
		defer func() {
			_ = gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmColor))
			_ = gdi32dll.DeleteObject(w32.HGDIOBJ(iInfo.HbmMask))
		}()
	}

	bmp := w32.Bitmap{}
	{
		// Create BITMAP by ICONINFO
		gdi32dll.GetObject(w32.HANDLE(iInfo.HbmColor), int32(unsafe.Sizeof(bmp)), uintptr(unsafe.Pointer(&bmp)))
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
		bmpSize := ((bmp.Width*int32(bitmapInfoHeader.BitCount) + 31) / 32) * 4 /* uint32 */ * bmp.Height // see the wiki: https://en.wikipedia.org/wiki/BMP_file_format#Pixel_storage

		sizeofDIB := 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)) + uint32(bmpSize)
		bitmapFileHeader = w32.BitmapFileHeader{
			Type:       0x4D42,    // BM. // B: 42, M: 4D  //  All of the integer values are stored in little-endian format
			Size:       sizeofDIB, // HEADER + INFO + DATA
			OffsetBits: 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)),
		}

		hdc := user32dll.GetDC(0)

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
		outputBmpPath := "testdata/temp001.bmp"
		// Write: FileHeader, DIPHeader, bitmapData
		{
			f, _ := os.Create(outputBmpPath)
			defer func() {
				_ = os.Remove(outputBmpPath) // Remove test data. If you want to see the result, delete this line to see the final data.
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
}

func ExampleGdi32DLL_EnumFonts() {
	gdi32dll := w32.NewGdi32DLL()
	user32dll := w32.NewUser32DLL()
	var fontEnumProc w32.FONTENUMPROC
	fontEnumProc = func(lpLF *w32.LOGFONT, lpTM *w32.TEXTMETRIC, dwType uint32, lpData w32.LPARAM) int32 {
		log.Println(lpLF.GetFaceName())
		return 1
	}

	hwndTarget := user32dll.FindWindow("Notepad", "")
	if hwndTarget == 0 {
		log.Println("Notepad not found. Using the screen instead.")
	}
	hdcTarget := user32dll.GetDC(hwndTarget)
	defer func() {
		user32dll.ReleaseDC(hwndTarget, hdcTarget)
	}()
	// gdi32dll.EnumFonts(hdcTarget, "", fontEnumProc, 0) // List All
	gdi32dll.EnumFonts(hdcTarget, "Arial", fontEnumProc, 0) // 列出所有FaceName含Arial(宋體)的項目
	// Output:
}

func ExampleGdi32DLL_EnumFontFamilies() {
	gdi32dll := w32.NewGdi32DLL()
	user32dll := w32.NewUser32DLL()
	hdc := user32dll.GetDC(0)
	defer func() {
		user32dll.ReleaseDC(0, hdc)
	}()

	var enumFontFamProc w32.EnumFontFamProc
	enumFontFamProc = func(logFont *w32.ENUMLOGFONT, textMetric *w32.TEXTMETRIC, fontType uint32, lParam w32.LPARAM) int32 {
		log.Println(logFont.GetFullName(), "|", logFont.GetStyle(), "|",
			logFont.LogFont.GetFaceName(), "|",
			"W:", logFont.LogFont.LfWeight,
			"Italic:", logFont.LogFont.IsItalic(),
			"Strike:", logFont.LogFont.IsStrikeOut(),
			"U:", logFont.LogFont.IsUnderline())
		return 1
	}

	// gdi32dll.EnumFontFamilies(hdc, "", enumFontFamProc, 0) // Enum All
	gdi32dll.EnumFontFamilies(hdc, "Arial", enumFontFamProc, 0)

	// Output:
}
