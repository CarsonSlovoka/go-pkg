package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
)

func ExampleGET_X_LPARAM() {
	lParam := uintptr(38800252)
	xPos := w32.GET_X_LPARAM(lParam)
	yPos := w32.GET_Y_LPARAM(lParam)
	fmt.Println(xPos, yPos)
	// Output:
	// 2940 592
}
