//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNExtractIcon ProcName = "ExtractIconW"

	PNShellExecute    ProcName = "ShellExecuteW"
	PNShellExecuteEx  ProcName = "ShellExecuteExW"
	PNShellNotifyIcon ProcName = "Shell_NotifyIconW"
)

type ShellDLL struct {
	*dLL
}

// NewShellDLL You can import the interested process only instead of the whole process.
// We will load the entire process list if you do not give any parameters.
// We encourage you to fill in only the items you are using, as importing the all item will cause unnecessary waste.
func NewShellDLL(procList ...ProcName) *ShellDLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNExtractIcon,

			PNShellExecute,
			PNShellExecuteEx,
			PNShellNotifyIcon,
		}
	}
	dll := newDll(DNShell32, procList)
	return &ShellDLL{dll}
}

// ExtractIcon 可以獲得應用程式的HICON等相關資源
// https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-extracticonw
// exeFileName: 類型可以是{*.exe, *.dll, *.ico}皆可
// nIconIndex:
// - 0 返回第一個圖標句柄
// - -1 返回圖標總數. 如果是exe, dll返回RT_GROUP_ICON資源數量, .ico文件返回1
// - 對於其他不等於-1的負數，表示要取得的圖標資源下標值，例如-3表示取得第三個圖標句柄
func (dll *ShellDLL) ExtractIcon(hInst uintptr, // 透過哪一個對象來呼叫此dll函數，一般用本身應用程式自身0就可以了
	exeFileName string, // {相對路徑, 絕對路徑, 只有運用程式名稱(要系統路徑能找到)}，這三類都可以
	nIconIndex int, // 雖然我們用的是int，但它不影響轉成uintptr的結果: https://go.dev/play/p/kv17S1IfWGB
) HICON {
	proc := dll.mustProc(PNExtractIcon)
	hIcon, _, _ := syscall.SyscallN(proc.Addr(),
		hInst,
		UintptrFromStr(exeFileName),
		uintptr(nIconIndex),
	)
	return HICON(hIcon)
}

// ShellExecute https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew
// If the function succeeds, it returns a value greater than 32 (Hinstance)
// 如果錯誤Hinstance的數值其實和Errno的數值是一樣的
func (dll *ShellDLL) ShellExecute(hwnd HWND,
	operation, // edit, explore, find, print, runas, NULL
	file,
	paras,
	wkDir string, // If this value is NULL or "", the current working directory is used
	showCmd int32, // SW_SHOWNORMAL
) (HINSTANCE, syscall.Errno) {
	proc := dll.mustProc(PNShellExecute)
	r1, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(hwnd),
		UintptrFromStr(operation),
		UintptrFromStr(file),
		UintptrFromStr(paras),
		UintptrFromStr(wkDir),
		uintptr(showCmd),
	)
	return HINSTANCE(r1), eno
}

// ShellExecuteEx https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecuteexw
// Returns TRUE if successful; otherwise, FALSE.
func (dll *ShellDLL) ShellExecuteEx(info *ShellExeCuteInfo) syscall.Errno {
	proc := dll.mustProc(PNShellExecuteEx)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(info)),
	)
	// return r1 != 0, eno
	return eno
}

// ShellNotifyIcon https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
// Returns TRUE if successful, or FALSE otherwise.
func (dll *ShellDLL) ShellNotifyIcon(dwMessage uint32, nid PNOTIFYICONDATA) bool {
	proc := dll.mustProc(PNShellNotifyIcon)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(dwMessage),
		uintptr(unsafe.Pointer(nid)),
	)
	return r1 != 0
}
