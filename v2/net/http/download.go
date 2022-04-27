package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
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
	var outFile *os.File
	if reflect.TypeOf(out).String() == "string" {
		// Create the file
		outFile, err = os.Create(fmt.Sprintf("%v", out))
		if err != nil {
			return err
		}
		defer outFile.Close()
	} else {
		outFile = any(out).(*os.File)
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
