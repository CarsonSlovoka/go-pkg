//go:build windows

package w32

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	PNBeginUpdateResource ProcName = "BeginUpdateResourceW"

	PNCloseHandle ProcName = "CloseHandle"

	PNCopyFile ProcName = "CopyFileW"

	PNCreateFile               ProcName = "CreateFileW"
	PNCreateMutex              ProcName = "CreateMutexW"
	PNCreateProcess            ProcName = "CreateProcessW"
	PNCreateToolHelp32Snapshot ProcName = "CreateToolhelp32Snapshot"

	PNEndUpdateResource ProcName = "EndUpdateResourceW"

	PNFindResource ProcName = "FindResourceW"

	PNFreeLibrary ProcName = "FreeLibrary"

	PNGetCurrentThread     ProcName = "GetCurrentThread"
	PNGetCurrentThreadId   ProcName = "GetCurrentThreadId"
	PNGetExitCodeProcess   ProcName = "GetExitCodeProcess"
	PNGetLastError         ProcName = "GetLastError"
	PNGetModuleFileName    ProcName = "GetModuleFileNameW"
	PNGetModuleHandle      ProcName = "GetModuleHandleW"
	PNGetNativeSystemInfo  ProcName = "GetNativeSystemInfo"
	PNGetThreadDescription ProcName = "GetThreadDescription"
	PNGetUserDefaultLCID   ProcName = "GetUserDefaultLCID"

	PNGlobalAlloc  ProcName = "GlobalAlloc"
	PNGlobalFree   ProcName = "GlobalFree"
	PNGlobalLock   ProcName = "GlobalLock"
	PNGlobalUnlock ProcName = "GlobalUnlock"

	PNLoadLibrary  ProcName = "LoadLibraryW"
	PNLoadResource ProcName = "LoadResource"
	PNLockResource ProcName = "LockResource"

	PNProcess32First ProcName = "Process32FirstW"
	PNProcess32Next  ProcName = "Process32NextW"

	PNReadDirectoryChanges ProcName = "ReadDirectoryChangesW"

	PNSetLastError         ProcName = "SetLastError"
	PNSetThreadDescription ProcName = "SetThreadDescription"

	PNSizeofResource ProcName = "SizeofResource"

	PNOpenProcess ProcName = "OpenProcess"

	PNTerminateProcess ProcName = "TerminateProcess"

	PNUpdateResource ProcName = "UpdateResourceW"

	PNWaitForSingleObject ProcName = "WaitForSingleObject"

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
			PNCreateProcess,
			PNCreateToolHelp32Snapshot,

			PNEndUpdateResource,

			PNFindResource,

			PNFreeLibrary,

			PNGetCurrentThread,
			PNGetCurrentThreadId,
			PNGetExitCodeProcess,
			PNGetLastError,
			PNGetModuleFileName,
			PNGetModuleHandle,
			PNGetNativeSystemInfo,
			PNGetThreadDescription,
			PNGetUserDefaultLCID,

			PNGlobalAlloc,
			PNGlobalFree,
			PNGlobalLock,
			PNGlobalUnlock,

			PNLoadLibrary,
			PNLoadResource,
			PNLockResource,

			PNProcess32First,
			PNProcess32Next,

			PNReadDirectoryChanges,

			PNSetLastError,
			PNSetThreadDescription,

			PNSizeofResource,

			PNOpenProcess,

			PNTerminateProcess,

			PNUpdateResource,

			PNWaitForSingleObject,

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
func (dll *Kernel32DLL) CloseHandle(handle HANDLE) syscall.Errno {
	proc := dll.mustProc(PNCloseHandle)
	// r1, _, err := proc.Call(handle) // 其為syscall.SyscallN的封裝(多了檢查的動作)，如果已經確定，可以直接用syscall.SyscallN會更有效率
	_, _, errno := syscall.SyscallN(proc.Addr(), uintptr(handle)) // Returns TRUE if successful or FALSE otherwise. 不需要特別管回傳值，如果成功errno會是0
	return errno
}

// CopyFile https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-copyfilew
// - bFailIfExists: TRUE在已經存在時，會引發錯誤；FALSE如果引經存在則會覆蓋
// If this parameter is TRUE and the new file specified by lpNewFileName already exists, the function fails.
// If this parameter is FALSE and the new file already exists, the function overwrites the existing file and succeeds.
//
// Returns TRUE if successful or FALSE otherwise.
func (dll *Kernel32DLL) CopyFile(existingFileName string, newFileName string, bFailIfExists bool) syscall.Errno {
	proc := dll.mustProc(PNCopyFile)
	_, _, errno := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(existingFileName),
		UintptrFromStr(newFileName),
		UintptrFromBool(bFailIfExists),
	)
	return errno
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

// CreateProcess https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessw
func (dll *Kernel32DLL) CreateProcess(applicationName string, cmd string,
	processAtt *SECURITY_ATTRIBUTES, threadAttr *SECURITY_ATTRIBUTES,
	isInherit bool,
	creationFlags uint32,
	lpEnvironment uintptr,
	currentDirectory string,
	startupInfo *STARTUPINFO,
	processInformation *PROCESS_INFORMATION,
) syscall.Errno {
	proc := dll.mustProc(PNCreateProcess)
	_, _, errno := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(applicationName),
		UintptrFromStr(cmd),
		uintptr(unsafe.Pointer(processAtt)),
		uintptr(unsafe.Pointer(threadAttr)),
		UintptrFromBool(isInherit),
		uintptr(creationFlags),
		lpEnvironment,
		UintptrFromStr(currentDirectory),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(processInformation)),
	)
	return errno
}

// CreateToolHelp32Snapshot https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
// If the function fails, it returns INVALID_HANDLE_VALUE.
func (dll *Kernel32DLL) CreateToolHelp32Snapshot(dwFlags uint32, th32ProcessID uint32) (HANDLE, syscall.Errno) {
	proc := dll.mustProc(PNCreateToolHelp32Snapshot)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)
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
func (dll *Kernel32DLL) FreeLibrary(hLibModule HMODULE) syscall.Errno {
	proc := dll.mustProc(PNFreeLibrary)
	_, _, errno := syscall.SyscallN(proc.Addr(), uintptr(hLibModule))
	return errno
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

// GetExitCodeProcess https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess?redirectedfrom=MSDN
// If the function succeeds, the return value is nonzero.
func (dll *Kernel32DLL) GetExitCodeProcess(hProcess HANDLE) (uint32, syscall.Errno) {
	var exitCode uint32
	proc := dll.mustProc(PNGetExitCodeProcess)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&exitCode)),
	)
	return exitCode, eno
}

func (dll *Kernel32DLL) GetLastError() uint32 {
	proc := dll.mustProc(PNGetLastError)
	ret, _, _ := syscall.SyscallN(proc.Addr())
	return uint32(ret)
}

// GetModuleFileName https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (dll *Kernel32DLL) GetModuleFileName(hModule HMODULE) (string, syscall.Errno) {
	proc := dll.mustProc(PNGetModuleFileName)

	modPath := make([]uint16, MAX_PATH)

	// If the function succeeds, the return value is the length of the string that is copied to the buffer, in characters, not including the terminating null character.
	n, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&modPath[0])),
		uintptr(MAX_PATH),
	)
	if eno != 0 {
		return "", eno
	}
	// return syscall.UTF16ToString(modPath), eno
	return string(utf16.Decode(modPath[:n])), eno
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

// GetUserDefaultLCID https://learn.microsoft.com/en-us/windows/win32/api/winnls/nf-winnls-getuserdefaultlcid
// 1033: en-US, 1028: Chinese-Taiwan, ...
// https://learn.microsoft.com/en-us/openspecs/office_standards/ms-oe376/6c085406-a698-4e12-9d4d-c3b0ee3dbc4a
// LcID locale identifier
// LOCALE_SYSTEM_DEFAULT
func (dll *Kernel32DLL) GetUserDefaultLCID() LCID {
	proc := dll.mustProc(PNGetUserDefaultLCID)
	r1, _, _ := syscall.SyscallN(proc.Addr())
	return LCID(r1)
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
// If the memory object is still locked after decrementing the lock count
// If the memory object is unlocked after decrementing the lock count, the function returns zero. and GetLastError returns NO_ERROR. 0也有可能是成功
// If the function fails, the return value is zero and GetLastError returns a value other than NO_ERROR.
// 所以要檢查是否有錯誤，用Errno!=0為主
func (dll *Kernel32DLL) GlobalUnlock(hMem HGLOBAL) (int32, syscall.Errno) {
	proc := dll.mustProc(PNGlobalUnlock)
	r, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hMem),
	)
	return int32(r), errno
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

// Process32First https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
// Returns TRUE if the next entry of the process list has been copied to the buffer or FALSE otherwise
// The ERROR_NO_MORE_FILES error value is returned by the GetLastError function
func (dll *Kernel32DLL) Process32First(hSnapshot HANDLE,
	lppe *PROCESSENTRY32W, // [out]
) (int32, syscall.Errno) {
	proc := dll.mustProc(PNProcess32First)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lppe)),
	)
	return int32(r1), errno
}

// Process32Next https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32nextw
// Returns TRUE if the next entry of the process list has been copied to the buffer or FALSE otherwise
// The ERROR_NO_MORE_FILES error value is returned by the GetLastError function
func (dll *Kernel32DLL) Process32Next(hSnapshot HANDLE, lppe *PROCESSENTRY32W) (int32, syscall.Errno) {
	proc := dll.mustProc(PNProcess32Next)
	r1, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lppe)),
	)
	return int32(r1), errno
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
) syscall.Errno {
	proc := dll.mustProc(PNReadDirectoryChanges)
	_, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hDirectory),
		lpBuffer,
		uintptr(nBufferLength),
		UintptrFromBool(bWatchSubtree),
		uintptr(dwNotifyFilter),
		uintptr(unsafe.Pointer(lpBytesReturned)), // [out]
		uintptr(unsafe.Pointer(lpOverlapped)),
		lpCompletionRoutine,
	)
	return errno
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

// OpenProcess https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess
// If the function fails, the return value is NULL.
func (dll *Kernel32DLL) OpenProcess(desiredAccess uint32, isInheritHandle bool, processID uint32) (HANDLE, syscall.Errno) {
	proc := dll.mustProc(PNOpenProcess)
	r, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(desiredAccess),
		UintptrFromBool(isInheritHandle),
		uintptr(processID),
	)
	return HANDLE(r), errno
}

// TerminateProcess https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminateprocess
func (dll *Kernel32DLL) TerminateProcess(handle HANDLE, exitCode uint32) syscall.Errno {
	proc := dll.mustProc(PNTerminateProcess)
	_, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(handle),
		uintptr(exitCode),
	)
	return errno
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

// WaitForSingleObject https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
// return value: WAIT_ABANDONED, WAIT_OBJECT_0, WAIT_TIMEOUT, WAIT_FAILED
func (dll *Kernel32DLL) WaitForSingleObject(handle HANDLE, milliseconds uint32) uint32 {
	proc := dll.mustProc(PNWaitForSingleObject)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(handle),
		uintptr(milliseconds),
	)
	return uint32(r)
}

// WriteFile https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
// If the function succeeds, the return value is nonzero (TRUE).
func (dll *Kernel32DLL) WriteFile(hFile HANDLE,
	lpBuffer uintptr,
	nNumberOfBytesToWrite uint32,
	lpNumberOfBytesWritten *uint32, // out
	lpOverlapped *OVERLAPPED,
) syscall.Errno {
	proc := dll.mustProc(PNWriteFile)
	_, _, errno := syscall.SyscallN(proc.Addr(),
		uintptr(hFile),
		lpBuffer,
		uintptr(nNumberOfBytesToWrite),
		uintptr(unsafe.Pointer(lpNumberOfBytesWritten)),
		uintptr(unsafe.Pointer(lpOverlapped)),
	)
	return errno
}
