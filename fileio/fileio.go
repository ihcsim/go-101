package main

import (
	"bufio"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "fileio: ", log.Ldate)

func main() {
	_, err := os.Open("nonexistent.txt")
	if err != nil {
		logger.Println(err)
	}

	file, err := os.Open("sample_license.txt")
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()

	filestat, _ := file.Stat()
	fileContent := make([]byte, filestat.Size())
	bufReader := bufio.NewReader(file)
	counts, err := bufReader.Read(fileContent)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Read %d bytes from sample_license.txt\n\n************\nFile Content\n************\n\n%s\n", counts, fileContent)
}
