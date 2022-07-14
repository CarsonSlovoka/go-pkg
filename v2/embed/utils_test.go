package embed

import (
	embed0 "embed"
	"fmt"
	"io"
	"os"
)

//go:embed testData
var embedBuildBat embed0.FS

// Extract the file from the embed path (src) to the output path (dst).
func ExampleExtractFile() {
	outputFilePath := "./output.txt"
	if err := ExtractFile(&embedBuildBat, "testData/hello.txt", outputFilePath); err != nil {
		panic(err)
	}
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		panic(err)
	}
	defer func() {
		if err := os.Remove(outputFilePath); err != nil {
			panic(err)
		}
	}()
	f, _ := os.Open(outputFilePath)
	defer f.Close()
	dataBytes, _ := io.ReadAll(f)
	fmt.Println(string(dataBytes))

	// Output: Hello World!
}

func ExamplePromiseExtractFile() {
	outputFilePath := "./output.txt"
	if err := PromiseExtractFile(&embedBuildBat, "testData/hello.txt", outputFilePath,
		func(dst string) error {
			f, _ := os.Open(dst)
			dataBytes, _ := io.ReadAll(f)
			_ = f.Close()
			fmt.Println(string(dataBytes))
			return os.Remove(outputFilePath)
		}, nil); err != nil {
		panic(err)
	}

	if err := PromiseExtractFile(&embedBuildBat, "not_exist_file.txt", outputFilePath,
		nil, func(err error) error {
			fmt.Println("got an error")
			return err
		}); err == nil {
		panic(err)
	}

	// Output:
	// Hello World!
	// got an error
}
