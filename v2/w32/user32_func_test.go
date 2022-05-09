package w32

import (
	"fmt"
	"syscall"
	"testing"
)

func TestGetActiveWindow(t *testing.T) {
	user32DLL := syscall.NewLazyDLL("User32.dll")
	procGetForegroundWindow := User32LazyProc[HWND]{user32DLL.NewProc("GetForegroundWindow")}
	values, err := procGetForegroundWindow.Run()
	if err != nil {
		t.Fatal(err)
	}
	e := values[1]
	if e.Interface() != nil {
		t.Fatal(e.Interface().(error))
	}
	curHandle := values[0].Interface().(uintptr)
	fmt.Println("current window HWND:", curHandle) // 當前窗口的識別號

	procGetClassNameW := User32LazyProc[uintptr]{user32DLL.NewProc("GetClassNameW")}
	clsName, err := procGetClassNameW.GetClassNameW(curHandle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("window class Name:", clsName)

	procGetWindowTextW := User32LazyProc[HWND]{user32DLL.NewProc("GetWindowTextW")}
	values, err = procGetWindowTextW.Run(HWND(curHandle))
	if err != nil {
		t.Fatal(err)
	}
	winText := values[0].String()
	fmt.Println("window text Name:", winText)
}
