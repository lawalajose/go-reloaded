package main

import (
	"fmt"
	"go-reloaded/internal/transformers"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	originalText := os.Args[1]
	convertedText := os.Args[2]

	content, err := os.ReadFile(originalText)
	if err != nil {
		err = fmt.Errorf("could not read file %w", err)
		fmt.Println(err)
	}

	converted := transformers.Transform(string(content))

	os.WriteFile(convertedText, []byte(converted), 0644)

}
