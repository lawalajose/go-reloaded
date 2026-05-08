package transformers

import "testing"

func TestAlphabets(t *testing.T) {

	

	expectedOutput := "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"
	test := "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house? (up, 2)" 
	output := alphabets(test)
		if output != expectedOutput {
		t.Errorf("Output [%s]), not the same as expected output (%s)", output,expectedOutput)
	}
}
