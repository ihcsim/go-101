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
			expected: `File1=Music/David Bowie/Singles 1/01-Space Oddity.ogg
Title1=David Bowie - Space Oddity
Length1=315
`},
		{input: NewSongRecord(
			2,
			"Music/David Bowie/Singles 1/02-Changes.ogg",
			"David Bowie - Changes",
			"-1"),
			expected: `File2=Music/David Bowie/Singles 1/02-Changes.ogg
Title2=David Bowie - Changes
Length2=-1
`},
	}

	for _, test := range tests {
		if actual, _ := test.input.ToPls(); test.expected != actual {
			t.Errorf("Expected PLS format to be:\n%q\n\nBut got:\n%q", test.expected, actual)
		}
	}
}

func TestToPls_GivenRecordsWithEmptyProperties_PropertiesAreMarkedAsUnknown(t *testing.T) {
	var tests = []struct {
		input    *SongRecord
		expected string
	}{
		{input: NewSongRecord(0, "", "", ""),
			expected: `File0=UNKNOWN
Title0=UNKNOWN
Length0=-1
`},
		{input: NewSongRecord(1, "Music/David Bowie/Singles 1/10-Sorrow.ogg", "", ""),
			expected: `File1=Music/David Bowie/Singles 1/10-Sorrow.ogg
Title1=UNKNOWN
Length1=-1
`},
		{input: NewSongRecord(2, "", "David Bowie - Sorrow", ""),
			expected: `File2=UNKNOWN
Title2=David Bowie - Sorrow
Length2=-1
`},
		{input: NewSongRecord(0, "", "", "174"),
			expected: `File0=UNKNOWN
Title0=UNKNOWN
Length0=174
`},
		{input: NewSongRecord(2, "", "David Bowie - Sorrow", "174"),
			expected: `File2=UNKNOWN
Title2=David Bowie - Sorrow
Length2=174
`},
		{input: NewSongRecord(1, "Music/David Bowie/Singles 1/10-Sorrow.ogg", "David Bowie - Sorrow", "-1"),
			expected: `File1=Music/David Bowie/Singles 1/10-Sorrow.ogg
Title1=David Bowie - Sorrow
Length1=-1
`},
	}

	for _, test := range tests {
		actual, _ := test.input.ToPls()
		if actual != test.expected {
			t.Errorf("Expected PLS format to be:\n%q\n\nBut got:\n%q", test.expected, actual)
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
