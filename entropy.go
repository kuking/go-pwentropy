package go_pwentropy

import (
	"math"
)

// Given a provided password, it will return the number of entropy bits. It is calculated estimating the symbols classes
// used in the password, i.e. if there are only lower case characters, if there are lower and upper cases, if it contains
// numbers, etc.
func EstimateEntropyByClasses(pw string) float64 {
	clsUppers := false
	clsLowers := false
	clsNumbers := false
	clsOthers := false
	for _, ch := range pw {
		if ch >= 'A' && ch <= 'Z' {
			clsUppers = true
		} else if ch >= 'a' && ch <= 'z' {
			clsLowers = true
		} else if ch >= '0' && ch <= '9' {
			clsNumbers = true
		} else {
			clsOthers = true
		}
	}
	symbols := 0
	if clsUppers {
		symbols += 26
	}
	if clsLowers {
		symbols += 26
	}
	if clsNumbers {
		symbols += 10
	}
	if clsOthers {
		symbols += 40
	}
	return entropy(symbols, len(pw))
}

// Calculates the entropy of a password by counting the unique symbols in the password and its length, this is conservative
// lower bound to the entropy.
func UniqueSymbolsEntropy(pw string) float64 {
	uniques := map[rune]bool{}
	for _, ch := range pw {
		uniques[ch] = true
	}
	return entropy(len(uniques), len(pw))
}

// Calculates a "fair" entropy, using the average between EstimateEntropyByClasses and UniqueSymbolsEntropy
func FairEntropy(pw string) float64 {
	return (EstimateEntropyByClasses(pw) + UniqueSymbolsEntropy(pw)) / 2
}

func entropy(uniqueSymbols int, length int) float64 {
	if uniqueSymbols == 0 || length == 0 {
		return 0
	}
	return math.Log2(math.Pow(float64(uniqueSymbols), float64(length-1)))
}
