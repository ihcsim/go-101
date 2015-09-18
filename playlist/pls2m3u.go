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

// Parse takes an input of string type, attempts to extract the song filepath, title and duration
// from input, and returns a SongRecord type.
// It returns an error if any of the required song properties are missing.
func Parse(input string) (*SongRecord, error) {
	if valid, err := validate(input); !valid && err != nil {
		return nil, err
	}

	newSongRecord := NewSongRecord("", "", "")
	for _, property := range strings.Split(input, "\n") {
		p := extractAndTrim(property)
		if strings.HasPrefix(strings.TrimSpace(property), "File") {
			newSongRecord.setFilepath(p)
		} else if strings.HasPrefix(strings.TrimSpace(property), "Title") {
			newSongRecord.setTitle(p)
		} else if strings.HasPrefix(strings.TrimSpace(property), "Length") {
			if err := newSongRecord.setDuration(p); err != nil {
				return nil, err
			}
		}
	}

	return newSongRecord, nil
}

func validate(properties string) (bool, error) {
	filepathRx := regexp.MustCompile(`File[\d]+[\s]*=`)
	titleRx := regexp.MustCompile(`Title[\d]+[\s]*=`)
	durationRx := regexp.MustCompile(`Length[\d]+[\s]*=`)

	missingProperties := make([]string, 0, 3)
	if !filepathRx.Match([]byte(properties)) {
		missingProperties = append(missingProperties, "filepath")
	}

	if !titleRx.Match([]byte(properties)) {
		missingProperties = append(missingProperties, "song title")
	}

	if !durationRx.Match([]byte(properties)) {
		missingProperties = append(missingProperties, "song duration")
	}

	if len(missingProperties) > 0 {
		return false, errors.New(fmt.Sprintf("Failed to convert record to PLS format. Missing required properties: %s.", strings.Join(missingProperties, ", ")))
	}

	return true, nil
}

func extractAndTrim(input string) string {
	startAt := strings.Index(input, "=")
	rx := regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)
	trimmedValue := strings.TrimSpace(rx.ReplaceAllLiteralString(input[startAt+1:], " "))
	return trimmedValue
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
	if len(filename) == 0 {
		filename = "UNKNOWN"
	}

	title = strings.TrimSpace(rx.ReplaceAllLiteralString(title, " "))
	if len(title) == 0 {
		title = "UNKNOWN"
	}

	duration = strings.TrimSpace(rx.ReplaceAllLiteralString(duration, " "))
	if len(duration) == 0 {
		duration = "-1"
	}

	pls := "#EXTINF:" + duration + "," + title + "\n" + filename
	return pls, nil
}

// use bytes.Buffer to write
