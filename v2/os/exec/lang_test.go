package exec

import (
	"fmt"
)

func ExampleGetLocalLangLoc() {
	lang, loc := GetLocalLangLoc("en", "US")
	fmt.Println(lang, loc)
}
