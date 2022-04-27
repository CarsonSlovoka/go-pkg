package rand

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.FailNow()
	}
	// BE76F2EC-F918-7FE8-41D2-83CF5A321988
	fmt.Println(uuid)
}
