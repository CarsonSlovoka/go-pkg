package rand

import (
	"fmt"
	"log"
)

func ExampleNewUUID() {
	uuid := NewUUID()
	log.Println(uuid)
	fmt.Println(len(uuid) == len("BE76F2EC-F918-7FE8-41D2-83CF5A321988"))
	// Output: true
}
