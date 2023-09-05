package w32_test

import (
	"encoding/binary"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"os/exec"
	"reflect"
	"sync"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func ExampleRGB() {
	rgb := w32.RGB(0xff, 0x80, 0x40) // 255, 128, 64
	fmt.Println(rgb.R(), rgb.G(), rgb.B())
	fmt.Println(0x4080ff == rgb)
	fmt.Println(w32.GetRValue(rgb), w32.GetGValue(rgb), w32.GetBValue(rgb))

	// Output:
	// 255 128 64
	// true
	// 255 128 64
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

	errno := user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_FONTCHANGE, 0, 0)
	if errno != 0 {
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

	if errno := user32dll.PostMessage(w32.HWND_BROADCAST, w32.WM_FONTCHANGE, 0, 0); errno != 0 {
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
	if en := user32dll.GetClientRect(hwndNotepad, &rcClient); en != 0 {
		log.Fatalf("GetClientRect errno:%d\n", en)
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
		if en := gdi32dll.BitBlt(hdcMemNotepad,
			0, 0,
			rcClient.Right-rcClient.Left, rcClient.Bottom-rcClient.Top,
			hdcNotepad,
			0, 0,
			w32.SRCCOPY); en != 0 {
			log.Fatalf("Bit-block has failed. errno: %s\n", en) // errno有個好處，它有處理字串，所以不需要用%d，不然還要再去查數字的意思反而麻煩。
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
	defer kernel32dll.GlobalFree(hDIB)

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
	_, _ = kernel32dll.GlobalUnlock(hDIB)

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
	// 寫檔，寫入FileHeader, DIBHeader (Device Independent Bitmap), bitmapData
	{
		f, err := os.Create(outputBmpPath)
		if err != nil {
			log.Fatal(err)
		}
		// FileHeader
		_ = binary.Write(f, binary.LittleEndian, bitmapFileHeader)

		// DIB Header
		_ = binary.Write(f, binary.LittleEndian, bitmapInfoHeader)

		// 其實可以直接透過以下這段把數值也順便寫入，即可完成。但我們因為要展示kernel32dll.CreateFile，所以寫入data的部分還是交由它去完成
		if false {
			// bitmapData
			bmpData := make([]byte, bmpSize)
			var offset uint32
			for offset = 0; offset < uint32(bmpSize); offset++ {
				curByteAddr := unsafe.Pointer(uintptr(lpBitmap) + uintptr(offset)) // 計算當前要寫入的byte位址在哪 // 我們是一個byte一個byte寫入，所以大小都是1
				bmpData[offset] = *(*byte)(curByteAddr)
			}

			/* 如果不想要用for慢慢一個一個給，可以用以下的方法一次賦值完畢
			sliceHeader := reflect.SliceHeader{
				Data: uintptr(lpBitmap),
				Len:  int(bmpSize),
				Cap:  int(bmpSize),
			}
			bmpData := *(*[]byte)(unsafe.Pointer(&sliceHeader))
			*/

			_ = binary.Write(f, binary.LittleEndian, bmpData)
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
	// DIB HEADER 不行用以下的方法寫，會有endian的問題
	// _, _ = kernel32dll.WriteFile(hFile, uintptr(unsafe.Pointer(&bitmapInfoHeader)), uint32(unsafe.Sizeof(bitmapInfoHeader)), &dwBytesWritten, nil)
	// DATA
	_ = kernel32dll.WriteFile(hFile, uintptr(lpBitmap), uint32(bmpSize), &dwBytesWritten, nil)
	_ = kernel32dll.CloseHandle(hFile)

	fmt.Println("ok")

	// Output:
	// ok
}

func saveHBitmap(outputPath string, hBitmap w32.HBITMAP) error {
	var bitmap w32.Bitmap
	gdiDll.GetObject(w32.HANDLE(hBitmap), int32(unsafe.Sizeof(bitmap)), uintptr(unsafe.Pointer(&bitmap)))
	hdc := userDll.GetDC(0)
	defer userDll.ReleaseDC(0, hdc)
	hdcMem := gdiDll.CreateCompatibleDC(hdc)
	defer gdiDll.DeleteDC(hdcMem)
	gdiDll.SelectObject(hdcMem, w32.HGDIOBJ(hBitmap))

	bitCount := uint16(32)
	bmpSize := ((bitmap.Width*int32(bitCount) + 31) / 32) * 4 * bitmap.Height
	hDIB, _ := kernelDll.GlobalAlloc(w32.GHND, w32.SIZE_T(bmpSize))
	defer kernelDll.GlobalFree(hDIB)

	var lpBitmap w32.LPVOID
	lpBitmap, _ = kernelDll.GlobalLock(hDIB)

	bitmapInfo := &w32.BitmapInfo{Header: w32.BitmapInfoHeader{
		Size:  40,
		Width: bitmap.Width, Height: bitmap.Height,
		Planes:      1,
		BitCount:    bitCount,
		Compression: w32.BI_RGB,
	}}

	gdiDll.GetDIBits(
		hdcMem, hBitmap, 0,
		w32.UINT(bitmap.Height),
		lpBitmap, // [out]
		bitmapInfo,
		w32.DIB_RGB_COLORS,
	)
	_, _ = kernelDll.GlobalUnlock(hDIB)

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_ = binary.Write(f, binary.LittleEndian, w32.BitmapFileHeader{
		Type:       0x4D42,
		Size:       14 + 40 + uint32(bmpSize), // HEADER + INFO + DATA
		OffsetBits: 14 + 40,
	})
	_ = binary.Write(f, binary.LittleEndian, bitmapInfo.Header)

	bmpData := make([]byte, bmpSize)
	sliceHeader := reflect.SliceHeader{
		Data: uintptr(lpBitmap),
		Len:  int(bmpSize),
		Cap:  int(bmpSize),
	}
	bmpData = *(*[]byte)(unsafe.Pointer(&sliceHeader))
	_, err = f.Write(bmpData)
	return err
}

func ExampleGdi32DLL_CreateDIBSection() {
	width := int32(30)
	height := int32(50)
	bmi := w32.BitmapInfo{
		Header: w32.BitmapInfoHeader{
			Size:     40,
			Width:    width,
			Height:   height,
			Planes:   1,
			BitCount: 32,
		},
	}
	var lpBits unsafe.Pointer
	hBitmap := gdiDll.CreateDIBSection(0, &bmi, w32.DIB_RGB_COLORS, &lpBits, 0, 0)
	bmpSize := ((width*int32(bmi.Header.BitCount) + 31) / 32) * 4 * height
	pixels := make([]byte, bmpSize)
	sliceHeader := reflect.SliceHeader{
		Data: uintptr(lpBits),
		Len:  int(bmpSize),
		Cap:  int(bmpSize),
	}
	pixels = *(*[]byte)(unsafe.Pointer(&sliceHeader))

	for i := int32(0); i < bmpSize; i += 4 {
		copy(pixels[i:i+4], []byte{
			// b, g, r, a
			255, 255, 0, 0,
		}[:])
	}

	if err := saveHBitmap("testdata/test4.bmp", hBitmap); err == nil {
		_ = os.Remove("testdata/test4.bmp")
	}
	// Output:
}

func Example_savePixelsAsImage() {
	width := int32(30)
	height := int32(50)
	bitCount := uint16(32)
	bmpSize := ((width*int32(bitCount) + 31) / 32) * 4 * height
	pixels := make([]byte, bmpSize)
	var x, y int32
	var i int32
	var b byte
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			if y < height/2 {
				b = 255
			} else {
				b = 0
			}
			copy(pixels[i:i+4], []byte{
				// b, g, r, a
				b, 255, 0, 0,
			}[:])
			i += 4
		}
	}
	outputPath := "testdata/test5.bmp"
	f, _ := os.Create(outputPath)
	defer func() {
		_ = f.Close()
		_ = os.Remove(outputPath)
	}()
	_ = binary.Write(f, binary.LittleEndian, w32.BitmapFileHeader{
		Type:       0x4D42,
		Size:       14 + 40 + uint32(bmpSize), // HEADER + INFO + DATA
		OffsetBits: 14 + 40,
	})
	_ = binary.Write(f, binary.LittleEndian, w32.BitmapInfoHeader{
		Size:     40,
		Width:    width,
		Height:   height,
		Planes:   1,
		BitCount: bitCount,
	})
	_, _ = f.Write(pixels)

	// Output:
}

func TestGdi32DLL_CreateDIBSection(t *testing.T) {
	width := int32(30)
	height := int32(50)
	bmi := w32.BitmapInfo{
		Header: w32.BitmapInfoHeader{
			Size:     40,
			Width:    width,
			Height:   height,
			Planes:   1,
			BitCount: 32,
		},
	}
	var lpBits unsafe.Pointer
	hBitmap := gdiDll.CreateDIBSection(0, &bmi, w32.DIB_RGB_COLORS, &lpBits, 0, 0)

	// 設定顏色
	var (
		r, g, b, a byte
		x, y       int32
	)
	bmpSize := ((width*int32(bmi.Header.BitCount) + 31) / 32) * 4 * height

	// pixels := (*[1 << 30]byte)(unsafe.Pointer(lpBits)) // 這可行，可以直接指到一個大的區塊，理論上如果圖片沒有那麼大是可行的
	pixels := make([]byte, bmpSize)
	// pixels = *(*[]byte)(unsafe.Pointer(&lpBits)) // 錯誤，會不知道界線在哪,所以要透過sliceHeader告知長度來幫忙
	sliceHeader := reflect.SliceHeader{
		Data: uintptr(lpBits),
		Len:  int(bmpSize),
		Cap:  int(bmpSize),
	}
	pixels = *(*[]byte)(unsafe.Pointer(&sliceHeader))

	i := int32(0)
	for y = int32(0); y <= height; y++ {
		if i >= bmpSize {
			break
		}
		for x = int32(0); x < width; x++ {
			if y < height/2 {
				r = 255
				g = 0
				b = 0
				a = 0
			} else {
				r = 255
				g = 255
				b = 0
				a = 0
			}
			if x > width/2 {
				b = 128
			}
			// copy(pixels[i:i+4], []byte{b, g, r, a}[:]) // 也可以用copy來幫忙
			pixels[i+0] = b
			pixels[i+1] = g
			pixels[i+2] = r
			pixels[i+3] = a
			i += 4
		}
	}

	// 以下為存檔，一種直接寫入數據資料，另一種透過HBITMAP來存檔
	outputFile1Path := "testdata/temp1.png"
	outputFile2Path := "testdata/temp2.png"
	{
		f, err := os.Create(outputFile1Path)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			_ = f.Close()
			_ = os.Remove(outputFile1Path)
		}()
		_ = binary.Write(f, binary.LittleEndian, w32.BitmapFileHeader{
			Type:       0x4D42,
			Size:       14 + 40 + uint32(bmpSize), // HEADER + INFO + DATA
			OffsetBits: 14 + 40,
		})
		_ = binary.Write(f, binary.LittleEndian, w32.BitmapInfoHeader{
			Size:        40,
			Width:       width,
			Height:      height,
			Planes:      1,
			BitCount:    32,
			Compression: w32.BI_RGB,
		})
		_, _ = f.Write(pixels)
	}

	// 這是另一種存檔方法
	if err := saveHBitmap(outputFile2Path, hBitmap); err != nil {
		t.Fatal(err)
	} else {
		_ = os.Remove(outputFile2Path)
	}
}

// 以鼠標為中心，依據輸入範圍，將矩形框內的範圍複製到剪貼簿之中
func Example_copy2clipboard() {
	// 鼠標中心矩形的框的範圍
	var rectSize int32 = 100

	var cursorPos w32.POINT
	if errno := userDll.GetCursorPos(&cursorPos); errno != 0 {
		log.Printf("GetCursorPos %s", errno)
	}

	// 計算矩形左上角和右下角的座標
	centerX := cursorPos.X
	centerY := cursorPos.Y
	left := centerX - rectSize
	top := centerY - rectSize
	right := centerX + rectSize
	bottom := centerY + rectSize

	screenDC := userDll.GetDC(0)
	defer userDll.ReleaseDC(0, screenDC)
	memDC := gdiDll.CreateCompatibleDC(screenDC)
	defer gdiDll.DeleteObject(w32.HGDIOBJ(memDC))

	hBitmapMem := gdiDll.CreateCompatibleBitmap(screenDC, right-left, bottom-top)
	defer gdiDll.DeleteObject(w32.HGDIOBJ(hBitmapMem))

	gdiDll.SelectObject(memDC, w32.HGDIOBJ(hBitmapMem))

	// 複製到位圖之中
	_ = gdiDll.BitBlt(memDC, 0, 0, right-left, bottom-top, screenDC, left, top, w32.SRCCOPY)

	_ = userDll.OpenClipboard(0)

	_ = userDll.EmptyClipboard()

	if _, eno := userDll.SetClipboardData(w32.CF_BITMAP, w32.HANDLE(hBitmapMem)); eno != 0 {
		log.Println(eno)
	}

	_ = userDll.CloseClipboard()

	// Output:
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
		if user32dll.GetIconInfo(hIcon, &iInfo) != 0 {
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
		// 列資料對齊: 每一列數據資料都以4byte對齊，如果不能被4整除，就會在結尾填充0直到能被整除
		bmpSize := ((bmp.Width*int32(bitmapInfoHeader.BitCount) + 31) / 32) * 4 /* uint32 */ * bmp.Height // see the wiki: https://en.wikipedia.org/wiki/BMP_file_format#Pixel_storage

		sizeofDIB := 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)) + uint32(bmpSize)
		bitmapFileHeader = w32.BitmapFileHeader{
			Type:       0x4D42,    // BM. // B: 42, M: 4D  //  All the integer values are stored in little-endian format
			Size:       sizeofDIB, // HEADER + INFO + DATA
			OffsetBits: 14 + uint32(unsafe.Sizeof(bitmapInfoHeader)),
		}

		hdc := user32dll.GetDC(0)

		var lpBitmap w32.LPVOID
		hDIB, _ := kernel32dll.GlobalAlloc(w32.GHND, w32.SIZE_T(bmpSize))
		lpBitmap, _ = kernel32dll.GlobalLock(hDIB)
		defer kernel32dll.GlobalFree(hDIB)

		gdi32dll.GetDIBits(
			hdc, iInfo.HbmColor,
			0,
			w32.UINT(bmp.Height),
			lpBitmap, // [out]
			&w32.BitmapInfo{Header: bitmapInfoHeader},
			w32.DIB_RGB_COLORS,
		)
		count, _ := kernel32dll.GlobalUnlock(hDIB)
		_ = count
		outputBmpPath := "testdata/temp001.bmp"
		// Write: FileHeader, DIBHeader, bitmapData
		{
			f, _ := os.Create(outputBmpPath)
			defer func() {
				_ = os.Remove(outputBmpPath) // Remove test data. If you want to see the result, delete this line to see the final data.
			}()

			// FileHeader
			_ = binary.Write(f, binary.LittleEndian, bitmapFileHeader)

			// DIB Header
			_ = binary.Write(f, binary.LittleEndian, bitmapInfoHeader)

			// bitmapData
			bmpData := make([]byte, bmpSize)
			for offset := uint32(0); offset < uint32(bmpSize); offset += 1 {
				curByteAddr := unsafe.Pointer(uintptr(lpBitmap) + uintptr(offset))
				bmpData[offset] = *(*byte)(curByteAddr)
			}
			_ = binary.Write(f, binary.LittleEndian, bmpData)

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
	gdi32dll.EnumFonts(hdcTarget, "Arial", fontEnumProc, 0) // 列出所有FaceName含Arial(sans-serif)的項目
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
			"W:", logFont.LogFont.Weight,
			"Italic:", logFont.LogFont.IsItalic(),
			"Strike:", logFont.LogFont.IsStrikeOut(),
			"U:", logFont.LogFont.IsUnderline())
		return 1
	}

	// gdi32dll.EnumFontFamilies(hdc, "", enumFontFamProc, 0) // Enum All
	gdi32dll.EnumFontFamilies(hdc, "Arial", enumFontFamProc, 0)

	// Output:
}

func ExampleGdi32DLL_CreateFont() {
	gdi32dll := w32.NewGdi32DLL()
	user32dll := w32.NewUser32DLL()
	hwnd := user32dll.FindWindow("Notepad", "")
	hdc := user32dll.GetDC(hwnd)
	defer func() {
		user32dll.ReleaseDC(hwnd, hdc)
	}()

	var (
		hFontSystem w32.HFONT
		hFontArial  w32.HFONT
	)

	// hFont: System 使用CreateFont來建立
	{
		hFontSystem = gdi32dll.CreateFont(
			48, 0, 0, 0,
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
			if gdi32dll.DeleteObject(w32.HGDIOBJ(hFontSystem)) {
				log.Println("DeleteObject HFONT")
			}
		}()
	}

	// hFont: Arial 使用CreateFontIndirect來建立
	{
		var logFont w32.LOGFONT
		{
			gdi32dll.EnumFonts(hdc, "Arial",
				func(lpLF *w32.LOGFONT, lpTM *w32.TEXTMETRIC, dwType uint32, lpData w32.LPARAM) int32 {
					logFont = *lpLF
					return 0 // DO NOT CONTINUE // 找到一個就停止
				},
				0)
		}
		logFont.Height = 66 // 可以再做調整
		hFontArial = gdi32dll.CreateFontIndirect(&logFont)
		defer func() {
			if gdi32dll.DeleteObject(w32.HGDIOBJ(hFontArial)) {
				log.Println("DeleteObject HFONT")
			}
		}()
	}

	gdi32dll.SetBkColor(hdc, w32.RGB(128, 128, 128))
	gdi32dll.SetBkMode(hdc, w32.OPAQUE) // 預設就是不透明，可以不用設定
	// gdi32dll.SetBkMode(hdc, w32.TRANSPARENT) // 如果設定成透明，那麼BkColor的顏色不會顯示

	var rect w32.RECT
	gdi32dll.SelectObject(hdc, w32.HGDIOBJ(hFontArial))
	_ = user32dll.GetClientRect(hwnd, &rect)
	user32dll.DrawText(hdc, "Hello World 您好 世界", -1, &rect, w32.DT_NOCLIP)

	gdi32dll.SelectObject(hdc, w32.HGDIOBJ(hFontSystem))
	user32dll.SetRect(&rect, 100, 100, 700, 200)
	if gdi32dll.SetTextColor(hdc, w32.RGB(0, 255, 255)) == w32.CLR_INVALID {
		fmt.Println("SetTextColor error")
	}
	user32dll.DrawText(hdc, "Hello World 您好 世界", -1, &rect, w32.DT_NOCLIP)

	// Output:
}

// 點擊滑鼠左鍵，可以顯示當前鼠標位置其像素的顏色
// 點擊右鍵可以終止程式
// 可以在此網頁點擊不同的顏色進行測試: https://www.w3schools.com/colors/colors_rgb.asp
func ExampleGdi32DLL_GetPixel() {
	gdi32dll := w32.NewGdi32DLL()
	user32dll := w32.NewUser32DLL()

	// 🕹️ 如果您要運行久一點，請調整此常數
	const inputRunTimeSecond = 2

	var (
		hwnd          w32.HWND
		hdc           w32.HDC
		hMemDC        w32.HDC
		hBitmapMem    w32.HBITMAP
		width, height int32
	)

	// init
	{
		hwnd = user32dll.GetDesktopWindow()
		hdc = user32dll.GetDC(hwnd)
		defer user32dll.ReleaseDC(hwnd, hdc)

		hMemDC = gdi32dll.CreateCompatibleDC(hdc)
		defer func() {
			if gdi32dll.DeleteDC(hMemDC) {
				log.Println("DeleteDC OK")
			}
		}()
		var rect w32.RECT
		if errno := user32dll.GetWindowRect(hwnd, &rect); errno != 0 {
			fmt.Printf("%s\n", errno)
			return
		}
		// hBitmapMem = gdi32dll.CreateCompatibleBitmap(hdc, rect.Width(), rect.Height()) // 需要初始化width, height, 這兩個屬性，此範例是全域變數，沒有設定在GetPixel之中會都是0，導致抓不到東西
		width = rect.Width()
		height = rect.Height()
		hBitmapMem = gdi32dll.CreateCompatibleBitmap(hdc, width, height)
		defer func() {
			if gdi32dll.DeleteObject(w32.HGDIOBJ(hBitmapMem)) {
				log.Println("Delete hBitmapMem OK")
			}
		}()
	}

	GetPixel := func(x, y int32) (w32.COLORREF, error) {
		// hdc = user32dll.GetDC(hwnd) // hdc表示裝置資訊，這種資訊不需要每次都要求要獲取，可以在外層初始化即可，在最後不用時在銷毀即可

		hObjOld := gdi32dll.SelectObject(hMemDC, w32.HGDIOBJ(hBitmapMem))
		// 開始傳輸，當BitBlt完成之後hBitmapMem的數據才會有資料
		if errno := gdi32dll.BitBlt(hMemDC, 0, 0, width, height, hdc, 0, 0, w32.SRCCOPY); errno != 0 {
			return 0, fmt.Errorf("%s", errno)
		}

		defer gdi32dll.SelectObject(hMemDC, hObjOld) // 不用之後可以考慮選回之前的物件

		return gdi32dll.GetPixel(hMemDC, x, y), nil
	}

	// 手動觸發
	if color, err := GetPixel(0, 0); err == nil {
		log.Println(w32.GetRValue(color), w32.GetGValue(color), w32.GetBValue(color))
	}

	ch := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(chQuit chan<- bool) {
		var (
			hLLMouseHook     w32.HHOOK
			hLLMouseHookProc w32.HOOKPROC
		)

		hLLMouseHookProc = func(nCode int32, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			if nCode < 0 {
				return user32dll.CallNextHookEx(hLLMouseHook, nCode, wParam, lParam)
			}

			if nCode == w32.HC_ACTION {
				mouseMsgID := wParam
				// msLLHookStruct := *(*w32.MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
				switch mouseMsgID {
				case w32.WM_LBUTTONDOWN:
					var pos w32.POINT
					if errno := user32dll.GetCursorPos(&pos); errno != 0 {
						fmt.Printf("GetCursorPos %s", errno)
					}
					log.Println(pos.X, pos.Y)
					color, err := GetPixel(pos.X, pos.Y)
					if err != nil {
						fmt.Println(err)
					}
					log.Println(w32.GetRValue(color), w32.GetGValue(color), w32.GetBValue(color))
				case w32.WM_RBUTTONUP:
					wg.Done()
				}
			}
			return user32dll.CallNextHookEx(hLLMouseHook, nCode, wParam, lParam)
		}

		var errno syscall.Errno
		kernel32dll := w32.NewKernel32DLL(w32.PNGetModuleHandle)
		hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
		if hLLMouseHook, errno = user32dll.SetWindowsHookEx(w32.WH_MOUSE_LL, hLLMouseHookProc, hInstance, 0); hLLMouseHook == 0 {
			log.Printf("Error SetWindowsHookEx [WH_MOUSE_LL] %s", errno)
			wg.Done()
			return
		}

		defer func() {
			if user32dll.UnhookWindowsHookEx(hLLMouseHook) != 0 {
				log.Printf("UnhookWindowsHookEx OK")
			}
			close(chQuit)
		}()

		go func() {
			// 一定要觸發GetMessage，之後Hook才會開始有作用，由於GetMessage之後就會一直鎖住，所以我們把它寫在另一個goRoutine，透過他的parent終止來關閉
			var msg w32.MSG
			if status, _ := user32dll.GetMessage(&msg, hwnd, 0, 0); status <= 0 {
				return
			}
		}()
		wg.Wait() // 這邊結束之後會自動把GetMessage的subRoutine給終止
	}(ch)

	for {
		select {
		case _, isOpen := <-ch:
			if !isOpen {
				return
			}
		case <-time.After(inputRunTimeSecond * time.Second):
			log.Println("[MaxRuntime] quit.")
			// return // 不直接結束，讓goroutine內的內容可以被完整運行完畢
			wg.Done()
		}
	}

	// Output:
}

func ExampleUser32DLL_GetWindowDC() {
	user32dll := w32.NewUser32DLL()
	hwnd := user32dll.GetDesktopWindow()
	hdc := user32dll.GetWindowDC(hwnd)
	defer user32dll.ReleaseDC(hwnd, hdc)
	// Output:
}

// 擷取畫面, 投放到記事本上
func ExampleGdi32DLL_BitBlt() {
	user32dll := w32.NewUser32DLL()
	gdi32dll := w32.NewGdi32DLL()

	var (
		// Screen
		// 來源資料 (投放內容)
		hwndS w32.HWND
		hdcS  w32.HDC
		rectS w32.RECT

		// Notepad
		// 目的資料 (放置被投放的內容)
		hwndN       w32.HWND
		hdcN        w32.HDC
		hdcMemN     w32.HDC
		hbitmapMemN w32.HBITMAP // 表明此hbitmap是仰賴hdcMemN所生成
	)

	// init
	{
		// hwndS = user32dll.FindWindow("ApplicationFrameWindow", "小算盤") // 擷取視窗除了desktopWindow以外的都會是黑畫面
		hwndS = user32dll.GetDesktopWindow()

		// hwndNotepad
		hwndN = user32dll.FindWindow("Notepad", "")
		if hwndN == 0 {
			return
		}

		hdcS = user32dll.GetWindowDC(hwndS)
		defer user32dll.ReleaseDC(hwndS, hdcS)

		if errno := user32dll.GetClientRect(hwndS, &rectS); errno != 0 {
			log.Printf("%s\n", errno)
			return
		}

		// hdcN = user32dll.GetDC(hwndN)
		hdcN = user32dll.GetWindowDC(hwndN) // 我們想連menu, scroll bar都填充
		defer user32dll.ReleaseDC(hwndN, hdcN)

		hdcMemN = gdi32dll.CreateCompatibleDC(hdcN)
		defer gdi32dll.DeleteDC(hdcMemN)
	}

	// 投放到notepad上
	// 注意這種方法如果投影的不是GetDesktopWindow的對象，會得到黑畫面(即時用了SRCCOPY|CAPTUREBLT也是一樣)
	{
		// 方法一: 直接投射
		_ = gdi32dll.BitBlt(hdcN,
			0, 0, rectS.Width(), rectS.Height(), // dst
			hdcS, 0, 0,
			w32.SRCCOPY,
		)

		_ = gdi32dll.BitBlt(hdcN,
			0, 0, rectS.Width(), rectS.Height(), // dst
			hdcS, 100, 200, // 可以想成先0,0之後100, 200以上的部分都會被截掉
			w32.SRCCOPY,
		)

		// 方法二: 帶有伸縮的投放方式
		previousMode := gdi32dll.SetStretchBltMode(hdcN, w32.HALFTONE)
		gdi32dll.StretchBlt(hdcN, 0, 0, 100, 200, // 會自動調整尺寸以符合100, 200
			hdcS, 0, 0, user32dll.GetSystemMetrics(w32.SM_CXSCREEN), user32dll.GetSystemMetrics(w32.SM_CYSCREEN),
			w32.SRCCOPY,
		)
		gdi32dll.SetStretchBltMode(hdcN, previousMode)
	}

	// 以下我們想把圖畫在記憶體之中
	// init2
	{
		hbitmapMemN = gdi32dll.CreateCompatibleBitmap(hdcN, rectS.Width(), rectS.Height())
		defer gdi32dll.DeleteObject(w32.HGDIOBJ(hbitmapMemN))
	}

	// 所有的CompatibleDC都要先指定物件，之後的動作才會知道是修改該物件的哪一部分{HBITMAP, HPEN, ...}
	gdi32dll.SelectObject(hdcMemN, w32.HGDIOBJ(hbitmapMemN))

	// 此動作完成之後hdcMemN.hbitmapN該資料區就會被寫入
	_ = gdi32dll.BitBlt(hdcMemN,
		0, 0, rectS.Width(), rectS.Height(),
		hdcS, 0, 0,
		w32.SRCCOPY,
	)

	// 為了印證該記憶體的圖資料已經被寫入，我們把來源改由記憶體，再重新用兩種畫法畫在notepad上
	{
		_ = gdi32dll.BitBlt(hdcN, 0, 0, rectS.Width(), rectS.Height(), // dst
			hdcMemN, 0, 0,
			w32.SRCCOPY,
		)

		// gdi32dll.SetStretchBltMode(hdcN, w32.HALFTONE) // 當我們省略衍伸模式，預設的模式會按照來源圖與目的圖，1對1複製，來源超過的部分將會自動被截掉
		gdi32dll.StretchBlt(hdcN, 0, 0, rectS.Width(), rectS.Height(),
			hdcMemN, 0, 0, user32dll.GetSystemMetrics(w32.SM_CXSCREEN), user32dll.GetSystemMetrics(w32.SM_CYSCREEN),
			w32.SRCCOPY,
		)
	}

	// 如果您的目的不是單純投放，而是想取得圖片點集的資訊，就要考慮使用以下內容(GetDIBits)來獲得點資料
	{
		// HBITMAP TO BITMAP
		var bitmapN w32.Bitmap
		gdi32dll.GetObject(w32.HANDLE(hbitmapMemN), int32(unsafe.Sizeof(bitmapN)), uintptr(unsafe.Pointer(&bitmapN)))

		// 第一次呼叫GetDIBits取得BitmapInfo的資料
		var bitmapInfo w32.BitmapInfo
		gdi32dll.GetDIBits(hdcMemN, hbitmapMemN, 0, 0, 0, &bitmapInfo, w32.DIB_RGB_COLORS) // DDB to DIB

		// 第二次呼叫GetDIBits取得圖的資料內容
		var lpBitmapData w32.LPVOID // 這個資料包含三樣東西{BitmapFileHeader, BitmapInfoHeader, 點集資料}
		gdi32dll.GetDIBits(hdcMemN, hbitmapMemN, 0, w32.UINT(bitmapInfo.Header.Height), lpBitmapData, &bitmapInfo, w32.DIB_RGB_COLORS)

		// 以下我們只對點集的資料數據有興趣，header, info都不是我們所關心的
		{
			bmpSize := ((bitmapN.Width*int32(bitmapInfo.Header.BitCount) + 31) / 32) * 4 * bitmapN.Height

			var bitmapFileHeader w32.BitmapFileHeader
			sizeofDIB := 14 + uint32(unsafe.Sizeof(bitmapInfo.Header)) + uint32(bmpSize)
			bitmapFileHeader = w32.BitmapFileHeader{
				Type:       0x4D42,    // BM. // B: 42, M: 4D  // 因為BitmapFile所有的描述都要用"little-endian"讀取，所以要反過來寫4D42
				Size:       sizeofDIB, // HEADER + INFO + DATA
				OffsetBits: 14 + uint32(unsafe.Sizeof(bitmapInfo.Header)),
			}

			var offset uint32
			bmpPointsDatas := make([]byte, bitmapFileHeader.OffsetBits) // 排除了header, info等資訊
			for offset = 14 + uint32(unsafe.Sizeof(bitmapInfo.Header)); offset < bitmapFileHeader.OffsetBits; offset += 1 {
				bmpPointsDatas[offset] = *(*byte)(unsafe.Pointer(uintptr(lpBitmapData) + uintptr(offset)))
			}
		}
	}

	// Output:
}

// https://learn.microsoft.com/en-us/windows/win32/inputdev/using-mouse-input#drawing-lines-with-the-mouse
func ExampleGdi32DLL_LineTo() {
	opt := &w32.WindowOptions{}
	var (
		hdc w32.HDC

		rectClient w32.RECT
		ptClientLT w32.POINT // client left top corner
		ptClientRB w32.POINT // client right bottom corner

		ptBegin   w32.POINT
		ptEnd     w32.POINT
		ptPrevEnd w32.POINT

		isLineMode bool // 當前的狀態是否處於畫線
	)

	hPen := gdiDll.CreatePen(w32.PS_SOLID, 5, w32.RGB(255, 0, 0))
	defer func() {
		if !gdiDll.DeleteObject(w32.HGDIOBJ(hPen)) {
			log.Println("DeleteObject hpen error")
		}
	}()

	opt.WndProc = func(hwnd w32.HWND, uMsg uint32, wParam w32.WPARAM, lParam w32.LPARAM) uintptr {
		switch uMsg {
		case w32.WM_CREATE:
			userDll.ShowWindow(hwnd, w32.SW_SHOW)
			go func() {
				<-time.After(5 * time.Second) // 如果要測試，可以自行延長秒數
				_, _, _ = userDll.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)
			}()

		// 限制鼠標不可超過繪圖區，以及紀錄開始繪製的起點
		case w32.WM_LBUTTONDOWN:
			userDll.SetCapture(hwnd)

			_ = userDll.GetClientRect(hwnd, &rectClient)
			log.Printf("client rect: %+v", rectClient)
			ptClientLT.X = rectClient.Left
			ptClientLT.Y = rectClient.Top
			ptClientRB.X = rectClient.Right + 1 // +1 是為了包含到邊界，因為GetClientRect並不包含右下邊界線
			ptClientRB.Y = rectClient.Bottom + 1

			userDll.ClientToScreen(hwnd, &ptClientLT)
			userDll.ClientToScreen(hwnd, &ptClientRB)

			// 重新設定rectClient
			userDll.SetRect(&rectClient,
				ptClientLT.X, ptClientLT.Y,
				ptClientRB.X, ptClientRB.Y,
			)
			log.Printf("client to screen: %+v", rectClient)

			if eno := userDll.ClipCursor(&rectClient); eno != 0 {
				log.Println(eno)
				break
			}
			ptBegin = w32.MakePoint(lParam)
			return 0
		case w32.WM_MOUSEMOVE: // 移動且左鍵被按下，則畫線
			if wParam != w32.MK_LBUTTON {
				break
			}

			hdc = userDll.GetDC(hwnd)
			gdiDll.SelectObject(hdc, w32.HGDIOBJ(hPen))
			gdiDll.SetROP2(hdc, w32.R2_NOTXORPEN)

			// 當我們按住左鍵移動的時候，如果繞一個弧線，此時我們在最後彈開的位置才需要畫線，因此我們要讓弧線被劃過的部分都消去
			if isLineMode {
				// 在畫一次會消去
				gdiDll.MoveToEx(hdc, ptBegin.X, ptBegin.Y, nil)
				gdiDll.LineTo(hdc, ptPrevEnd.X, ptPrevEnd.Y)
			}

			// 底下才是最後被畫的線
			ptEnd = w32.MakePoint(lParam)
			gdiDll.MoveToEx(hdc, ptBegin.X, ptBegin.Y, nil)
			gdiDll.LineTo(hdc, ptEnd.X, ptEnd.Y)
			isLineMode = true
			ptPrevEnd = ptEnd
			userDll.ReleaseDC(hwnd, hdc)
		case w32.WM_LBUTTONUP:
			isLineMode = false
			_ = userDll.ClipCursor(nil)
			if eno := userDll.ReleaseCapture(); eno != 0 {
				log.Println(eno)
			}
		case w32.WM_DESTROY:
			userDll.PostQuitMessage(0)
			return 0
		}
		return uintptr(userDll.DefWindowProc(hwnd, w32.UINT(uMsg), wParam, lParam))
	}

	wnd, err := createWindow("LineTo demo", opt)
	if err != nil {
		log.Fatal(err)
	}
	wnd.Run(nil)

	// Output:
}

func ExampleGdi32DLL_MoveToEx() {
	var pt w32.POINT
	hdc := userDll.GetDC(0)
	defer func() {
		if userDll.ReleaseDC(0, hdc) == 0 {
			log.Println("ReleaseDC error")
		}
	}()
	gdiDll.MoveToEx(hdc, 100, 200, &pt) // 如果之前都沒有移動過，那麼一開始得到的pt位置是0, 0. 移動到(100, 200)的位置並獲取移動前的位置存放在pt中
	fmt.Println(pt)
	gdiDll.MoveToEx(hdc, 200, 300, &pt) // 前一個點位置在100, 200
	fmt.Println(pt)

	// Output:
	// {0 0}
	// {100 200}
}
