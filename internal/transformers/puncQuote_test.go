package transformers

import "testing"

func TestPunc(t *testing.T) {

	expectedOutput := "As Elton John said: 'I am the most well-known homosexual in the world'"

	test := "As Elton John said: ' I am the most well-known homosexual in the world '"

	output := punQuote(test)

	if output != expectedOutput {
		t.Errorf("Output (%s), not the same as expected output (%s)", output, expectedOutput)
	}
}
