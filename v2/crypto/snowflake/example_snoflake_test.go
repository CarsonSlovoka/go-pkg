package snowflake_test

import (
	"fmt"
	. "github.com/CarsonSlovoka/go-pkg/v2/crypto/snowflake"
)

var gN *Node // global Node // 方便其他測試案例使用

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
}

func ExampleID_Base2() {
	id := gN.Generate()
	fmt.Println(id.Base2())
}

func ExampleID_String() {
	id := gN.Generate()
	fmt.Println(id.String())
}
