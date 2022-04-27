package http

import (
	"fmt"
	"os"
	"testing"
)

func TestDownload2Memory(t *testing.T) {
	dataBytes, err := Download2Memory("https://www.google.com.tw/")
	if err != nil {
		t.Fatalf("%s", err)
	}
	_ = fmt.Sprintf("%s", dataBytes)
}

func TestDownloadFile(t *testing.T) {
	// para string
	if err := DownloadFile("temp.html", "https://www.google.com.tw/"); err != nil {
		t.Fatalf("%s", err)
	}
	if err := os.Remove("temp.html"); err != nil {
		t.Fatalf("%s", err)
	}

	// para: File
	writer, err := os.Create("temp2.html")
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer func() {
		_ = writer.Close()
		if err := os.Remove("temp2.html"); err != nil {
			t.Fatalf("%s", err)
		}
	}()
	if err := DownloadFile(writer, "https://www.google.com.tw/"); err != nil {
		t.Fatalf("%s", err)
	}
}
