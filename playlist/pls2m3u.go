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
