package fmt_test

import (
	"fmt"
	. "github.com/CarsonSlovoka/go-pkg/v2/fmt"
	"log"
	"testing"
)

func TestColorPrinter(t *testing.T) {
	pYellow := NewColorPrinter(0, 0, 0, 255, 255, 0)
	pYellow.Println(12, "ab", "cd")
	pYellow.Println(123)
	pYellow.Printf("%d %s %q\n", 123, "Hi", "quote")
	pYellow.Printf("%d\n", 123)
	pErr := NewColorPrinter(255, 255, 255, 255, 0, 0)
	fmt.Printf("%s\n", pErr.Errorf("ErrorF test"))
}

func ExampleColorPrinter_style() {
	pYellow := NewColorPrinter(0, 0, 0, 255, 255, 0)
	pHighlight := NewColorPrinter(255, 0, 255, 255, 255, 0)
	pMark := NewColorPrinter(255, 0, 255, 255, 255, 0)
	pOK := NewColorPrinter(0, 0, 0, 0, 255, 0)
	pErr := NewColorPrinter(255, 255, 255, 255, 0, 0)
	pInfo := NewColorPrinter(255, 255, 255, 0, 0, 255)

	for _, p := range []*ColorPrinter{
		pYellow,
		pHighlight,
		pOK,
		pErr,
		pMark,
		pInfo,
	} {
		log.Println(p.Sprintf("%d %s %q", 123, "Hi", "quote"))
	}

	// Output:
}

func ExampleColorPrinter() {
	p := NewColorPrinter(0, 0, 0, 255, 255, 0)
	log.Println(p.Sprintf("%d %s %q", 123, "Hi", "quote"))

	p.SetBGColor(0, 255, 0)
	log.Println(p.Sprintf("%d %s %q", 123, "Hi", "quote"))

	p.SetFGColor(0, 0, 255)
	log.Println(p.Sprintf("%d %s %q", 123, "Hi", "quote"))
	// Output:
}
