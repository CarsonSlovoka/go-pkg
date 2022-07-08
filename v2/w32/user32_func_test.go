package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
)

func ExampleUser32DLL_GetWindowText() {
	user32dll := w32.NewUser32DLL([]w32.ProcName{
		w32.PNGetForegroundWindow,
		w32.PNGetClassName,
		w32.PNGetWindowText,
	})

	curHwnd, err := user32dll.GetForegroundWindow()
	fmt.Println("current window HWND:", curHwnd) // 當前窗口的識別號

	clsName, err := user32dll.GetClassName(curHwnd)
	if err != nil {
		panic(err)
	}
	fmt.Println("window class Name:", clsName)

	winText, err := user32dll.GetWindowText(curHwnd)
	if err != nil {
		panic(err)
	}
	fmt.Println("window text Name:", winText)
}

func ExampleUser32DLL_MessageBox() {
	user32dll := w32.NewUser32DLL([]w32.ProcName{
		w32.PNMessageBox,
	})

	hwndTop := uintptr(w32.HWND_TOP)
	response, _ := user32dll.MessageBox(hwndTop, "title", "body message", w32.MB_OK)
	switch response {
	case w32.IDYES:
		fmt.Println("Yes")
	case w32.IDNO:
		fmt.Println("No")
	default: // 基本上沒辦法有這個出現，對話框只有Yes,No可以選，右上角的打X也不能按
		fmt.Println("Unknown")
	}

	messageBox := user32dll.MessageBox
	_, _ = messageBox(hwndTop, "Test", "OK", w32.MB_OK)
	_, _ = messageBox(hwndTop, "Test", "Yes No Cancel", w32.MB_YESNOCANCEL)
	_, _ = messageBox(hwndTop, "Test", "OK", w32.MB_OK)
	_, _ = messageBox(hwndTop, "Test", "Help button", w32.MB_HELP)
	_, _ = messageBox(hwndTop, "Test", "OK CANCEL", w32.MB_OKCANCEL)
	_, _ = messageBox(hwndTop, "Test", "ABORT RETRY IGNORE", w32.MB_ABORTRETRYIGNORE)
	_, _ = messageBox(hwndTop, "Test", "RETRY CANCEL", w32.MB_RETRYCANCEL)
	_, _ = messageBox(hwndTop, "Test", "CANCEL TRY CONTINUE", w32.MB_CANCELTRYCONTINUE)

	// newline
	_, _ = messageBox(hwndTop, "Test", "row1\nrow2\nrow3", w32.MB_OK)

	body := `r1
r2
...
rn`
	_, _ = messageBox(hwndTop, "Test", body, w32.MB_OK)

	// Icon
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONSTOP|
		w32.MB_RIGHT| // text right-justified
		// w32.MB_TOPMOST,
		w32.MB_SYSTEMMODAL, // 比使用MB_TOPMOST好
	)
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONQUESTION)
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONWARNING)
	_, _ = messageBox(0, "Test", "OK", w32.MB_OK|w32.MB_ICONINFORMATION)
	// Output
}
