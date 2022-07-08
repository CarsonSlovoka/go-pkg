package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

// Download2Memory https://stackoverflow.com/a/21351456/9935654
func Download2Memory(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("ReadFile: Size of download: %d\n", len(d))
	return d, err
}

// DownloadFile https://stackoverflow.com/a/33853856/9935654
func DownloadFile[T *os.File | string](out T, url string) (err error) {
	return DownloadFileWithTimeout(out, url, 0)
}

func DownloadFileWithTimeout[T *os.File | string](out T, url string, timeout time.Duration) (err error) {
	// Get the data

	// resp, err := http.Get(url) // 使用這種方法沒辦法設定timeout
	var client *http.Client
	if timeout != 0 {
		client = &http.Client{Timeout: timeout}
	} else {
		client = &http.Client{}
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Prepare the output file.
	var outFile *os.File
	if reflect.TypeOf(out).String() == "string" {
		outFile, err = os.Create(fmt.Sprintf("%v", out))
		if err != nil {
			return err
		}
		defer outFile.Close()
	} else {
		outFile = any(out).(*os.File)
	}

	// Writer the body to file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
