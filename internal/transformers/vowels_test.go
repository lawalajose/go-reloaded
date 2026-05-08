package transformers

import "testing"

func TestVowels(t *testing.T) {
	expectedOutput := "There it was. An amazing rock!"

	output := vowels("There it was. A amazing rock!")

	if output != expectedOutput {
		t.Errorf("Output (%s), not the same as expected output (%s)", output, expectedOutput)
	}
}
