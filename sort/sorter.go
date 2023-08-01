package sort

type Sorter func(string) string

func NewSortString() Sorter {
	return func(str string) string {
		charCount := make([]int, 26)
		for _, char := range str {
			// NOTE: treads as different the word containing special char
			if char == '"' || char == '\'' {
				continue
			}
			charCount[char-'a']++
		}

		sortedRunes := make([]rune, 0, len(str))
		for i := 0; i < 26; i++ {
			for j := 0; j < charCount[i]; j++ {
				sortedRunes = append(sortedRunes, rune(i+'a'))
			}
		}

		return string(sortedRunes)
	}
}
