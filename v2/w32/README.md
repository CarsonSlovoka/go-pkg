## Go使用windows dll的技巧

```go
package main

import (
  "log"
  "syscall"
  "unsafe"
)

func main() {
  dll := syscall.NewLazyDLL("Shell32.dll")                           // 載入dll
  proc := dll.NewProc("ExtractIconW")                                // 抓取dll的方法
  lpcwstrExeFileName, _ := syscall.UTF16PtrFromString("notepad.exe") // 準備該方法中需要的參數
  hIcon, _, _ := syscall.SyscallN(proc.Addr(), // 使用該方法
    0,
    uintptr(unsafe.Pointer(lpcwstrExeFileName)),
    0,
  )
  log.Println(hIcon)
}
```

## [型別與GO的對照](https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types)

| C++ | Go  |
|-----|-----|
BOOLEAN | byte
BYTE | byte
UCHAR | byte
| |
USHORT | uint16
WORD | uint16
WORD | uint16
WCHAR  xxx[128] | [128]uint16
LPWSTR | *uint16
LPCWSTR | *uint16
| |
**BOOL** | int32 你會看到文檔在該回傳值為此類型時，常常只有提到0表示錯誤，但不會說TRUE代表什麼，因為非0的數值不見得一定都是1，有可能是其他的數值。
DWORD | uint32
LPDWORD   | *uint32
DWORD_PTR | *uint32
| |
DWORD64 | uint64
LARGE_INTEGER | int64
| |
ULONGLONG | uint64
| |
CHAR | int8
LPCSTR | *uint8
| |
LONG | int32
| |
LPVOID | uintptr
HANDLE | uintptr
SIZE_T | uintptr

## 使用go來調用windows api的專案

### [kbinani/win](https://github.com/kbinani/win)

`lxn/win`沒有把dll中的所有方法都實現，相較於`lxn/win`這個較為齊全。

### [lxn/win]

常數數值大部分都是參考此專案[lxn/win]，這專案很棒，省去您花很多時間去找微軟文件所定義的該數值為何

- [user32.go]
- [gdi32.go]
- [win.go]
- [kernel32.go-1], [kernel32.go-2]

不用擔心宣告太多變數執行檔會變得很大，實際上就[user32.go]的這些內容，打包出來的執行檔差異也只有10KB左右<sup>就算以上四項全部包含進去也才影響14KB</sup>，影響很小

至於struct的定義，宣告再多都不會影響執行檔大小(看的是用了多少)

const的內容在`-ldflags -s -w`啟用的狀態下，也不會影響(仍然是有用到才會計算)，而就算沒啟用，實際上所增加的大小也非常小

如果要把`lxn/win`所有的init也包含進去<sup>
例如其中之一的[ user32.go.init](https://github.com/lxn/win/blob/7a0e89e/user32.go#L1903-L2059)</sup>

此包的所有init都涵蓋近來約莫會增加368KB，好處是使用上更加直接，缺點就是執行檔變大一些

[lxn/win]: https://github.com/lxn/win

[user32.go]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/user32.go#L18-L1744

[gdi32.go]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/gdi32.go#L16-L1038

[win.go]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/win.go#L15-L40

[kernel32.go-1]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/kernel32.go#L15-L54

[kernel32.go-2]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/kernel32.go#L92-L140

### ⭐[gonutz/w32](https://github.com/gonutz/w32)

## [build](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

GO中針對不同平台的建置方法有以下幾種:

1. 使用檔名(filename suffixes)來判斷，例如:

- `xxx_GOOS_GOARCH.go`
- `xxx_windows_[arch].go`
- `xxx_windows_arm64.go`

以上GOOS和GOARCH都是可選上，可加也可不加，通常不太會去特別註明arch

2. 明確使用build標籤:
   - `//go:build windows` 只在windows平台build
   - `//go:build !windows` 只要是`非`windows平台就build
   - `//go:build darwin`
   - `//+build windows,amd64` 只在windows且arch為amd64才構建
   - `//+build darwin linux dragonfly js,wasm` 前面三個為os，最後一個`js,wasm`這要看成一個，意思為OS:js, arch為wasm

    要查看有哪些OS, ARCH可用，可以使用指令
    > go tool dist list
    > 它會呈現出: `OS/ARCH`的列表 (list all supported platforms):

    ```
    aix/ppc64
    android/386
    android/amd64
    android/arm
    android/arm64
    darwin/amd64
    darwin/arm64
    dragonfly/amd64
    freebsd/386
    ...
    illumos/amd64
    ios/amd64
    ios/arm64
    js/wasm
    linux/386
    linux/amd64
    ...
    windows/386
    windows/amd64
    windows/arm
    windows/arm64
    ```

3. 使用`GOOS`, `GOARCH`，如果1,2都沒有，就會依照這兩個變數而定
