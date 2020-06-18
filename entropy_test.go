package go_pwentropy

import (
	"fmt"
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
	assertEstimateEntropyByClasses(t, "!\"'-_=+£$%*()[]{}:;@´#~|\\/?,.<>`¬|", entropy(40, 37))
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
	assertUniqueSymbolsEntropy(t, "A common password", entropy(12, 17)) // dupes: ' 'mos
}

func TestFairEntropy(t *testing.T) {
	pw := "this is a dictionary password"
	clsEntropy := EstimateEntropyByClasses(pw)
	uniqEntropy := UniqueSymbolsEntropy(pw)
	fairEntropy := (clsEntropy + uniqEntropy) / 2

	// 250k words in the average english speaker vocabulary, 5 words + space
	//dictionaryEntropy := entropy(250000, 5+1)
	//fmt.Println("class entropy:", clsEntropy)
	//fmt.Println("unique entropy:", uniqEntropy)
	//fmt.Println("fair entropy:", fairEntropy)
	//fmt.Println("dictionary entropy:", dictionaryEntropy)

	fmt.Printf("%2.4f\n", entropy(4, 5))
	fmt.Printf("%2.4f\n", entropy(26, 5))
	fmt.Printf("%2.4f\n", entropy(250000, 1))
	fmt.Printf("%2.4f\n", entropy(25000, 1))
	fmt.Printf("%2.4f\n", entropy(2, 1))

	if FairEntropy(pw)-fairEntropy > 0.001 {
		t.Error("Fair entropy not calculated as average of both Unique and Class entropy")
	}
}

func assertEstimateEntropyByClasses(t *testing.T, pw string, expectedEntropy float64) {
	retEntropy := EstimateEntropyByClasses(pw)
	if math.Abs(retEntropy-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' calculated entropy %2.2f, but expected ~%2.2f", pw, retEntropy, expectedEntropy)
	}
}

func assertUniqueSymbolsEntropy(t *testing.T, pw string, expectedEntropy float64) {
	retEntropy := UniqueSymbolsEntropy(pw)
	if math.Abs(retEntropy-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' calculated entropy %2.2f, but expected ~%2.2f", pw, retEntropy, expectedEntropy)
	}
}
