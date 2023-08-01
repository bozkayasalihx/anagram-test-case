package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bozkayasalihx/anagramfinder/logger"
	"github.com/bozkayasalihx/anagramfinder/scan"
	"github.com/bozkayasalihx/anagramfinder/sort"
)

func genMockInput() string {
	inputData := "listen\nsilent\nenlist\nword\n"
	return inputData
}

func genMockScanner() *scan.ScannerFromReader {
	mockScanner := scan.NewScannerFromReader(bytes.NewReader([]byte(genMockInput())))
	return mockScanner
}

func TestAddWord(t *testing.T) {
	anagramFinder := NewAnagramFinder(sort.NewSortString(), nil, nil)

	anagramFinder.AddWord("listen", "eilnst")
	anagramFinder.AddWord("silent", "eilnst")
	anagramFinder.AddWord("enlist", "eilnst")
	anagramFinder.AddWord("word", "dorw")

	expectedAnagrams := map[string][]string{
		"eilnst": {"listen", "silent", "enlist"},
		"dorw":   {"word"},
	}

	assert.Equal(t, anagramFinder.anagrams, expectedAnagrams)
}

func TestPrintCombinedAnagrams(t *testing.T) {
	anagramFinder := NewAnagramFinder(sort.NewSortString(), nil, &logger.DefaultLogger{})

	anagramFinder.anagrams = map[string][]string{
		"eilnst": {"listen", "silent", "enlist"},
		"dorw":   {"word"},
	}

	expectedOutput := "listen, silent, enlist\n"

	buffer := &bytes.Buffer{}
	anagramFinder.printCombinedAnagrams(buffer)

	assert.Equal(t, expectedOutput, buffer.String())
}

func TestAnagramFinder(t *testing.T) {
	mockScanner := genMockScanner()
	outputBuffer := &bytes.Buffer{}
	anagramFinder := NewAnagramFinder(sort.NewSortString(), mockScanner, &logger.DefaultLogger{})

	anagramFinder.ScanLoop()
	anagramFinder.printCombinedAnagrams(outputBuffer)

	expectedOutput := "listen, silent, enlist\n"
	assert.Equal(t, expectedOutput, outputBuffer.String())
}

func WriteToTestFile(inputData string, tempFile string) (*os.File, error) {
	file, err := os.CreateTemp("", tempFile)
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(inputData)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func TestRun(t *testing.T) {
	inputFile := "https:/sjlfjsdlfjsdlfj.com/blob/text.txt"
	assert.Panics(t, func() { Run(inputFile) })
}

func TestMain_ExistingInputFile(t *testing.T) {
	inputData := "listen\nsilent\nenlist\nword"

	file, err := WriteToTestFile(inputData, "input.txt")
	assert.NoError(t, err)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"main", file.Name()}

	realStdout := os.Stdout

	tempFile, err := os.CreateTemp("", "output.txt")
	assert.NoError(t, err)

	defer func() {
		os.Remove(file.Name())
		file.Close()

		tempFile.Close()
		os.Stdout = realStdout
		os.Remove(tempFile.Name())
	}()

	os.Stdout = tempFile

	main()

	capturedOutput, err := os.ReadFile(tempFile.Name())
	assert.NoError(t, err)
	expectedOutput := "listen, silent, enlist\n"
	assert.Equal(t, expectedOutput, string(capturedOutput))
}
