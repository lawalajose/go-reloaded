package transformers

import (
	"strings"
)

func punc(sentence string) string {
	runes := []rune(sentence)
	var strToBuild strings.Builder

	for i := 0; i < len(runes); i++ {

		if pun(runes[i]) {

			if i > 0 && runes[i-1] == ' ' {
				strSoFar := strToBuild.String()
				strToBuild.Reset()

				strToBuild.WriteString(strSoFar[:len(strSoFar)-1])

			}
			strToBuild.WriteRune(runes[i])

			if i+1 < len(runes) && runes[i+1] != ' ' && !pun(runes[i+1]) {
				strToBuild.WriteRune(' ')
			}

		} else {
			strToBuild.WriteRune(runes[i])
		}
	}

	return strToBuild.String()
}

func pun(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ':', ';':
		return true
	default:
		return false
	}
}
