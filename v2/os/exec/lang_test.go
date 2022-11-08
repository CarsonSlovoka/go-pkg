package exec

import (
	"log"
)

func ExampleGetLocalLangLoc() {
	lang, loc := GetLocalLangLoc("en", "US")
	log.Println(lang, loc)

	// Output:
}
