package main

import "testing"
import "time"

func TestToPls_CanConvertToPls(t *testing.T) {
	var tests = []struct {
		input    *SongRecord
		expected string
	}{
		{input: NewSongRecord(
			1,
			"Music/David Bowie/Singles 1/01-Space Oddity.ogg",
			"David Bowie - Space Oddity",
			"315"),
			expected: `#EXTINF:315,David Bowie - Space Oddity
Music/David Bowie/Singles 1/01-Space Oddity.ogg`},
		{input: NewSongRecord(
			2,
			"Music/David Bowie/Singles 1/02-Changes.ogg",
			"David Bowie - Changes",
			"-1"),
			expected: `#EXTINF:-1,David Bowie - Changes
Music/David Bowie/Singles 1/02-Changes.ogg`},
	}

	for _, test := range tests {
		if actual, _ := test.input.ToPls(); test.expected != actual {
			t.Errorf("Expected PLS format to be:\n%s\n\nBut got:\n%s", test.expected, actual)
		}
	}
}

func TestToPls_GivenRecordsWithEmptyProperties_PropertiesAreMarkedAsUnknown(t *testing.T) {
	var tests = []struct {
		input    *SongRecord
		expected string
	}{
		{input: NewSongRecord(0, "", "", ""),
			expected: `#EXTINF:-1,UNKNOWN
UNKNOWN`},
		{input: NewSongRecord(1, "Music/David Bowie/Singles 1/10-Sorrow.ogg", "", ""),
			expected: `#EXTINF:-1,UNKNOWN
Music/David Bowie/Singles 1/10-Sorrow.ogg`},
		{input: NewSongRecord(2, "", "David Bowie - Sorrow", ""),
			expected: `#EXTINF:-1,David Bowie - Sorrow
UNKNOWN`},
		{input: NewSongRecord(0, "", "", "174"),
			expected: `#EXTINF:174,UNKNOWN
UNKNOWN`},
		{input: NewSongRecord(2, "", "David Bowie - Sorrow", "174"),
			expected: `#EXTINF:174,David Bowie - Sorrow
UNKNOWN`},
		{input: NewSongRecord(1, "Music/David Bowie/Singles 1/10-Sorrow.ogg", "David Bowie - Sorrow", "-1"),
			expected: `#EXTINF:-1,David Bowie - Sorrow
Music/David Bowie/Singles 1/10-Sorrow.ogg`},
	}

	for _, test := range tests {
		actual, _ := test.input.ToPls()
		if actual != test.expected {
			t.Errorf("Expected PLS format to be:\n%s\n\nBut got:\n%s", test.expected, actual)
		}
	}
}

func TestNewSongRecord(t *testing.T) {
	var tests = []struct {
		index            int
		filepath         string
		title            string
		duration         string
		expectedIndex    int
		expectedFilepath string
		expectedTitle    string
		expectedDuration time.Duration
	}{
		{
			index:            1,
			filepath:         "Music/David Bowie/Singles 1/01-Space Oddity.ogg",
			title:            "David Bowie - Space Oddity",
			duration:         "315",
			expectedIndex:    1,
			expectedFilepath: "Music/David Bowie/Singles 1/01-Space Oddity.ogg",
			expectedTitle:    "David Bowie - Space Oddity",
			expectedDuration: 315 * time.Second,
		},
		{
			index:            0,
			filepath:         "",
			title:            "",
			duration:         "",
			expectedIndex:    0,
			expectedFilepath: "UNKNOWN",
			expectedTitle:    "UNKNOWN",
			expectedDuration: -1 * time.Second,
		},
	}

	for _, test := range tests {
		s := NewSongRecord(test.index, test.filepath, test.title, test.duration)
		if s.index != test.expectedIndex {
			t.Errorf("Expected song index to be %d, but got %d", test.expectedIndex, s.index)
		}

		if s.filepath != test.expectedFilepath {
			t.Errorf("Expected song filepath to be %s, but got %s", test.expectedFilepath, s.filepath)
		}

		if s.title != test.expectedTitle {
			t.Errorf("Expected song title to be %s, but got %s", test.expectedTitle, s.title)
		}

		if s.duration != test.expectedDuration {
			t.Errorf("Expected song duration to be %s, but got %s", test.expectedDuration, s.duration)
		}
	}
}
