//go:build windows

package w32

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PNAppendMenu ProcName = "AppendMenuW"

	PNBeginPaint ProcName = "BeginPaint"

	PNCallNextHookEx ProcName = "CallNextHookEx"

	PNCloseWindow ProcName = "CloseWindow"

	PNCopyImage ProcName = "CopyImage"

	PNCreateIconFromResourceEx ProcName = "CreateIconFromResourceEx"
	PNCreatePopupMenu          ProcName = "CreatePopupMenu"
	PNCreateWindowEx           ProcName = "CreateWindowExW"

	PNDestroyMenu   ProcName = "DestroyMenu"
	PNDestroyWindow ProcName = "DestroyWindow"

	PNDefWindowProc ProcName = "DefWindowProcW"

	PNDispatchMessage ProcName = "DispatchMessageW"

	PNDrawIcon   ProcName = "DrawIcon"
	PNDrawIconEx ProcName = "DrawIconEx"
	PNDrawText   ProcName = "DrawTextW"

	PNEndPaint ProcName = "EndPaint"

	PNFindWindow   ProcName = "FindWindowW"
	PNFindWindowEx ProcName = "FindWindowExW"

	PNGetActiveWindow          ProcName = "GetActiveWindow"
	PNGetClassName             ProcName = "GetClassNameW"
	PNGetClientRect            ProcName = "GetClientRect"
	PNGetCursorPos             ProcName = "GetCursorPos"
	PNGetDC                    ProcName = "GetDC"
	PNGetDesktopWindow         ProcName = "GetDesktopWindow"
	PNGetForegroundWindow      ProcName = "GetForegroundWindow"
	PNGetIconInfo              ProcName = "GetIconInfo"
	PNGetMessage               ProcName = "GetMessageW"
	PNGetSystemMetrics         ProcName = "GetSystemMetrics"
	PNGetWindowLongPtr         ProcName = "GetWindowLongPtrW"
	PNGetWindowRect            ProcName = "GetWindowRect"
	PNGetWindowText            ProcName = "GetWindowTextW"
	PNGetWindowThreadProcessId ProcName = "GetWindowThreadProcessId"

	PNLoadCursor ProcName = "LoadCursorW"
	PNLoadIcon   ProcName = "LoadIconW"
	PNLoadImage  ProcName = "LoadImageW"

	PNLookupIconIdFromDirectoryEx ProcName = "LookupIconIdFromDirectoryEx"

	PNMapVirtualKey ProcName = "MapVirtualKeyW"
	PNMessageBox    ProcName = "MessageBoxW"

	PNPostMessage     ProcName = "PostMessageW"
	PNPostQuitMessage ProcName = "PostQuitMessage"

	PNRegisterClass  ProcName = "RegisterClassW"
	PNRegisterHotKey ProcName = "RegisterHotKey"

	PNReleaseDC ProcName = "ReleaseDC"

	PNSetActiveWindow     ProcName = "SetActiveWindow"
	PNSetForegroundWindow ProcName = "SetForegroundWindow"
	PNSetMenuItemInfo     ProcName = "SetMenuItemInfoW"
	PNSetRect             ProcName = "SetRect"
	PNSetWindowLongPtr    ProcName = "SetWindowLongPtrW"
	PNSetWindowsHookEx    ProcName = "SetWindowsHookExW"

	PNSendInput   ProcName = "SendInput"
	PNSendMessage ProcName = "SendMessageW"

	PNShowWindow ProcName = "ShowWindow"

	PNTrackPopupMenu   ProcName = "TrackPopupMenu"
	PNTranslateMessage ProcName = "TranslateMessage"

	PNUnhookWindowsHookEx ProcName = "UnhookWindowsHookEx"

	PNUnregisterClass  ProcName = "UnregisterClassW"
	PNUnregisterHotKey ProcName = "UnregisterHotKey"
)

type User32DLL struct {
	*dLL
}

// NewUser32DLL You can import the interested process only instead of the whole process.
// We will load the entire process list if you do not give any parameters.
// We encourage you to fill in only the items you are using, as importing the all item will cause unnecessary waste.
func NewUser32DLL(procList ...ProcName) *User32DLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNAppendMenu,

			PNBeginPaint,

			PNCallNextHookEx,

			PNCloseWindow,

			PNCopyImage,

			PNCreateIconFromResourceEx,
			PNCreatePopupMenu,
			PNCreateWindowEx,

			PNDestroyMenu,
			PNDestroyWindow,

			PNDefWindowProc,

			PNDispatchMessage,

			PNDrawIcon,
			PNDrawIconEx,
			PNDrawText,

			PNEndPaint,

			PNFindWindow,
			PNFindWindowEx,

			PNGetActiveWindow,
			PNGetClassName,
			PNGetClientRect,
			PNGetCursorPos,
			PNGetDC,
			PNGetDesktopWindow,
			PNGetForegroundWindow,
			PNGetIconInfo,
			PNGetMessage,
			PNGetSystemMetrics,
			PNGetWindowLongPtr,
			PNGetWindowRect,
			PNGetWindowText,
			PNGetWindowThreadProcessId,

			PNLoadCursor,
			PNLoadIcon,
			PNLoadImage,

			PNLookupIconIdFromDirectoryEx,

			PNMapVirtualKey,
			PNMessageBox,

			PNPostMessage,
			PNPostQuitMessage,

			PNRegisterClass,
			PNRegisterHotKey,

			PNReleaseDC,

			PNSetActiveWindow,
			PNSetForegroundWindow,
			PNSetMenuItemInfo,
			PNSetRect,
			PNSetWindowLongPtr,
			PNSetWindowsHookEx,

			PNSendInput,
			PNSendMessage,

			PNShowWindow,

			PNTrackPopupMenu,
			PNTranslateMessage,

			PNUnhookWindowsHookEx,

			PNUnregisterClass,
			PNUnregisterHotKey,
		}
	}
	dll := newDll(DNUser32, procList)
	return &User32DLL{dll}
}

// AppendMenu https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-appendmenuw
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) AppendMenu(hMenu HMENU, uFlags uint32, uIDNewItem UINT_PTR, lpNewItem string) (bool, syscall.Errno) {
	proc := dll.mustProc(PNAppendMenu)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(uIDNewItem),
		UintptrFromStr(lpNewItem),
	)
	return r1 != 0, errno
}

// BeginPaint https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint
// Sets the update region of a window to NULL. This clears the region, preventing it from generating subsequent WM_PAINT messages.
// If an application processes a WM_PAINT message but does not call BeginPaint or otherwise clear the update region,
// the application continues to receive WM_PAINT messages as long as the region is not empty.
// In all cases, an application must clear the update region before returning from WM_PAINT message.
//
// If the function fails, the return value is NULL, indicating that no display device context is available.
func (dll *User32DLL) BeginPaint(hWnd HWND, lpPaint *PAINTSTRUCT /* out */) HDC {
	proc := dll.mustProc(PNBeginPaint)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPaint)),
	)
	return HDC(r1)
}

// CallNextHookEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-callnexthookex
func (dll *User32DLL) CallNextHookEx(hhk HHOOK, nCode int32, wParam WPARAM, lParam LPARAM) LRESULT {
	proc := dll.mustProc(PNCallNextHookEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(r1)
}

// CloseWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closewindow
// Minimizes (but does not destroy) the specified window.
// If the function fails, the return value is zero.
func (dll *User32DLL) CloseWindow(hWnd HWND) (bool, syscall.Errno) {
	proc := dll.mustProc(PNCloseWindow)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
	)
	return r1 != 0, errno
}

// CopyImage Creates a new image (icon, cursor, or bitmap)
// 通常我們會使用它，接著後面會用GetObject(hwnd, size, out)來取得資料
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyimage
// If the function fails, the return value is NULL.
func (dll *User32DLL) CopyImage(h HANDLE, imgType uint32, cx, cy int32,
	flags uint32, // This parameter can be one or more of the following values. {LR_DEFAULTCOLOR, LR_DEFAULTSIZE, ...}
) (HANDLE, syscall.Errno) {
	proc := dll.mustProc(PNCopyImage)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(h),
		uintptr(imgType),
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
	)
	return HANDLE(r1), errno
}

// CreateIconFromResourceEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
// flags:
// - LR_DEFAULTCOLOR: Uses the default color format.
// - LR_DEFAULTSIZE
// - LR_MONOCHROME: 單色
// - LR_SHARED
//
// return:
// If the function succeeds, the return value is a handle to the icon or cursor.
// If the function fails, the return value is NULL.
func (dll *User32DLL) CreateIconFromResourceEx(
	presBits uintptr,
	dwResSize uint32,
	fIcon bool,
	dwVer uint32, // must be greater than or equal to 0x00020000 and less than or equal to 0x00030000. This parameter is generally set to 0x00030000.
	cxDesired int,
	cyDesired int,
	flags uint, // combination of the following values:
) HICON {
	proc := dll.mustProc(PNCreateIconFromResourceEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		presBits,
		uintptr(dwResSize),
		UintptrFromBool(fIcon),
		uintptr(dwVer),
		uintptr(cxDesired), uintptr(cyDesired),
		uintptr(flags),
	)
	return HICON(r1)
}

// CreatePopupMenu https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createpopupmenu
// If the function fails, the return value is NULL.
func (dll *User32DLL) CreatePopupMenu() HMENU {
	proc := dll.mustProc(PNCreatePopupMenu)
	r1, _, _ := syscall.SyscallN(proc.Addr())
	return HMENU(r1)
}

// CreateWindowEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
// If the function succeeds, the return value is a handle to the new window.
// If the function fails, the return value is NULL.
func (dll *User32DLL) CreateWindowEx(
	dwExStyle DWORD,
	lpClassName string, lpWindowName string,
	dwStyle DWORD,
	x int32, y int32, nWidth int32, nHeight int32,
	hWndParent HWND,
	hMenu HMENU,
	hInstance HINSTANCE, // A handle to the instance of the module to be associated with the window.
	lpParam uintptr,
) (HWND, syscall.Errno) {
	proc := dll.mustProc(PNCreateWindowEx)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(dwExStyle),
		UintptrFromStr(lpClassName),
		UintptrFromStr(lpWindowName),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam),
	)
	return HWND(r1), errno
}

// DestroyMenu https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroymenu
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) DestroyMenu(hMenu HMENU) (bool, syscall.Errno) {
	proc := dll.mustProc(PNDestroyMenu)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hMenu),
	)
	return r1 != 0, errno
}

// DestroyWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
// If the function fails, the return value is zero.
func (dll *User32DLL) DestroyWindow(hWnd HWND) (bool, syscall.Errno) {
	proc := dll.mustProc(PNDestroyWindow)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
	)
	return r1 != 0, errno
}

// DefWindowProc https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
func (dll *User32DLL) DefWindowProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	proc := dll.mustProc(PNDefWindowProc)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(r1)
}

// DispatchMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessagew
func (dll *User32DLL) DispatchMessage(lpMsg /*const*/ *MSG) LRESULT {
	proc := dll.mustProc(PNDispatchMessage)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return LRESULT(r1)
}

// DrawIcon https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) DrawIcon(hdc HDC, x, y int, hIcon HICON) (bool, syscall.Errno) {
	proc := dll.mustProc(PNDrawIcon)
	hwnd, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(hIcon),
	)

	return hwnd != 0, errno
}

// DrawIconEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawiconex
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) DrawIconEx(hdc HDC,
	xLeft int32, yTop int32, // 作圖位置
	hIcon HICON, // 來源圖像
	cxWidth int32, cyWidth int32, // 目標大小. 當flag設定為DI_DEFAULTSIZE, 會使用SM_CXICON, SM_CYICON來代替. 如果DI_DEFAULTSIZE沒有設定且此數值為0，那麼會用原始圖像大小來取代
	istepIfAniCur uint32,
	hbrFlickerFreeDraw HBRUSH,
	diFlags uint32, // DI_COMPAT, DI_DEFAULTSIZE, DI_IMAGE, ...
) (bool, syscall.Errno) {
	proc := dll.mustProc(PNDrawIconEx)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(hIcon),
		uintptr(cxWidth),
		uintptr(cyWidth),
		uintptr(istepIfAniCur),
		uintptr(hbrFlickerFreeDraw),
		uintptr(diFlags))
	return r1 != 0, errno
}

// DrawText https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawtextw
// If the function fails, the return value is zero.
func (dll *User32DLL) DrawText(hdc HDC, text string,
	textLength int32, // If textLength is -1, then the "text" parameter is assumed to be a pointer to a null-terminated string and DrawText computes the character count automatically.
	lprc *RECT, format UINT) int32 {
	proc := dll.mustProc(PNDrawText)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		UintptrFromStr(text),
		uintptr(textLength),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(format),
	)
	return int32(r1)
}

// EndPaint https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endpaint
// If the caret was hidden by BeginPaint, EndPaint restores the caret to the screen.
// The return value is always nonzero.
func (dll *User32DLL) EndPaint(hWnd HWND, lpPaint *PAINTSTRUCT) bool {
	proc := dll.mustProc(PNEndPaint)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPaint)),
	)
	return r1 != 0
}

// FindWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
// If the function fails, the return value is NULL.
func (dll *User32DLL) FindWindow(className, windowName string) HWND {
	proc := dll.mustProc(PNFindWindow)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(className),
		UintptrFromStr(windowName),
	)
	return HWND(r1)
}

// FindWindowEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindowexw
// If the function fails, the return value is NULL.
func (dll *User32DLL) FindWindowEx(hWndParent, hWndChildAfter uintptr, className, windowName string) HWND {
	proc := dll.mustProc(PNFindWindowEx)
	// lpClassName, _ := syscall.UTF16PtrFromString(className)
	// lpWindowName, _ := syscall.UTF16PtrFromString(windowName)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		hWndParent,
		hWndChildAfter,
		UintptrFromStr(className), // uintptr(unsafe.Pointer(lpClassName)) // 這樣其實也可以，不過如果是NULL就會有問題，要給0
		UintptrFromStr(windowName),
	)
	return HWND(r1)
}

// GetActiveWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
// The return value is the handle to the active window attached to the calling thread's message queue. Otherwise, the return value is NULL.
func (dll *User32DLL) GetActiveWindow() HWND {
	proc := dll.mustProc(PNGetActiveWindow)
	r1, _, _ := syscall.SyscallN(proc.Addr())
	return HWND(r1)
}

// GetClassName If the function succeeds, the return value is the number of characters copied to the buffer
// win32api定義此函數的「參數」說明
// 1. 傳入 hwnd (您想要取得視窗名稱的.hwnd 可以想成是一個窗口id)
// 2. [out] 要輸出的對象, 傳入一個記憶體位址
// 3. 長度大小  由於參數2只告訴電腦取得道的字串該放到哪個記憶體位址，但可以放的長度未知，因此需要參數3來說明
// 回傳的數值為取得到的className長度
// 而我做了以下調整:
// 1. 統一最大長度: 256
// 2. 輸出的記憶體位址自動找
// 回傳名稱而不是長度(如果有需要長度在自己用len即可)
// https://docs.microsoft.com/zh-tw/windows/win32/api/winuser/nf-winuser-getclassname
// https://go.dev/play/p/dKueOJv9Sx
func (dll *User32DLL) GetClassName(hwnd HWND) (name string, err error) {
	proc := dll.mustProc(PNGetClassName)

	maxCount := 256
	clsName := make([]uint16, maxCount)
	pClsName := &clsName[0]

	/* 用syscall.SyscallN效率似乎比較好
	r0, _, err = proc.Call(uintptr(hwnd),
		uintptr(unsafe.Pointer(pClsName)),
		uintptr(len(clsName)),
	)
	*/
	r0, _, errno := syscall.SyscallN(proc.Addr(), uintptr(hwnd),
		uintptr(unsafe.Pointer(pClsName)),
		uintptr(len(clsName)), // 等同len(maxCount)
	)

	if r0 == 0 {
		if errno != 0 {
			// err = error(errno) // syscall.errno也屬於error，不需要特別轉型
			return "", errno
		}
		return "", syscall.EINVAL
	}
	name = syscall.UTF16ToString(clsName)
	return // 由於我們的回傳值皆以具名，故當省略回傳項目時，會直接以具名變數取代
}

// GetClientRect https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (dll *User32DLL) GetClientRect(hwnd HWND, lpRect *RECT /* out */) (bool, syscall.Errno) {
	proc := dll.mustProc(PNGetClientRect)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpRect)),
	)
	return r1 != 0, errno
}

// GetCursorPos https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
// Returns nonzero if successful or zero otherwise.
func (dll *User32DLL) GetCursorPos(
	lpPoint *POINT, // [out]
) (bool, syscall.Errno) {
	proc := dll.mustProc(PNGetCursorPos)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpPoint)),
	)
	return r1 != 0, errno
}

// GetDC LoadIcon https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
// 🧙 Call ReleaseDC(hwnd, hdc) when you are not used.
// the default font is "System"
func (dll *User32DLL) GetDC(hwnd HWND) HDC {
	proc := dll.mustProc(PNGetDC)
	hdc, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
	)
	return HDC(hdc)
}

// GetDesktopWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdesktopwindow
func (dll *User32DLL) GetDesktopWindow() HWND {
	proc := dll.mustProc(PNGetDesktopWindow)
	r1, _, _ := syscall.SyscallN(proc.Addr())
	return HWND(r1)
}

// GetForegroundWindow User32.dll 此函數可以獲得當前窗口的HWND
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
// The return value is a handle to the foreground window.
// The foreground window can be NULL in certain circumstances, such as when a window is losing activation.
func (dll *User32DLL) GetForegroundWindow() HWND {
	proc := dll.mustProc(PNGetForegroundWindow)
	hwnd, _, _ := proc.Call()
	return HWND(hwnd)
}

// GetIconInfo https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
// Returns TRUE if successful or FALSE otherwise.
func (dll *User32DLL) GetIconInfo(hIcon HICON, pIconInfo *ICONINFO) bool {
	proc := dll.mustProc(PNGetIconInfo)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(pIconInfo)),
	)
	return r1 != 0
}

// GetMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
// If the function retrieves a message other than WM_QUIT, the return value is nonzero.
// If the function retrieves the WM_QUIT message, the return value is zero
// If there is an error, the return value is -1, To get extended error information, call GetLastError.
func (dll *User32DLL) GetMessage(lpMsg *MSG, hWnd HWND, wMsgFilterMin uint32, wMsgFilterMax uint32) (int32, syscall.Errno) {
	proc := dll.mustProc(PNGetMessage)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return int32(r1), errno
}

// GetSystemMetrics 依據所傳入的參數回傳您所要查詢的數值資料
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics#parameters
// If the function fails, the return value is 0
func (dll *User32DLL) GetSystemMetrics(targetIdx int32) int32 {
	proc := dll.mustProc(PNGetSystemMetrics)
	r0, _, _ := syscall.SyscallN(proc.Addr(), uintptr(targetIdx))
	return int32(r0)
}

// GetWindowLongPtr https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
// If the function fails, the return value is zero.
func (dll *User32DLL) GetWindowLongPtr(hWnd HWND, nIndex int32) (uintptr, syscall.Errno) {
	proc := dll.mustProc(PNGetWindowLongPtr)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(nIndex),
	)
	return r1, errno
}

// GetWindowRect https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
// If the function fails, the return value is zero.
func (dll *User32DLL) GetWindowRect(hWnd HWND,
	lpRect *RECT, // [out]
) (bool, syscall.Errno) {
	proc := dll.mustProc(PNGetWindowRect)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
	)
	return r1 != 0, errno
}

// GetWindowText
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
func (dll *User32DLL) GetWindowText(hwnd HWND) (string, error) {
	proc := dll.mustProc(PNGetWindowText)

	maxCount := 256
	textName := make([]uint16, maxCount)
	pTextName := &textName[0]

	r0, _, errno := syscall.SyscallN(proc.Addr(), uintptr(hwnd),
		uintptr(unsafe.Pointer(pTextName)),
		uintptr(len(textName)),
	)

	if r0 == 0 {
		return "", errno
	}
	return syscall.UTF16ToString(textName), nil
}

// GetWindowThreadProcessId https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowthreadprocessid
// The return value is the identifier of the thread that created the window.
func (dll *User32DLL) GetWindowThreadProcessId(hWnd HWND, lpdwProcessId *uint32) uint32 {
	proc := dll.mustProc(PNGetWindowThreadProcessId)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpdwProcessId)),
	)
	return uint32(r1)
}

// LoadCursor https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (dll *User32DLL) LoadCursor(hInstance HINSTANCE, lpCursorName *uint16) (HCURSOR, syscall.Errno) {
	proc := dll.mustProc(PNLoadCursor)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpCursorName)),
		0)
	return HCURSOR(r1), errno
}

func (dll *User32DLL) MustLoadCursor(hInstance HINSTANCE, lpCursorName *uint16) HCURSOR {
	r1, errno := dll.LoadCursor(hInstance, lpCursorName)
	if r1 == 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return r1
}

// LoadIcon https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
// If the function fails, the return value is NULL.
func (dll *User32DLL) LoadIcon(hInstance HINSTANCE, lpIconName *uint16) (HICON, syscall.Errno) {
	proc := dll.mustProc(PNLoadIcon)
	hwnd, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
	)
	return HICON(hwnd), errno
}

func (dll *User32DLL) MustLoadIcon(hInstance HINSTANCE, lpIconName *uint16) HICON {
	r1, errno := dll.LoadIcon(hInstance, lpIconName)
	if r1 == 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return r1
}

// LoadImage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
// If the function succeeds, the return value is the handle of the newly loaded image.
// If the function fails, the return value is NULL
func (dll *User32DLL) LoadImage(hInst HINSTANCE, name string, aType uint32, cx int32, cy int32, fuLoad uint32) (HANDLE, syscall.Errno) {
	proc := dll.mustProc(PNLoadImage)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hInst),
		UintptrFromStr(name),
		uintptr(aType),
		uintptr(cx),
		uintptr(cy),
		uintptr(fuLoad),
	)
	return HANDLE(r1), errno
}

func (dll *User32DLL) MustLoadImage(hInst HINSTANCE, name string, aType uint32, cx int32, cy int32, fuLoad uint32) HANDLE {
	handle, errno := dll.LoadImage(hInst, name, aType, cx, cy, fuLoad)
	if handle == 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return handle
}

// LookupIconIdFromDirectoryEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-lookupiconidfromdirectoryex
// If the function succeeds, the return value is an integer resource identifier for the icon or cursor that best fits the current display device.
// If the function fails, the return value is zero.
func (dll *User32DLL) LookupIconIdFromDirectoryEx(presBits uintptr,
	fIcon bool, // Indicates whether an icon or a cursor is sought. If this parameter is TRUE, the function is searching for an icon; if the parameter is FALSE, the function is searching for a cursor.
	cxDesired, // The desired width, in pixels, of the icon. If this parameter is zero, the function uses the SM_CXICON or SM_CXCURSOR system metric value.
	cyDesired int, // 0, SM_CYICON, SM_CYCURSOR
	flags uint, // LR_DEFAULTCOLOR or LR_MONOCHROME
) int {
	proc := dll.mustProc(PNLookupIconIdFromDirectoryEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		presBits,
		UintptrFromBool(fIcon),
		uintptr(cxDesired), uintptr(cyDesired),
		uintptr(flags),
	)
	return int(r1)
}

// MapVirtualKey https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeyw
// 可以區分、不區分左右的shift等等
// If there is no translation, the return value is zero.
func (dll *User32DLL) MapVirtualKey(uCode uint32, uMapType uint32) uint32 {
	proc := dll.mustProc(PNMapVirtualKey)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(uCode),
		uintptr(uMapType),
	)
	return uint32(r1)
}

// MessageBox https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
// If the function fails, the return value is zero.
func (dll *User32DLL) MessageBox(hwnd HWND, text, caption string, btnFlag uint32) (clickBtnValue uintptr, errno syscall.Errno) {
	proc := dll.mustProc(PNMessageBox)
	clickBtnValue, _, errno = syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
		UintptrFromStr(text),
		UintptrFromStr(caption),
		uintptr(btnFlag),
	)
	return clickBtnValue, errno
}

// PostMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) PostMessage(hwnd HWND, wmMsgID uint32, wParam, lParam uintptr) (bool, syscall.Errno) {
	proc := dll.mustProc(PNPostMessage)
	r1, _, errno := syscall.SyscallN(proc.Addr(), uintptr(hwnd), uintptr(wmMsgID), wParam, lParam)
	return r1 != 0, errno
}

// PostQuitMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func (dll *User32DLL) PostQuitMessage(nExitCode int32) {
	proc := dll.mustProc(PNPostQuitMessage)
	_, _, _ = syscall.SyscallN(proc.Addr(),
		uintptr(nExitCode),
	)
}

// RegisterClass https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassw
// If the function fails, the return value is zero.
func (dll *User32DLL) RegisterClass(lpWndClass /* const */ *WNDCLASS) (ATOM, syscall.Errno) {
	proc := dll.mustProc(PNRegisterClass)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpWndClass)),
	)
	return ATOM(r1), errno
}

// RegisterHotKey https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerhotkey?redirectedfrom=MSDN
func (dll *User32DLL) RegisterHotKey(hWnd HWND, id int32, fsModifiers uint32, vk uint32) (bool, syscall.Errno) {
	proc := dll.mustProc(PNRegisterHotKey)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vk),
	)
	return r1 != 0, errno
}

// ReleaseDC https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
// If the DC was released, the return value is 1 or 0 otherwise.
func (dll *User32DLL) ReleaseDC(hwnd HWND, hdc HDC) int32 {
	proc := dll.mustProc(PNReleaseDC)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
		uintptr(hdc),
	)
	return int32(r1)
}

// SetActiveWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setactivewindow
// If the function fails, the return value is NULL.
func (dll *User32DLL) SetActiveWindow(hWnd HWND) (HWND, syscall.Errno) {
	proc := dll.mustProc(PNSetForegroundWindow)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
	)
	return HWND(r1), errno
}

// SetForegroundWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
// If the window was brought to the foreground, the return value is nonzero.
// If the window was not brought to the foreground, the return value is zero.
func (dll *User32DLL) SetForegroundWindow(hWnd HWND) bool {
	proc := dll.mustProc(PNSetForegroundWindow)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
	)
	return r1 != 0
}

// SetMenuItemInfo https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmenuiteminfow
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) SetMenuItemInfo(hmenu HMENU, item UINT, fByPosition bool, lpmii /*const*/ *MENUITEMINFO) (bool, syscall.Errno) {
	proc := dll.mustProc(PNSetMenuItemInfo)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hmenu),
		uintptr(item),
		UintptrFromBool(fByPosition),
		uintptr(unsafe.Pointer(lpmii)),
	)
	return r1 != 0, errno
}

// SetRect https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setrect
// If the function fails, the return value is zero.
func (dll *User32DLL) SetRect(lprc *RECT, // [out]
	xLeft int32, yTop int32, xRight int32, yBottom int32) bool {
	proc := dll.mustProc(PNSetRect)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(xRight),
		uintptr(yBottom),
	)
	return r1 != 0
}

// SetWindowLongPtr https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongptrw
// If the function fails, the return value is zero.
func (dll *User32DLL) SetWindowLongPtr(hWnd HWND, nIndex int32, dwNewLong uintptr) (uintptr, syscall.Errno) {
	proc := dll.mustProc(PNSetWindowLongPtr)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(nIndex),
		dwNewLong,
	)
	return r1, errno
}

// SetWindowsHookEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
// If the function fails, the return value is NULL.
func (dll *User32DLL) SetWindowsHookEx(idHook int32, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId uint32) (HHOOK, syscall.Errno) {
	proc := dll.mustProc(PNSetWindowsHookEx)
	lpfnCallback := syscall.NewCallback(func(codeRawArg int32, wParamRawArg WPARAM, lParamRawArg LPARAM) uintptr {
		ret := lpfn(codeRawArg, wParamRawArg, lParamRawArg)
		return uintptr(ret)
	})
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(idHook),
		lpfnCallback,
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(r1), errno
}

// SendInput https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendinput
// If the function returns zero
func (dll *User32DLL) SendInput(arraySize uint32,
	pInputs *INPUT, // An "array" of INPUT structures.
	cbSize int32, // The size, in bytes, of "an" INPUT structure
) (uint32, syscall.Errno) {
	proc := dll.mustProc(PNSendInput)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(arraySize),
		uintptr(unsafe.Pointer(pInputs)),
		uintptr(cbSize))
	return uint32(r1), errno
}

// SendMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessage
// 注意，他會等待回應，會造成程式當掉的錯覺 https://social.msdn.microsoft.com/Forums/en-US/6900f74f-6ece-47da-88fc-f9c8bcd40206/sendmessage-api-slow?forum=wpf
func (dll *User32DLL) SendMessage(hwnd HWND, wmMsgID uint32, wParam, lParam uintptr) (r1, r2 uintptr, err error) {
	proc := dll.mustProc(PNSendMessage)
	return syscall.SyscallN(proc.Addr(), uintptr(hwnd), uintptr(wmMsgID), wParam, lParam)
}

// ShowWindow https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
func (dll *User32DLL) ShowWindow(hWnd HWND, nCmdShow int32) bool {
	proc := dll.mustProc(PNShowWindow)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(nCmdShow),
	)
	return r1 != 0
}

// TrackPopupMenu https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
// 當TPM_RETURNCMD有設定的時候，且回傳值>0，表示該wParam的識別號，即:
// cmd := TrackPopupMenu(hMenu, TPM_RETURNCMD, x, y, 0, hwnd, 0)
// if (cmd) { SendMessage(hwnd, WM_COMMAND, cmd, 0) }
// 當TPM_RETURNCMD沒有設定，回傳值!=0時，表示成功，且會自動發送命令，不需要再靠SendMessage去發送消息, 若回傳值0表示失敗，沒任何失敗訊息
func (dll *User32DLL) TrackPopupMenu(hMenu HMENU, uFlags uint32, x int32, y int32, nReserved int32, hWnd HWND,
	prcRect /*const*/ *RECT, // [in, optional] Ignored.
) (int32, syscall.Errno) {
	proc := dll.mustProc(PNTrackPopupMenu)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(x),
		uintptr(y),
		uintptr(nReserved),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(prcRect)),
	)
	return int32(r1), errno
}

// TranslateMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
// If the message is not translated, the return value is zero.
func (dll *User32DLL) TranslateMessage(lpMsg /*const*/ *MSG) bool {
	proc := dll.mustProc(PNTranslateMessage)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return r1 != 0
}

// UnhookWindowsHookEx https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unhookwindowshookex
// If the function fails, the return value is zero.
func (dll *User32DLL) UnhookWindowsHookEx(hhk HHOOK) (bool, syscall.Errno) {
	proc := dll.mustProc(PNUnhookWindowsHookEx)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hhk),
	)
	return r1 != 0, errno
}

// UnregisterClass https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) UnregisterClass(lpClassName string, hInstance HINSTANCE) (bool, syscall.Errno) {
	proc := dll.mustProc(PNUnregisterClass)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(lpClassName),
		uintptr(hInstance),
	)
	return r1 != 0, errno
}

// UnregisterHotKey https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterhotkey
func (dll *User32DLL) UnregisterHotKey(hWnd HWND, id int32) (bool, syscall.Errno) {
	proc := dll.mustProc(PNUnregisterHotKey)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hWnd),
		uintptr(id),
	)
	return r1 != 0, errno
}
