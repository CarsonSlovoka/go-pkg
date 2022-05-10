package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"testing"
)

func TestGetActiveWindow(t *testing.T) {
	user32dll := w32.NewDll[w32.User32DLL](w32.DN_USER32, []w32.ProcName{
		w32.PCGetForegroundWindow,
		w32.PCGetClassName,
		w32.PCGetWindowText,
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
