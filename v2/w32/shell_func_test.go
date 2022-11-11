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

func ExampleShellDLL_ShellNotifyIcon() {
	user32dll := w32.NewUser32DLL()
	shell32dll := w32.NewShellDLL()

	// prepare test data
	var myHICON w32.HICON
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
	}

	guid := *(*w32.GUID)(unsafe.Pointer(&([]byte("abcdef1234567890"))[0]))
	hwndTarget := user32dll.GetForegroundWindow()
	// hwndTarget := user32dll.FindWindow("powershell", "")
	// hwndTarget := user32dll.GetActiveWindow()
	if hwndTarget == 0 {
		return
	}

	var notifyIconData w32.NOTIFYICONDATA
	// log.Println(uint32(unsafe.Sizeof(notifyIconData))) 976這種算法有問題，要手動計算
	// 建立常規屬性
	notifyIconData = w32.NOTIFYICONDATA{
		CbSize:   968,
		HWnd:     hwndTarget,
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
			log.Fatalf("NIM_DELETE ERROR")
		}
	}()

	if !shell32dll.ShellNotifyIcon(w32.NIM_ADD, &notifyIconData) {
		log.Fatalf("NIM_ADD ERROR")
		return
	}

	// UFlags 以下標示依需求決定
	var (
		enableInfo        bool
		enableMessage     bool
		enableTooltip     bool
		enableIcon        bool
		enableBalloonIcon bool
	)

	flag.BoolVar(&enableInfo, "eInfo", true, "")
	flag.BoolVar(&enableMessage, "eMsg", true, "")
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

	MYCallBackMsgID := w32.WM_APP + 1
	if enableMessage {
		notifyIconData.UFlags |= w32.NIF_MESSAGE // msg有設定UCallbackMessage就會生效 // The uCallbackMessage member is valid.
		notifyIconData.UCallbackMessage = uint32(MYCallBackMsgID)
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
		log.Fatalf("NIM_MODIFY ERROR")
	}
	// Output:
}
