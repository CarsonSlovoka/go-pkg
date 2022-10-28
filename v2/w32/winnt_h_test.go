package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
)

// https://renenyffenegger.ch/notes/Windows/development/Internationalization/language
func ExampleMakeLangID() {
	for _, v := range []uint16{
		w32.MakeLangID(w32.LANG_ENGLISH, w32.SUBLANG_ENGLISH_US),                    // 1<<10 + 9
		w32.MakeLangID(w32.LANG_CHINESE, w32.SUBLANG_CHINESE_TRADITIONAL),           // 4100
		w32.MakeLangID(w32.LANG_CHINESE_SIMPLIFIED, w32.SUBLANG_CHINESE_SIMPLIFIED), // 2052
		//
		w32.MakeLangID(w32.LANG_GERMAN, w32.SUBLANG_GERMAN),
		w32.MakeLangID(w32.LANG_GERMAN, w32.SUBLANG_GERMAN_SWISS),
		w32.MakeLangID(w32.LANG_GERMAN, w32.SUBLANG_GERMAN_AUSTRIAN),
		w32.MakeLangID(w32.LANG_GERMAN, w32.SUBLANG_GERMAN_LUXEMBOURG),
		w32.MakeLangID(w32.LANG_GERMAN, w32.SUBLANG_GERMAN_LIECHTENSTEIN),
	} {
		fmt.Println(v)
	}

	// Output:
	// 1033
	// 4100
	// 2052
	// 1031
	// 2055
	// 3079
	// 4103
	// 5127
}
