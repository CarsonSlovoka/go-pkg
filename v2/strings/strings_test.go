package strings

import (
	"fmt"
	"testing"
)

func TestIsUpper(t *testing.T) {
	for _, d := range []struct {
		actual   bool
		expected bool
	}{
		{IsUpperU("中文A1"), false}, // be careful!
		{IsUpper("中文A1"), true},

		{IsUpper("中文a1"), false},
		{IsUpperU("中文a1"), false},
	} {
		if d.actual != d.expected {
			t.Fatal()
		}
	}
}

func ExampleIsUpper() {
	fmt.Println(IsUpper("中文A1"))
	fmt.Println(IsUpper("中文a1"))
	// Output:
	// true
	// false
}

// 建議您檢驗的對象為Latin1的字集在使用此判斷，不然可能會不如您預期
func ExampleIsUpperU() {
	fmt.Println(IsUpperU("ABC123"))
	fmt.Println(IsUpperU("中文A1")) // be careful!
	fmt.Println(IsUpperU("中文a1"))
	// Output:
	// true
	// false
	// false
}

func TestIsLower(t *testing.T) {
	for idx, d := range []struct {
		actual   bool
		expected bool
	}{
		{IsLowerU("中文a1"), false}, // be careful!
		{IsLower("中文a1"), true},

		{IsLower("中文A1"), false},
		{IsLowerU("中文A1"), false},
	} {
		if d.actual != d.expected {
			t.Fatal(idx)
		}
	}
}

func ExampleIsLower() {
	fmt.Println(IsLower("中文a1"))
	fmt.Println(IsLower("中文A1"))
	// Output:
	// true
	// false
}

func ExampleIsLowerU() {
	fmt.Println(IsLowerU("a1"))
	fmt.Println(IsLowerU("中文a1")) // be careful!
	fmt.Println(IsLowerU("中文A1"))
	// Output:
	// true
	// false
	// false
}
