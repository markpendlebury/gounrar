# go unrar
[![Go Reference](https://pkg.go.dev/badge/github.com/markpendlebury/gounrar.svg)](https://pkg.go.dev/github.com/markpendlebury/gounrar)

Work with rar files in go

# Install
```
go get -v https://github.com/mpendlebury/gounrar
```

# Usage
```
import "github.com/markpendlebury/gounrar"

func Unrar {
	err := unrar.RarExtractor("myrarfile.rar", "folder/to/extract/to/")
	if err != nil{
		fmt.ErrorF("Error unpacking rar file: %v", err)
	}
}

func GetFilenameFromWithinArchive {
	filename, err := unrar.GetRarContents("myrarfile.rar")
	if err != nil {
		fmt.ErrorF("Error getting rar contents: %v", err)
	}
	fmt.Println("Found: " + filename)
}
```
  

## TODO

 - [x] Unpack rar file
 - [x] Get first filename from within an archive 
 - [ ] Get a list of files from within an archive
 - [ ] Get file size(s) from within an archive

  

## Credits
[Jagadeesh Kotra](https://github.com/jkotra/gorar/) for the base concept
[Nicholas Waples](https://github.com/nwaples/rardecode) for rardecode