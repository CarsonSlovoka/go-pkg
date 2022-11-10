package w32

import "unsafe"

// OVERLAPPED https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-overlapped
type OVERLAPPED struct {
	Internal     ULONG_PTR
	InternalHigh ULONG_PTR
	/*
		union { // 表示結構可以拆分，或者可以統一看成某個東西 union一次只會儲存一個變數資料, 而且會已宣告的變數型態size最大的變數空間作為記憶體空間。
		    struct {
		      DWORD Offset     // 0-4
		      DWORD OffsetHigh // 4-8
		    }
		    PVOID Pointer      // 0-8
		  }
	*/
	union1 [8]byte
	HEvent HANDLE
}

func (o *OVERLAPPED) Offset() *DWORD {
	return (*DWORD)(unsafe.Pointer(&o.union1[0]))
}
func (o *OVERLAPPED) OffsetHigh() *DWORD {
	return (*DWORD)(unsafe.Pointer(&o.union1[4]))
}
func (o *OVERLAPPED) Pointer() *PVOID {
	return (*PVOID)(unsafe.Pointer(&o.union1[0]))
}
