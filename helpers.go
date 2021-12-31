package gounrar

import (
	"fmt"
	"os"
	"strings"
)

func mkdir(path string, dirMode os.FileMode) error {
	err := os.MkdirAll(path, dirMode)
	if err != nil {
		return fmt.Errorf("%s: creating directory: %v", path, err)
	}
	return nil
}

// CheckExt ensures the file extension matches the format.
func CheckExt(filename string) error {
	if !strings.HasSuffix(filename, ".rar") {
		return fmt.Errorf("filename must have a .rar extension")
	}
	return nil
}
