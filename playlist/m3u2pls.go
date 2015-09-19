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
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs)
	}
}

func readM3uPlaylist(data string) (songs []*SongRecord) {
	var songIndex int
	var filepath, title, duration string
	for index, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}

		if strings.HasPrefix(line, "#EXTINF:") {
			songIndex += 1
			title, duration = parseExtinfLine(line)
		} else {
			filepath = strings.Map(mapPlatformDirSeparator, line)
		}

		if index%2 == 0 && filepath != "" && title != "" && duration != "" {
			songs = append(songs, NewSongRecord(songIndex, filepath, title, duration))
			filepath, title, duration = "", "", ""
		}
	}
	return songs
}

func parseExtinfLine(line string) (title, duration string) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			duration = line[:j]
		}
	}
	return
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlsPlaylist(songs []*SongRecord) error {
	fmt.Println("[playlist]")
	for _, song := range songs {
		if m3u, err := song.ToM3u(); err != nil {
			return err
		} else {
			fmt.Printf(m3u)
		}
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
	return nil
}
