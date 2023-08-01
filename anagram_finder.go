package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/bozkayasalihx/anagramfinder/logger"
	"github.com/bozkayasalihx/anagramfinder/scan"
	"github.com/bozkayasalihx/anagramfinder/sort"
	"github.com/bozkayasalihx/anagramfinder/utils"
)

type AnagramFinder struct {
	anagrams map[string][]string
	Sort     sort.Sorter
	scanner  scan.WordScanner
	logger   logger.Logger
}

func NewAnagramFinder(s sort.Sorter, w scan.WordScanner, l logger.Logger) *AnagramFinder {
	return &AnagramFinder{
		anagrams: make(map[string][]string),
		Sort:     s,
		scanner:  w,
		logger:   l,
	}
}

func (a *AnagramFinder) AddWord(word, sortedWord string) {
	a.anagrams[sortedWord] = append(a.anagrams[sortedWord], word)
}

func (a *AnagramFinder) printCombinedAnagrams(w io.Writer) {
	for _, words := range a.anagrams {
		if len(words) > 1 {
			a.logger.FPrintln(w, strings.Join(words, ", "))
		}
	}
}

func (af *AnagramFinder) ScanLoop() {
	for {
		word, err := af.scanner.ScanWord()
		if err == io.EOF {
			break
		}
		if err != nil {
			af.logger.Printf("Error reading input source: %v", err)
			return
		}

		af.AddWord(word, af.Sort(strings.ToLower(utils.SenitizeWord(word))))
	}
}

func StartFileVersion(inputFileOrUrl string) {
	file, err := os.Open(inputFileOrUrl)
	if err != nil {
		log.Fatalf("couldn't open the %s: %v", inputFileOrUrl, err)
	}
	defer file.Close()

	scanner := scan.NewScannerFromReader(file)
	sortString := sort.NewSortString()

	anagramFinder := NewAnagramFinder(sortString, scanner, &logger.DefaultLogger{})
	anagramFinder.ScanLoop()

	anagramFinder.printCombinedAnagrams(os.Stdout)
}

func Run(inputFileOrUrl string) {
	// NOTE: extend this case below if you want to find anagrams from different sources
	switch {
	case utils.IsFile(inputFileOrUrl):
		StartFileVersion(inputFileOrUrl)
	default:
		panic("please implement the other file types")
	}
}

func main() {
	inputFileOrUrl := utils.ParseArgs()
	Run(inputFileOrUrl)
}
