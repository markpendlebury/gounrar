package unrar

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func mkdir(path string, dirMode os.FileMode) error {
	err := os.MkdirAll(path, dirMode)
	if err != nil {
		return fmt.Errorf("%s: creating directory: %v", path, err)
	}
	return nil
}

func checkExt(filename string) error {
	if !strings.HasSuffix(filename, ".rar") {
		return fmt.Errorf("filename must have a .rar extension")
	}
	return nil
}

func writeNewFile(path string, in io.Reader, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("%s: creating directory for file: %v", path, err)
	}

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", path, err)
	}
	defer out.Close()

	err = out.Chmod(mode)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", path, err)
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", path, err)
	}
	return nil
}
