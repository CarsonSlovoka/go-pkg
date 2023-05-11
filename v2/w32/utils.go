// 本腳本內所提供的內容都非正統win32api，只是為了方便創建所使用

package w32

type WindowOptions struct {
	ClassName  string
	IconPath   string
	X, Y       int32
	Width      int32
	Height     int32
	ClassStyle uint32 // window class styles: CS_NOCLOSE
	ExStyle    uint32 // Extended window style: WS_EX_TOOLWINDOW
	Style      uint32 // window styles: WS_OVERLAPPEDWINDOW

	WndProc func(hwnd HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) uintptr
}
