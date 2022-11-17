//go:build windows

package w32

import "unsafe"

type (
	ATOM          uint16
	HANDLE        uintptr // Handle 是代表系統的內核對像，如文件句柄，線程句柄，進程句柄。
	HGLOBAL       HANDLE  // HGLOBAL是兼容以前windows 3.x時代的16位內存分配API的數據，表示一個內存塊，這個內存塊是GlobalAlloc分配的，需要GlobalLock才能取得內存塊的指針。 // 不過在Win32下，這個數據和Global內存分配函數沒有用了，保留下來只是為了兼容。而且在Win32下，HGLOBAL就是一個void *的指針，可以不用GlobalLock直接使用來訪問內存。
	HINSTANCE     HANDLE  // HINSTANCE 在win32下與HMODULE是相同的東西，在Win32下還存在主要是因為win16程序使用HINSTANCE來區別task。
	LCID          uint32
	LCTYPE        uint32
	LANGID        uint16
	HMODULE       uintptr // HMODULE 是代表應用程序載入的模塊
	HWINEVENTHOOK HANDLE
	HRSRC         uintptr
	LPVOID        uintptr
)

type (
	HACCEL    HANDLE
	HCURSOR   HANDLE
	HDWP      HANDLE
	HICON     HANDLE
	HMENU     HANDLE
	HMONITOR  HANDLE
	HRAWINPUT HANDLE
	HWND      HANDLE
	HLOCAL    HANDLE
)

type (
	CHAR    byte
	DWORD   uint32
	LPDWORD *uint32
	UINT    uint32
	// INT int32
	UINT8    uint8
	UINT_PTR uintptr

	SHORT int16
	LONG  int32

	ULONG     uint32
	ULONG64   uint64
	ULONGLONG uint64
	ULONG_PTR *uint32

	WORD  uint16
	WCHAR uint16
	PWSTR *WCHAR
	PVOID uintptr

	SCODE LONG
)

type (
	LPCVOID unsafe.Pointer
	LRESULT uintptr
	WPARAM  uintptr
	LPARAM  uintptr
	LPSTR   *CHAR
	LPCSTR  *byte
	LPCWSTR *uint16
	LPWSTR  *uint16
)

type HOOKPROC func(code int32, wParam WPARAM, lParam LPARAM) LRESULT
type HHOOK HANDLE

/* 這些的定義還是有特別的意思在，如果都把它們混為一談，不太恰當
type Handle interface {
	~uintptr // uintptr | HWND | HANDLE //不需要這樣寫，使用~，其表示只要「原型」是如此就算
}

type Hwnd interface {
	~uintptr
}
*/
