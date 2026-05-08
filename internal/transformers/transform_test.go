package transformers

import "testing"


var inputsOutput = []struct {
	input string
	expectedOutput string

} {
	{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
	{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure","I have to pack 5 outfits. Packed 26 just to be sure"},
	{"Don not be sad ,because sad backwards is das . And das not good","Don not be sad, because sad backwards is das. And das not good"},
	{"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '","Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
}

func TestTransform(t *testing.T) {

	for _, inOut:= range inputsOutput {

	expectedOutput := inOut.expectedOutput
	output:= Transform(inOut.input)



	if output != expectedOutput {
		t.Errorf("Output (%s), not the same as expected output (%s)", output, expectedOutput)
	}
}

	}

