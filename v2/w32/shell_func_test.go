package w32_test

import (
	"flag"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

// 使用ExtractIcon取得該應用程式的HICON
func ExampleShellDLL_ExtractIcon() {
	shell32dll := w32.NewShellDLL(w32.PNExtractIcon)

	exePath := filepath.Join(os.Getenv("windir"), "System32/fontview.exe")
	// exePath := "powershell.exe" // 系統路徑可以找到的執行檔也可以(不需要再標明路徑位置)，副檔名不可以省略
	// exePath := "../myXXX.exe" // 相對路徑也可以

	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		log.Printf("not found %q", exePath)
	}

	hIcon := shell32dll.ExtractIcon(0, exePath, 0)

	if hIcon == 0 {
		return
	}

	// 以下只是把hicon畫在notepad.exe上而已
	{
		user32dll := w32.NewUser32DLL(
			w32.PNFindWindow,
			w32.PNGetDC,
			w32.PNReleaseDC,
			w32.PNDrawIcon,
		)

		hwndNotepad := user32dll.FindWindow("Notepad", "")
		if hwndNotepad == 0 {
			log.Println("notepad.exe not found")
			return
		}
		curHDC := user32dll.GetDC(hwndNotepad)
		defer func() {
			if curHDC != 0 {
				if user32dll.ReleaseDC(hwndNotepad, curHDC) == 0 {
					log.Fatal("ReleaseDC")
				}
			}
		}()

		if ok, errno := user32dll.DrawIcon(curHDC, 50, 100, hIcon); !ok {
			log.Fatalf("%s", errno)
		}
	}
	// Output:
}

// 使用ExtractIcon來計算該檔案擁有的圖標數量
func ExampleShellDLL_ExtractIcon_count() {
	shell32dll := w32.NewShellDLL(
		w32.PNExtractIcon,
	)

	const exeFileName = "powershell.exe"
	numIcon := shell32dll.ExtractIcon(0, exeFileName,
		-1, // 抓出所有icon的數量
	)

	if numIcon == 0 {
		return
	}

	// 作圖在notepad.exe上
	{
		user32dll := w32.NewUser32DLL(
			w32.PNFindWindow,
			w32.PNGetDC,
			w32.PNReleaseDC,
			w32.PNDrawIcon,
		)

		hwndNotepad := user32dll.FindWindow("Notepad", "")
		if hwndNotepad == 0 {
			log.Println("notepad.exe not found")
			return
		}
		curHDC := user32dll.GetDC(hwndNotepad)
		defer func() {
			if curHDC != 0 {
				if user32dll.ReleaseDC(hwndNotepad, curHDC) == 0 {
					log.Fatal("ReleaseDC")
				}
			}
		}()
		for iconIdx := 0; iconIdx < int(numIcon); iconIdx++ {
			hicon := shell32dll.ExtractIcon(0, exeFileName, iconIdx)

			if ok, errno := user32dll.DrawIcon(curHDC, 50, 50*(iconIdx+1), hicon); !ok {
				log.Fatalf("%s", errno)
			}
		}
	}
	// Output:
}

func TestGetGuid(t *testing.T) {
	var guidBuf []byte
	guidBuf = []byte("abcdef1234567890")
	guid1 := *(*w32.GUID)(unsafe.Pointer(&guidBuf[0])) // 他會以您要輸出的型別來自動抓取合適的byte數量，與餵入的資料(guidBuf)沒有關係

	var guidBuf2 = [16]byte{
		'a', 'b', 'c', 'd', 'e', 'f',
		'1', '2', '3', '4', '5',
		'6', '7', '8', '9', '0',
	}
	guid2 := *(*w32.GUID)(unsafe.Pointer(&guidBuf2[0]))

	if fmt.Sprintf("%+v", guid1) != fmt.Sprintf("%+v", guid2) {
		t.Fatalf("should not get the error.")
	}
}

// SetTimeout uTimeout is valid only in Windows 2000 and Windows XP.
func ExampleNOTIFYICONDATA_SetTimeout() {
	var n w32.NOTIFYICONDATA
	n.SetTimeout(3500) // 3.5 sec
}

// https://www.programcreek.com/python/?code=IronLanguages%2Fironpython2%2Fironpython2-master%2FSrc%2FStdLib%2FLib%2Fsite-packages%2Fwin32%2FDemos%2Fwin32gui_taskbar.py
//
// 如果因為不正常結束而有殘留的ShellNotifyIcon存在，可以用以下方式刪掉
// 1. HKEY_CURRENT_USER\SOFTWARE\Classes\Local Settings\Software\Microsoft\Windows\CurrentVersion\TrayNotify
// 2. (備份整個TrayNotify資料夾，以防萬一)
// 3. 刪除IconStreams, PastIconsStream兩個機碼數值
// 4. 開啟工作管理員(taskmgr.exe)，刪除所有explorer.exe的項目
// 5. 重新執行explorer.exe
// 做完之後無效的圖示應該就會被清除
func ExampleShellDLL_ShellNotifyIcon() {
	user32dll := w32.NewUser32DLL()
	shell32dll := w32.NewShellDLL()
	kernel32dll := w32.NewKernel32DLL(w32.PNGetModuleHandle)
	gdi32dll := w32.NewGdi32DLL(w32.PNGetObject, w32.PNDeleteObject)

	// prepare test data
	var (
		myHICON w32.HICON
		iInfo   w32.ICONINFO
	)

	{
		// myHICON = shell32dll.ExtractIcon(0, "notepad", 0) // 不是隨便載入一個HICON都可以被順利呈現
		// myHICON, _ := user32dll.LoadIcon(0, w32.MakeIntResource(w32.IDI_QUESTION)) // <- 使用這個系統圖標測試是可行的
		myHICON = w32.HICON(user32dll.MustLoadImage( // returns a HANDLE so we have to cast to HICON
			0,                         // hInstance must be NULL when loading from a file
			"testdata/img/golang.ico", // the icon file name
			w32.IMAGE_ICON,            // specifies that the file is an icon

			0, // width of the image (we'll specify default later on) // 如果您省略了LR_DEFAULTSIZE，想自己設定尺寸，則尺寸也不能亂給(給32*32是可行的, 16*16不行)，不然ShellNotifyIcon無法建立成功
			0, // height of the image

			w32.LR_LOADFROMFILE| // we want to load a file (as opposed to a resource)
				w32.LR_DEFAULTSIZE| // default metrics based on the type (IMAGE_ICON, 32x32)
				w32.LR_SHARED, // let the system release the handle when it's no longer used
		))

		// get IconInfo, 為了SetMenuItemInfo須提供HBITMAP才可以設置成功
		{
			if !user32dll.GetIconInfo(myHICON, &iInfo) {
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
		}
	}

	// variable used for createWindow
	chanWin := make(chan w32.HWND)
	const (
		wndClassName  = "classShellNotifyIcon"
		wndWindowName = "windowShellNotifyIcon"
	)
	hInstance := w32.HINSTANCE(kernel32dll.GetModuleHandle(""))
	const WMNotifyIconMsg = w32.WM_APP + 123 // 定義notifyIcon會觸發的訊息ID
	// Create a window https://github.com/CarsonSlovoka/go-pkg/blob/efe1c50fa40229c299232fe3b236135b1046ef35/v2/w32/user32_func_test.go#L457-L659
	go func(ch chan<- w32.HWND) {
		// 定義訊息處理函數
		wndProcFuncPtr := syscall.NewCallback(w32.WNDPROC(func(hwnd w32.HWND, uMsg w32.UINT, wParam w32.WPARAM, lParam w32.LPARAM) w32.LRESULT {
			switch uMsg {
			case w32.WM_CLOSE:
				if wParam != 123 {
					log.Println("縮小視窗，不真的結束")
					user32dll.ShowWindow(hwnd, w32.SW_HIDE)
					return 0
				}
			case w32.WM_DESTROY: // 這個訊息要寫，不能倚靠DefWindowProc，否則會關閉不掉
				log.Println("WM_DESTROY")
				user32dll.PostQuitMessage(0)
				return 0
			case WMNotifyIconMsg:
				switch lParam {
				case w32.WM_LBUTTONUP:
					log.Println("WMNotifyIconMsg->WM_LBUTTONUP")
					if wParam == 123 {
						log.Println("WMNotifyIconMsg->WM_LBUTTONUP")
					}
				case w32.WM_LBUTTONDBLCLK:
					log.Println("WMNotifyIconMsg->WM_LBUTTONDBLCLK")
					user32dll.ShowWindow(hwnd, w32.SW_SHOWNORMAL) // SW_MAXIMIZE
				case w32.WM_RBUTTONDBLCLK:
					if ok, errno := user32dll.DestroyWindow(hwnd); !ok {
						log.Printf("%s", errno)
						return 0
					}
				case w32.WM_RBUTTONUP:
					log.Println("WMNotifyIconMsg->WM_RBUTTONUP")
					hMenu := user32dll.CreatePopupMenu()
					_, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1023, "Display Dialog")
					_, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1024, "Say Hello 哈囉！")
					var menuItemInfo w32.MENUITEMINFO
					menuItemInfo.CbSize = uint32(unsafe.Sizeof(menuItemInfo))
					menuItemInfo.FMask = w32.MIIM_BITMAP
					menuItemInfo.HbmpItem = iInfo.HbmColor
					_, _ = user32dll.SetMenuItemInfo(hMenu, 1024, false, &menuItemInfo)
					_, _ = user32dll.AppendMenu(hMenu, w32.MF_STRING, 1025, "Exit program")

					defer func() {
						// gdi32dll.DeleteObject(w32.HGDIOBJ(menuItemInfo.HbmpItem)) // 不需要把這個HBITMAP銷毀，否則下一次在新建就看不到該HICON了
						if ok, errno := user32dll.DestroyMenu(hMenu); !ok { // 因為每次右鍵都會新增一個HMENU，所以不用之後要在銷毀，避免一直累積
							log.Printf("%s\n", errno)
						}
					}()

					var pos w32.POINT
					if ok, errno := user32dll.GetCursorPos(&pos); !ok {
						fmt.Printf("%s", errno)
						return 1
					}
					user32dll.SetForegroundWindow(hwnd)
					_, _ = user32dll.TrackPopupMenu(hMenu, w32.TPM_LEFTALIGN, pos.X, pos.Y, 0, hwnd, nil)
					// _, _ = user32dll.PostMessage(hwnd, w32.WM_NULL, 0, 0)
				}
				return 1 // 讓消息循環繼續處理其他訊息(>0即可)
			case w32.WM_COMMAND:
				id := w32.LOWORD(wParam)
				switch id {
				case 1023:
					_, _ = user32dll.PostMessage(hwnd, WMNotifyIconMsg, 0, w32.WM_LBUTTONDBLCLK)
				case 1024:
					log.Println("hello")
				case 1025:
					log.Println("1025")
					_, _ = user32dll.PostMessage(hwnd, w32.WM_DESTROY, 0, 0)
				}
			}
			return user32dll.DefWindowProc(hwnd, uMsg, wParam, lParam)
		}))

		// 類別名稱註冊
		pUTF16ClassName, _ := syscall.UTF16PtrFromString(wndClassName)
		if atom, errno := user32dll.RegisterClass(&w32.WNDCLASS{
			Style:         w32.CS_HREDRAW | w32.CS_HREDRAW,
			HbrBackground: w32.COLOR_WINDOW,
			LpfnWndProc:   wndProcFuncPtr,
			HInstance:     hInstance,
			HIcon:         myHICON,
			LpszClassName: pUTF16ClassName,
		}); atom == 0 {
			fmt.Printf("%s", errno)
			chanWin <- 0
			return
		}

		// 創建視窗
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
			ch <- hwnd
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
			close(ch)
		}()

		ch <- hwnd

		// 消息循環
		var msg w32.MSG
		for {
			if status, _ := user32dll.GetMessage(&msg, 0, 0, 0); status <= 0 {
				break
			}
			user32dll.TranslateMessage(&msg)
			user32dll.DispatchMessage(&msg)
		}
	}(chanWin)

	// 不能隨便找hwnd過來，除非這些hwnd會處理您的自訂訊息WMNotifyIconMsg，因此我們需要自己創建視窗
	// hwndTarget := user32dll.GetForegroundWindow()
	// hwndTarget := user32dll.FindWindow("powershell", "")
	// hwndTarget := user32dll.GetActiveWindow()
	hwndTarget := <-chanWin
	if hwndTarget == 0 {
		return
	}

	var notifyIconData w32.NOTIFYICONDATA
	guid := *(*w32.GUID)(unsafe.Pointer(&([]byte("abcdef12345678zr"))[0]))
	{
		notifyIconData = w32.NOTIFYICONDATA{
			CbSize:   968,
			HWnd:     hwndTarget,   // 注意！不是用您創建的window的HWND，是要用當前app的hwnd
			UFlags:   w32.NIF_GUID, // NIF_GUID有設定就可以讓GuidItem生效
			GuidItem: guid,
		}
		notifyIconData.SetVersion(w32.NOTIFYICON_VERSION_4)
		notifyIconDataCopy := notifyIconData // 這個只是用來驗證，刪除的資訊與其他資訊無關，它依照GuidItem的內容去刪

		// 確保沒有殘留的資料, 如果程式有不正常結束，那麼殘留的對象會影響，使的NIM_ADD會一直沒辦法被刪除
		if shell32dll.ShellNotifyIcon(w32.NIM_DELETE, &notifyIconDataCopy) {
			log.Println("clear previous data.")
		}

		defer func() {
			// 刪除認的是NIF_GUID，所以只要NOTIFYICONDATA中的GuidItem相同，就會被刪掉
			if !shell32dll.ShellNotifyIcon(w32.NIM_DELETE, &notifyIconDataCopy) {
				log.Println("NIM_DELETE ERROR")
			}
		}()
	}

	// 掛勾訊息處理
	notifyIconData.UFlags |= w32.NIF_MESSAGE // msg有設定UCallbackMessage就會生效 // The uCallbackMessage member is valid.
	notifyIconData.UCallbackMessage = uint32(WMNotifyIconMsg)

	if !shell32dll.ShellNotifyIcon(w32.NIM_ADD, &notifyIconData) {
		// 關閉所建立的背景視窗
		_, _, _ = user32dll.SendMessage(hwndTarget, w32.WM_CLOSE, 123, 0)

		// 等待背景視窗確實被關閉
		<-chanWin
		log.Println("NIM_ADD ERROR")
		return
	}

	// UFlags 以下標示依需求決定
	var (
		enableInfo        bool
		enableTooltip     bool
		enableIcon        bool
		enableBalloonIcon bool
	)

	flag.BoolVar(&enableInfo, "eInfo", true, "")
	flag.BoolVar(&enableTooltip, "eTooltip", true, "")
	flag.BoolVar(&enableIcon, "eIcon", true, "")
	flag.BoolVar(&enableBalloonIcon, "eBalloonIcon", true, "")
	flag.Parse()

	if !enableInfo {
		notifyIconData.UFlags |= w32.NIF_INFO // The szInfo, szInfoTitle, dwInfoFlags, and uTimeout members are valid. Note that uTimeout is valid only in Windows 2000 and Windows XP.
		infoMsg, _ := syscall.UTF16FromString("SzInfo")
		infoTitle, _ := syscall.UTF16FromString("SzInfoTitle")
		copy(notifyIconData.SzInfo[:], infoMsg)
		copy(notifyIconData.SzInfoTitle[:], infoTitle)

		// DwInfoFlags
		if enableBalloonIcon {
			notifyIconData.DwInfoFlags |= w32.NIIF_USER | w32.NIIF_LARGE_ICON // Windows XP: hIcon as the notification, Windows Vista and later: hBalloonIcon as the notification balloon's title icon.
			// 所以建議把hIcon與hBalloonIcon都設定就不用擔心是XP還是新版本的問題
			notifyIconData.HIcon = myHICON
			notifyIconData.HBalloonIcon = myHICON
		}
	}

	if enableTooltip {
		notifyIconData.UFlags |= w32.NIF_TIP // tip有設定, SzTip就會生效 The szTip member is valid.
		utf16Title, _ := syscall.UTF16FromString("SzTip")
		copy(notifyIconData.SzTip[:], utf16Title) // windows幾乎都採用utf16編碼
	}

	if enableIcon {
		notifyIconData.UFlags |= w32.NIF_ICON // The hIcon member is valid.
		notifyIconData.HIcon = myHICON
	}

	if !shell32dll.ShellNotifyIcon(w32.NIM_MODIFY, &notifyIconData) {
		_, _, _ = user32dll.SendMessage(hwndTarget, w32.WM_DESTROY, 0, 0)
		<-chanWin
		log.Fatalf("NIM_MODIFY ERROR")
	}

	// 傳送自定義的訊息到notifyIcon上
	{
		_, _ = user32dll.PostMessage(hwndTarget, WMNotifyIconMsg, 123, w32.WM_LBUTTONUP)
		_, _ = user32dll.PostMessage(hwndTarget, WMNotifyIconMsg, 0, w32.WM_RBUTTONUP) // 選單測試

		time.Sleep(time.Second) // 我們想要自動結束，又想測試選單功能，所以採用PostMessage，否則使用SendMessage必須要選擇選單中的內容(或者取消)才可再往下進行。而如果PostMessage緊接著就關閉視窗，可能也會導致選單訊息還在處理之前就被關閉了，所以才會等待1秒鐘
	}

	// 🕹️ 如果您要手動嘗試，請把以下的SendMessage.WM_CLOSE註解掉，手動雙擊右鍵(兩下)即可結束
	// 自動發送關閉訊息, 作為github.action的測試，我們不做其他UI控件的處理，僅確認ShellNotifyIcon有被創建成功即可
	_, _, _ = user32dll.SendMessage(hwndTarget, w32.WM_CLOSE, 123 /* 自定義內容 */, 0) // 不要用WM_DESTROY 不然UnregisterClass會因為視窗還存在而跑出錯誤Class still has open windows.

	<-chanWin
	// Output:
}
