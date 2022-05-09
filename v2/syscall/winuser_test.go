package syscall

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
	"testing"
)

func TestMessageBoxOK(t *testing.T) {
	user32DLL := syscall.NewLazyDLL("user32.dll")
	procMessageBox := user32DLL.NewProc("MessageBoxW")

	response := MessageBoxYesNo(procMessageBox, w32.HWND_TOP, "Test", "Yes or No")
	switch response {
	case w32.IDYES:
		fmt.Println("Yes")
	case w32.IDNO:
		fmt.Println("No")
	default: // 基本上沒辦法有這個出現，對話框只有Yes,No可以選，右上角的打X也不能按
		fmt.Println("Unknown")
	}

	_ = MessageBoxOK(procMessageBox, w32.HWND_TOP, "Test", "OK")
	_ = MessageBoxYesNoCancel(procMessageBox, w32.HWND_TOP, "Test", "Yes No Cancel")
	_ = MessageBox(procMessageBox, w32.HWND_TOP, "Test", "OK", w32.MB_OK)
	_ = MessageBoxW(procMessageBox, w32.HWND_TOP, "Test", "Help button", w32.MB_HELP)
	_ = MessageBoxW(procMessageBox, w32.HWND_TOP, "Test", "OK CANCEL", w32.MB_OKCANCEL)
	_ = MessageBoxW(procMessageBox, w32.HWND_TOP, "Test", "ABORT RETRY IGNORE", w32.MB_ABORTRETRYIGNORE)
	_ = MessageBoxW(procMessageBox, w32.HWND_TOP, "Test", "RETRY CANCEL", w32.MB_RETRYCANCEL)
	_ = MessageBoxW(procMessageBox, w32.HWND_TOP, "Test", "CANCEL TRY CONTINUE", w32.MB_CANCELTRYCONTINUE)
}