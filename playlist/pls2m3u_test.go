package main

import (
	"errors"
	"testing"
	"time"
)

func TestWriteHeader_ReturnM3UHeader(t *testing.T) {
	expected := "EXTM3U"
	if actual := writeHeader(); expected == actual {
		t.Errorf("Expected new file header to be %b, but got %b", expected, actual)
	}
}

func TestNewSongRecord(t *testing.T) {
	var tests = []struct {
		filepath         string
		title            string
		duration         string
		expectedFilepath string
		expectedTitle    string
		expectedDuration time.Duration
	}{
		{
			filepath:         "Music/David Bowie/Singles 1/01-Space Oddity.ogg",
			title:            "David Bowie - Space Oddity",
			duration:         "315",
			expectedFilepath: "Music/David Bowie/Singles 1/01-Space Oddity.ogg",
			expectedTitle:    "David Bowie - Space Oddity",
			expectedDuration: 315 * time.Second,
		},
		{
			filepath:         "",
			title:            "",
			duration:         "",
			expectedFilepath: "UNKNOWN",
			expectedTitle:    "UNKNOWN",
			expectedDuration: -1 * time.Second,
		},
	}

	for _, test := range tests {
		s := NewSongRecord(test.filepath, test.title, test.duration)
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

func TestParse_GivenWellFormedStringInput_CanCreateSongRecord(t *testing.T) {
	var tests = []struct {
		input    string
		expected *SongRecord
	}{
		{input: `File1=Music/David Bowie/Singles 1/01-Space Oddity.ogg
Title1=David Bowie - Space Oddity
Length1=315`,
			expected: NewSongRecord(
				"Music/David Bowie/Singles 1/01-Space Oddity.ogg",
				"David Bowie - Space Oddity",
				"315"),
		},
		{input: `File2=Music/David Bowie/Singles 1/02-Changes.ogg
Title2=David Bowie - Changes
Length2=-1`,
			expected: NewSongRecord(
				"Music/David Bowie/Singles 1/02-Changes.ogg",
				"David Bowie - Changes",
				"-1"),
		},
	}

	for _, test := range tests {
		actual, err := Parse(test.input)
		if err != nil {
			t.Errorf("Tests failed with unexpected error: %s", err)
		}

		expectedSong := test.expected
		if *expectedSong != *actual {
			t.Errorf("Expected song to be:\n%+v\n\nBut got:\n%+v", expectedSong, actual)
		}
	}
}

func TestParse_GivenRecordsWithIrregularSpacing_CanTrimAndCreateSongRecord(t *testing.T) {
	var tests = []struct {
		input    string
		expected *SongRecord
	}{
		{input: `      File3=Music/David Bowie/Singles 1/03-Starman.ogg   
   Title3=David Bowie - Starman   
       Length3=258     `,
			expected: NewSongRecord(
				"Music/David Bowie/Singles 1/03-Starman.ogg",
				"David Bowie - Starman",
				"258"),
		},
		{input: `   File4   =   Music/David    Bowie/Singles     1/04-Ziggy     Stardust.ogg     
  Title4 =    David Bowie   -    Ziggy Stardust
 Length4 = 194`,
			expected: NewSongRecord(
				"Music/David Bowie/Singles 1/04-Ziggy Stardust.ogg",
				"David Bowie - Ziggy Stardust",
				"194"),
		},
	}

	for _, test := range tests {
		actual, _ := Parse(test.input)
		expectedSong := test.expected
		if *expectedSong != *actual {
			t.Errorf("Expected song to be:\n%+v\n\nBut got:\n%+v", expectedSong, actual)
		}
	}
}

func TestParse_GivenMalformedRecordsWithMissingFields_ReturnsAnError(t *testing.T) {
	var tests = []struct {
		input         string
		expectedError error
	}{
		{input: "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath, song title, song duration."),
		},
		{input: `Title5=David Bowie - Suffragette City
Length5=206`,
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath."),
		},
		{input: `File6=Music/David Bowie/Singles 1/07-The Jean Genie.ogg
Length6=247`,
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: song title."),
		},
		{input: `File7=Music/David Bowie/Singles 1/09-Life On Mars.ogg
Title7=David Bowie - Life On Mars?`,
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: song duration."),
		},
		{input: `File6=Music/David Bowie/Singles 1/07-The Jean Genie.ogg`,
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: song title, song duration."),
		},
		{input: `Length6=247`,
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath, song title."),
		},
		{input: `Title7=David Bowie - Life On Mars?`,
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath, song duration."),
		},
	}

	for _, test := range tests {
		_, err := Parse(test.input)
		if err.Error() != test.expectedError.Error() {
			t.Errorf("Expected Parse() to return an error with message:\n\"%s\"\n\nBut got:\n\"%s\"", test.expectedError, err)
		}
	}
}

func TestParse_GivenInputWithEmptyProperties_CanCreateSongRecordWithEmptyProperties(t *testing.T) {
	var tests = []struct {
		input    string
		expected *SongRecord
	}{
		{input: `File8=
Title8=
Length8=`,
			expected: NewSongRecord("UNKNOWN", "UNKNOWN", "-1.0"),
		},
		{input: `File8=Music/David Bowie/Singles 1/10-Sorrow.ogg
Title8=
Length8=`,
			expected: NewSongRecord("Music/David Bowie/Singles 1/10-Sorrow.ogg", "UNKNOWN", "-1.0"),
		},
		{input: `File8=
Title8=David Bowie - Sorrow
Length8=`,
			expected: NewSongRecord("UNKNOWN", "David Bowie - Sorrow", "-1.0"),
		},
		{input: `File8=
Title8=
Length8=174`,
			expected: NewSongRecord("UNKNOWN", "UNKNOWN", "174"),
		},
		{input: `File8=
Title8=David Bowie - Sorrow
Length8=174`,
			expected: NewSongRecord("UNKNOWN", "David Bowie - Sorrow", "174"),
		},
		{input: `File8=Music/David Bowie/Singles 1/10-Sorrow.ogg
Title8=David Bowie - Sorrow
Length8=`,
			expected: NewSongRecord("Music/David Bowie/Singles 1/10-Sorrow.ogg", "David Bowie - Sorrow", "-1.0"),
		},
	}

	for _, test := range tests {
		actual, _ := Parse(test.input)
		expectedSong := test.expected
		if *expectedSong != *actual {
			t.Errorf("Expected song to be:\n%+v\n\nBut got:\n%+v", expectedSong, actual)
		}
	}
}
