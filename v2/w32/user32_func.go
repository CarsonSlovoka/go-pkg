package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNFindWindow          ProcName = "FindWindow"
	PNGetForegroundWindow ProcName = "GetForegroundWindow"
	PNGetClassName        ProcName = "GetClassNameW"
	PNGetWindowText       ProcName = "GetWindowTextW"
	PNMessageBox          ProcName = "MessageBoxW"
)

type User32DLL struct {
	*dLL
}

func NewUser32DLL(procList []ProcName) *User32DLL {
	dll := newDll(DN_USER32, procList)
	// dll.mustProc = ...
	return &User32DLL{dll}
}

func (dll *User32DLL) FindWindow(className, windowName string) (hwnd uintptr, err error) {
	proc := dll.mustProc(PNFindWindow)
	lpClassName, _ := syscall.UTF16PtrFromString(className)
	lpWindowName, _ := syscall.UTF16PtrFromString(windowName)
	hwnd, _, err = proc.Call(
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0)
	return
}

// GetForegroundWindow User32.dll 此函數可以獲得當前窗口的HWND
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func (dll *User32DLL) GetForegroundWindow() (hwnd uintptr, err error) {
	proc := dll.mustProc(PNGetForegroundWindow)
	hwnd, _, err = proc.Call()
	if hwnd == 0 {
		return hwnd, err
	}
	return hwnd, nil
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
func (dll *User32DLL) GetClassName(hwnd uintptr) (name string, err error) {
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
func (dll *User32DLL) GetWindowText(hwnd uintptr) (string, error) {
	proc := dll.mustProc(PNGetWindowText)

	maxCount := 256
	textName := make([]uint16, maxCount)
	pTextName := &textName[0]

	r0, _, errno := syscall.SyscallN(proc.Addr(), hwnd,
		uintptr(unsafe.Pointer(pTextName)),
		uintptr(len(textName)),
	)

	if r0 == 0 {
		return "", errno
	}
	return syscall.UTF16ToString(textName), nil
}

// MessageBox
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
func (dll *User32DLL) MessageBox(hwnd uintptr, caption, text string, btnFlag uintptr) (clickBtnValue uintptr, errno error) {
	proc := dll.mustProc(PNMessageBox)
	pCaption, _ := syscall.UTF16PtrFromString(caption)
	pText, _ := syscall.UTF16PtrFromString(text)
	clickBtnValue, _, errno = syscall.SyscallN(proc.Addr(),
		hwnd,
		uintptr(unsafe.Pointer(pText)),
		uintptr(unsafe.Pointer(pCaption)),
		btnFlag,
	)
	return
}
