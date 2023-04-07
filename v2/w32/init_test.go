package w32_test

import "github.com/CarsonSlovoka/go-pkg/v2/w32"

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
