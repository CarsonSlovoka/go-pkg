package embed

import (
	"embed"
	"io"
	"os"
)

func ExtractFile(fs *embed.FS, srcPath, dstPath string) error {
	return PromiseExtractFile(fs, srcPath, dstPath, nil, nil)
}

// PromiseExtractFile 與ExtractFile類似，但可以在成功或者失敗的時候在做某些事情
func PromiseExtractFile(fs *embed.FS, srcPath, dstPath string,
	successFunc func(dst string) error,
	errFunc func(err error) error,
) (err error) {

	defer func() { // defer有後進先出的特性，因此這個defer會最後執行
		if err == nil && successFunc != nil {
			err = successFunc(dstPath)
		} else if errFunc != nil {
			err = errFunc(err)
		}
	}()

	srcFile, err := fs.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() {
		err = srcFile.Close()
	}()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func() {
		err = dstFile.Close()
	}()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}
