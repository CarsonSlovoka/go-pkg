//go:build windows

package w32

const (
	S_OK           = 0x00000000
	S_FALSE        = 0x00000001
	E_UNEXPECTED   = 0x8000FFFF
	E_NOTIMPL      = 0x80004001
	E_OUTOFMEMORY  = 0x8007000E
	E_INVALIDARG   = 0x80070057
	E_NOINTERFACE  = 0x80004002
	E_POINTER      = 0x80004003
	E_HANDLE       = 0x80070006
	E_ABORT        = 0x80004004
	E_FAIL         = 0x80004005
	E_ACCESSDENIED = 0x80070005
	E_PENDING      = 0x8000000A
)

const (
	FALSE = 0
	TRUE  = 1
)

type (
	BOOL    int32
	HRESULT int32
)

const (
	PIPE_CLIENT_END       uint32 = 0x00000000
	PIPE_SERVER_END              = 0x00000001
	FILE_FLAG_OVERLAPPED         = 0x40000000
	FILE_ATTRIBUTE_NORMAL        = 0x00000080
	FILE_SHARE_READ              = 0x00000001
	FILE_SHARE_WRITE             = 0x00000002

	GENERIC_READ                  uint32 = 0x80000000
	GENERIC_WRITE                        = 0x40000000
	GENERIC_EXECUTE                      = 0x20000000
	GENERIC_ALL                          = 0x10000000
	CREATE_NEW                           = 1
	CREATE_ALWAYS                        = 2
	OPEN_EXISTING                        = 3
	OPEN_ALWAYS                          = 4
	TRUNCATE_EXISTING                    = 5
	FILE_FLAG_FIRST_PIPE_INSTANCE        = 0x00080000

	INVALID_HANDLE_VALUE int = -1
)
