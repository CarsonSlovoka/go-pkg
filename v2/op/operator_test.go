package op

import "testing"

func TestTernary(t *testing.T) {
	i := 30
	if Ternary(i%2 == 0, "even", "odd").(string) != "even" {
		t.FailNow()
	}

	if Ternary(i > 30, 300, -1).(int) != -1 {
		t.FailNow()
	}
}
