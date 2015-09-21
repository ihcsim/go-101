package main

import "fmt"
import "time"

// SongRecord represents a song, with the song filepath, title and duration in seconds.
type SongRecord struct {
	index    int
	filepath string
	title    string
	duration time.Duration
}

// NewSongRecord returns a pointer to a SongRecord.
func NewSongRecord(index int, filepath, title string, duration string) *SongRecord {
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
		index:    index,
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

func (s *SongRecord) ToPls() (string, error) {
	return fmt.Sprintf("File%d=%s\nTitle%d=%s\nLength%d=%d\n",
		s.index, s.filepath,
		s.index, s.title,
		s.index, s.duration/time.Second), nil
}

func (s *SongRecord) ToM3u() (string, error) {
	return fmt.Sprintf("#EXTINF:%d,%s\n%s\n", s.duration/time.Second, s.title, s.filepath), nil
}
