//go:build windows

package w32

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PNCreateMutex         ProcName = "CreateMutexW"
	PNCloseHandle         ProcName = "CloseHandle"
	PNGetNativeSystemInfo ProcName = "GetNativeSystemInfo"
	PNGetModuleHandle     ProcName = "GetModuleHandleW"
	PNFreeLibrary         ProcName = "FreeLibrary"
	PNGetLastError        ProcName = "GetLastError"
	PNCreateFile          ProcName = "CreateFileW"
	PNWriteFile           ProcName = "WriteFile"
	PNCopyFile            ProcName = "CopyFileW"
	PNFindResource        ProcName = "FindResourceW"
	PNLoadLibrary         ProcName = "LoadLibraryW"
	PNSizeofResource      ProcName = "SizeofResource"
	PNBeginUpdateResource ProcName = "BeginUpdateResourceW"
	PNUpdateResource      ProcName = "UpdateResourceW"
	PNEndUpdateResource   ProcName = "EndUpdateResourceW"
	PNLoadResource        ProcName = "LoadResource"
	PNLockResource        ProcName = "LockResource"
	PNGlobalAlloc         ProcName = "GlobalAlloc"
	PNGlobalLock          ProcName = "GlobalLock"
	PNGlobalUnlock        ProcName = "GlobalUnlock"
	PNGlobalFree          ProcName = "GlobalFree"
)

type Kernel32DLL struct {
	*dLL
}

func NewKernel32DLL(procList ...ProcName) *Kernel32DLL {
	dll := newDll(DNKernel32, procList)
	return &Kernel32DLL{dll}
}

// CloseHandle Closes an open object handle.
// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle?redirectedfrom=MSDN
// Returns TRUE if successful or FALSE otherwise.
func (dll *Kernel32DLL) CloseHandle(handle HANDLE) (bool, syscall.Errno) {
	proc := dll.mustProc(PNCloseHandle)
	// r1, _, err := proc.Call(handle) // 其為syscall.SyscallN的封裝(多了檢查的動作)，如果已經確定，可以直接用syscall.SyscallN會更有效率
	r1, _, errno := syscall.SyscallN(proc.Addr(), uintptr(handle))
	return r1 != 0, errno
}

// CreateMutex You can use it to restrict to a single instance of executable
// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createmutexW#return-value
// If the function fails, the return value is NULL.
func (dll *Kernel32DLL) CreateMutex(lpMutexAttributes *SECURITY_ATTRIBUTES, bInitialOwner bool, name string) (HANDLE, syscall.Errno) {
	proc := dll.mustProc(PNCreateMutex)
	// r1, _, errno := syscall.SyscallN(proc.Addr(), 0, 0, uintptr(unsafe.Pointer(lpName)))
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(lpMutexAttributes)),
		UintptrFromBool(bInitialOwner),
		UintptrFromStr(name),
	)
	/*
		handle, _, err = proc.Call(
			0,
			0,
			uintptr(unsafe.Pointer(lpName)),
		)
		if err.(syscall.Errno) == 0 {
			return handle, nil
		}
	*/
	return HANDLE(r1), errno
}

// GetNativeSystemInfo
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getnativesysteminfo
func (dll *Kernel32DLL) GetNativeSystemInfo() *SYSTEM_INFO {
	info := new(SYSTEM_INFO)
	proc := dll.mustProc(PNGetNativeSystemInfo)
	// _, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(&info))) // 適用回傳 (info SYSTEM_INFO)
	_, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(info)))
	return info
}

// GetModuleHandle https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
// If the function fails, the return value is NULL.
// do not pass a handle returned by GetModuleHandle to the FreeLibrary function.
// Doing so can cause a DLL module to be unmapped prematurely.
// 如果您要取得自己，傳入空字串即可。UintptrFromStr("")會回傳傳0
func (dll *Kernel32DLL) GetModuleHandle(moduleName string) HMODULE {
	proc := dll.mustProc(PNGetModuleHandle)
	hModule, _, _ := syscall.SyscallN(proc.Addr(), UintptrFromStr(moduleName))
	return HMODULE(hModule)
}

// FreeLibrary https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (dll *Kernel32DLL) FreeLibrary(hLibModule HMODULE) bool {
	proc := dll.mustProc(PNFreeLibrary)
	r1, _, _ := syscall.SyscallN(proc.Addr(), uintptr(hLibModule))
	return r1 != 0
}

func (dll *Kernel32DLL) GetLastError() uint32 {
	proc := dll.mustProc(PNGetLastError)
	ret, _, _ := syscall.SyscallN(proc.Addr())
	return uint32(ret)
}

// CreateFile https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
// 不能單靠errno來判斷到底有沒有創建成功，errno應該視為取得更多的創建資訊。
// 如果創建失敗，那麼r1所回傳的數值一定是: INVALID_HANDLE_VALUE (-1)
// 注意！ 不用的時候記得呼叫CloseHandle來關閉
func (dll *Kernel32DLL) CreateFile(lpFileName string, dwDesiredAccess, dwShareMode uint32,
	lpSecurityAttributes uintptr,
	dwCreationDisposition, dwFlagsAndAttributes uint32,
	hTemplateFile uintptr,
) (HANDLE, syscall.Errno) {
	proc := dll.mustProc(PNCreateFile)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(lpFileName),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		lpSecurityAttributes,
		uintptr(dwCreationDisposition),
		uintptr(dwFlagsAndAttributes),
		hTemplateFile)
	return HANDLE(r1), errno
}

// WriteFile https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
// If the function succeeds, the return value is nonzero (TRUE).
func (dll *Kernel32DLL) WriteFile(hFile HANDLE,
	lpBuffer uintptr,
	nNumberOfBytesToWrite uint32,
	lpNumberOfBytesWritten *uint32, // out
	lpOverlapped *OVERLAPPED,
) (bool, syscall.Errno) {
	proc := dll.mustProc(PNWriteFile)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hFile),
		lpBuffer,
		uintptr(nNumberOfBytesToWrite),
		uintptr(unsafe.Pointer(lpNumberOfBytesWritten)),
		uintptr(unsafe.Pointer(lpOverlapped)),
	)
	return r1 != 0, errno
}

// CopyFile https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-copyfilew
// - bFailIfExists: TRUE在已經存在時，會引發錯誤；FALSE如果引經存在則會覆蓋
// If this parameter is TRUE and the new file specified by lpNewFileName already exists, the function fails.
// If this parameter is FALSE and the new file already exists, the function overwrites the existing file and succeeds.
//
// Returns TRUE if successful or FALSE otherwise.
func (dll *Kernel32DLL) CopyFile(existingFileName string, newFileName string, bFailIfExists bool) bool {
	proc := dll.mustProc(PNCopyFile)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(existingFileName),
		UintptrFromStr(newFileName),
		UintptrFromBool(bFailIfExists),
	)
	return r1 != 0
}

// FindResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-findresourcew
// If the function fails, the return value is NULL.
// lpName: 資源的ID或者名稱 MakeIntResource(150)
// lpType: w32.MakeIntResource(w32.RT_GROUP_ICON)
// Resource的資料可能有以下這些，而在每一個分類底下，又有該資源的各個ID
// Icon: RT_ICON
// Icon Group: RT_GROUP_ICON
//
// Version Info: 使用RT_VERSION抓取
//
//	1: 1033
//
// Manifest: 使用RT_MANIFEST來表示
//
//	1: 1033 (ID: 1 語系對應1033即英文)
//
// ...其他的資源類型以此類推
func (dll *Kernel32DLL) FindResource(hModule HMODULE, lpName, lpType *uint16) (HRSRC, syscall.Errno) {
	proc := dll.mustProc(PNFindResource)
	ret, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpType)), // https://learn.microsoft.com/en-us/windows/win32/menurc/resource-types
	)
	return HRSRC(ret), errno
}

// LoadLibrary https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
// If the function fails, the return value is NULL
func (dll *Kernel32DLL) LoadLibrary(lpLibFileName string) HMODULE {
	proc := dll.mustProc(PNLoadLibrary)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(lpLibFileName),
	)
	return HMODULE(r1)
}

// SizeofResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-sizeofresource
// If the function fails, the return value is NULL (0)
func (dll *Kernel32DLL) SizeofResource(hModule HMODULE, hResInfo HRSRC) (uint32, syscall.Errno) {
	proc := dll.mustProc(PNSizeofResource)
	ret, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hModule),
		uintptr(hResInfo),
	)
	return uint32(ret), errno
}

func (dll *Kernel32DLL) MustSizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
	r1, errno := dll.SizeofResource(hModule, hResInfo)
	if r1 == 0 {
		panic(fmt.Sprintf("%s", errno))
	}
	return r1
}

// BeginUpdateResource https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-beginupdateresourceW
// Returns handle if successful or FALSE otherwise.
func (dll *Kernel32DLL) BeginUpdateResource(filePath string, bDeleteExistingResources bool) HANDLE {
	proc := dll.mustProc(PNBeginUpdateResource)

	utf16ptrFilepath, err := syscall.UTF16PtrFromString(filePath) // *uint16, err
	if err != nil {
		panic(err)
	}
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(utf16ptrFilepath)),
		UintptrFromBool(bDeleteExistingResources),
	)
	return HANDLE(r1)
}

// UpdateResource https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-updateresourcew
// Returns TRUE if successful or FALSE otherwise.
// Example: https://learn.microsoft.com/en-us/windows/win32/menurc/using-resources
func (dll *Kernel32DLL) UpdateResource(handle HANDLE,
	lpType uint16, // RT_FONT, RT_DIALOG, ...
	lpName *uint16, // id代號，隨便您定. MakeIntResource(123)
	wLanguage uint16, // MakeLangID(w32.LANG_ENGLISH, w32.SUBLANG_ENGLISH_US)
	lpData uintptr, // ptr to resource info
	cb uint32, // size of resource info. SizeofResource(hExe, hRes)
) bool {
	proc := dll.mustProc(PNUpdateResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(handle),
		uintptr(lpType),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(wLanguage),
		lpData,
		uintptr(cb),
	)
	return r1 == 1
}

// EndUpdateResource https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-endupdateresourcew
// fDiscard: FALSE會真的更新, TRUE僅是測試用，不會更新. Indicates whether to write the resource updates to the file. If this parameter is TRUE, no changes are made. If it is FALSE, the changes are made: the resource updates will take effect.
// Returns TRUE if function succeeds; FALSE otherwise.
func (dll *Kernel32DLL) EndUpdateResource(hUpdate HANDLE, fDiscard bool) bool {
	proc := dll.mustProc(PNEndUpdateResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hUpdate),
		UintptrFromBool(fDiscard),
	)
	return r1 == 1
}

// LoadResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadresource
// If the function fails, the return value is NULL (0)
func (dll *Kernel32DLL) LoadResource(hModule HMODULE, hResInfo HRSRC) (HGLOBAL, syscall.Errno) {
	proc := dll.mustProc(PNLoadResource)
	ret, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hModule),
		uintptr(hResInfo),
	)
	return HGLOBAL(ret), errno
}

// LockResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-lockresource
func (dll *Kernel32DLL) LockResource(hResData HGLOBAL) uintptr {
	proc := dll.mustProc(PNLockResource)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hResData),
	)
	return ret
}

// GlobalAlloc https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
// If the function fails, the return value is NULL.
func (dll *Kernel32DLL) GlobalAlloc(uFlags UINT, dwBytes SIZE_T) (HGLOBAL, syscall.Errno) {
	proc := dll.mustProc(PNGlobalAlloc)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(uFlags),
		uintptr(dwBytes),
	)
	return HGLOBAL(r1), errno
}

// GlobalLock https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
// If the function fails, the return value is NULL.
func (dll *Kernel32DLL) GlobalLock(hMem HGLOBAL) (LPVOID, syscall.Errno) {
	proc := dll.mustProc(PNGlobalLock)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hMem),
	)
	return LPVOID(r1), errno
}

// GlobalUnlock https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
// If the function fails, the return value is zero
func (dll *Kernel32DLL) GlobalUnlock(hMem HGLOBAL) bool {
	proc := dll.mustProc(PNGlobalUnlock)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hMem),
	)
	return r1 != 0
}

// GlobalFree https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
// If the function **succeeds**, the return value is NULL.
// If the function fails, the return value is equal to a handle to the global memory object.
func (dll *Kernel32DLL) GlobalFree(hMem HGLOBAL) HGLOBAL {
	proc := dll.mustProc(PNGlobalFree)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hMem),
	)
	return HGLOBAL(r1)
}
