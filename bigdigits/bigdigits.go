// Copyright © 2010-12 Qtrac Ltd.
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
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	readFlags()

	if showUsage {
		showUsageAndExit(0)
	}

	if insufficientArgs() {
		showUsageAndExit(1)
	}

	input := os.Args[len(os.Args)-1]
	for row := 0; row < DIGIT_HEIGHT; row++ {
		line := ""
		for _, char := range input {
			if digit, valid := toDigit(char); valid {
				line += toBigDigit(digit, row)
			} else {
				log.Fatal("invalid whole number")
			}
		}
		print(line, row)
	}
}

var showUsage, showBars bool

func readFlags() {
	flag.BoolVar(&showUsage, "help", false, "Show usages")
	flag.BoolVar(&showUsage, "h", false, "Show usages (shorthand)")
	flag.BoolVar(&showBars, "bar", false, "Display header and footer bars in output")
	flag.BoolVar(&showBars, "b", false, "Display header and footer bars in output (shorthand)")
	flag.Parse()
}

func insufficientArgs() bool {
	return len(os.Args) <= 1
}

func showUsageAndExit(exitStatus int) {
	fmt.Printf("usage: bigdigits [-b | --bar] <whole-number>\n\n-b --bar draw an underbar and an overbar\n")
	os.Exit(exitStatus)
}

func toDigit(char rune) (rune, bool) {
	digit := char - '0'
	if 0 <= digit && digit <= 9 {
		return digit, true
	}
	return 0, false
}

func toBigDigit(digit rune, row int) string {
	return bigDigits[digit][row] + "  "
}

func print(line string, row int) {
	if row == 0 && showBars {
		fmt.Println(strings.Repeat("*", len(line)))
	}

	fmt.Println(line)

	if row == DIGIT_HEIGHT-1 && showBars {
		fmt.Println(strings.Repeat("*", len(line)))
	}
}

// magnified output of each digit
const DIGIT_HEIGHT = 7

var bigDigits = [][DIGIT_HEIGHT]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
