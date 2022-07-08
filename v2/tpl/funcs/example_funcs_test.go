package funcs

import (
	"bytes"
	"fmt"
	"text/template"
)

func ExampleGetUtilsFuncMap() {
	content := `
# {{.Title}}

## Dict
{{- $MyDict := dict "Name" "Carson" "ID" "001" }}
Hi {{ $MyDict.Name -}}. Data {{ (dict "A" "65" "B" "66").B }} {{ index (dict "A" "65" "C" "67") "C" }}
// Hi Carson. Data 66 67

## Slice
{{- $Numbers := list 1 3 5 -}}
{{- $Items := list "aa" "bb" "cc" }}
{{ index $Numbers 2 }} {{ index $Items 1 }} {{ index (list "a" "b" "c") 2 }}
// 5 bb c

{{ index (split "2022/05/23" "/") 0 }}
// 2022

{{ replace "Hi xxx" "xxx" "Carson" -1 }}
// Hi Carson

{{ replace "Hi XXX XXX" "X" "0" 2 }}
// Hi 00X XXX

## Math
可以善用:
replace:
\{\{.*\}\} = (.*)
$1 = $1
來把左側取代為右側內容
{{add 3 5}} = 8
{{add 3 0.5}} = 3.5
{{sub 1.5 -0.7}} = 2.2
{{mul 3 2}} = 6
{{mul 1.5 0.7}} = 1.0499999999999998
{{div 3 2}} = 1.5
{{mul 0 -0.7}} = -0
{{mul 0 0.7}} = 0
{{ceil 0}} = 0
{{ceil 0.1}} = 1
{{ceil -1.1}} = -1
{{floor 0.1}} = 0
{{floor -0.1}} = -1
{{ln 1}} = 0
{{ln 2}} = 0.6931471805599453
{{log 1 10}} = 0
{{log 2 10}} = 0.30102999566398114
{{sqrt 0}} = 0
{{sqrt 4}} = 2
{{mod 9 5}} = 4
{{mod -5 2}} = -1
{{mod 2.8 2}} = 0
{{modBool 9 5}} = false
{{modBool 0 5}} = true
{{modBool 2.8 2}} = true
{{pow 2 3}} = 8
{{pow 0 2}} = 0
{{pow 1.2 0}} = 1
{{round -1.4}} = -1
{{round 1.5000}} = 2
`

	t := template.Must(new(template.Template).
		Funcs(GetUtilsFuncMap()).
		Parse(content))

	buffer := bytes.NewBuffer(make([]byte, 0))
	_ = t.Execute(buffer, map[string]string{"Title": "Example"})

	fmt.Println(buffer)
	// Output:
	//
	//# Example
	//
	//## Dict
	//Hi Carson. Data 66 67
	//// Hi Carson. Data 66 67
	//
	//## Slice
	//5 bb c
	//// 5 bb c
	//
	//2022
	//// 2022
	//
	//Hi Carson
	//// Hi Carson
	//
	//Hi 00X XXX
	//// Hi 00X XXX
	//
	//## Math
	//可以善用:
	//replace:
	//\{\{.*\}\} = (.*)
	//$1 = $1
	//來把左側取代為右側內容
	//8 = 8
	//3.5 = 3.5
	//2.2 = 2.2
	//6 = 6
	//1.0499999999999998 = 1.0499999999999998
	//1.5 = 1.5
	//-0 = -0
	//0 = 0
	//0 = 0
	//1 = 1
	//-1 = -1
	//0 = 0
	//-1 = -1
	//0 = 0
	//0.6931471805599453 = 0.6931471805599453
	//0 = 0
	//0.30102999566398114 = 0.30102999566398114
	//0 = 0
	//2 = 2
	//4 = 4
	//-1 = -1
	//0 = 0
	//false = false
	//true = true
	//true = true
	//8 = 8
	//0 = 0
	//1 = 1
	//-1 = -1
	//2 = 2
}
