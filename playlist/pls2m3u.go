package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const fileHeader = "#EXTM3U"

func writeHeader() string {
	return fileHeader
}

func toPls(m3uRecord string) (string, error) {
	filename := ""
	title := ""
	duration := ""

	foundFilename := false
	foundTitle := false
	foundDuration := false

	rx := regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)
	record := string(m3uRecord)
	for _, line := range strings.Split(record, "\n") {
		line = strings.TrimSpace(line)
		startAt := strings.Index(line, "=")
		if strings.HasPrefix(line, "File") {
			foundFilename = true
			filename = line[startAt+1:]
		} else if strings.HasPrefix(line, "Title") {
			foundTitle = true
			title = line[startAt+1:]
		} else if strings.HasPrefix(line, "Length") {
			foundDuration = true
			duration = line[startAt+1:]
		}
	}

	if !foundFilename || !foundTitle || !foundDuration {
		missingProperties := make([]string, 0, 3)
		if !foundFilename {
			missingProperties = append(missingProperties, "filepath")
		}

		if !foundTitle {
			missingProperties = append(missingProperties, "song title")
		}

		if !foundDuration {
			missingProperties = append(missingProperties, "song duration")
		}

		errMsg := fmt.Sprintf("Failed to convert record to PLS format. Missing required properties: %s.", strings.Join(missingProperties, ", "))
		err := errors.New(errMsg)
		return "", err
	}

	filename = strings.TrimSpace(rx.ReplaceAllLiteralString(filename, " "))
	title = strings.TrimSpace(rx.ReplaceAllLiteralString(title, " "))
	duration = strings.TrimSpace(rx.ReplaceAllLiteralString(duration, " "))

	pls := "#EXTINF:" + duration + "," + title + "\n" + filename
	return pls, nil
}

// use bytes.Buffer to write
