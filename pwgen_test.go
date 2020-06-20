package go_pwentropy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestPwGen(t *testing.T) {
//	for i := 0; i < 10000; i++ {
//		pw := PwGen(FORMAT_EASY, STRENGTH_96)
//		fmt.Println(pw, FairEntropy(pw))
//	}
//}

func TestEntropyForEasyFormat(t *testing.T) {
	assert.Greater(t, FairEntropy(PwGen(FormatEasy, Strength96)), 96.0)
	assert.Greater(t, FairEntropy(PwGen(FormatEasy, Strength128)), 128.0)
	assert.Greater(t, FairEntropy(PwGen(FormatEasy, Strength192)), 192.0)
	assert.Greater(t, FairEntropy(PwGen(FormatEasy, Strength256)), 256.0)
}

func TestEntropyForNumbersFormat(t *testing.T) {
	assert.Greater(t, FairEntropy(PwGen(FormatNumbers, Strength96)), 96.0)
	assert.Greater(t, FairEntropy(PwGen(FormatNumbers, Strength128)), 128.0)
	assert.Greater(t, FairEntropy(PwGen(FormatNumbers, Strength192)), 192.0)
	assert.Greater(t, FairEntropy(PwGen(FormatNumbers, Strength256)), 256.0)
}

func TestEntropyForComplexFormat(t *testing.T) {
	assert.Greater(t, FairEntropy(PwGen(FormatComplex, Strength96)), 96.0)
	assert.Greater(t, FairEntropy(PwGen(FormatComplex, Strength128)), 128.0)
	assert.Greater(t, FairEntropy(PwGen(FormatComplex, Strength192)), 192.0)
	assert.Greater(t, FairEntropy(PwGen(FormatComplex, Strength256)), 256.0)
}
