package w32_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
)

// 不提供外部直接對其union做設定
// 應用在w32裡面所使用的結構
// 此範例只是為了增加測試覆蓋率所寫
func ExampleOVERLAPPED_OffsetHigh() {
	o := new(w32.OVERLAPPED)
	log.Println(o.Offset()) // 記憶體位址，例如: 0xc0000a6450
	log.Println(o.OffsetHigh())
	log.Println(o.Pointer())
	// Output:
}
