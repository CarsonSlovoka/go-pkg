package w32_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"os"
	"path/filepath"
)

// 使用ExtractIcon取得該應用程式的HICON
func ExampleShellDLL_ExtractIcon() {
	shell32dll := w32.NewShellDLL([]w32.ProcName{
		w32.PNExtractIcon,
	})

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
				if err := user32dll.ReleaseDC(hwndNotepad, curHDC); err != nil {
					log.Fatal(err)
				}
			}
		}()

		if err := user32dll.DrawIcon(curHDC, 50, 100, hIcon); err != nil {
			panic(err)
		}
	}
	// Output:
}

// 使用ExtractIcon來計算該檔案擁有的圖標數量
func ExampleShellDLL_ExtractIcon_count() {
	shell32dll := w32.NewShellDLL([]w32.ProcName{
		w32.PNExtractIcon,
	})

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
				if err := user32dll.ReleaseDC(hwndNotepad, curHDC); err != nil {
					log.Fatal(err)
				}
			}
		}()
		for iconIdx := 0; iconIdx < int(numIcon); iconIdx++ {
			hicon := shell32dll.ExtractIcon(0, exeFileName, iconIdx)

			if err := user32dll.DrawIcon(curHDC, 50, 50*(iconIdx+1), hicon); err != nil {
				panic(err)
			}
		}
	}
	// Output:
}
