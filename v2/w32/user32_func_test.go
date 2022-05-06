package w32

import (
	"fmt"
	"syscall"
	"testing"
)

func TestGetActiveWindow(t *testing.T) {
	user32DLL := syscall.NewLazyDLL("User32.dll")

	procGetForegroundWindow := user32DLL.NewProc("GetForegroundWindow")
	curHandle, err := GetForegroundWindow(procGetForegroundWindow)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("current window HWND:", curHandle) // 當前窗口的識別號

	procGetClassName := user32DLL.NewProc("GetClassNameW")
	clsName, err := GetClassNameW(procGetClassName, curHandle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("window class Name:", clsName)

	procGetWindowTextW := user32DLL.NewProc("GetWindowTextW")
	winText, err := GetWindowTextW(procGetWindowTextW, curHandle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("window text Name:", winText)
}
