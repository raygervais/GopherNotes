package main

import (
	"io/ioutil"
	"os"
)

func CheckFileErrors(e error) {
	if e != nil {
		panic(e)
	}
}

func ValidateFilePath(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
	CheckFileErrors(err)

	return string(content)
}

func WriteFile(file *os.File, data []byte) *os.File {
	_, err := file.Write(data)
	CheckFileErrors(err)

	return file
}

func CreateNotesFile(path string) *os.File {
	file, err := os.Create(path)
	CheckFileErrors(err)

	return file
}
