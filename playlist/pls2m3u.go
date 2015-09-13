package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const fileHeader = "#EXTM3U"

func writeHeader() string {
	return fileHeader
}

// SongRecord represents a song, with the song filepath, title and duration in seconds.
type SongRecord struct {
	filepath string
	title    string
	duration time.Duration
}

// NewSongRecord returns a pointer to a SongRecord.
func NewSongRecord(filepath, title string, duration string) *SongRecord {
	if len(filepath) == 0 {
		filepath = "UNKNOWN"
	}

	if len(title) == 0 {
		title = "UNKNOWN"
	}

	if len(duration) == 0 {
		duration = "-1"
	}

	durationInSeconds, _ := time.ParseDuration(duration + "s")
	return &SongRecord{
		filepath: filepath,
		title:    title,
		duration: durationInSeconds,
	}
}

func (s *SongRecord) setFilepath(filepath string) {
	if len(filepath) > 0 {
		s.filepath = filepath
	} else {
		s.filepath = "UNKNOWN"
	}
}

func (s *SongRecord) setTitle(title string) {
	if len(title) > 0 {
		s.title = title
	} else {
		s.title = "UNKNOWN"
	}
}

func (s *SongRecord) setDuration(duration string) error {
	if len(duration) > 0 {
		d, err := time.ParseDuration(duration + "s")
		if err != nil {
			return err
		}
		s.duration = d
	} else {
		s.duration = -1 * time.Second
	}

	return nil
}

// Parse takes an input of string type, attempts to extract the song filepath, title and duration
// from input, and returns a SongRecord type.
// It returns an error if any of the required song properties are missing.
func Parse(input string) (*SongRecord, error) {
	newSongRecord := NewSongRecord("", "", "")

	for _, line := range strings.Split(input, "\n") {
		property := extractAndTrimProperty(line)
		if strings.HasPrefix(strings.TrimSpace(line), "File") {
			newSongRecord.setFilepath(property)
		} else if strings.HasPrefix(strings.TrimSpace(line), "Title") {
			newSongRecord.setTitle(property)
		} else if strings.HasPrefix(strings.TrimSpace(line), "Length") {
			if err := newSongRecord.setDuration(property); err != nil {
				return nil, err
			}
		}
	}

	return newSongRecord, nil
}

func extractAndTrimProperty(input string) string {
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
