package scan

import (
	"bufio"
	"io"
)

type WordScanner interface {
	ScanWord() (string, error)
}

type ScannerFromReader struct {
	scanner *bufio.Scanner
}

func NewScannerFromReader(r io.Reader) *ScannerFromReader {
	return &ScannerFromReader{
		scanner: bufio.NewScanner(r),
	}
}

func (s *ScannerFromReader) ScanWord() (string, error) {
	if !s.scanner.Scan() {
		if err := s.scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return s.scanner.Text(), nil
}
