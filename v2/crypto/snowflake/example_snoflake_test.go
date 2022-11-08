package snowflake_test

import (
	. "github.com/CarsonSlovoka/go-pkg/v2/crypto/snowflake"
	"log"
)

var gN *Node // global Node // 方便其他測試案例使用 (缺點是在pkg.go.dev的網站上，在片段範例之中仍然不會給出任何提示

func init() {
	var err error
	gN, err = NewNode(0, BaseT, 10, 12)
	if err != nil {
		panic(err)
	}
}

func ExampleNode_Generate() {
	// BaseT := time.Date(2022, 7, 1, 16, 10, 54, 0, time.UTC)
	n, err := NewNode(0, BaseT, 10, 12)
	if err != nil {
		panic(err)
	}
	n.Generate()
	// ...
	n.Generate()

	log.Println(n.MaskStep())
	// Output:
}

func ExampleID_Base2() {
	id := gN.Generate()
	log.Println(id.Base2())
	// Output:
}

func ExampleID_String() {
	id := gN.Generate()
	log.Println(id.String())
	log.Println(id.Step(1024))
	// Output:
}
