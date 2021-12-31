package gounrar

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nwaples/rardecode"
)

// Extract Rar files from the source into destination
func RarExtractor(source string, destination string) error {

	rr, err := rardecode.OpenReader(source, "")

	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}

	for {
		header, err := rr.Next()
		if err == io.EOF {
			break
		}

		if header.IsDir {
			err = mkdir(filepath.Join(destination, header.Name), 0755)
			if err != nil {
				return err
			}
			continue
		}
		err = mkdir(filepath.Dir(filepath.Join(destination, header.Name)), 0755)
		if err != nil {
			return err
		}

		err = WriteNewFile(filepath.Join(destination, header.Name), rr, header.Mode())
		if err != nil {
			return err
		}

	}

	return nil
}

// Get filename(s) from within the Archive
func GetRarContents(source string) (string, error) {

	rr, err := rardecode.OpenReader(source, "")

	if err != nil {
		return "", fmt.Errorf("read: failed to create reader: %v", err)
	}

	header, err := rr.Next()
	if err == io.EOF {
		return "", fmt.Errorf("archive is empty: %v", err)
	}
	return header.Name, nil
}

func WriteNewFile(path string, in io.Reader, mode os.FileMode) error {
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
