package fmt

import (
	"fmt"
	"os"
)

// ColorPrinter ANSI Escape
// https://stackoverflow.com/a/69924820/9935654
type ColorPrinter struct {
	fr, fg, fb int // fore color: rgb
	br, bg, bb int // background color
	prefix     string
	suffix     string
}

func NewColorPrinter(fr, fg, fb, br, bg, bb int) *ColorPrinter {
	c := &ColorPrinter{fr: fr, fg: fg, fb: fb, br: br, bg: bg, bb: bb}
	c.prefix = fmt.Sprintf("\u001B[48;2;%d;%d;%dm\u001B[38;2;%d;%d;%dm", c.br, c.bg, c.bb, c.fr, c.fg, c.fb)
	c.suffix = "\u001B[0m"
	return c
}

func (c *ColorPrinter) SetFGColor(fr, fg, fb int) {
	c.fr = fr
	c.fg = fg
	c.fb = fb
	c.updatePrefix()
}

func (c *ColorPrinter) SetBGColor(br, bg, bb int) {
	c.br = br
	c.bg = bg
	c.bb = bb
	c.updatePrefix()
}

func (c *ColorPrinter) SetColor(fr, fg, fb, br, bg, bb int) {
	c.fr = fr
	c.fg = fg
	c.fb = fb
	c.br = br
	c.bg = bg
	c.bb = bb
	c.updatePrefix()
}

func (c *ColorPrinter) updatePrefix() {
	c.prefix = fmt.Sprintf("\u001B[48;2;%d;%d;%dm\u001B[38;2;%d;%d;%dm", c.br, c.bg, c.bb, c.fr, c.fg, c.fb)
}

func (c *ColorPrinter) Sprintf(format string, a ...any) string {
	return c.prefix + fmt.Sprintf(format, a...) + c.suffix
}

func (c *ColorPrinter) addFix(a ...any) []any {
	s := make([]any, len(a)+2)
	s[0] = c.prefix
	for i, e := range a {
		s[i+1] = e
	}
	s[len(a)+1] = c.suffix
	return s
}

func (c *ColorPrinter) Sprintln(a ...any) string {
	s := c.addFix(a)
	return fmt.Sprintln(s...)
}

func (c *ColorPrinter) Println(a ...any) {
	s := c.addFix(a)
	_, _ = fmt.Fprintln(os.Stdout, s...)
}

func (c *ColorPrinter) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stdout, c.prefix+format+c.suffix, a...)
}

func (c *ColorPrinter) Errorf(format string, a ...any) error {
	return fmt.Errorf(c.prefix+format+c.suffix, a...)
}
