package embed

import (
	embed0 "embed"
	"io"
	"os"
	"testing"
)

//go:embed testData
var embedBuildBat embed0.FS

func TestExtractFile(t *testing.T) {
	outputFilePath := "./output.txt"
	if err := ExtractFile(&embedBuildBat, "testData/hello.txt", outputFilePath); err != nil {
		t.Fatalf(err.Error())
	}
	if _, exists := os.Stat(outputFilePath); os.IsNotExist(exists) {
		t.FailNow()
	}
	defer func() {
		if err := os.Remove(outputFilePath); err != nil {
			t.Fatalf(err.Error())
		}
	}()
	f, _ := os.Open(outputFilePath)
	defer f.Close()
	dataBytes, _ := io.ReadAll(f)
	if string(dataBytes) != "Hello World!" {
		t.FailNow()
	}
}
