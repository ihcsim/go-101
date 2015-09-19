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
	"strings"
)

func parseM3UPlaylist(data string) (songs []*SongRecord) {
	var songIndex int
	var newSong *SongRecord

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if !isSongRecord(line) {
			continue
		}

		if isSongRecordHeader(line) {
			songIndex += 1
			newSong = extractSongRecordHeader(songIndex, line)
		} else {
			newSong.setFilepath(line)
			songs = append(songs, newSong)
		}
	}
	return songs
}

func isSongRecord(input string) bool {
	return input != "" && !strings.HasPrefix(input, "#EXTM3U")
}

func isSongRecordHeader(input string) bool {
	return strings.HasPrefix(input, "#EXTINF:")
}

func extractSongRecordHeader(index int, line string) (s *SongRecord) {
	var title, duration string
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			duration = line[:j]
		}
	}

	return NewSongRecord(index, "", title, duration)
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
