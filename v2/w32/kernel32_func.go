//go:build windows

package w32

import (
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
	PNCopyFile            ProcName = "CopyFileW"
	PNFindResource        ProcName = "FindResourceW"
	PNLoadLibrary         ProcName = "LoadLibraryW"
	PNSizeofResource      ProcName = "SizeofResource"
	PNBeginUpdateResource ProcName = "BeginUpdateResourceW"
	PNUpdateResource      ProcName = "UpdateResourceW"
	PNEndUpdateResource   ProcName = "EndUpdateResourceW"
	PNLoadResource        ProcName = "LoadResource"
	PNLockResource        ProcName = "LockResource"
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
func (dll *Kernel32DLL) CloseHandle(handle uintptr) bool {
	proc := dll.mustProc(PNCloseHandle)
	// r1, _, err := proc.Call(handle) // 其為syscall.SyscallN的封裝(多了檢查的動作)，如果已經確定，可以直接用syscall.SyscallN會更有效率
	r1, _, _ := syscall.SyscallN(proc.Addr(), handle)
	return r1 != 0
}

// CreateMutex You can use it to restrict to a single instance of executable
// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createmutexW#return-value
func (dll *Kernel32DLL) CreateMutex(name string) (handle uintptr, err error) {
	proc := dll.mustProc(PNCreateMutex)
	lpName, _ := syscall.UTF16PtrFromString(name) // LPCWSTR
	handle, _, errno := syscall.SyscallN(proc.Addr(), 0, 0, uintptr(unsafe.Pointer(lpName)))
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
	if errno == 0 {
		return handle, nil
	}
	return handle, errno
}

// GetNativeSystemInfo
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getnativesysteminfo
func (dll *Kernel32DLL) GetNativeSystemInfo() (info SYSTEM_INFO) {
	proc := dll.mustProc(PNGetNativeSystemInfo)
	_, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(&info)))
	return
}

// GetModuleHandle https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
// If the function fails, the return value is NULL.
// do not pass a handle returned by GetModuleHandle to the FreeLibrary function.
// Doing so can cause a DLL module to be unmapped prematurely.
func (dll *Kernel32DLL) GetModuleHandle(lpModuleName *uint16) (hmodule uintptr) {
	proc := dll.mustProc(PNGetModuleHandle)
	hmodule, _, _ = syscall.SyscallN(proc.Addr(), uintptr(unsafe.Pointer(lpModuleName)))
	return hmodule
}

// FreeLibrary https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (dll *Kernel32DLL) FreeLibrary(hLibModule uintptr) bool {
	proc := dll.mustProc(PNFreeLibrary)
	r1, _, _ := syscall.SyscallN(proc.Addr(), hLibModule)
	return r1 != 0
}

func (dll *Kernel32DLL) GetLastError() uint32 {
	proc := dll.mustProc(PNGetLastError)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		0,
		0,
		0)
	return uint32(ret)
}

// CreateFile https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
// error => r1 == INVALID_HANDLE_VALUE
// 不用的時候記得呼叫CloseHandle來關閉
func (dll *Kernel32DLL) CreateFile(lpFileName string, dwDesiredAccess, dwShareMode uint32,
	lpSecurityAttributes uintptr,
	dwCreationDisposition, dwFlagsAndAttributes uint32,
	hTemplateFile uintptr,
) uintptr {
	proc := dll.mustProc(PNCreateFile)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(lpFileName),
		uintptr(dwDesiredAccess),
		uintptr(dwShareMode),
		lpSecurityAttributes,
		uintptr(dwCreationDisposition),
		uintptr(dwFlagsAndAttributes),
		hTemplateFile)
	return r1
}

// CopyFile https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-copyfilew
// - bFailIfExists: TRUE在已經存在時，會引發錯誤；FALSE如果引經存在則會覆蓋
// If this parameter is TRUE and the new file specified by lpNewFileName already exists, the function fails.
// If this parameter is FALSE and the new file already exists, the function overwrites the existing file and succeeds.
//
// Returns TRUE if successful or FALSE otherwise.
func (dll *Kernel32DLL) CopyFile(existingFileName string, newFileName string, bFailIfExists bool) bool {
	proc := dll.mustProc(PNCopyFile)
	var failIfExists uintptr
	if bFailIfExists {
		failIfExists = 1
	} else {
		failIfExists = 0
	}
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(existingFileName),
		UintptrFromStr(newFileName),
		failIfExists,
	)
	return r1 != 0
}

// FindResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-findresourcew
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
func (dll *Kernel32DLL) FindResource(hModule uintptr, lpName, lpType *uint16) HRSRC {
	proc := dll.mustProc(PNFindResource)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		hModule,
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpType)), // https://learn.microsoft.com/en-us/windows/win32/menurc/resource-types
	)
	return HRSRC(ret)
}

// LoadLibrary https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
// If the function fails, the return value is NULL
func (dll *Kernel32DLL) LoadLibrary(lpLibFileName string) (handle uintptr) {
	proc := dll.mustProc(PNLoadLibrary)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(lpLibFileName),
	)
	return r1
}

// SizeofResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-sizeofresource
// If the function fails, the return value is NULL (0)
func (dll *Kernel32DLL) SizeofResource(hModule uintptr, hResInfo HRSRC) uint32 {
	proc := dll.mustProc(PNSizeofResource)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		hModule,
		uintptr(hResInfo),
		0,
	)
	return uint32(ret)
}

// BeginUpdateResource https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-beginupdateresourceW
// Returns handle if successful or FALSE otherwise.
func (dll *Kernel32DLL) BeginUpdateResource(filePath string, bDeleteExistingResources bool) (handle uintptr) {
	proc := dll.mustProc(PNBeginUpdateResource)

	utf16ptrFilepath, err := syscall.UTF16PtrFromString(filePath) // *uint16, err
	if err != nil {
		panic(err)
	}
	var uintptrDeleteExistingResources uintptr
	if bDeleteExistingResources {
		uintptrDeleteExistingResources = 1
	} else {
		uintptrDeleteExistingResources = 0
	}
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(utf16ptrFilepath)),
		uintptrDeleteExistingResources,
	)
	return r1
}

// UpdateResource https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-updateresourcew
// Returns TRUE if successful or FALSE otherwise.
// Example: https://learn.microsoft.com/en-us/windows/win32/menurc/using-resources
func (dll *Kernel32DLL) UpdateResource(handle uintptr,
	lpType uint16, // RT_FONT, RT_DIALOG, ...
	lpName *uint16, // id代號，隨便您定. MakeIntResource(123)
	wLanguage uint16, // MakeLangID(w32.LANG_ENGLISH, w32.SUBLANG_ENGLISH_US)
	lpData uintptr, // ptr to resource info
	cb uint32, // size of resource info. SizeofResource(hExe, hRes)
) bool {
	proc := dll.mustProc(PNUpdateResource)
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		handle,
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
func (dll *Kernel32DLL) EndUpdateResource(hUpdate uintptr, fDiscard bool) bool {
	proc := dll.mustProc(PNEndUpdateResource)
	var uintptrForceDiscard uintptr
	if fDiscard {
		uintptrForceDiscard = 1
	} else {
		uintptrForceDiscard = 0
	}
	r1, _, _ := syscall.SyscallN(proc.Addr(),
		hUpdate,
		uintptrForceDiscard,
	)
	return r1 == 1
}

// LoadResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadresource
// If the function fails, the return value is NULL (0)
func (dll *Kernel32DLL) LoadResource(hModule uintptr, hResInfo HRSRC) (hglobal HGLOBAL) {
	proc := dll.mustProc(PNLoadResource)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		hModule,
		uintptr(hResInfo),
	)
	return HGLOBAL(ret)
}

// LockResource https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-lockresource
func (dll *Kernel32DLL) LockResource(hResData HGLOBAL) uintptr {
	proc := dll.mustProc(PNLockResource)
	ret, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(hResData),
		0,
		0,
	)
	return ret
}
