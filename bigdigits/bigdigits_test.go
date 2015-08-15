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
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBigDigits(t *testing.T) {
	command := setCmd()
	actual, err := execCmd(command)
	if err != nil {
		t.Fatalf("Failed to execute command %s: %s\n", actual, err)
	}

	path, _ := os.Getwd()
	expected, err := readExpectedOutput(filepath.Join(path, "0123456789.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(actual, expected) != 0 {
		t.Errorf("Expected output didn't match.\nExpected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestHelp(t *testing.T) {
	cmd := setCmd("-help")
	actual, err := execCmd(cmd)
	if err != nil {
		t.Fatalf("Failed to execute command: (%s)\n%s", err, actual)
	}

	expected := "usage: bigdigits [-b | --bar] <whole-number>\n\n-b --bar draw an underbar and an overbar\n"
	if string(actual) != expected {
		t.Errorf("Expected usage message didn't match.\nExpected:\n%s\nGot:\n%s", expected, string(actual))
	}
}

func TestHelp_WithShortHand(t *testing.T) {
	cmd := setCmd("-h")
	actual, err := execCmd(cmd)
	if err != nil {
		t.Fatalf("Failed to execute command: (%s)\n%s", err, actual)
	}

	expected := "usage: bigdigits [-b | --bar] <whole-number>\n\n-b --bar draw an underbar and an overbar\n"
	if string(actual) != expected {
		t.Errorf("Expected usage message didn't match.\nExpected:\n%s\nGot:\n%s", expected, string(actual))
	}
}

func TestBar(t *testing.T) {
	cmd := setCmd("-bar")
	actual, err := execCmd(cmd)
	if err != nil {
		t.Fatalf("Failed to execute command: (%s)\n%s", err, actual)
	}

	path, _ := os.Getwd()
	expected, err := readExpectedOutput(filepath.Join(path, "0123456789_WithBars.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(actual, expected) != 0 {
		t.Errorf("Expected output didn't match.\nExpected:\n%s\nGot:\n%s", expected, actual)
	}
}

func setCmd(args ...string) *exec.Cmd {
	executable := filepath.Join(os.Getenv("GOPATH"), "bin/bigdigits")
	input := "0123456789"
	command := exec.Command(executable, strings.Join(args, " "), input)
	return command
}

func execCmd(command *exec.Cmd) ([]byte, error) {
	reader, writer, _ := os.Pipe()
	defer reader.Close()

	command.Stdout = writer

	var stderr bytes.Buffer
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		return stderr.Bytes(), err
	}

	writer.Close()
	return ioutil.ReadAll(reader)
}

func readExpectedOutput(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}
