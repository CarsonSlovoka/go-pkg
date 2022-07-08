package http

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDownload2Memory(t *testing.T) {
	dataBytes, err := Download2Memory("https://www.google.com.tw/")
	if err != nil {
		t.Fatalf("%s", err)
	}
	_ = fmt.Sprintf("%s", dataBytes)
}

func ExampleDownload2Memory() {
	dataBytes, err := Download2Memory("https://www.google.com.tw/")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(dataBytes) > 0)
	// Output: true
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

func TestDownloadFileWithTimeout(t *testing.T) {
	// We give it very little time to prepare, so a timeout error is bound to occur.
	err := DownloadFileWithTimeout("temp.html", "https://www.google.com.tw/", time.Nanosecond)
	if err == nil ||
		!strings.Contains(err.Error(), "Timeout") { // https://github.com/golang/go/blob/1930977/src/net/http/client.go#L718-L723
		t.Fatal("must error")
	}
}

// Output support string or os.File both OK!
func ExampleDownloadFile() {
	// Output type: string
	if err := DownloadFile("temp.html", "https://www.google.com.tw/"); err != nil {
		panic(err)
	}

	// Output type: writer
	writer, err := os.Create("temp2.html")
	if err != nil {
		panic(err)
	}
	if err = DownloadFile(writer, "https://www.google.com.tw/"); err != nil {
		panic(err)
	}
}

func ExampleDownloadFileWithTimeout() {
	// Output type: string
	if err := DownloadFileWithTimeout("temp.html", "https://www.google.com.tw/", 5*time.Second); err != nil {
		panic(err)
	}

	// Output type: writer
	writer, err := os.Create("temp2.html")
	if err != nil {
		panic(err)
	}
	if err = DownloadFileWithTimeout(writer, "https://www.google.com.tw/", 5*time.Second); err != nil {
		panic(err)
	}
}
