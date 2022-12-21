package w32

// VARENUM
const (
	VT_EMPTY            VT = 0
	VT_NULL             VT = 1
	VT_I2               VT = 2
	VT_I4               VT = 3
	VT_R4               VT = 4
	VT_R8               VT = 5
	VT_CY               VT = 6
	VT_DATE             VT = 7
	VT_BSTR             VT = 8
	VT_DISPATCH         VT = 9
	VT_ERROR            VT = 10
	VT_BOOL             VT = 11
	VT_VARIANT          VT = 12
	VT_UNKNOWN          VT = 13
	VT_DECIMAL          VT = 14
	VT_I1               VT = 16
	VT_UI1              VT = 17
	VT_UI2              VT = 18
	VT_UI4              VT = 19
	VT_I8               VT = 20
	VT_UI8              VT = 21
	VT_INT              VT = 22
	VT_UINT             VT = 23
	VT_VOID             VT = 24
	VT_HRESULT          VT = 25
	VT_PTR              VT = 26
	VT_SAFEARRAY        VT = 27
	VT_CARRAY           VT = 28
	VT_USERDEFINED      VT = 29
	VT_LPSTR            VT = 30
	VT_LPWSTR           VT = 31
	VT_RECORD           VT = 36
	VT_INT_PTR          VT = 37
	VT_UINT_PTR         VT = 38
	VT_FILETIME         VT = 64
	VT_BLOB             VT = 65
	VT_STREAM           VT = 66
	VT_STORAGE          VT = 67
	VT_STREAMED_OBJECT  VT = 68
	VT_STORED_OBJECT    VT = 69
	VT_BLOB_OBJECT      VT = 70
	VT_CF               VT = 71
	VT_CLSID            VT = 72
	VT_VERSIONED_STREAM VT = 73
	VT_BSTR_BLOB        VT = 0xfff
	VT_VECTOR           VT = 0x1000
	VT_ARRAY            VT = 0x2000
	VT_BYREF            VT = 0x4000
	VT_RESERVED         VT = 0x8000
	VT_ILLEGAL          VT = 0xffff
	VT_ILLEGALMASKED    VT = 0xfff
	VT_TYPEMASK         VT = 0xfff
)
