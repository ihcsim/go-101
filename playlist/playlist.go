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
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		m3uToPls(rawBytes)
	}
}

func m3uToPls(b []byte) {
	songs := readM3uPlaylist(string(b))
	writePlsPlaylist(songs)
}
