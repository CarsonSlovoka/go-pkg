package funcs_test

import (
	"bytes"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs"
	htmlTemplate "html/template"
	"testing"
)

func TestMath(t *testing.T) {
	type D struct {
		content  string
		expected string
	}
	for _, d := range []struct {
		funcMap map[string]any
		cases   []D
	}{
		{
			map[string]any{"add": funcs.Add},
			[]D{
				{`{{add 3 2}}`, "5"},
				{`{{add 3 0.5}}`, "3.5"},
				{`{{add 0.5 0.5}}`, "1"},
				{`{{add 1.5 0.7}}`, "2.2"},
				{`{{add 1.5 -0.7}}`, "0.8"},
				{`{{add -1.5 -0.7}}`, "-2.2"},
				{`{{add 0 -0.7}}`, "-0.7"},
				{`{{add 0 0.7}}`, "0.7"},
			},
		},
		{
			map[string]any{"sub": funcs.Sub},
			[]D{
				{`{{sub 3 2}}`, "1"},
				{`{{sub 3 0.5}}`, "2.5"},
				{`{{sub 0.5 0.5}}`, "0"},
				{`{{sub 1.5 0.7}}`, "0.8"},
				{`{{sub 1.5 -0.7}}`, "2.2"},
				{`{{sub -1.5 -0.7}}`, "-0.8"},
				{`{{sub 0 -0.7}}`, "0.7"},
				{`{{sub 0 0.7}}`, "-0.7"},
			},
		},
		{
			map[string]any{"mul": funcs.Mul},
			[]D{
				{`{{mul 3 2}}`, "6"},
				{`{{mul 3 0.5}}`, "1.5"},
				{`{{mul 0.5 0.5}}`, "0.25"},
				{`{{mul 1.5 0.7}}`, "1.0499999999999998"},
				{`{{mul 1.5 -0.7}}`, "-1.0499999999999998"},
				{`{{mul -1.5 -0.7}}`, "1.0499999999999998"},
				{`{{mul 0 -0.7}}`, "-0"},
				{`{{mul 0 0.7}}`, "0"},
			},
		},
		{
			map[string]any{"div": funcs.Div},
			[]D{
				{`{{div 3 2}}`, "1.5"},
				{`{{div 3 0.5}}`, "6"},
				{`{{div 0.5 0.5}}`, "1"},
				{`{{div 1.5 0.7}}`, "2.142857142857143"},
				{`{{div 1.4 -0.7}}`, "-2"},
				{`{{div -1.4 -0.7}}`, "2"},
				{`{{div 0 -0.7}}`, "-0"},
				{`{{div 0 0.7}}`, "0"},
			},
		},
		{
			map[string]any{"ceil": funcs.Ceil},
			[]D{
				{`{{ceil 0}}`, "0"},
				{`{{ceil 0.1}}`, "1"},
				{`{{ceil 1.1}}`, "2"},
				{`{{ceil -0.1}}`, "-0"},
				{`{{ceil -1.1}}`, "-1"},
				{`{{ceil 5}}`, "5"},
				{`{{ceil -5}}`, "-5"},
			},
		},
		{
			map[string]any{"floor": funcs.Floor},
			[]D{
				{`{{floor 0}}`, "0"},
				{`{{floor 0.1}}`, "0"},
				{`{{floor 1.1}}`, "1"},
				{`{{floor -0.1}}`, "-1"},
				{`{{floor -1.1}}`, "-2"},
				{`{{floor 5}}`, "5"},
				{`{{floor -5}}`, "-5"},
			},
		},
		{
			map[string]any{"ln": funcs.Ln},
			[]D{
				{`{{ln 1}}`, "0"},
				{`{{ln 2}}`, "0.6931471805599453"},
				{`{{ln 3}}`, "1.0986122886681096"},
				{`{{ln 5}}`, "1.6094379124341003"},
			},
		},
		{
			map[string]any{"log": funcs.Log},
			[]D{
				{`{{log 1 10}}`, "0"},
				{`{{log 2 10}}`, "0.30102999566398114"},
				{`{{log 3 10}}`, "0.4771212547196623"},
				{`{{log 5 10}}`, "0.6989700043360187"},

				{`{{log 8 2}}`, "3"},
			},
		},
		{
			map[string]any{"sqrt": funcs.Sqrt},
			[]D{
				{`{{sqrt 0}}`, "0"},
				{`{{sqrt 4}}`, "2"},
				{`{{sqrt 9}}`, "3"},
				{`{{sqrt 16}}`, "4"},
				{`{{sqrt 121}}`, "11"},
			},
		},
		{
			map[string]any{"mod": funcs.Mod},
			[]D{
				{`{{mod 9 5}}`, "4"},
				{`{{mod 0 5}}`, "0"},
				{`{{mod 0 -5}}`, "0"},
				{`{{mod -5 2}}`, "-1"},
				{`{{mod -5 -3}}`, "-2"},
				{`{{mod 2.8 2}}`, "0"},
				// {`{{mod 2.8 0.9}}`, "0.1"}, error calling mod: the number can't be divided by zero at modulo operation
				// {`{{mod 2.8 1.9}}`, "0"},
			},
		},
		{
			map[string]any{"modBool": funcs.ModBool},
			[]D{
				{`{{modBool 9 5}}`, "false"},
				{`{{modBool 0 5}}`, "true"},
				{`{{modBool 0 -5}}`, "true"},
				{`{{modBool -5 2}}`, "false"},
				{`{{modBool -5 -3}}`, "false"},
				{`{{modBool 2.8 2}}`, "true"},
			},
		},
		{
			map[string]any{"pow": funcs.Pow},
			[]D{
				{`{{pow 2 3}}`, "8"},
				{`{{pow -2 3}}`, "-8"},
				{`{{pow 1.1 2}}`, "1.2100000000000002"},
				{`{{pow -1.1 2}}`, "1.2100000000000002"},
				{`{{pow 0 2}}`, "0"},
				{`{{pow 1 0}}`, "1"},
				{`{{pow 1.2 0}}`, "1"},
			},
		},
		{
			map[string]any{"round": funcs.Round},
			[]D{
				{`{{round -1.4}}`, "-1"},
				{`{{round -1.5}}`, "-2"},
				{`{{round -0.4}}`, "-0"},
				{`{{round -0.5}}`, "-1"},
				{`{{round 0}}`, "0"},
				{`{{round 0.4}}`, "0"},
				{`{{round 0.5}}`, "1"},
				{`{{round 1.4}}`, "1"},
				{`{{round 1.5}}`, "2"},
				{`{{round 1.44444}}`, "1"},
				{`{{round 1.500000}}`, "2"},
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
			if err := tmpl.Execute(buffer, nil); err != nil {
				t.Fatal(err)
			}
			actual := buffer.String()
			if expected := curCase.expected; actual != expected {
				t.Fatalf("%s\n%s\n%s", curCase.content, actual, expected)
			}
		}
	}
}
