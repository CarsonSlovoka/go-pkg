## lxn/win

常數數值大部分都是參考此專案[lxn/win]，這專案很棒，省去您花很多時間去找微軟文件所定義的該數值為何

- [user32.go]
- [gdi32.go]
- [win.go]
- [kernel32.go-1], [kernel32.go-2]

不用擔心宣告太多變數執行檔會變得很大，實際上就[user32.go]的這些內容，打包出來的執行檔差異也只有10KB左右<sub>就算以上四項全部包含進去也才影響14KB</sub>，影響很小

至於struct的定義，宣告再多都不會影響執行檔大小(看的是用了多少)

[lxn/win]: https://github.com/lxn/win
[user32.go]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/user32.go#L18-L1744
[gdi32.go]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/gdi32.go#L16-L1038
[win.go]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/win.go#L15-L40
[kernel32.go-1]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/kernel32.go#L15-L54
[kernel32.go-2]: https://github.com/lxn/win/blob/a377121e959e22055dd01ed4bb2383e5bd02c238/kernel32.go#L92-L140
