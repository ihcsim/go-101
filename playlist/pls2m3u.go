package main

import (
	"regexp"
	"strings"
)

const fileHeader = "#EXTM3U"

func writeHeader() []byte {
	return []byte(fileHeader)
}

func toPls(m3uRecord []byte) []byte {
	filename := ""
	title := ""
	duration := ""

	rx := regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)
	record := string(m3uRecord)
	for _, line := range strings.Split(record, "\n") {
		line = strings.TrimSpace(line)
		startAt := strings.Index(line, "=")
		if strings.HasPrefix(line, "File") {
			filename = line[startAt+1:]
		} else if strings.HasPrefix(line, "Title") {
			title = line[startAt+1:]
		} else if strings.HasPrefix(line, "Length") {
			duration = line[startAt+1:]
		}
	}
	filename = strings.TrimSpace(rx.ReplaceAllLiteralString(filename, " "))
	title = strings.TrimSpace(rx.ReplaceAllLiteralString(title, " "))
	duration = strings.TrimSpace(rx.ReplaceAllLiteralString(duration, " "))

	return []byte("#EXTINF:" + duration + "," + title + "\n" + filename)
}

// use bytes.Buffer to write
