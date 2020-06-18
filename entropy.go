package go_pwentropy

// Given a provided password, it will return the number of entropy bits. It is calculated estimating the symbols classes
// used in the password, i.e. if there are only lower case characters, if there are lower and upper cases, if it contains
// numbers, etc.
func EstimateEntropyByClasses(pw string) float64 {

	return 0
}

// Calculates the entropy of a password by counting the unique symbols in the password and its length, this is conservative
// lower bound to the entropy.
func UniqueSymbolsEntropy(pw string) float64 {

	return 0
}

// Calculates a "fair" entropy, using the average between EstimateEntropyByClasses and UniqueSymbolsEntropy
func FairEntropy(pw string) float64 {
	return 0
}

func LongestCommonWords(pw string) int {
	return 0
}
