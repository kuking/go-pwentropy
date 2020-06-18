package go_pwentropy

import (
	"math"
	"testing"
)

func TestEmpty(t *testing.T) {
	assertEstimateEntropyByClasses(t, "", 0)
	assertUniqueSymbolsEntropy(t, "", 0)
}

func TestEstimateEntropyByClasses_OneClass(t *testing.T) {
	assertEstimateEntropyByClasses(t, "A", entropy(26, 1))
	assertEstimateEntropyByClasses(t, "AAAA", entropy(26, 4))

	assertEstimateEntropyByClasses(t, "a", entropy(26, 1))
	assertEstimateEntropyByClasses(t, "aa", entropy(26, 2))
	assertEstimateEntropyByClasses(t, "aaa", entropy(26, 3))
	assertEstimateEntropyByClasses(t, "aaaa", entropy(26, 4))

	assertEstimateEntropyByClasses(t, "0", entropy(10, 1))
	assertEstimateEntropyByClasses(t, "12", entropy(10, 2))

	assertEstimateEntropyByClasses(t, "!", entropy(40, 1))
	assertEstimateEntropyByClasses(t, "!\"'-_=+£$%*()[]{}:;@´#~|\\/?,.<>`¬|", entropy(40, 34))
}

func TestEstimateEntropyByClasses_MultipleClasses(t *testing.T) {
	assertEstimateEntropyByClasses(t, "aA", entropy(26+26, 2))
	assertEstimateEntropyByClasses(t, "abcABC", entropy(26+26, 6))
	assertEstimateEntropyByClasses(t, "abc123", entropy(26+10, 6))
	assertEstimateEntropyByClasses(t, "abc123", entropy(26+10, 6))
	assertEstimateEntropyByClasses(t, "abcABC123", entropy(26+26+10, 9))
	assertEstimateEntropyByClasses(t, "abcABC123!", entropy(26+26+10+40, 10))
	assertEstimateEntropyByClasses(t, "ABC123!", entropy(26+10+40, 7))
	assertEstimateEntropyByClasses(t, "ABC!", entropy(26+40, 4))
}

func TestUniqueSymbolsEntropy_OneSymbol(t *testing.T) {
	assertUniqueSymbolsEntropy(t, "A", entropy(1, 1))
	assertUniqueSymbolsEntropy(t, "AAAA", entropy(1, 4))
	assertUniqueSymbolsEntropy(t, "a", entropy(1, 1))
	assertUniqueSymbolsEntropy(t, "aa", entropy(1, 2))
	assertUniqueSymbolsEntropy(t, "aaa", entropy(1, 3))
	assertUniqueSymbolsEntropy(t, "aaaa", entropy(1, 4))
}

func TestUniqueSymbolsEntropy_Multiple(t *testing.T) {
	assertUniqueSymbolsEntropy(t, "aAbB", entropy(4, 4))
	assertUniqueSymbolsEntropy(t, "A common password", entropy(13, 17)) // dupes: ' 'mos
}

func entropy(uniqueSymbols int, length int) float64 {
	return math.Log2(math.Pow(float64(uniqueSymbols), float64(length-1)))
}

func assertEstimateEntropyByClasses(t *testing.T, pw string, expectedEntropy float64) {
	if math.Abs(EstimateEntropyByClasses(pw)-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' has an entropy of ~%2.2f", pw, expectedEntropy)
	}
}

func assertUniqueSymbolsEntropy(t *testing.T, pw string, expectedEntropy float64) {
	if math.Abs(UniqueSymbolsEntropy(pw)-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' has an entropy of ~%2.2f", pw, expectedEntropy)
	}
}
