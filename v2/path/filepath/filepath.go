package filepath

import (
	"embed"
	"github.com/CarsonSlovoka/go-pkg/v2/slices"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func CollectFiles(dir string, excludeList []string) (fileList []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if regexp.MustCompile(strings.Join(excludeList, "|")).Match([]byte(path)) {
			return nil
		}

		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileList, nil
}

// CollectFilesFromFS 蒐集FS中特定資料夾路徑底下內各個檔案路徑位置
func CollectFilesFromFS(fs embed.FS, dirName string, excludeExtName []string, isRecursive bool) (filepathList []string, err error) {
	dirEntryList, err := fs.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	for _, dirEntry := range dirEntryList {
		if dirEntry.IsDir() {
			if isRecursive {
				fpList, _ := CollectFilesFromFS(fs, path.Join(dirName, dirEntry.Name()), excludeExtName, isRecursive)
				filepathList = append(filepathList, fpList...)
			}
			continue
		}
		if slices.Any([]string{path.Ext(dirEntry.Name())}, excludeExtName...) {
			continue
		}
		filepathList = append(filepathList, path.Join(dirName, dirEntry.Name()))
	}
	return
}
