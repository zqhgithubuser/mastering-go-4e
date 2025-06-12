package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed static
var f embed.FS

var searchString string

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path=%q, isDir=%v\n", path, d.IsDir())
	return nil
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
		fileInfo, err := fs.Stat(f, path)
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size())
		return nil
	}
	return nil
}

func list(f embed.FS) error {
	return fs.WalkDir(f, ".", walkFunction)
}

func search(f embed.FS) error {
	return fs.WalkDir(f, ".", walkSearch)
}

func extract(f embed.FS, filepath string) ([]byte, error) {
	s, err := fs.ReadFile(f, filepath)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func write(s []byte, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	n, err := fd.Write(s)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

func main() {
	// List all files
	err := list(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		Path=".", isDir=true
		Path="static", isDir=true
		Path="static/file.txt", isDir=false
		Path="static/image.png", isDir=false
		Path="static/temporary.png", isDir=false
		Path="static/textfile", isDir=false
	*/

	// Search
	searchString = "file.txt"
	err = search(f) // Found static/file.txt with size 13
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract into a byte slice
	buffer, err := extract(f, "static/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save it to an actual file
	err = write(buffer, "static/IOFS.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
