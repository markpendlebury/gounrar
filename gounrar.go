package unrar

import (
	"fmt"
	"io"
	"path/filepath"

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
