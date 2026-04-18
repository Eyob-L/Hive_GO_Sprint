package sprint

func ReverseAlphabetValue(ch rune) rune {
	gap := ch - 'a'
	revgap := 25 - gap
	return 'a' + revgap
}
