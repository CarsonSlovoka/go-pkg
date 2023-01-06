package w32_test

import "github.com/CarsonSlovoka/go-pkg/v2/w32"

var (
	oleDll    *w32.Ole32DLL
	kernelDll *w32.Kernel32DLL
	userDll   *w32.User32DLL
	oleAutDll *w32.OleAut32DLL
)

func init() {
	oleDll = w32.NewOle32DLL()
	kernelDll = w32.NewKernel32DLL()
	oleAutDll = w32.NewOleAut32DLL()
	userDll = w32.NewUser32DLL()
}

func getTestHwnd() w32.HWND {
	hwnd := userDll.FindWindow("Notepad", "")
	if hwnd == 0 {
		hwnd = userDll.GetDesktopWindow()
	}
	return hwnd
}
