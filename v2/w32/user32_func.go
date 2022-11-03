//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNFindWindow                  ProcName = "FindWindowW"
	PNFindWindowEx                ProcName = "FindWindowExW"
	PNGetForegroundWindow         ProcName = "GetForegroundWindow"
	PNGetClassName                ProcName = "GetClassNameW"
	PNGetWindowText               ProcName = "GetWindowTextW"
	PNMessageBox                  ProcName = "MessageBoxW"
	PNGetSystemMetrics            ProcName = "GetSystemMetrics"
	PNLoadIcon                    ProcName = "LoadIconW"
	PNGetDC                       ProcName = "GetDC"
	PNReleaseDC                   ProcName = "ReleaseDC"
	PNDrawIcon                    ProcName = "DrawIcon"
	PNGetIconInfo                 ProcName = "GetIconInfo"
	PNPostMessage                 ProcName = "PostMessageW"
	PNSendMessage                 ProcName = "SendMessageW"
	PNLookupIconIdFromDirectoryEx ProcName = "LookupIconIdFromDirectoryEx"
	PNCreateIconFromResourceEx    ProcName = "CreateIconFromResourceEx"
	PNCopyImage                   ProcName = "CopyImage"
	PNGetClientRect               ProcName = "GetClientRect"
)

type User32DLL struct {
	*dLL
}

func NewUser32DLL(procList ...ProcName) *User32DLL {
	dll := newDll(DNUser32, procList)
	// dll.mustProc = ...
	return &User32DLL{dll}
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

// GetForegroundWindow User32.dll 此函數可以獲得當前窗口的HWND
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
// The return value is a handle to the foreground window.
// The foreground window can be NULL in certain circumstances, such as when a window is losing activation.
func (dll *User32DLL) GetForegroundWindow() HWND {
	proc := dll.mustProc(PNGetForegroundWindow)
	hwnd, _, _ := proc.Call()
	return HWND(hwnd)
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

// GetSystemMetrics 依據所傳入的參數回傳您所要查詢的數值資料
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics#parameters
// If the function fails, the return value is 0
func (dll *User32DLL) GetSystemMetrics(targetIdx int32) int32 {
	proc := dll.mustProc(PNGetSystemMetrics)
	r0, _, _ := syscall.SyscallN(proc.Addr(), uintptr(targetIdx))
	return int32(r0)
}

// LoadIcon https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
// If the function fails, the return value is NULL.
func (dll *User32DLL) LoadIcon(hInstance HINSTANCE, lpIconName *uint16) (hIcon HICON, err syscall.Errno) {
	proc := dll.mustProc(PNLoadIcon)
	hwnd, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
	)

	return HICON(hwnd), errno
}

// GetDC LoadIcon https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func (dll *User32DLL) GetDC(hwnd HWND) HDC {
	proc := dll.mustProc(PNGetDC)
	hdc, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
	)
	return HDC(hdc)
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

// PostMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
// If the function succeeds, the return value is nonzero.
func (dll *User32DLL) PostMessage(hwnd HWND, wmMsgID uint32, wParam, lParam uintptr) (bool, syscall.Errno) {
	proc := dll.mustProc(PNPostMessage)
	r1, _, errno := syscall.SyscallN(proc.Addr(), uintptr(hwnd), uintptr(wmMsgID), wParam, lParam)
	return r1 != 0, errno
}

// SendMessage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessage
// 注意，他會等待回應，會造成程式當掉的錯覺 https://social.msdn.microsoft.com/Forums/en-US/6900f74f-6ece-47da-88fc-f9c8bcd40206/sendmessage-api-slow?forum=wpf
func (dll *User32DLL) SendMessage(hwnd HWND, wmMsgID uint32, wParam, lParam uintptr) (r1, r2 uintptr, err error) {
	proc := dll.mustProc(PNSendMessage)
	return syscall.SyscallN(proc.Addr(), uintptr(hwnd), uintptr(wmMsgID), wParam, lParam)
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

// CopyImage https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyimage
// If the function fails, the return value is NULL.
func (dll *User32DLL) CopyImage(h HANDLE, imgType uint32, cx, cy int32, flags uint32) (HANDLE, syscall.Errno) {
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

// GetClientRect https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclientrect
func (dll *User32DLL) GetClientRect(hwnd HWND, lpRect *RECT) (bool, syscall.Errno) {
	proc := dll.mustProc(PNGetClientRect)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpRect)),
	)
	return r1 != 0, errno
}
