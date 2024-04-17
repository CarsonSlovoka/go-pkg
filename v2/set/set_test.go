package set_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/set"
	"testing"
)

func ExampleNewMapSet() {
	s := set.NewMapSet[string]()
	s.Add("Apple")
	s.Add("Foo")
	s.Add("Foo")
	s.Add("Bar")
	fmt.Println(s.Size())   // 3
	fmt.Println(s.String()) // Apple Bar Foo
	fmt.Println(s)          // same as above
	s.Delete("Foo")
	fmt.Println(s.Has("Foo"))
	fmt.Println(s.Has("Bar"))
	fmt.Println(s) // Apple Bar

	// Output:
	// 3
	// Apple Bar Foo
	// Apple Bar Foo
	// false
	// true
	// Apple Bar
}

func ExampleMapSet_Clear() {
	s := set.NewMapSet([]int{1, 8, 3, 5, 8, 8}...)
	s.Clear()
	fmt.Println(s.Size())
	s.Add(200, 100)
	fmt.Println(s)
	// Output:
	// 0
	// 100 200
}

func ExampleNewMapSet_slices() {
	s := set.NewMapSet([]uint8{1, 8, 3, 5, 8, 8}...)
	fmt.Println(s)
	// Output:
	// 1 3 5 8
}

func ExampleNewMapSet_string() {
	s := set.NewMapSet("apple", "foo", "bar")
	s.Delete("foo")
	fmt.Println(s)
	// Output:
	// apple bar
}

func TestMapSet(t *testing.T) {
	s1 := set.NewMapSet(1, 2, 3)
	t.Run("clone", func(tt *testing.T) {
		s2 := s1.Clone()
		s2.Delete(2)
		s2.Add(5)
		if s2.String() != "1 3 5" {
			tt.Fatal()
		}
		if s1.String() != "1 2 3" {
			tt.Fatal()
		}
	})

	t.Run("union", func(tt *testing.T) {
		s2 := set.NewMapSet(2, 5, 8)
		s3 := s1.Union(s2)
		if s3.String() != "1 2 3 5 8" {
			tt.Fatal()
		}
	})

	t.Run("diff", func(tt *testing.T) {
		s2 := set.NewMapSet(2, 5, 8)
		s3 := s1.Diff(s2)
		if s3.String() != "1 3 5 8" {
			tt.Fatal()
		}
	})

	t.Run("contains", func(tt *testing.T) {
		if !s1.Contains(1, 3) { // true
			tt.Fatal()
		}

		if s1.Contains(1, 8) { // false 8不符合
			tt.Fatal()
		}
	})
}
