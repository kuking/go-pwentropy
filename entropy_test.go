package go_pwentropy

import (
	"math"
	"testing"
)

func TestEmpty(t *testing.T) {
	assertEntropyByClasses(t, "", 0)
	assertEntropyByUniqueSymbols(t, "", 0)
}

func TestEntropyByClasses_OneClass(t *testing.T) {
	assertEntropyByClasses(t, "A", entropy(26, 1))
	assertEntropyByClasses(t, "AAAA", entropy(26, 4))

	assertEntropyByClasses(t, "a", entropy(26, 1))
	assertEntropyByClasses(t, "aa", entropy(26, 2))
	assertEntropyByClasses(t, "aaa", entropy(26, 3))
	assertEntropyByClasses(t, "aaaa", entropy(26, 4))

	assertEntropyByClasses(t, "0", entropy(10, 1))
	assertEntropyByClasses(t, "12", entropy(10, 2))

	assertEntropyByClasses(t, "!", entropy(40, 1))
	assertEntropyByClasses(t, "!\"'-_=+£$%*()[]{}:;@´#~|\\/?,.<>`¬|", entropy(40, 37))
}

func TestEntropyByClasses_MultipleClasses(t *testing.T) {
	assertEntropyByClasses(t, "aA", entropy(26+26, 2))
	assertEntropyByClasses(t, "abcABC", entropy(26+26, 6))
	assertEntropyByClasses(t, "abc123", entropy(26+10, 6))
	assertEntropyByClasses(t, "abc123", entropy(26+10, 6))
	assertEntropyByClasses(t, "abcABC123", entropy(26+26+10, 9))
	assertEntropyByClasses(t, "abcABC123!", entropy(26+26+10+40, 10))
	assertEntropyByClasses(t, "ABC123!", entropy(26+10+40, 7))
	assertEntropyByClasses(t, "ABC!", entropy(26+40, 4))
}

func TestEntropyByUniqueSymbols_OneSymbol(t *testing.T) {
	assertEntropyByUniqueSymbols(t, "A", entropy(1, 1))
	assertEntropyByUniqueSymbols(t, "AAAA", entropy(1, 4))
	assertEntropyByUniqueSymbols(t, "a", entropy(1, 1))
	assertEntropyByUniqueSymbols(t, "aa", entropy(1, 2))
	assertEntropyByUniqueSymbols(t, "aaa", entropy(1, 3))
	assertEntropyByUniqueSymbols(t, "aaaa", entropy(1, 4))
}

func TestEntropyByUniqueSymbols_Multiple(t *testing.T) {
	assertEntropyByUniqueSymbols(t, "aAbB", entropy(4, 4))
	assertEntropyByUniqueSymbols(t, "A common password", entropy(12, 17)) // dupes: ' 'mos
}

func TestEntropyByUniqueExclCommonSeqs(t *testing.T) {
	assertEntropyByUniqueExclCommonSeqs(t, "abc", entropy(3, 3))
	assertEntropyByUniqueExclCommonSeqs(t, "abcdictionary", entropy(10, 3))
	assertEntropyByUniqueExclCommonSeqs(t, "dictionary", entropy(9, 0))
	assertEntropyByUniqueExclCommonSeqs(t, "dictionarydictionary", entropy(9, 0))
	assertEntropyByUniqueExclCommonSeqs(t, "dictionarydictionarydictionary", entropy(9, 0))
	assertEntropyByUniqueExclCommonSeqs(t, "dictionary dictionary dictionary", entropy(10, 2))
}

func TestFairEntropy(t *testing.T) {
	pws := []string{"this is a dictionary password",
		"hello",
		"dictionary",
		"dictionarydictionary", "dictionary dictionary",
		"dictionarydictionarydictionary", "dictionary dictionary dictionary",
		"4HAGK-RMYKQ", "WP96N-BTY7X-DSNFQ-VAAWH", "WP96N-BTY7X-DSNFQ-VAAWH-QSTH5-AE7E7-VED5E-7TMWD"}

	for _, pw := range pws {
		clsEntropy := EntropyByClasses(pw)
		uniqEntropy := EntropyByUniqueSymbols(pw)
		uniqExclCommonEntropy := EntropyByUniqueExclCommonSeqs(pw)
		fairEntropy := uniqEntropy*0.5 + clsEntropy*0.25 + uniqExclCommonEntropy*0.25
		if FairEntropy(pw)-fairEntropy > 0.001 {
			t.Error("Fair entropy not calculated as average of both Unique and Class entropy")
		}
		//fmt.Printf("| %v | %2.2f | %2.2f | %2.2f | %2.2f |\n", pw, uniqEntropy, clsEntropy, uniqExclCommonEntropy, fairEntropy)
	}
}

func TestHowManyCommonCharSeqs(t *testing.T) {
	if HowManyCommonCharSeqs("showmethemoney") != 14 ||
		HowManyCommonCharSeqs("its complicated to go from buenosaires to singapore") != 30 {
		t.Fail()
	}
}

// -------------------------------------------------------------------------------------------------------------------

func assertEntropyByClasses(t *testing.T, pw string, expectedEntropy float64) {
	retEntropy := EntropyByClasses(pw)
	if math.Abs(retEntropy-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' calculated entropy %2.2f, but expected ~%2.2f", pw, retEntropy, expectedEntropy)
	}
}

func assertEntropyByUniqueSymbols(t *testing.T, pw string, expectedEntropy float64) {
	retEntropy := EntropyByUniqueSymbols(pw)
	if math.Abs(retEntropy-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' calculated entropy %2.2f, but expected ~%2.2f", pw, retEntropy, expectedEntropy)
	}
}

func assertEntropyByUniqueExclCommonSeqs(t *testing.T, pw string, expectedEntropy float64) {
	retEntropy := EntropyByUniqueExclCommonSeqs(pw)
	if math.Abs(retEntropy-expectedEntropy) > 0.001 {
		t.Errorf("Password '%v' calculated entropy %2.2f, but expected ~%2.2f", pw, retEntropy, expectedEntropy)
	}

}
