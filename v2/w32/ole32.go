package w32

type OLECHAR WCHAR

type BSTR *OLECHAR
type LPOLESTR *OLECHAR

type DISPID LONG

type IRecordInfo struct {
	lpVtbl uintptr
}
