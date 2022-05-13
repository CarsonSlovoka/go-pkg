package exec

import (
	"fmt"
	"testing"
)

func TestGetLocalLangLoc(t *testing.T) {
	lang, loc := GetLocalLangLoc("en", "US")
	fmt.Println(lang, loc)
}
