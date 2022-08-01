package filepath_test

import (
	"embed"
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/path/filepath"
)

// 蒐集指定資料夾內的所有檔案, 本範例的排除項目為{sub和sub2資料夾內所有的md}
func ExampleCollectFiles() {
	fileList, err := filepath.CollectFiles("./testdata", []string{`testdata\\sub\\.*\.md`, `testdata\\sub2\\.*\.md`})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(fileList))
	fmt.Printf("%v", fileList)
	// Output:
	// 5
}

//go:embed testdata/*
var testFS embed.FS

// 蒐集embed中特定資料夾路徑其包含的所有檔案路徑
func ExampleCollectFilesFromFS() {
	fileList, err := filepath.CollectFilesFromFS(testFS, []string{"testdata"}, []string{".md"}, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(fileList))

	fileList, err = filepath.CollectFilesFromFS(testFS, []string{"testdata"}, []string{".md"}, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(fileList))

	fileList, _ = filepath.CollectFilesFromFS(testFS, []string{"testdata/sub"}, nil, true)
	fmt.Println(len(fileList))

	fileList, _ = filepath.CollectFilesFromFS(testFS, []string{"testdata/sub", "testdata/sub2"}, nil, false)
	fmt.Println(len(fileList))
	// Output:
	// 4
	// 2
	// 4
	// 6
}
