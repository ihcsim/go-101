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
		if isFilepath(property) {
			if s, err := initSong(property); err != nil {
				return nil, err
			} else {
				newSongRecord = s
			}
		} else if isTitle(property) {
			newSongRecord.setTitle(extractAndTrim(property))
		} else if isDuration(property) {
			if err := newSongRecord.setDuration(extractAndTrim(property)); err != nil {
				return nil, err
			} else {
				songRecords = append(songRecords, newSongRecord)
			}
		}
	}
	return songRecords, nil
}

func isFilepath(input string) bool {
	return strings.HasPrefix(strings.TrimSpace(input), "File")
}

func initSong(property string) (*SongRecord, error) {
	songIndex, err := extractSongIndexFromFilepath(property)
	if err != nil {
		return nil, err
	}
	return NewSongRecord(songIndex, extractAndTrim(property), "", ""), nil
}

func extractSongIndexFromFilepath(filepathProperty string) (int, error) {
	indexRx := regexp.MustCompile(`File[\d]+`)
	loc := indexRx.FindIndex([]byte(filepathProperty))
	index, err := strconv.Atoi(filepathProperty[loc[0]+4 : loc[1]])
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

func isTitle(input string) bool {
	return strings.HasPrefix(strings.TrimSpace(input), "Title")
}

func isDuration(input string) bool {
	return strings.HasPrefix(strings.TrimSpace(input), "Length")
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

func writeM3uPlaylist(songRecords []*SongRecord) error {
	fmt.Println(writeHeader())
	for _, songRecord := range songRecords {
		if r, err := songRecord.ToM3u(); err != nil {
			return err
		} else {
			fmt.Print(r)
		}
	}

	return nil
}
