package utils

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

type FileScanner struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func (fs *FileScanner) Close() {
	fs.File.Close()
}

func getFileName(day int) string {
	if testing.Testing() {
		return "example"
	} else {
		return fmt.Sprintf("day%02d/input", day)
	}
}

func ScanInput(day int) *FileScanner {
	file, err := os.Open(getFileName(day))
	CheckErr(err)

	return &FileScanner{File: file, Scanner: bufio.NewScanner(file)}
}
