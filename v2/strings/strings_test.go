package strings

import (
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
