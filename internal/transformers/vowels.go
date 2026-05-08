package transformers

import "strings"

func vowels(sentence string) string {
	newSlice := strings.Fields(sentence)

	for i, word := range newSlice {
		if isVowel(word[0]) {

			if newSlice[i-1] == "a" {
				newSlice[i-1] = "an"
			}

			if newSlice[i-1] == "A" {
				newSlice[i-1] = "An"
			}
		}
	}

	return strings.Join(newSlice, " ")

}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
