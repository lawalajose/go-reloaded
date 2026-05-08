package transformers

import "strings"

func punQuote(sentence string) string {
	runes := []rune(sentence)

	var strToBuild strings.Builder
	isInQuote := false

	for i := 0; i < len(runes); i++ {

		if i > 0 && runes[i] == '\'' {
			strToBuild.WriteRune('\'')
			isInQuote = !isInQuote
			continue

		}

		if isInQuote && runes[i] == ' ' {

			if runes[i-1] == '\'' {
				continue
			}

			if runes[i+1] == '\'' {
				continue
			}
		}
		strToBuild.WriteRune(runes[i])
	}
	return strToBuild.String()
}
