package template_test

import (
	"embed"
	"fmt"
	filepath2 "github.com/CarsonSlovoka/go-pkg/v2/path/filepath"
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/template"
	htmlTemplate "html/template"
	"os"
	"path/filepath"
	"testing"
)

//go:embed testdata/tmpl/*
//go:embed testdata/pages/*
var dataFS embed.FS

func TestGetAllTmplName(t *testing.T) {
	tplList, err := filepath2.CollectFilesFromFS(dataFS, []string{"testdata/tmpl"}, nil, true)
	if err != nil {
		t.Fatal(err)
	}
	filterTmpl, err := template.GetAllTmplName(os.ReadFile, "testdata/pages/demo.gohtml", tplList)
	if err != nil {
		t.Fatal(err)
	}
	if len(filterTmpl) != 4 {
		t.Fatal()
	}
}

func ExampleGetAllTmplName() {
	tplList, err := filepath2.CollectFilesFromFS(dataFS, []string{"testdata/tmpl"}, nil, true)
	if err != nil {
		panic(err)
	}

	var filterTmpl []string
	src := "testdata/pages/demo.gohtml" // 當前所要渲染的頁面
	filterTmpl, err = template.GetAllTmplName(dataFS.ReadFile, src, tplList)
	if err != nil {
		panic(err)
	}
	filterTmpl = append(filterTmpl, src) // 當前的文件也要包含
	fmt.Println(len(filterTmpl))

	// render
	_ = func() {
		tmpl := htmlTemplate.New(filepath.Base(src)).Funcs(nil)
		tmpl, err = tmpl.ParseFS(dataFS, filterTmpl...)
		// tmpl, err = tmpl.ParseFiles(filterTmpl...)
		if err = tmpl.Execute(os.Stdout, nil); err != nil {
			panic(err)
		}
	}
	// Output:
	// 5
}
