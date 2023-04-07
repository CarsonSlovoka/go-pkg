//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

type PsApiDLL struct {
	*dLL
}

const (
	PNGetModuleFileNameExA ProcName = "GetModuleFileNameExA" // 得到的是ANSI的字元集
	PNGetModuleFileNameExW ProcName = "GetModuleFileNameExW"

	PNEnumProcesses        ProcName = "EnumProcesses"
	PNEnumProcessModules   ProcName = "EnumProcessModules"
	PNEnumProcessModulesEx ProcName = "EnumProcessModulesEx"
)

func NewPsApiDLL(procList ...ProcName) *PsApiDLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNGetModuleFileNameExA,
			PNGetModuleFileNameExW,

			PNEnumProcesses,
			PNEnumProcessModules,
			PNEnumProcessModulesEx,
		}
	}
	dll := newDll(DNPsApi, procList)
	return &PsApiDLL{dll}
}

// GetModuleFileNameExA https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulefilenameexa
// hProcess: The handle must have the PROCESS_QUERY_INFORMATION and PROCESS_VM_READ
// hModule: A handle to the module. If this parameter is NULL, GetModuleFileNameEx returns the path of the executable file of the process specified in hProcess.
// A所用的字元集是ANSI，Go的字串是Unicode字元集，所以如果有非ASCII的字符，此結果會有問題。
// 不過一般程式名稱都會用英文取名，所以不影響結果，用A會比W的效率好一些
func (dll *PsApiDLL) GetModuleFileNameExA(
	hProcess HANDLE,
	hModule HMODULE,
	size uint32, // The size of the name buffer, in characters.
) (filename string, eno syscall.Errno) {
	proc := dll.mustProc(PNGetModuleFileNameExA)

	name := make([]byte, size)

	_, _, eno = syscall.SyscallN(proc.Addr(),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&name[0])), // [out]
		uintptr(size),
	)
	if eno != 0 {
		return "", eno
	}

	// return string(name), 0 // 可能會打印出很多空字元
	var end uint32
	for end = uint32(0); end < size; end++ {
		if name[end] == 0 {
			break
		}
	}
	return string(name[:end]), 0
}

// GetModuleFileNameExW https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulefilenameexw
// hProcess: The handle must have the PROCESS_QUERY_INFORMATION and PROCESS_VM_READ
// hModule: A handle to the module. If this parameter is NULL, GetModuleFileNameEx returns the path of the executable file of the process specified in hProcess.
func (dll *PsApiDLL) GetModuleFileNameExW(
	hProcess HANDLE,
	hModule HMODULE, // If this parameter is NULL, GetModuleFileNameEx returns the path of the executable file of the process specified in hProcess.
	size uint32, // The size of the name buffer, in characters.
) (filename string, eno syscall.Errno) {
	proc := dll.mustProc(PNGetModuleFileNameExW)

	name := make([]uint16, size)

	_, _, eno = syscall.SyscallN(proc.Addr(),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&name[0])), // [out]
		uintptr(size),
	)
	if eno != 0 {
		return "", eno
	}

	return syscall.UTF16ToString(name), 0
}

// EnumProcesses https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocesses
// 如果只需要找視窗類的項目，可以使用 User32DLL.EnumWindows
// To obtain process handles for the processes whose identifiers you have just obtained, call the OpenProcess function.
func (dll *PsApiDLL) EnumProcesses(size uint32, // 當傳入的size不足時，會自動添加到足夠為止，一次增加1024個. 通常一般使用者都在400以內。
) (
	processIDs []uint32, // 陣列的大小依據參數size而定
	numProcesses uint32, // 實際的長度.
	eno syscall.Errno,
) {
	proc := dll.mustProc(PNEnumProcesses)

	for {
		processIDs = make([]uint32, size)
		var sizeBytes = size * 4
		var bytesReturned uint32
		_, _, eno = syscall.SyscallN(proc.Addr(), // 如果找不到，有可能是在kernel32.dll之中. Kernel32.dll on Windows 7 and Windows Server 2008 R2; Psapi.dll (if PSAPI_VERSION=1) on Windows 7 and Windows Server 2008 R2; Psapi.dll on Windows Server 2008, Windows Vista, Windows Server 2003 and Windows XP
			uintptr(unsafe.Pointer(&processIDs[0])), // [out] A pointer to an array that receives the list of process identifiers.
			uintptr(sizeBytes),                      // The size of the pProcessIds array, in bytes.
			uintptr(unsafe.Pointer(&bytesReturned)), // LpDWORD // [out] The number of bytes returned in the pProcessIds array.
		)
		if eno != 0 {
			return nil, 0, eno
		}
		numProcesses = bytesReturned / 4

		if numProcesses == size { // 假設您的size給1024，但實際上用到的proc只有237個，那麼此數值為237而非1024; 但是當size比實際的長度小或者等於的時候，則會回傳size個，因此當size==numProcesses若要得到確切的資料，會需要重算
			size += 1024
			continue
		}
		break
	}

	return processIDs, numProcesses, 0
}

// EnumProcessModules https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocessmodules
// To control whether a 64-bit application enumerates 32-bit modules, 64-bit modules, or both types of modules, use the EnumProcessModulesEx function.
// ❌ Do not call CloseHandle on any of the handles returned by this function.
// size: 很少程序會超過256個module，至於超過400個更是少之又少，如果不想要重算，可以給512或者乾脆1024 一般使用者占用最多的應該是explorer.exe通常在300左右. 會用差不多63個dll
func (dll *PsApiDLL) EnumProcessModules(hProcess HANDLE,
	size uint32,
) (
	hModules []HMODULE, // 第一筆資料是放該執行檔的路徑, 接著分別為dll, exe
	numModules uint32, // 得到的結果有可能是0
	eno syscall.Errno,
) {
	proc := dll.mustProc(PNEnumProcessModules)

	var actualSize uint32

	for {
		hModules = make([]HMODULE, size)

		_, _, eno = syscall.SyscallN(proc.Addr(),
			uintptr(hProcess),
			uintptr(unsafe.Pointer(&hModules[0])), // [out]
			uintptr(size),
			uintptr(unsafe.Pointer(&actualSize)), // [out]
		)
		if eno != 0 {
			return
		}
		numModules = actualSize / uint32(unsafe.Sizeof(HMODULE(0))) // To determine how many modules were enumerated by the call to EnumProcessModules, divide the resulting value in the lpcbNeeded parameter by sizeof(HMODULE).

		if numModules <= size { // 出來的結果確實有可能會超過給定size，此時如果還想要得到完整的結果就需要重算
			break
		}
		size += 256 // 在不夠的情況下，每次增加256個之後再重跑一次
	}

	return hModules, numModules, 0
}

// EnumProcessModulesEx https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocessmodules
// flag: LIST_MODULES_32BIT, LIST_MODULES_64BIT, LIST_MODULES_ALL, LIST_MODULES_DEFAULT
// ❌ Do not call CloseHandle on any of the handles returned by this function.
func (dll *PsApiDLL) EnumProcessModulesEx(hProcess HANDLE, size, flag uint32) (
	hModules []HMODULE, numModules uint32, eno syscall.Errno,
) {
	proc := dll.mustProc(PNEnumProcessModulesEx)

	var actualSize uint32
	hModules = make([]HMODULE, size)

	_, _, eno = syscall.SyscallN(proc.Addr(),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&hModules[0])), // [out]
		uintptr(size),
		uintptr(unsafe.Pointer(&actualSize)), // [out]
		uintptr(flag),
	)
	if eno != 0 {
		return
	}
	numModules = actualSize / 4
	return hModules, numModules, 0
}
