package transformers

import (
	"fmt"
	"strconv"
	"strings"
)

var numberVariables = []struct {
	name string
	base int
}{
	{"(hex)", 16},
	{"(bin)", 2},
}

func numbers(sentence string) string {
	newSlice := strings.Fields(sentence)

	for _, nv := range numberVariables {

		for i := 0; i < len(newSlice); i++ {
			if i > 0 && newSlice[i] == nv.name {

				number, err := strconv.ParseInt(newSlice[i-1], nv.base, 64)
				if err != nil {
					err := fmt.Errorf("could not parse int %w", err)
					fmt.Println(err)
				}
				num := strconv.Itoa(int(number))

				newSlice[i-1] = num
				newSlice = append(newSlice[:i], newSlice[i+1:]...)

			}
		}

		sentence = strings.Join(newSlice, " ")
	}

	return sentence
}
