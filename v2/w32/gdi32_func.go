//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNAddFontMemResourceEx    ProcName = "AddFontMemResourceEx"
	PNRemoveFontMemResourceEx ProcName = "RemoveFontMemResourceEx"
	PNAddFontResource         ProcName = "AddFontResourceW"
	PNAddFontResourceEx       ProcName = "AddFontResourceExW"
	PNRemoveFontResource      ProcName = "RemoveFontResourceW"
	PNRemoveFontResourceEx    ProcName = "RemoveFontResourceExW"
)

type Gdi32DLL struct {
	*dLL
}

func NewGdi32DLL(procList ...ProcName) *Gdi32DLL {
	dll := newDll(DNGdi32, procList)
	return &Gdi32DLL{dll}
}

// AddFontResource https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourcew
// If the function succeeds, the return value specifies the number of fonts added.
// If the function fails, the return value is zero.
// No extended error information is available.
// 此函數可以添加字型，如果您沒有再調用RemoveFontResource，那麼已經添加的字型會一直等待下次重開機後才會被清除
//
// 如果要安裝永久字型有以下兩種方法:
// 1. HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts : 寫入到裡面的字型，如果省略絕對路徑，則表示此字型應該在系統字型資料夾之中: 即%winDir%\Fonts之中可以找到。相關專案參考: https://github.com/CarsonSlovoka/font-install/blob/8b9fb28d0b299ca0ac061e0d0eefc03faf4ea7ad/install_windows.go#L68-L79
// 2. HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts : 這裡面的數值，必須寫入字型檔案的絕對路徑 https://github.com/CarsonSlovoka/font-install/blob/8b9fb28d0b299ca0ac061e0d0eefc03faf4ea7ad/install_windows.go#L50-L66
func (dll *Gdi32DLL) AddFontResource(fontPath string) int {
	proc := dll.mustProc(PNAddFontResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(), StrToLPCWSTR(fontPath))
	return int(r1)
}

// RemoveFontResource https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-removefontresourcew
// If the function succeeds, the return value is nonzero.
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) RemoveFontResource(name string) int {
	proc := dll.mustProc(PNRemoveFontResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		StrToLPCWSTR(name),
		0,
		0)
	return int(r1)
}

// AddFontResourceEx 使用此函數，可以讓字型只能被自己使用，其他程式不能訪問其資源(用FR_PRIVATE而非FR_NOT_ENUM) https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourceexw
// 一般來說，所有試用版本的字型，它們如果加載到記憶體之中，都不會使用這個函數，因為這樣除了自己以外的應用程式都沒辦法使用到該字型
// 因此使用者取得該試用字型，如果不能在自己熟悉的應用程式中查看效果，將大大降低使用者的體驗。
// reserved: Reserved. Must be zero.
//
// Return value
//
//	If the function succeeds, the return value specifies the number of fonts added.
//	If the function fails, the return value is zero. No extended error information is available.
func (dll *Gdi32DLL) AddFontResourceEx(fontPath string, flag uint32, reserved uintptr) int {
	proc := dll.mustProc(PNAddFontResourceEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		StrToLPCWSTR(fontPath),
		uintptr(flag),
		reserved,
	)
	return int(r1)
}

// RemoveFontResourceEx https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-removefontresourceexw
// If the function succeeds, the return value is nonzero.
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) RemoveFontResourceEx(name string, flag uint32, reserved uintptr) int {
	proc := dll.mustProc(PNRemoveFontResourceEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		StrToLPCWSTR(name),
		uintptr(flag),
		reserved,
	)
	return int(r1)
}

// AddFontMemResourceEx https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontmemresourceex
// 文檔有誤pNumFonts是一個out的項目而非in
func (dll *Gdi32DLL) AddFontMemResourceEx(pFileView uintptr, cjSize uint32, pvReserved unsafe.Pointer, pNumFonts *uint32) HANDLE {
	proc := dll.mustProc(PNAddFontMemResourceEx)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		pFileView,
		uintptr(cjSize),
		uintptr(pvReserved),
		uintptr(unsafe.Pointer(pNumFonts)),
		0,
		0)
	return HANDLE(ret)
}

// RemoveFontMemResourceEx https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-removefontmemresourceex
// If the function succeeds, the return value is nonzero.
// If the function fails, the return value is zero. No extended error information is available.
func (dll *Gdi32DLL) RemoveFontMemResourceEx(h HANDLE) bool {
	proc := dll.mustProc(PNRemoveFontMemResourceEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(h),
		0,
		0)
	return r1 != 0
}

// NewFontMemResource 這不是屬於winapi正統的函數，是一個包裝，方便使用AddFontMemResourceEx
// resourceID: 您的字型資源8(RT_FONT)資源下，要取得其子項目的ID代號
// 如果您的resourceID是字串，請使用syscall.UTF16PtrFromString(resourceName)即可轉成*uint16
func NewFontMemResource(hModule uintptr, resourceID *uint16) (*FontMemResource, error) {
	kernel32dll := NewKernel32DLL(
		PNFindResource,
		PNSizeofResource,
		PNLoadResource,
		PNLockResource,
	)

	hRes := kernel32dll.FindResource(hModule,
		resourceID,
		MakeIntResource(RT_FONT), // 此函數針對Font，所以直接這邊寫死
	)
	if hRes == HRSRC(0) {
		return nil, lastError("FindResource")
	}

	size := kernel32dll.SizeofResource(hModule, hRes) // 如果它顯示149008，其實代表149008bytes=>145KB
	if size == 0 {
		return nil, lastError("SizeofResource")
	}

	hLoadRes := kernel32dll.LoadResource(hModule, hRes)
	if hLoadRes == HGLOBAL(0) {
		return nil, lastError("LoadResource")
	}

	ptr := kernel32dll.LockResource(hLoadRes)
	if ptr == 0 {
		return nil, lastError("LockResource")
	}

	numFonts := uint32(0) // 回傳值

	gdi32dll := NewGdi32DLL(PNAddFontMemResourceEx)
	hFontResource := gdi32dll.AddFontMemResourceEx(ptr, size, nil, &numFonts)

	if hFontResource == HANDLE(0) || numFonts == 0 {
		return nil, lastError("AddFontMemResource")
	}

	return &FontMemResource{hFontResource: hFontResource}, nil
}

// Remove removes the font resource from memory
func (fmr *FontMemResource) Remove() error {
	if fmr.hFontResource != 0 {
		gdi32dll := NewGdi32DLL(PNRemoveFontMemResourceEx)
		if ok := gdi32dll.RemoveFontMemResourceEx(fmr.hFontResource); !ok {
			return lastError("RemoveFontMemResourceEx")
		}
		fmr.hFontResource = 0
	}
	return nil
}
