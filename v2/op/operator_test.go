package op

import (
	"fmt"
	"testing"
)

func TestTernary(t *testing.T) {
	i := 30
	if Ternary(i%2 == 0, "even", "odd").(string) != "even" {
		t.FailNow()
	}

	if Ternary(i > 30, 300, -1).(int) != -1 {
		t.FailNow()
	}
}

func ExampleIf() {
	i := 30
	fmt.Println(If(i >= 60, "OK", "Fail"))
	i = 60
	fmt.Println(If(i >= 60, 600, 0))

	// 如果If沒有使用generic，這種用法會報錯: invalid operation: cannot call non-function msgFunc (variable of type any)
	msgFunc := If(i >= 60, func() {
		fmt.Println("123")
	}, func() {
		fmt.Println("Fail")
	})
	msgFunc()

	// Output:
	// Fail
	// 600
	// 123
}
