//go:build windows

package w32

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PNAddFontMemResourceEx ProcName = "AddFontMemResourceEx"
	PNAddFontResource      ProcName = "AddFontResourceW"
	PNAddFontResourceEx    ProcName = "AddFontResourceExW"

	PNBitBlt ProcName = "BitBlt"

	PNCreateCompatibleBitmap ProcName = "CreateCompatibleBitmap"
	PNCreateCompatibleDC     ProcName = "CreateCompatibleDC"
	PNCreateFontIndirect     ProcName = "CreateFontIndirectW"
	PNCreateFont             ProcName = "CreateFontW"
	PNCreateRectRgnIndirect  ProcName = "CreateRectRgnIndirect"
	PNCreateSolidBrush       ProcName = "CreateSolidBrush"

	PNDeleteObject ProcName = "DeleteObject"

	PNEnumFontFamilies ProcName = "EnumFontFamiliesW"
	PNEnumFonts        ProcName = "EnumFontsW"

	PNFillRgn ProcName = "FillRgn"

	PNGetDIBits ProcName = "GetDIBits"
	PNGetObject ProcName = "GetObjectW"

	PNRemoveFontMemResourceEx ProcName = "RemoveFontMemResourceEx"
	PNRemoveFontResource      ProcName = "RemoveFontResourceW"
	PNRemoveFontResourceEx    ProcName = "RemoveFontResourceExW"

	PNSelectObject ProcName = "SelectObject"

	PNSetBkColor        ProcName = "SetBkColor"
	PNSetBkMode         ProcName = "SetBkMode"
	PNSetStretchBltMode ProcName = "SetStretchBltMode"
	PNSetTextColor      ProcName = "SetTextColor"

	PNStretchBlt ProcName = "StretchBlt"

	PNTextOut ProcName = "TextOutW"
)

type Gdi32DLL struct {
	*dLL
}

// NewGdi32DLL You can import the interested process only instead of the whole process.
// We will load the entire process list if you do not give any parameters.
// We encourage you to fill in only the items you are using, as importing the all item will cause unnecessary waste.
func NewGdi32DLL(procList ...ProcName) *Gdi32DLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNAddFontMemResourceEx,
			PNAddFontResource,
			PNAddFontResourceEx,

			PNBitBlt,

			PNCreateCompatibleBitmap,
			PNCreateCompatibleDC,
			PNCreateFontIndirect,
			PNCreateFont,
			PNCreateRectRgnIndirect,
			PNCreateSolidBrush,

			PNDeleteObject,

			PNEnumFontFamilies,
			PNEnumFonts,

			PNFillRgn,

			PNGetDIBits,
			PNGetObject,

			PNRemoveFontMemResourceEx,
			PNRemoveFontResource,
			PNRemoveFontResourceEx,

			PNSelectObject,

			PNSetBkColor,
			PNSetBkMode,
			PNSetStretchBltMode,
			PNSetTextColor,

			PNStretchBlt,

			PNTextOut,
		}
	}
	dll := newDll(DNGdi32, procList)
	return &Gdi32DLL{dll}
}

func RGB(r, g, b byte) COLORREF {
	return COLORREF(r) | (COLORREF(g) << 8) | (COLORREF(b) << 16)
}

// NewFontMemResource 這不是屬於winapi正統的函數，是一個包裝，方便使用AddFontMemResourceEx
// resourceID: 您的字型資源8(RT_FONT)資源下，要取得其子項目的ID代號
// 如果您的resourceID是字串，請使用syscall.UTF16PtrFromString(resourceName)即可轉成*uint16
func NewFontMemResource(hModule HMODULE, resourceID *uint16) (*FontMemResource, syscall.Errno) {
	kernel32dll := NewKernel32DLL(
		PNFindResource,
		PNSizeofResource,
		PNLoadResource,
		PNLockResource,
	)

	hRes, errno := kernel32dll.FindResource(hModule,
		resourceID,
		MakeIntResource(RT_FONT), // 此函數針對Font，所以直接這邊寫死
	)
	if hRes == HRSRC(0) {
		return nil, errno
	}

	size, errno := kernel32dll.SizeofResource(hModule, hRes) // 如果它顯示149008，其實代表149008bytes=>145KB
	if size == 0 {
		return nil, errno
	}

	hLoadRes, errno := kernel32dll.LoadResource(hModule, hRes)
	if hLoadRes == HGLOBAL(0) {
		return nil, errno
	}

	ptr := kernel32dll.LockResource(hLoadRes)
	if ptr == 0 {
		return nil, 0
	}

	numFonts := uint32(0) // 回傳值

	gdi32dll := NewGdi32DLL(PNAddFontMemResourceEx)
	hFontResource := gdi32dll.AddFontMemResourceEx(ptr, size, nil, &numFonts)

	if hFontResource == HANDLE(0) || numFonts == 0 {
		return nil, 0
	}

	return &FontMemResource{hFontResource: hFontResource}, 0
}

// Remove removes the font resource from memory
func (fmr *FontMemResource) Remove() error {
	if fmr.hFontResource != 0 {
		gdi32dll := NewGdi32DLL(PNRemoveFontMemResourceEx)
		if ok := gdi32dll.RemoveFontMemResourceEx(fmr.hFontResource); !ok {
			return fmt.Errorf("RemoveFontMemResourceEx")
		}
		fmr.hFontResource = 0
	}
	return nil
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

// AddFontResource https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourcew
// If the function succeeds, the return value specifies the number of fonts added.
// If the function fails, the return value is zero.
// No extended error information is available.
// 此函數可以添加字型，如果您沒有再調用RemoveFontResource，那麼已經添加的字型會一直等待下次重開機(或登出)後才會被清除
//
// 如果要安裝永久字型有以下兩種方法:
// 1. HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts : 寫入到裡面的字型，如果省略絕對路徑，則表示此字型應該在系統字型資料夾之中: 即%winDir%\Fonts之中可以找到。相關專案參考: https://github.com/CarsonSlovoka/font-install/blob/8b9fb28d0b299ca0ac061e0d0eefc03faf4ea7ad/install_windows.go#L68-L79
// 2. HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts : 這裡面的數值，必須寫入字型檔案的絕對路徑 https://github.com/CarsonSlovoka/font-install/blob/8b9fb28d0b299ca0ac061e0d0eefc03faf4ea7ad/install_windows.go#L50-L66
func (dll *Gdi32DLL) AddFontResource(fontPath string) int {
	proc := dll.mustProc(PNAddFontResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(), UintptrFromStr(fontPath))
	return int(r1)
}

// AddFontResourceEx https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourceexw
// Return value:
// If the function succeeds, the return value specifies the number of fonts added.
// If the function fails, the return value is zero. No extended error information is available.
func (dll *Gdi32DLL) AddFontResourceEx(fontPath string,
	flag uint32, // 可以是FR_PRIVATE或FR_NOT_ENUM,又或者為0，用0與沒有Ex效果相同
	reserved uintptr, // Reserved. Must be zero.
) int {
	proc := dll.mustProc(PNAddFontResourceEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(fontPath),
		uintptr(flag),
		reserved,
	)
	return int(r1)
}

// BitBlt https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-bitblt
// 將src的點集資料傳送到dst中 (類似把圖複製到dst中去)
func (dll *Gdi32DLL) BitBlt(
	dstHDC HDC,
	dstX int32, dstY int32, dstCx int32, dstCy int32,
	srcHDC HDC,
	srcX int32, srcY int32,
	rasterOperation DWORD,
) (bool, syscall.Errno) {
	proc := dll.mustProc(PNBitBlt)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(dstHDC),
		uintptr(dstX),
		uintptr(dstY),
		uintptr(dstCx),
		uintptr(dstCy),
		uintptr(srcHDC),
		uintptr(srcX),
		uintptr(srcY),
		uintptr(rasterOperation),
	)
	return r1 != 0, errno
}

// CreateCompatibleBitmap https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
// If the function fails, the return value is NULL.
func (dll *Gdi32DLL) CreateCompatibleBitmap(hdc HDC, cx int32, cy int32) HBITMAP {
	proc := dll.mustProc(PNCreateCompatibleBitmap)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(cx),
		uintptr(cy),
	)
	return HBITMAP(r1)
}

// CreateCompatibleDC creates a memory device context (DC) compatible with the specified device.
// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
// If the function fails, the return value is NULL.
func (dll *Gdi32DLL) CreateCompatibleDC(hdc HDC) HDC {
	proc := dll.mustProc(PNCreateCompatibleDC)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
	)
	return HDC(r1)
}

// CreateFontIndirect https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
// If the function fails, the return value is NULL.
func (dll *Gdi32DLL) CreateFontIndirect(logFont *LOGFONT) HFONT {
	proc := dll.mustProc(PNCreateFontIndirect)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(logFont)),
	)
	return HFONT(r1)
}

// CreateFont https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontw
// If the function fails, the return value is NULL.
func (dll *Gdi32DLL) CreateFont(
	cHeight int32,
	cWidth int32, // If this value is zero, the font mapper chooses a closest match value.
	cEscapement int32, cOrientation int32, cWeight int32,
	bItalic uint32, bUnderline uint32, bStrikeOut uint32,
	iCharSet uint32, iOutPrecision uint32, iClipPrecision uint32, iQuality uint32, iPitchAndFamily uint32,
	pszFaceName string,
) HFONT {
	proc := dll.mustProc(PNCreateFont)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(cHeight),
		uintptr(cWidth),
		uintptr(cEscapement),
		uintptr(cOrientation),
		uintptr(cWeight),
		uintptr(bItalic),
		uintptr(bUnderline),
		uintptr(bStrikeOut),
		uintptr(iCharSet),
		uintptr(iOutPrecision),
		uintptr(iClipPrecision),
		uintptr(iQuality),
		uintptr(iPitchAndFamily),
		UintptrFromStr(pszFaceName),
	)
	return HFONT(r1)
}

// CreateRectRgnIndirect https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
// If the function fails, the return value is NULL.
func (dll *Gdi32DLL) CreateRectRgnIndirect(lpRect *RECT) HRGN {
	proc := dll.mustProc(PNCreateRectRgnIndirect)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpRect)),
	)
	return HRGN(r1)
}

// DeleteObject https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
// hObject: A handle to a logical pen, brush, font, bitmap, region, or palette.
// If the function succeeds, the return value is nonzero.
// If the specified handle is not valid or is currently selected into a DC, the return value is zero.
func (dll *Gdi32DLL) DeleteObject(hObject HGDIOBJ) bool {
	proc := dll.mustProc(PNDeleteObject)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hObject),
	)
	return r1 != 0
}

// CreateSolidBrush https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
// If the function fails, the return value is NULL.
func (dll *Gdi32DLL) CreateSolidBrush(color COLORREF) HBRUSH {
	proc := dll.mustProc(PNCreateSolidBrush)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(color),
	)
	return HBRUSH(r1)
}

// EnumFontFamilies https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-enumfontfamiliesw
// ENUMLOGFONT
// The return value is the last value returned by the callback function. Its meaning is implementation specific.
func (dll *Gdi32DLL) EnumFontFamilies(hdc HDC,
	lpLogfont string, // NULL => selects and enumerates one font of "each" available type family.
	lpProc EnumFontFamProc, lParam LPARAM) int32 {
	proc := dll.mustProc(PNEnumFontFamilies)

	// https://learn.microsoft.com/en-us/previous-versions/dd162621(v=vs.85)
	lpProcCallback := syscall.NewCallback(func(logFont *ENUMLOGFONT, textMetric *TEXTMETRIC, fontType uint32, lParam LPARAM) uintptr {
		ret := lpProc(logFont, textMetric, fontType, lParam)
		return uintptr(ret)
	})

	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		UintptrFromStr(lpLogfont),
		lpProcCallback,
		uintptr(lParam),
	)
	return int32(r1)
}

// EnumFonts https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-enumfontsw
// LOGFONT
// The return value is the last value returned by the callback function. Its meaning is defined by the application
func (dll *Gdi32DLL) EnumFonts(hdc HDC,
	lpLogfont string, // If NULL => enumerates one font of "each" available typeface.
	lpProc FONTENUMPROC, lParam LPARAM) int32 {
	proc := dll.mustProc(PNEnumFonts)

	// https://learn.microsoft.com/en-us/previous-versions/dd162623(v=vs.85)
	lpProcCallback := syscall.NewCallback(func(logFont *LOGFONT, textMetric *TEXTMETRIC, fontType uint32, lpData LPARAM) uintptr {
		ret := lpProc(logFont, textMetric, fontType, lpData)
		return uintptr(ret)
	})

	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		UintptrFromStr(lpLogfont),
		lpProcCallback,
		uintptr(lParam),
	)
	return int32(r1)
}

// FillRgn https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) FillRgn(hdc HDC, hrgn HRGN, hbr HBRUSH) bool {
	proc := dll.mustProc(PNFillRgn)
	ret1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(hrgn),
		uintptr(hbr))
	return ret1 != 0
}

// GetDIBits retrieves the bits of the specified compatible bitmap and copies them into a buffer as a DIB(Device-Independent Bitmap) using the specified format.
// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibits
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) GetDIBits(hdc HDC, hbm HBITMAP, start UINT, cLines UINT, lpvBits LPVOID, lpbmi *BitmapInfo, usage UINT) int32 {
	proc := dll.mustProc(PNGetDIBits)
	ret1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(hbm),
		uintptr(start),
		uintptr(cLines),
		uintptr(lpvBits),               // [out]
		uintptr(unsafe.Pointer(lpbmi)), // [out]
		uintptr(usage),
	)
	return int32(ret1)
}

// GetObject https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobjectw
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) GetObject(h HANDLE, size int32, output uintptr) int32 {
	proc := dll.mustProc(PNGetObject)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(h),
		uintptr(size),
		output)
	return int32(r1)
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

// RemoveFontResource https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-removefontresourcew
// If the function succeeds, the return value is nonzero.
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) RemoveFontResource(name string) int {
	proc := dll.mustProc(PNRemoveFontResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(name),
		0,
		0)
	return int(r1)
}

// RemoveFontResourceEx https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-removefontresourceexw
// 如果已經呼叫成功，卻還是看的到字型，重開機或者登出就不見了
// If the function succeeds, the return value is nonzero.
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) RemoveFontResourceEx(name string, flag uint32, reserved uintptr) bool {
	proc := dll.mustProc(PNRemoveFontResourceEx)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(name),
		uintptr(flag),
		reserved,
	)
	return r1 != 0
}

// SelectObject https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
// The SelectObject function selects an object into the specified device context (DC)
func (dll *Gdi32DLL) SelectObject(hdc HDC, h HGDIOBJ) HGDIOBJ {
	proc := dll.mustProc(PNSelectObject)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(h),
	)
	return HGDIOBJ(r1)
}

// SetBkColor https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
// If the function fails, the return value is CLR_INVALID.
func (dll *Gdi32DLL) SetBkColor(hdc HDC, color COLORREF) COLORREF {
	proc := dll.mustProc(PNSetBkColor)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(color),
	)
	return COLORREF(r1)
}

// SetBkMode https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
// If the function succeeds, the return value specifies the previous background mode.
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) SetBkMode(hdc HDC,
	mode int32, // OPAQUE, TRANSPARENT
) int32 {
	proc := dll.mustProc(PNSetBkMode)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(mode),
	)
	return int32(r1)
}

// SetStretchBltMode https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) SetStretchBltMode(hdc HDC, mode int32) int32 {
	proc := dll.mustProc(PNSetStretchBltMode)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(mode),
	)
	return int32(r1)
}

// SetTextColor https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextcolor
// If the function fails, the return value is CLR_INVALID.
func (dll *Gdi32DLL) SetTextColor(hdc HDC, color COLORREF) COLORREF {
	proc := dll.mustProc(PNSetTextColor)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(color),
	)
	return COLORREF(r1)
}

// StretchBlt https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-stretchblt
// 與BitBlt類似，但可以提供一些選項
func (dll *Gdi32DLL) StretchBlt(
	dstHDC HDC,
	dstX int32, dstY int32, dstW int32, dstH int32,
	srcHDC HDC,
	srcX int32, srcY int32, srcW int32, srcH int32,
	rasterOperation DWORD,
) bool {
	proc := dll.mustProc(PNStretchBlt)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(dstHDC),
		uintptr(dstX),
		uintptr(dstY),
		uintptr(dstW),
		uintptr(dstH),

		uintptr(srcHDC),
		uintptr(srcX),
		uintptr(srcY),
		uintptr(srcW),
		uintptr(srcH),
		uintptr(rasterOperation),
	)
	return r1 != 0
}

// TextOut https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-textoutw
// If the function fails, the return value is zero.
func (dll *Gdi32DLL) TextOut(hdc HDC, x int32, y int32, lpString string, length int32) bool {
	proc := dll.mustProc(PNTextOut)
	if length == 0 {
		length = int32((len(lpString) / 2) + 1) // 指的是utf16的個數，非utf8
	}
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		UintptrFromStr(lpString),
		uintptr(length),
	)
	return r1 != 0
}
