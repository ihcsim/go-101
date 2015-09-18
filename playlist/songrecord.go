package main

import "time"

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
