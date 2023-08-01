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
	sort     sort.Sorter
	scanner  scan.WordScanner
	logger   logger.Logger
}

func NewAnagramFinder(s sort.Sorter, w scan.WordScanner, l logger.Logger) *AnagramFinder {
	return &AnagramFinder{
		anagrams: make(map[string][]string),
		sort:     s,
		scanner:  w,
		logger:   l,
	}
}

func (af *AnagramFinder) AddWord(word, sortedWord string) {
	af.anagrams[sortedWord] = append(af.anagrams[sortedWord], word)
}

func (af *AnagramFinder) printCombinedAnagrams(w io.Writer) {
	for _, words := range af.anagrams {
		if len(words) > 1 {
			af.logger.FPrintln(w, strings.Join(words, ", "))
		}
	}
}

func (af *AnagramFinder) scanLoop() {
	for {
		word, err := af.scanner.ScanWord()
		if err == io.EOF {
			break
		}
		if err != nil {
			af.logger.Printf("error reading input source: %v", err)
			return
		}

		af.AddWord(word, af.sort(strings.ToLower(utils.SenitizeWord(word))))
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
	anagramFinder.scanLoop()

	anagramFinder.printCombinedAnagrams(os.Stdout)
}

func Run(inputFileOrUrl string) {
	// NOTE: extend this case below if you want to find anagrams from different sources
	switch {
	case utils.IsFile(inputFileOrUrl):
		StartFileVersion(inputFileOrUrl)
	default:
		panic("please implement the other file types!!!")
	}
}

func main() {
	inputFileOrUrl := utils.ParseArgs()
	Run(inputFileOrUrl)
}
