package transformers

import "testing"

func TestPuncQuote(t *testing.T) {

	expectedOutput := "I was sitting over there, and then BAMM!!"

	test := "I was sitting over there ,and then BAMM !!" 

	output := punc(test)

	if output != expectedOutput {
		t.Errorf("Output (%s), not the same as expected output (%s)", output, expectedOutput)
	}
}
