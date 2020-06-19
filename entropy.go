package go_pwentropy

import (
	"math"
	"strings"
)

// Given a provided password, it will return the number of entropy bits. It is calculated estimating the symbols classes
// used in the password, i.e. if there are only lower case characters, if there are lower and upper cases, if it contains
// numbers, etc.
func EntropyByClasses(pw string) float64 {
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

// Calculate how many unique symbols in a string
func UniqueSymbols(str string) int {
	uniques := map[rune]bool{}
	for _, ch := range str {
		uniques[ch] = true
	}
	return len(uniques)
}

// Calculates the entropy of a password by counting the unique symbols in the password and its length, this is
// conservative lower bound to the entropy. (as far as the password is not trivially generated, i.e. based on a dictionary.)
func EntropyByUniqueSymbols(pw string) float64 {
	return entropy(UniqueSymbols(pw), len(pw))
}

// Calculates the entropy like UniqueSymbols but removing common words from the length
func EntropyByUniqueExclCommonSeqs(pw string) float64 {
	return entropy(UniqueSymbols(pw), len(pw)-HowManyCommonCharSeqs(pw))
}

// Calculates a "fair" entropy, using the average between EntropyByClasses and EntropyByUniqueSymbols
func FairEntropy(pw string) float64 {
	return EntropyByUniqueSymbols(pw)*0.5 + EntropyByClasses(pw)*0.25 + EntropyByUniqueExclCommonSeqs(pw)*0.25
}

// Returns how many characters in the password are common char sequences
func HowManyCommonCharSeqs(pw string) (r int) {
	r = 0
	loPw := strings.ToLower(pw)
	for _, w := range COMMON_CHARSEQS {

		for i := strings.Index(loPw, w); i >= 0; i = strings.Index(loPw, w) {
			loPw = loPw[:i] + loPw[i+len(w):]
			r += len(w)
		}
	}
	return r
}

func entropy(uniqueSymbols int, length int) float64 {
	if uniqueSymbols == 0 || length == 0 {
		return 0
	}
	return math.Log2(math.Pow(float64(uniqueSymbols), float64(length)))
}
