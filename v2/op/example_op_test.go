package op_test

import (
	"fmt"
	. "github.com/CarsonSlovoka/go-pkg/v2/op"
)

// 範例說明: Ternary
func ExampleTernary() {
	i := 30
	fmt.Println(Ternary(i%2 == 0, "even", "odd").(string))
	fmt.Println(Ternary(i > 30, 300, -1).(int))
	// Output:
	// even
	// -1
}
