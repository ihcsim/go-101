package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const fileHeader = "#EXTM3U"

func writeHeader() string {
	return fileHeader
}

// parsePlsPlaylist takes an input of string type, attempts to extract all the song filepaths, titles and durations
// from input, and returns a slice of SongRecords.
// It returns an error if any of the required song properties are missing.
func parsePlsPlaylist(input string) ([]*SongRecord, error) {
	songRecords := make([]*SongRecord, 0)
	var newSongRecord *SongRecord

	for _, property := range strings.Split(input, "\n") {
		p := extractAndTrim(property)
		if strings.HasPrefix(strings.TrimSpace(property), "File") {
			songIndex, err := extractSongIndex(property)
			if err != nil {
				return nil, err
			}
			newSongRecord = NewSongRecord(songIndex, "", "", "")
			newSongRecord.setFilepath(p)
		} else if strings.HasPrefix(strings.TrimSpace(property), "Title") {
			newSongRecord.setTitle(p)
		} else if strings.HasPrefix(strings.TrimSpace(property), "Length") {
			if err := newSongRecord.setDuration(p); err != nil {
				return nil, err
			} else {
				songRecords = append(songRecords, newSongRecord)
			}
		}
	}
	return songRecords, nil
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

func extractSongIndex(input string) (int, error) {
	indexRx := regexp.MustCompile(`File[\d]+`)
	loc := indexRx.FindIndex([]byte(input))
	index, err := strconv.Atoi(input[loc[0]+4 : loc[1]])
	if err != nil {
		return 0, err
	}
	return index, nil
}

func extractAndTrim(input string) string {
	startAt := strings.Index(input, "=")
	rx := regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)
	trimmedValue := strings.TrimSpace(rx.ReplaceAllLiteralString(input[startAt+1:], " "))
	return trimmedValue
}
