package slices

import (
	"fmt"
	"testing"
)

type testData struct {
	actual   any
	expected any
}

func test(t *testing.T, data []testData) {
	for caseNumber, d := range data {
		if d.actual != d.expected {
			t.Fatalf("caseNumber: %d:\nactual: %+v\nexpected: %+v", caseNumber, d.actual, d.expected)
		}
	}
}

func TestIndex(t *testing.T) {
	test(t, []testData{
		{Index([]string{"A", "B", "C"}, "B"), 1},
		{Index([]int{1, 3, 5, 7}, 5), 2},
		{Index([]int{}, 0), -1}}, // slice為空，必定回傳-1
	)
}

func ExampleIndex() {
	fmt.Println(Index([]string{"A", "B", "C"}, "B"))
	fmt.Println(Index([]int{1, 3, 5, 7}, 7))

	// slice為空，必定回傳-1
	fmt.Println(Index([]int{}, 0))
	// Output:
	// 1
	// 3
	// -1
}

func TestContains(t *testing.T) {
	test(t, []testData{
		{Contains([]string{"A", "B", "C"}, "B"), true},
		{Contains([]string{"A", "B", "C"}, "Z"), false},
		{Contains([]int{1, 3, 5, 7}, 5), true},
		{Contains([]int{1, 3, 5, 7}, 9), false},
		{Contains([]int{}, 0), false}},
	)
}

func ExampleContains() {
	fmt.Println(Contains([]string{"A", "B", "C"}, "B"))
	fmt.Println(Contains([]string{"A", "B", "C"}, "Z"))
	fmt.Println(Contains([]int{1, 3, 5, 7}, 5))
	fmt.Println(Contains([]int{1, 3, 5, 7}, 9))
	fmt.Println(Contains([]int{}, 0))
	// Output:
	// true
	// false
	// true
	// false
	// false
}

func TestAny(t *testing.T) {
	test(t, []testData{
		{Any([]string{"AA", "B", "C"}, "AA", "C"), true},
		{Any([]string{"AA", "B", "C"}, "AK", "C"), true},
		{Any([]string{"AA", "B", "C"}, []string{"Z", "D"}...), false},
		{Any([]int{1, 2, 3}, []int{2, 8, 9}...), true},
		{Any([]int{1, 2, 3}, 5, 7, 9), false},
		{Any([]int{1, 2, 3}, 3), true},
	})
}

func ExampleAny() {
	fmt.Println(Any([]string{"AA", "B", "C"}, "AA", "C"))
	fmt.Println(Any([]string{"AA", "B", "C"}, "AK", "C"))
	fmt.Println(Any([]string{"AA", "B", "C"}, []string{"Z", "D"}...))
	fmt.Println(Any([]int{1, 2, 3}, []int{2, 8, 9}...))
	fmt.Println(Any([]int{1, 2, 3}, 5, 7, 9))
	fmt.Println(Any([]int{1, 2, 3}, 3))
	// Output:
	// true
	// true
	// false
	// true
	// false
	// true
}

func TestAll(t *testing.T) {
	test(t, []testData{
		{All([]string{"AA", "B", "C"}, "C"), true},
		{All([]string{"AA", "B", "C"}, "AA", "C"), true},
		{All([]string{"AA", "B", "C"}, "AK", "C"), false},
		{All([]string{"AA", "B", "C"}, []string{"B", "C"}...), true},
		{All([]int{1, 3, 5}, 3), true},
		{All([]int{1, 3, 5}, 2, 3), false},
		{All([]int{1, 3, 5}, 1, 5), true},
		{All([]int{}, 0), false},          // slice為空必定為false
		{All([]int{}, []int{}...), false}, // 此外檢驗空資料是否於某slice也視為false
	})
}

func ExampleAll() {
	fmt.Println(All([]string{"AA", "B", "C"}, "C"))
	fmt.Println(All([]string{"AA", "B", "C"}, "AA", "C"))
	fmt.Println(All([]string{"AA", "B", "C"}, "AK", "C"))
	fmt.Println(All([]string{"AA", "B", "C"}, []string{"B", "C"}...))
	fmt.Println(All([]int{1, 3, 5}, 3))
	fmt.Println(All([]int{1, 3, 5}, 2, 3))
	fmt.Println(All([]int{1, 3, 5}, 1, 5))

	// slice為空必定為false
	fmt.Println(All([]int{}, 0))

	// 此外檢驗空資料是否於某slice也視為false
	fmt.Println(All([]int{}, []int{}...))
	// Output:
	// true
	// true
	// false
	// true
	// true
	// false
	// true
	// false
	// false
}
