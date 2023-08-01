package utils

import (
	"os"
	"strings"
)

func ParseArgs() string {
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	inputFileOrUrl := os.Args[1]
	return inputFileOrUrl
}

func IsFile(str string) bool {
	if strings.Contains(str, "http") {
		return false
	}

	return true
}
