package w32

import (
	"syscall"
	"unsafe"
)

// PROCESSENTRY32W https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-processentry32w
// Size: 556
type PROCESSENTRY32W struct {
	DwSize              uint32
	cntUsage            uint32    // 🚮 This member is no longer used and is always set to zero.
	Th32ProcessID       uint32    // PID
	th32DefaultHeapID   ULONG_PTR // 🚮 This member is no longer used and is always set to zero.
	th32ModuleID        uint32    // 🚮This member is no longer used and is always set to zero.
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      int32
	dwFlags             uint32 // 🚮 This member is no longer used, and is always set to zero.
	SzExeFile           [MAX_PATH]uint16
}

func NewPROCESSENTRY32W() *PROCESSENTRY32W {
	var entry PROCESSENTRY32W
	entry.DwSize = uint32(unsafe.Sizeof(entry)) // https://www.autohotkey.com/boards/viewtopic.php?t=65745 // U64: 568, U32: 556
	return &entry
}

func (p *PROCESSENTRY32W) ExeFileName() string {
	return syscall.UTF16ToString(p.SzExeFile[:])
}

// https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
const (
	TH32CS_INHERIT      = 0x80000000
	TH32CS_SNAPALL      = TH32CS_SNAPHEAPLIST | TH32CS_SNAPMODULE | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
)
