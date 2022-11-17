//go:build windows

package w32

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PNBeginUpdateResource ProcName = "BeginUpdateResourceW"

	PNCloseHandle ProcName = "CloseHandle"

	PNCopyFile ProcName = "CopyFileW"

	PNCreateFile  ProcName = "CreateFileW"
	PNCreateMutex ProcName = "CreateMutexW"

	PNEndUpdateResource ProcName = "EndUpdateResourceW"

	PNFindResource ProcName = "FindResourceW"

	PNFreeLibrary ProcName = "FreeLibrary"

	PNGetCurrentThread     ProcName = "GetCurrentThread"
	PNGetCurrentThreadId   ProcName = "GetCurrentThreadId"
	PNGetLastError         ProcName = "GetLastError"
	PNGetModuleHandle      ProcName = "GetModuleHandleW"
	PNGetNativeSystemInfo  ProcName = "GetNativeSystemInfo"
	PNGetThreadDescription ProcName = "GetThreadDescription"

	PNGlobalAlloc  ProcName = "GlobalAlloc"
	PNGlobalFree   ProcName = "GlobalFree"
	PNGlobalLock   ProcName = "GlobalLock"
	PNGlobalUnlock ProcName = "GlobalUnlock"

	PNLoadLibrary  ProcName = "LoadLibraryW"
	PNLoadResource ProcName = "LoadResource"
	PNLockResource ProcName = "LockResource"

	PNReadDirectoryChanges ProcName = "ReadDirectoryChangesW"

	PNSetLastError         ProcName = "SetLastError"
	PNSetThreadDescription ProcName = "SetThreadDescription"

	PNSizeofResource ProcName = "SizeofResource"

	PNUpdateResource ProcName = "UpdateResourceW"

	PNWriteFile ProcName = "WriteFile"
)

type Kernel32DLL struct {
	*dLL
}

// NewKernel32DLL You can import the interested process only instead of the whole process.
// We will load the entire process list if you do not give any parameters.
// // We encourage you to fill in only the items you are using, as importing the all item will cause unnecessary waste.
func NewKernel32DLL(procList ...ProcName) *Kernel32DLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNBeginUpdateResource,

			PNCloseHandle,

			PNCopyFile,

			PNCreateFile,
			PNCreateMutex,

			PNEndUpdateResource,

			PNFindResource,

			PNFreeLibrary,

			PNGetCurrentThread,
			PNGetCurrentThreadId,
			PNGetLastError,
			PNGetModuleHandle,
			PNGetNativeSystemInfo,
			PNGetThreadDescription,

			PNGlobalAlloc,
			PNGlobalFree,
			PNGlobalLock,
			PNGlobalUnlock,

			PNLoadLibrary,
			PNLoadResource,
			PNLockResource,

			PNReadDirectoryChanges,

			PNSetLastError,
			PNSetThreadDescription,

			PNSizeofResource,

			PNUpdateResource,

			PNWriteFile,
		}
	}
	dll := newDll(DNKernel32, procList)
	return &Kernel32DLL{dll}
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

// CloseHandle Closes an open object handle.
// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle?redirectedfrom=MSDN
// Returns TRUE if successful or FALSE otherwise.
func (dll *Kernel32DLL) CloseHandle(handle HANDLE) (bool, syscall.Errno) {
	proc := dll.mustProc(PNCloseHandle)
	// r1, _, err := proc.Call(handle) // 其為syscall.SyscallN的封裝(多了檢查的動作)，如果已經確定，可以直接用syscall.SyscallN會更有效率
	r1, _, errno := syscall.SyscallN(proc.Addr(), uintptr(handle))
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

// FreeLibrary https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (dll *Kernel32DLL) FreeLibrary(hLibModule HMODULE) bool {
	proc := dll.mustProc(PNFreeLibrary)
	r1, _, _ := syscall.SyscallN(proc.Addr(), uintptr(hLibModule))
	return r1 != 0
}

// GetCurrentThread https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthread
// The return value is a pseudo handle for the current thread.
func (dll *Kernel32DLL) GetCurrentThread() HANDLE {
	proc := dll.mustProc(PNGetCurrentThread)
	r1, _, _ := syscall.SyscallN(proc.Addr())
	return HANDLE(r1)
}

func (dll *Kernel32DLL) GetCurrentThreadId() uint32 {
	proc := dll.mustProc(PNGetCurrentThreadId)
	r1, _, _ := syscall.SyscallN(proc.Addr())
	return uint32(r1)
}

func (dll *Kernel32DLL) GetLastError() uint32 {
	proc := dll.mustProc(PNGetLastError)
	ret, _, _ := syscall.SyscallN(proc.Addr())
	return uint32(ret)
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

// GetNativeSystemInfo
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getnativesysteminfo
func (dll *Kernel32DLL) GetNativeSystemInfo() *SYSTEM_INFO {
	info := new(SYSTEM_INFO)
	proc := dll.mustProc(PNGetNativeSystemInfo)
	// _, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(&info))) // 適用回傳 (info SYSTEM_INFO)
	_, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(info)))
	return info
}

// GetThreadDescription https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreaddescription
// Returns hResult>=0 if successful
// SUCCEEDED(hResult)
func (dll *Kernel32DLL) GetThreadDescription(hThread HANDLE, threadDesc *uint16) HRESULT {
	proc := dll.mustProc(PNGetThreadDescription)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hThread),
		uintptr(unsafe.Pointer(threadDesc)),
	)
	return HRESULT(r1)
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

// LoadLibrary https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
// If the function fails, the return value is NULL
func (dll *Kernel32DLL) LoadLibrary(lpLibFileName string) HMODULE {
	proc := dll.mustProc(PNLoadLibrary)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(lpLibFileName),
	)
	return HMODULE(r1)
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

// ReadDirectoryChanges https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-readdirectorychangesw?redirectedfrom=MSDN
// TODO: Unknown type(s): LPOVERLAPPED_COMPLETION_ROUTINE
// If the function succeeds, the return value is nonzero.
// If the function fails, the return value is zero.
func (dll *Kernel32DLL) ReadDirectoryChanges(hDirectory HANDLE,
	lpBuffer uintptr, nBufferLength uint32,
	bWatchSubtree bool,
	dwNotifyFilter uint32,
	lpBytesReturned *uint32, // [out]
	lpOverlapped *OVERLAPPED, lpCompletionRoutine uintptr,
) (bool, syscall.Errno) {
	proc := dll.mustProc(PNReadDirectoryChanges)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hDirectory),
		lpBuffer,
		uintptr(nBufferLength),
		UintptrFromBool(bWatchSubtree),
		uintptr(dwNotifyFilter),
		uintptr(unsafe.Pointer(lpBytesReturned)), // [out]
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutine,
	)
	return r1 != 0, errno
}

// SetLastError https://learn.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-setlasterror
func (dll *Kernel32DLL) SetLastError(dwErrCode uint32) {
	proc := dll.mustProc(PNSetLastError)
	_, _, _ = syscall.SyscallN(proc.Addr(), uintptr(dwErrCode))
}

// SetThreadDescription https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setthreaddescription
func (dll *Kernel32DLL) SetThreadDescription(hThread HANDLE, desc string) HRESULT {
	proc := dll.mustProc(PNSetThreadDescription)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hThread),
		UintptrFromStr(desc),
	)
	return HRESULT(r1)
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
