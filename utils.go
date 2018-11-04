package utils

import (
	"github.com/mholt/archiver"
	"path"
	"log"
	"io"
	"os"
)

// CreateFileInDirectory creates a file in a directory based on an io.Reader
func CreateFileInDirectory(directory string, filename string, reader io.Reader) error {

	fo, err := os.Create(path.Join(directory,filename))
	if err != nil {
		return err
	}

	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		if _, err := fo.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// MakeZip creates a zip archive file 
func MakeZip(zipName string, files []string) error {
	err := archiver.Zip.Make(zipName, files)
	if err != nil {
		return err
	}

	return nil
}
