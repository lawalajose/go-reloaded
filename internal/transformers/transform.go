package transformers

func Transform(sentence string) string{

	result := numbers(sentence)
	result  = alphabets(result)

	result = punc(result)
	result = punQuote(result)

	result = vowels(result)

	return result
}
