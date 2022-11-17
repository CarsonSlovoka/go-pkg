package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
)

func ExampleGET_KEYSTATE_WPARAM() {
	wParam := uintptr(4287102992)
	keyState := w32.GET_KEYSTATE_WPARAM(wParam)
	isMouseBtnDone := (keyState & w32.MK_MBUTTON) == w32.MK_MBUTTON
	fmt.Println(isMouseBtnDone)
	// Output:
	// true
}

func ExampleGET_WHEEL_DELTA_WPARAM() {
	wParam := uintptr(4287102992)
	wheelDelta := w32.GET_WHEEL_DELTA_WPARAM(wParam)
	fmt.Println(wheelDelta)
	if wheelDelta > 0 {
		fmt.Println("rotated forward, away from the user")
	} else {
		fmt.Println("rotated backward, toward the user")
	}
	// Output:
	// -120
	// rotated backward, toward the user
}
