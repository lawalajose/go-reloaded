package transformers

import (
	"fmt"
	"strconv"
	"strings"
)

var alphabetVariables = []struct {
	name     string
	name2    string
	function func(string) string
}{
	{"(up)", "(up,", strings.ToUpper},
	{"(low)", "(low,", strings.ToLower},
	{"(cap)", "(cap,", title},
}

func alphabets(sentence string) string {

	for _, av := range alphabetVariables {

		newSlice := strings.Fields(sentence)

		for i := 0; i < len(newSlice); i++ {
			if i > 0 && newSlice[i] == av.name {
				newSlice[i-1] = av.function(newSlice[i-1])
				newSlice = append(newSlice[:i], newSlice[i+1:]...)
			} else if i+1 > 0 && newSlice[i] == av.name2 {


				num := strings.TrimRight(newSlice[i+1], ")")
				number, err := strconv.Atoi(num)
				if err != nil {
					err = fmt.Errorf("could not covert string to int %w", err)
					fmt.Println(err)
				}

				for j := 1; j <= number; j++ {
						newSlice[i-j] = av.function(newSlice[i-j])
				}
				newSlice = append(newSlice[:i], newSlice[i+2:]...)
			}

		}
		sentence = strings.Join(newSlice, " ")

	}
	return sentence

}

func title(word string) string {
	return strings.ToUpper(string(word[0])) + word[1:]
}
