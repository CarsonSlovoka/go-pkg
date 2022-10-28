//go:build windows

package w32

import (
	"syscall"
)

const (
	PNExtractIcon ProcName = "ExtractIconW"
)

type ShellDLL struct {
	*dLL
}

func NewShellDLL(procList []ProcName) *ShellDLL {
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
	nIconIndex int,
) (hIcon uintptr) {
	proc := dll.mustProc(PNExtractIcon)
	hIcon, _, _ = syscall.SyscallN(proc.Addr(),
		hInst,
		UintptrFromStr(exeFileName),
		uintptr(nIconIndex),
	)
	return hIcon
}
