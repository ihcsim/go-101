package main

import "testing"

func TestToPls_CanConvertToPls(t *testing.T) {
	var tests = []struct {
		input    *SongRecord
		expected string
	}{
		{input: NewSongRecord(
			"Music/David Bowie/Singles 1/01-Space Oddity.ogg",
			"David Bowie - Space Oddity",
			"315"),
			expected: `#EXTINF:315,David Bowie - Space Oddity
Music/David Bowie/Singles 1/01-Space Oddity.ogg`},
		{input: NewSongRecord(
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
		{input: NewSongRecord("", "", ""),
			expected: `#EXTINF:-1,UNKNOWN
UNKNOWN`},
		{input: NewSongRecord("Music/David Bowie/Singles 1/10-Sorrow.ogg", "", ""),
			expected: `#EXTINF:-1,UNKNOWN
Music/David Bowie/Singles 1/10-Sorrow.ogg`},
		{input: NewSongRecord("", "David Bowie - Sorrow", ""),
			expected: `#EXTINF:-1,David Bowie - Sorrow
UNKNOWN`},
		{input: NewSongRecord("", "", "174"),
			expected: `#EXTINF:174,UNKNOWN
UNKNOWN`},
		{input: NewSongRecord("", "David Bowie - Sorrow", "174"),
			expected: `#EXTINF:174,David Bowie - Sorrow
UNKNOWN`},
		{input: NewSongRecord("Music/David Bowie/Singles 1/10-Sorrow.ogg", "David Bowie - Sorrow", "-1"),
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