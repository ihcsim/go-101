package main

import (
	"errors"
	"testing"
)

func TestWriteHeader_ReturnM3UHeader(t *testing.T) {
	expected := "EXTM3U"
	if actual := writeHeader(); expected == actual {
		t.Errorf("Expected new file header to be %b, but got %b", expected, actual)
	}
}

func TestToPls_GivenM3URecords_CanConvertToPls(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{input: `File1=Music/David Bowie/Singles 1/01-Space Oddity.ogg
Title1=David Bowie - Space Oddity
Length1=315`,
			expected: `#EXTINF:315,David Bowie - Space Oddity
Music/David Bowie/Singles 1/01-Space Oddity.ogg`},
		{input: `File2=Music/David Bowie/Singles 1/02-Changes.ogg
Title2=David Bowie - Changes
Length2=-1`,
			expected: `#EXTINF:-1,David Bowie - Changes
Music/David Bowie/Singles 1/02-Changes.ogg`},
	}

	for _, test := range tests {
		if actual, _ := toPls(test.input); test.expected != actual {
			t.Errorf("Expected PLS format to be:\n%s\n\nBut got:\n%s", test.expected, actual)
		}
	}
}

func TestToPls_GivenRecordsWithIrregularSpacing_CanTrimAndConvertToPls(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{input: `      File3=Music/David Bowie/Singles 1/03-Starman.ogg   
   Title3=David Bowie - Starman   
       Length3=258     `,
			expected: `#EXTINF:258,David Bowie - Starman
Music/David Bowie/Singles 1/03-Starman.ogg`},
		{input: `   File4   =   Music/David    Bowie/Singles     1/04-Ziggy     Stardust.ogg     
  Title4 =    David Bowie   -    Ziggy Stardust
 Length4 = 194`,
			expected: `#EXTINF:194,David Bowie - Ziggy Stardust
Music/David Bowie/Singles 1/04-Ziggy Stardust.ogg`},
	}

	for _, test := range tests {
		if actual, _ := toPls(test.input); test.expected != actual {
			t.Errorf("Expected PLS format to be:\n%s\n\nBut got:\n%s", test.expected, actual)
		}
	}
}

func TestToPls_GivenMalformedRecordsWithMissingFields_ReturnsAnError(t *testing.T) {
	var tests = []struct {
		input         string
		expected      string
		expectedError error
	}{
		{input: "",
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath, song title, song duration.")},
		{input: `Title5=David Bowie - Suffragette City
Length5=206`,
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath.")},
		{input: `File6=Music/David Bowie/Singles 1/07-The Jean Genie.ogg
Length6=247`,
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: song title.")},
		{input: `File7=Music/David Bowie/Singles 1/09-Life On Mars.ogg
Title7=David Bowie - Life On Mars?`,
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: song duration.")},
		{input: `File6=Music/David Bowie/Singles 1/07-The Jean Genie.ogg`,
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: song title, song duration.")},
		{input: `Length6=247`,
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath, song title.")},
		{input: `Title7=David Bowie - Life On Mars?`,
			expected:      "",
			expectedError: errors.New("Failed to convert record to PLS format. Missing required properties: filepath, song duration.")},
	}

	for _, test := range tests {
		actual, err := toPls(test.input)
		if actual != test.expected {
			t.Errorf("Expected toPls to return an empty string, but got %s", actual)
		}

		if err.Error() != test.expectedError.Error() {
			t.Errorf("Expected toPls to return an error with message:\n\"%s\"\n\nBut got:\n\"%s\"", test.expectedError, err)
		}
	}
}
