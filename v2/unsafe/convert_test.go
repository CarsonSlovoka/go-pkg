// go test -v -bench="^Benchmark" -run=none -benchmem

package unsafe_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/unsafe"
	"testing"
)

// const testStr = "hello world" // 不可以用常數，這樣編譯器會優化它，導致StrToBytes與一般的轉換[]byte("")結果差不多
var testStr = "hello world"
var testBytes = []byte("hello world")

func BenchmarkStrToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unsafe.StrToBytes(testStr)
	}
}

// 常規方法
func Benchmark_SafeStrToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testStr)
	}
}

func BenchmarkBytesToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unsafe.BytesToStr(testBytes)
	}
}

// 常規方法
func Benchmark_SafeBytesToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testBytes)
	}
}

func ExampleStrToBytes() {
	var bs []byte
	bs = unsafe.StrToBytes("hello")
	fmt.Printf("%v", bs)
	// Output:
	// [104 101 108 108 111]
}

func ExampleBytesToStr() {
	var s string
	s = unsafe.BytesToStr([]byte("hello"))
	fmt.Println(s)
	// Output:
	// hello
}
