package utils

import (
	"path"
	"log"
	"archive/zip"
	"io"
	"os"
	"strings"
)

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

func ZipFolder(zipName string, files []string, basePath string) error {
	newZipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		// to preserve the folder structure
		header.Name = strings.Replace(file, basePath, "", -1)
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}
