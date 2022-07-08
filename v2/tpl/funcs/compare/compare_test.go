package compare_test

import (
	"bytes"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs/compare"
	htmlTemplate "html/template"
	"testing"
	"text/template"
)

func TestCompare(t *testing.T) {
	type D struct {
		content  string
		context  any
		expected string
	}
	for _, d := range []struct {
		funcMap map[string]any
		cases   []D
	}{
		{map[string]any{"default": compare.Default},
			[]D{
				{`{{default "myDefault" .Data}}`,
					map[string]int{},
					"myDefault"},
				{`{{default "myDefault" .Data}}`,
					map[string]string{"Data": "Hello World"},
					"Hello World"},
				{`{{default "myDefault" .Data}}`,
					map[string]int{"Data": 123},
					"123"},

				{`{{index . "Data" | default "myDefault" }}`,
					map[string]string{"Data": "Hi"},
					"Hi"},

				{`{{index . "No" | default "myDefault" }}`,
					map[string]string{"Data": "Hi"},
					"myDefault"},

				{`{{.Data | default "myDefault" }}`,
					map[string]string{"Data": "Hi"},
					"Hi"},
				{`{{.No | default "myDefault" }}`,
					map[string]string{"Data": "Hi"},
					"myDefault"},

				{`{{ index .Data 1 | default "myDefault" }}`,
					map[string]any{"Data": []string{"a", "b"}},
					"b"},
				{`{{ .Data | default "myDefault" }}`,
					map[string]any{"Data": []string{}},
					"myDefault"},
				{`{{ default "myDefault" .Data }}`,
					map[string]any{"Data": []string{}},
					"myDefault"},
			},
		},
	} {
		for _, curCase := range d.cases {
			tmpl := htmlTemplate.Must(
				htmlTemplate.New("test").
					Funcs(d.funcMap).
					Parse(curCase.content),
			)
			buffer := bytes.NewBuffer(make([]byte, 0))
			if err := tmpl.Execute(buffer, curCase.context); err != nil {
				t.Fatal(err)
			}
			actual := buffer.String()
			if expected := curCase.expected; actual != expected {
				t.Fatalf("%s\n%s\n%s", curCase.content, actual, expected)
			}
		}
	}
}

func ExampleCompare() {
	t, err := new(template.Template).
		Funcs(map[string]any{"default": compare.Default}).
		Parse(`{{ .Data | default "myDefault" }}`)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(make([]byte, 0))
	_ = t.Execute(buffer, map[string]string{"Data": "Hello World"})
	fmt.Println(buffer)

	buffer.Truncate(0) // clear buffer
	_ = t.Execute(buffer, nil)
	fmt.Println(buffer)
	// Output:
	// Hello World
	// myDefault
}
