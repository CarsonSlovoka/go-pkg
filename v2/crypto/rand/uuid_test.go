package rand

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	_, err := NewUUID()
	if err != nil {
		t.FailNow()
	}
}

func ExampleNewUUID() {
	uuid, err := NewUUID()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(uuid) == len("BE76F2EC-F918-7FE8-41D2-83CF5A321988"))
	// Output: true
}
