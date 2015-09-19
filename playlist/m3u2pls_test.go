// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadM3uPlaylist(t *testing.T) {
	log.SetFlags(0)

	songs := parseM3uPlaylist(M3U)
	for i, song := range songs {
		if song.title != ExpectedSongs[i].title {
			t.Fatalf("%q != %q", song.title, ExpectedSongs[i].title)
		}
		if song.filepath != ExpectedSongs[i].filepath {
			t.Fatalf("%q != %q", song.filepath,
				ExpectedSongs[i].filepath)
		}
		if song.duration != ExpectedSongs[i].duration {
			t.Fatalf("%.0f != %.0f", song.duration,
				ExpectedSongs[i].duration)
		}
	}
}

func TestWritePlsPlaylist(t *testing.T) {
	songs := parseM3uPlaylist(M3U)
	var err error
	stdout := os.Stdout
	reader, writer := os.Stdin, os.Stdout
	if reader, writer, err = os.Pipe(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = writer
	writePlsPlaylist(songs)
	writer.Close()
	os.Stdout = stdout
	actual, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	reader.Close()
	if string(actual) != ExpectedPls {
		t.Fatal(fmt.Sprintf("actual != expected\nActual:%q\nExpected:%q\n", actual, ExpectedPls))
	}
}

const M3U = `#EXTM3U
#EXTINF:315,David Bowie - Space Oddity
Music/David Bowie/Singles 1/01-Space Oddity.ogg
#EXTINF:-1,David Bowie - Changes
Music/David Bowie/Singles 1/02-Changes.ogg
#EXTINF:258,David Bowie - Starman
Music/David Bowie/Singles 1/03-Starman.ogg`

var ExpectedSongs = []*SongRecord{
	NewSongRecord(
		1,
		"Music/David Bowie/Singles 1/01-Space Oddity.ogg",
		"David Bowie - Space Oddity",
		"315"),
	NewSongRecord(
		2,
		"Music/David Bowie/Singles 1/02-Changes.ogg",
		"David Bowie - Changes",
		"-1"),
	NewSongRecord(
		3,
		"Music/David Bowie/Singles 1/03-Starman.ogg",
		"David Bowie - Starman",
		"258"),
}

var ExpectedPls = `[playlist]
File1=Music/David Bowie/Singles 1/01-Space Oddity.ogg
Title1=David Bowie - Space Oddity
Length1=315
File2=Music/David Bowie/Singles 1/02-Changes.ogg
Title2=David Bowie - Changes
Length2=-1
File3=Music/David Bowie/Singles 1/03-Starman.ogg
Title3=David Bowie - Starman
Length3=258
NumberOfEntries=3
Version=2
`
