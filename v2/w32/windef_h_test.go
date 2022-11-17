package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
)

func ExampleMAKEWORD() {
	lParam := uintptr(18677761)
	wParam := uintptr(17)
	virtualKeyCode := w32.LOWORD(wParam)
	keyFlag := w32.HIWORD(lParam)
	var scanCode uint16
	scanCode = uint16(w32.LOBYTE(uintptr(keyFlag)))
	isExtendedKey := (keyFlag & w32.KF_EXTENDED) == w32.KF_EXTENDED // https://learn.microsoft.com/en-us/windows/win32/inputdev/about-keyboard-input?redirectedfrom=MSDN#extended-key-flag
	if isExtendedKey {
		// extended-key flag, 1 if scancode has 0xE0 prefix
		scanCode = w32.MAKEWORD(uintptr(scanCode), 0xE0)
	}

	user32dll := w32.NewUser32DLL(w32.PNMapVirtualKey)
	virtualKeyCode = w32.LOWORD(uintptr(user32dll.MapVirtualKey(uint32(scanCode), w32.MAPVK_VSC_TO_VK_EX)))
	if virtualKeyCode == w32.VK_RCONTROL {
		fmt.Println("VK_RCONTROL")
	}
	// Output:
	// VK_RCONTROL
}
