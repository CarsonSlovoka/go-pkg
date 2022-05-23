package funcs_test

import (
	"bytes"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs"
	htmlTemplate "html/template"
	"testing"
)

// TestDict
// TestList
// TestSplit
// TestReplace
// https://pkg.go.dev/text/template#hdr-Functions
func TestFunc(t *testing.T) {
	for _, d := range []struct {
		funcMap  map[string]any
		content  string
		expected string
	}{
		{
			map[string]any{"dict": funcs.Dict},
			`
{{- $MyDict := dict "Name" "Carson" "ID" "001" -}}
Hi {{ $MyDict.Name -}}. Data {{ (dict "A" "65" "B" "66").B }} {{ index (dict "A" "65" "C" "67") "C" -}}
`, "Hi Carson. Data 66 67",
		},
		{map[string]any{"list": funcs.List},
			`
{{- $Numbers := list 1 3 5 -}}
{{- $Items := list "aa" "bb" "cc" -}}
{{ index $Numbers 2 }} {{ index $Items 1 }} {{ index (list "a" "b" "c") 2 -}}
`,
			"5 bb c",
		},
		{map[string]any{"split": funcs.Split},
			`{{ index (split "2022/05/23" "/") 0 }}`,
			`2022`,
		},
		{map[string]any{"replace": funcs.Replace},
			`{{ replace "Hi xxx" "xxx" "Carson" -1 }}`,
			`Hi Carson`,
		},

		{map[string]any{"replace": funcs.Replace},
			`{{ replace "Hi XXX XXX" "X" "0" 2 }}`,
			`Hi 00X XXX`,
		},
	} {
		tmpl := htmlTemplate.Must(
			htmlTemplate.New("Note").
				Funcs(d.funcMap). // Funcs(funcs.GetUtilsFuncMap())
				Parse(d.content))
		buffer := bytes.NewBuffer(make([]byte, 0))
		if err := tmpl.Execute(buffer, nil); err != nil {
			t.Fatal(err)
		}
		actual := buffer.String()
		if actual != d.expected {
			t.Fatalf("%s\n%s", actual, d.expected)
		}
	}
}
