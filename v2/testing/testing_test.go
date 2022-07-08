package testing

import (
	"errors"
	"fmt"
	"testing"
)

func TestMustPanic(t *testing.T) {
	MustPanic(func() {
		panic("invalid memory address")
	}) // ok

	reason, isPanic := TestPanic(func() {
		MustPanic(func() { // 此函數不會引發任何錯誤，所以MustPanic就會有錯誤，我們才在外層再補上TestPanic捕獲此錯誤
			fmt.Println("")
		})
	})
	if !isPanic || reason != "should panic" {
		t.Fatalf("should panic")
	}
}

// If you know that you will get a panic, then you can call this function that will do recover.
func ExampleMustPanic() {
	MustPanic(func() {
		panic("raise a panic")
	})
	fmt.Println("I can run in here")
	// Output:
	// I can run in here
}

func TestPanicFunc(t *testing.T) {
	if reason, isPanic := TestPanic(func() {
		panic("invalid memory address")
	}); !isPanic || reason != "invalid memory address" { // reason.(error).Error() != "..."
		t.Fatalf(`did not panic or panic msg != invalid memory address`)
	}

	if _, isPanic := TestPanic(func() {
		_ = fmt.Sprintln("hello world")
	}); isPanic {
		t.Fatalf("It shouldn't cause panic.")
	}

	var ps *string
	reason, isPanic := TestPanic(func() {
		fmt.Print(*ps)
	})

	errMsg1 := "invalid memory address or nil pointer dereference"
	errMsg2 := "runtime error: invalid memory address or nil pointer dereference"
	for idx, d := range []struct {
		actual   interface{}
		expected interface{}
	}{
		{reason == errMsg1, false}, // 這個reason的錯誤其實是runtime.errorString, 它的Error= "runtime error: " + string(e)
		{reason == errMsg2, false}, // 還是錯誤，因為這等於是用一個interface{}和字串做比對，所以會認為兩者不同，應該要在同樣都是字串的基準下做比對
		{fmt.Sprintf("%v", reason) == errMsg2, true},
		{reason.(error).Error() == errMsg2, true},
		{isPanic, true},
	} {
		if d.actual != d.expected {
			t.Fatalf("%d", idx)
		}
	}
}

func ExampleTestPanic() {
	// example1
	reason, isPanic := TestPanic(func() {
		panic("invalid memory address")
	})
	fmt.Printf("%v | %v\n", reason, isPanic)

	// example2
	reason, isPanic = TestPanic(func() {
		panic(errors.New("reason type: error"))
	})
	fmt.Printf("%v | %v\n", reason.(error).Error(), isPanic)

	// example3
	reason, isPanic = TestPanic(func() {
		_ = fmt.Sprintln("hello world")
	})
	fmt.Printf("%v | %v\n", reason, isPanic)
	// Output:
	// invalid memory address | true
	// reason type: error | true
	// <nil> | false
}
