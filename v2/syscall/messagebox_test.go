package syscall

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"testing"
)

func TestMessageBoxOK(t *testing.T) {
	response := MessageBoxYesNo("Test", "Yes or No")
	switch response {
	case w32.IDYES:
		fmt.Println("Yes")
	case w32.IDNO:
		fmt.Println("No")
	default: // 基本上沒辦法有這個出現，對話框只有Yes,No可以選，右上角的打X也不能按
		fmt.Println("Unknown")
	}
	_ = MessageBoxOK("Test", "OK")
	_ = MessageBoxYesNoCancel("Test", "Yes No Cancel")
	_ = MessageBox("Test", "Help button", w32.MB_HELP)
	_ = MessageBox("Test", "OK CANCEL", w32.MB_OKCANCEL)
	_ = MessageBox("Test", "ABORT RETRY IGNORE", w32.MB_ABORTRETRYIGNORE)
	_ = MessageBox("Test", "RETRY CANCEL", w32.MB_RETRYCANCEL)
	_ = MessageBox("Test", "CANCEL TRY CONTINUE", w32.MB_CANCELTRYCONTINUE)
}
