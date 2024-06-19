package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"syscall"
)

var (
	oleDll    *w32.Ole32DLL
	kernelDll *w32.Kernel32DLL
	userDll   *w32.User32DLL
	oleAutDll *w32.OleAut32DLL
	shellDll  *w32.ShellDLL
	gdiDll    *w32.Gdi32DLL
	adApiDll  *w32.AdApiDLL
	psApiDll  *w32.PsApiDLL
)

func init() {
	oleDll = w32.NewOle32DLL()
	kernelDll = w32.NewKernel32DLL()
	oleAutDll = w32.NewOleAut32DLL()
	userDll = w32.NewUser32DLL()
	shellDll = w32.NewShellDLL()
	gdiDll = w32.NewGdi32DLL()
	adApiDll = w32.NewAdApi32DLL()
	psApiDll = w32.NewPsApiDLL()
}

func getTestHwnd() w32.HWND {
	hwnd := userDll.FindWindow("Notepad", "")
	if hwnd == 0 {
		hwnd = userDll.GetDesktopWindow()
	}
	return hwnd
}

type exampleWindow struct {
	hwnd w32.HWND
}

func (w *exampleWindow) Run(msgProc func(msg *w32.MSG) bool) {
	var msg w32.MSG
	for {
		if bRet, eno := userDll.GetMessage(&msg, 0, 0, 0); bRet == 0 {
			// WM_QUIT
			break
		} else if bRet == -1 {
			log.Printf("GetMessage error:%s\n", eno)
		} else {
			if msgProc != nil {
				if !msgProc(&msg) { // 若為假則不進行轉譯
					continue
				}
			}
			userDll.TranslateMessage(&msg)
			userDll.DispatchMessage(&msg)
		}
	}
}

// 方便創建測試用視窗
func createWindow(title string, opt *w32.WindowOptions) (*exampleWindow, error) {
	hInstance := w32.HINSTANCE(kernelDll.GetModuleHandle(""))

	if opt.WndProc == nil {
		// defaultWindowProc
		opt.WndProc = func(hwnd w32.HWND, uMsg uint32, wParam w32.WPARAM, lParam w32.LPARAM) uintptr {
			switch uMsg {
			case w32.WM_CREATE:
				userDll.ShowWindow(hwnd, w32.SW_SHOW)
			case w32.WM_DESTROY:
				userDll.PostQuitMessage(0)
				return 0
			}
			return uintptr(userDll.DefWindowProc(hwnd, w32.UINT(uMsg), wParam, lParam))
		}
	}

	// Register
	if opt.ClassName == "" {
		opt.ClassName = "example"
	}
	pUTF16ClassName, _ := syscall.UTF16PtrFromString(opt.ClassName)

	var hIcon w32.HANDLE
	if opt.IconPath != "" {
		hIcon, _ = userDll.LoadImage(0, // hInstance must be NULL when loading from a file
			opt.IconPath,
			w32.IMAGE_ICON, 0, 0, w32.LR_LOADFROMFILE|w32.LR_DEFAULTSIZE|w32.LR_SHARED)
	}

	if atom, errno := userDll.RegisterClass(&w32.WNDCLASS{
		Style:         opt.ClassStyle,
		HbrBackground: w32.COLOR_WINDOW,
		WndProc:       syscall.NewCallback(opt.WndProc),
		HInstance:     hInstance,
		HIcon:         w32.HICON(hIcon),
		ClassName:     pUTF16ClassName,
	}); atom == 0 {
		return nil, fmt.Errorf("[RegisterClass Error] %w", errno)
	}

	width := opt.Width
	if width == 0 {
		width = w32.CW_USEDEFAULT
	}
	height := opt.Height
	if height == 0 {
		height = w32.CW_USEDEFAULT
	}
	posX := opt.X
	if posX == 0 {
		posX = w32.CW_USEDEFAULT
	}
	posY := opt.Y
	if posY == 0 {
		posY = w32.CW_USEDEFAULT
	}

	if opt.Style == 0 {
		opt.Style = w32.WS_OVERLAPPEDWINDOW
	}

	// Create window
	hwnd, errno := userDll.CreateWindowEx(
		w32.DWORD(opt.ExStyle),
		opt.ClassName,
		title,
		w32.DWORD(opt.Style),

		// Size and position
		posX, posY, width, height,

		0, // Parent window
		0, // Menu
		hInstance,
		0, // Additional application data
	)

	if errno != 0 {
		if errno2 := userDll.UnregisterClass(opt.ClassName, hInstance); errno2 != 0 {
			fmt.Printf("Error UnregisterClass: %s", errno2)
		}
		return nil, errno
	}
	return &exampleWindow{hwnd}, nil
}

func saveImgEx(output string, width, height int32, bitmapBits []byte) {
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for y := int32(0); y < height; y++ {
		for x := int32(0); x < width; x++ {
			idx := (y*width + x) * 4
			b, g, r := bitmapBits[idx], bitmapBits[idx+1], bitmapBits[idx+2]
			img.Set(int(x), int(y), color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	saveImg(output, img)
}

func saveImg(outputPath string, img *image.RGBA) {
	file, err := os.Create(outputPath) // 建議用瀏覽器(chrome)來查看，可以觀察到alpha
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()
	if err = png.Encode(file, img); err != nil {
		panic(err)
	}
}

func newHFont(fontFamily string, height int32) w32.HFONT {
	lf := w32.LogFont{
		Height:         height, // 建議用負值
		Width:          0,
		Escapement:     0,
		Orientation:    0,
		Weight:         400,
		Italic:         0,
		Underline:      0,
		StrikeOut:      0,
		CharSet:        w32.DEFAULT_CHARSET,
		OutPrecision:   w32.OUT_TT_PRECIS,
		ClipPrecision:  w32.CLIP_DEFAULT_PRECIS,
		Quality:        w32.ANTIALIASED_QUALITY,
		PitchAndFamily: w32.FF_DONTCARE,
	}
	return gdiDll.CreateFont(
		lf.Height,
		lf.Width,
		lf.Escapement,
		lf.Orientation,
		lf.Weight,
		uint32(lf.Italic),
		uint32(lf.Underline),
		uint32(lf.StrikeOut),
		uint32(lf.CharSet),
		uint32(lf.OutPrecision),
		uint32(lf.ClipPrecision),
		uint32(lf.Quality),
		uint32(lf.PitchAndFamily),
		fontFamily, // 安裝到你電腦的字型 FontFamilyName, name.ID=1 都可以指定
	)
}
