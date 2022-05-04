package embed

import (
	"embed"
	"io"
	"os"
)

func ExtractFile(fs *embed.FS, srcPath, dstPath string) error {
	srcFile, err := fs.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}
