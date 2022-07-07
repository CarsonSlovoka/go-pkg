# Go-PKG

go標準庫的補強 (不使用第三方套件)

目錄結構同[go/src標準庫](https://github.com/golang/go/tree/master/src)

## [pkg.go.dev套件更新](https://pkg.go.dev/about#best-practices-h2)

要向他們發起請求，網站上面的包才會進行更新，請求範例
> https://pkg.go.dev/example.com/my/module@version

其中

- example:是您的託管的網站
- my: username
- module: your packageName(即go.mod所放的名稱)
- @version: 可選項，但建議一定要有，版號格式為`vX.X.X`

以本包為例:

> https://pkg.go.dev/github.com/CarsonSlovoka/go-pkg/v2@v2.1.0

如果不存在，那麼他會出現一個按鈕，讓您點擊，之後就會去生成了
