package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"testing"
)

func TestGetActiveWindow(t *testing.T) {
	user32dll := w32.NewUser32DLL([]w32.ProcName{
		w32.PNGetForegroundWindow,
		w32.PNGetClassName,
		w32.PNGetWindowText,
	})

	curHwnd, err := user32dll.GetForegroundWindow()
	fmt.Println("current window HWND:", curHwnd) // 當前窗口的識別號

	clsName, err := user32dll.GetClassName(curHwnd)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("window class Name:", clsName)

	winText, err := user32dll.GetWindowText(curHwnd)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("window text Name:", winText)
}
