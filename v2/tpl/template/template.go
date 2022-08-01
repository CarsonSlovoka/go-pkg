package template

import (
	"path/filepath"
	"regexp"
)

// GetAllTmplName 獲取文件中會使用到的所有樣板名稱(子嵌套也都會遞迴尋找)
func GetAllTmplName(
	ReadFile func(name string) ([]byte, error), // os.ReadFile, embed.FS.ReadFile
	filePath string,
	allTmpl []string,
) (filterTmpl []string, err error) {
	var content []byte
	content, err = ReadFile(filePath)
	reTmpl := regexp.MustCompile(`{{-? ?template \"(?P<Name>[^() ]*)\" ?.* ?-?}}`)
	matchList := reTmpl.FindAllStringSubmatch(string(content), -1)

	if len(matchList) == 0 {
		return
	}

	curTmplSet := map[string]string{} // 知道當前文件中，所有用到的tmpl的名稱
	for _, match := range matchList {
		tmplName := match[1]
		if _, exists := curTmplSet[tmplName]; exists {
			continue
		}
		curTmplSet[tmplName] = tmplName
	}

	for _, tmplFilepath := range allTmpl { // 從所有tmpl文件中，篩選名稱相符的
		_, exists := curTmplSet[filepath.Base(tmplFilepath)]
		if exists {
			filterTmpl = append(filterTmpl, tmplFilepath)
			fList, _ := GetAllTmplName(ReadFile, tmplFilepath, allTmpl) // 該模板也有可能再有模板，所以要再找
			if len(fList) > 0 {
				filterTmpl = append(filterTmpl, fList...)
			}
		}
	}
	return
}
