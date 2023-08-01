package utils

import "regexp"

func SenitizeWord(input string) string {
	re := regexp.MustCompile(`[\s\t\n]+`)
	return re.ReplaceAllString(input, "")
}
