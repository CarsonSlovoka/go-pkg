package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"syscall"
)

// 添加字型，不需要安裝。重開機或執行RemoveFontResource將會被移除
func ExampleGdi32DLL_AddFontResource() {
	ttfPath := "./testdata/fonts/teamviewer15.otf"
	gdi32dll := w32.NewGdi32DLL(w32.PNAddFontResource, w32.PNRemoveFontResource)
	user32dll := w32.NewUser32DLL([]w32.ProcName{w32.PNPostMessage})
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
	user32dll := w32.NewUser32DLL([]w32.ProcName{w32.PNPostMessage})
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
