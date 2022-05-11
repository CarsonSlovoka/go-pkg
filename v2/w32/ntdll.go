//go:build windows

package w32

// wProductType: https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/wdm/ns-wdm-_osversioninfoexw#remarks
const (
	VER_NT_DOMAIN_CONTROLLER = 0x0000002 // Server
	VER_NT_SERVER            = 0x0000003 // Server, 注意它是domain controller就會輸出VER_NT_DOMAIN_CONTROLLER，而不是VER_NT_SERVER
	VER_NT_WORKSTATION       = 0x0000001 // 如果非Server的版本就是這個
)

// RTL_OSVERSIONINFOEXW
// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/wdm/ns-wdm-_osversioninfoexw
// https://github.com/gonutz/w32
type RTL_OSVERSIONINFOEXW struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16
	ServicePackMajor  uint16
	ServicePackMinor  uint16
	SuiteMask         uint16
	ProductType       byte
	Reserved          byte
}

// wSuiteMask: https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
const (
	VER_SUITE_BACKOFFICE               = 0x00000004
	VER_SUITE_BLADE                    = 0x00000400
	VER_SUITE_COMPUTE_SERVER           = 0x00004000
	VER_SUITE_DATACENTER               = 0x00000080
	VER_SUITE_ENTERPRISE               = 0x00000002
	VER_SUITE_EMBEDDEDNT               = 0x00000040
	VER_SUITE_PERSONAL                 = 0x00000200
	VER_SUITE_SINGLEUSERTS             = 0x00000100
	VER_SUITE_SMALLBUSINESS            = 0x00000001
	VER_SUITE_SMALLBUSINESS_RESTRICTED = 0x00000020
	VER_SUITE_STORAGE_SERVER           = 0x00002000
	VER_SUITE_TERMINAL                 = 0x00000010
	VER_SUITE_WH_SERVER                = 0x00008000 // Windows Home Server is installed.
	VER_SUITE_MULTIUSERTS              = 0x00020000
)
