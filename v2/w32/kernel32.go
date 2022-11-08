//go:build windows

package w32

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type NUMBERFMT struct {
	NumDigits     uint32
	LeadingZero   uint32
	Grouping      uint32
	LpDecimalSep  *uint16
	LpThousandSep *uint16
	NegativeOrder uint32
}

type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

// ACTCTX
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-actctxw
type ACTCTX struct {
	size                  uint32
	Flags                 uint32
	Source                *uint16 // UTF-16 string
	ProcessorArchitecture uint16
	LangID                uint16
	AssemblyDirectory     *uint16 // UTF-16 string
	ResourceName          *uint16 // UTF-16 string
	ApplicationName       *uint16 // UTF-16 string
	Module                HMODULE
}

const MAX_PATH = 260

// GlobalAlloc flags
const (
	GHND          = 0x0042
	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
	GPTR          = GMEM_FIXED | GMEM_ZEROINIT
)

// Predefined locale ids
const (
	LOCALE_CUSTOM_DEFAULT     LCID = 0x0c00
	LOCALE_CUSTOM_UI_DEFAULT  LCID = 0x1400
	LOCALE_CUSTOM_UNSPECIFIED LCID = 0x1000
	LOCALE_INVARIANT          LCID = 0x007f
	LOCALE_USER_DEFAULT       LCID = 0x0400
	LOCALE_SYSTEM_DEFAULT     LCID = 0x0800
)

// LCTYPE constants
const (
	LOCALE_SDECIMAL          LCTYPE = 14
	LOCALE_STHOUSAND         LCTYPE = 15
	LOCALE_SISO3166CTRYNAME  LCTYPE = 0x5a
	LOCALE_SISO3166CTRYNAME2 LCTYPE = 0x68
	LOCALE_SISO639LANGNAME   LCTYPE = 0x59
	LOCALE_SISO639LANGNAME2  LCTYPE = 0x67
)

// SYSTEM_INFO https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type SYSTEM_INFO struct {
	/*
	  union {
	    DWORD dwOemId <- 我們不使用它
	    struct {
	      WORD wProcessorArchitecture // 直接寫下面兩個細項
	      WORD wReserved;
	    }
	  }
	*/
	// OemId uint32
	ProcessorArchitecture uint16 // PROCESSOR_ARCHITECTURE_AMD64, PROCESSOR_ARCHITECTURE_ARM64, PROCESSOR_ARCHITECTURE_IA64, ...
	Reserved              uint16

	PageSize                  uint32
	MinimumApplicationAddress LPCVOID
	MaximumApplicationAddress LPCVOID
	ActiveProcessorMask       *uint32
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

// https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstchangenotificationa
// https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-readdirectorychangesw#parameters
const (
	FILE_NOTIFY_CHANGE_FILE_NAME  uint32 = 0x00000001
	FILE_NOTIFY_CHANGE_DIR_NAME          = 0x00000002
	FILE_NOTIFY_CHANGE_ATTRIBUTES        = 0x00000004
	FILE_NOTIFY_CHANGE_SIZE              = 0x00000008
	FILE_NOTIFY_CHANGE_LAST_WRITE        = 0x00000010
	FILE_NOTIFY_CHANGE_CREATION          = 0x00000040 // Any change to the creation time of files in the watched directory or subtree causes a change notification wait operation to return.
	FILE_NOTIFY_CHANGE_SECURITY          = 0x00000100
)

// https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-file_notify_information
const (
	FILE_ACTION_ADDED            uint32 = 0x00000001
	FILE_ACTION_REMOVED                 = 0x00000002
	FILE_ACTION_MODIFIED                = 0x00000003
	FILE_ACTION_RENAMED_OLD_NAME        = 0x00000004
	FILE_ACTION_RENAMED_NEW_NAME        = 0x00000005
)

// FILE_NOTIFY_INFORMATION
// https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-file_notify_information
type FILE_NOTIFY_INFORMATION struct {
	NextEntryOffset uint32 // 指的是下一筆資料位於整體buffer所在位置的何處 (都是從buffer的頭開始算起)
	Action          uint32
	FileNameLength  uint32
	FileName1       uint16 // 指的是檔案名稱的第一個字, 完整的名稱取法: 為FileName1的下標值開始算起，取FileNameLength長度個 (如果您調用的函數為寬字串函數，那麼給的長度依舊是以byte算起，所以實際上的檔案長度要取其/2) [FileName1.idx:FileName1.idx+FileNameLength/2]
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-charprevexa
const (
	CP_ACP   uint32 = 0 // default ANSI code page.
	CP_OEMCP        = 1 // Use system default OEM code page.
	CP_MACCP        = 2 // Use the system default Macintosh code page.
)
