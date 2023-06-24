<p align="center">
  <a href="http://golang.org">
      <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go">
  </a>
  <a href="https://pkg.go.dev/github.com/CarsonSlovoka/go-pkg/v2">
      <img src="https://img.shields.io/badge/Documentation-go_pkg-blue.svg" alt="Documentation">
  </a>

  <img src="https://img.shields.io/github/go-mod/go-version/CarsonSlovoka/go-pkg?filename=go.mod" alt="Go Version">

  <a href="https://GitHub.com/CarsonSlovoka/go-pkg/releases/">
      <img src="https://img.shields.io/github/release/CarsonSlovoka/go-pkg" alt="Latest release">
  </a>
  <a href="https://github.com/CarsonSlovoka/go-pkg/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/CarsonSlovoka/go-pkg.svg" alt="License">
  </a>
</p>

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

----

當您剛發布的時候，有可能手動請求也還是出不來，這時候可以至[Open Source Insights](https://deps.dev/)，可參考下方連結

> https://deps.dev/go/github.com%2Fcarsonslovoka%2Fgo-pkg%2Fv2/

進入之後選擇您最新的版本，如果還在準備中會看到以下訊息

> Warning This is a recently released version, so we are still processing the data. Please check back later, or go to the latest version we have information for vx.x.x

所以要過一段時間之後才可以被抓到。

您也可以查看proxy的狀態，例如

> https://proxy.golang.org/github.com/carsonslovoka/go-pkg/v2/@v/v2.3.0.info
