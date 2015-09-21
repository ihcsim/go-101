package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || !isSupportedFormat(os.Args[1]) {
		fmt.Printf("usage: %s <file.m3u|file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filepath := os.Args[1]
	if rawBytes, err := ioutil.ReadFile(filepath); err != nil {
		log.Fatal(err)
	} else {
		if isM3uFile(filepath) {
			m3uToPls(rawBytes)
		} else if isPlsFile(filepath) {
			plsToM3u(rawBytes)
		}
	}
}

func isSupportedFormat(filepath string) bool {
	return isM3uFile(filepath) || isPlsFile(filepath)
}

func isM3uFile(filepath string) bool {
	return strings.HasSuffix(filepath, ".m3u")
}

func isPlsFile(filepath string) bool {
	return strings.HasSuffix(filepath, ".pls")
}

func m3uToPls(b []byte) {
	songRecords := parseM3uPlaylist(string(b))
	writePlsPlaylist(songRecords)
}

func plsToM3u(b []byte) {
	if songRecords, err := parsePlsPlaylist(string(b)); err != nil {
		log.Fatal(err)
	} else {
		writeM3uPlaylist(songRecords)
	}
}
