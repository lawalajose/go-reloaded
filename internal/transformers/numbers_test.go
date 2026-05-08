package transformers

import "testing"

func TestNumbers(t *testing.T) {

	expectedOutput := "Simply add 66 and 2 and you will see the result is 68."

	output := numbers("Simply add 42 (hex) and 10 (bin) and you will see the result is 68.")

	if output != expectedOutput {
		t.Errorf("Output (%s), not the same as expected output (%s)", output, expectedOutput)
	}
}
