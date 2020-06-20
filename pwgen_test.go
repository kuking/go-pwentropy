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
	assert.Greater(t, FairEntropy(PwGen(FORMAT_EASY, STRENGTH_96)), 96.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_EASY, STRENGTH_128)), 128.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_EASY, STRENGTH_192)), 192.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_EASY, STRENGTH_256)), 256.0)
}

func TestEntropyForNumbersFormat(t *testing.T) {
	assert.Greater(t, FairEntropy(PwGen(FORMAT_NUMBERS, STRENGTH_96)), 96.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_NUMBERS, STRENGTH_128)), 128.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_NUMBERS, STRENGTH_192)), 192.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_NUMBERS, STRENGTH_256)), 256.0)
}

func TestEntropyForComplexFormat(t *testing.T) {
	assert.Greater(t, FairEntropy(PwGen(FORMAT_COMPLEX, STRENGTH_96)), 96.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_COMPLEX, STRENGTH_128)), 128.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_COMPLEX, STRENGTH_192)), 192.0)
	assert.Greater(t, FairEntropy(PwGen(FORMAT_COMPLEX, STRENGTH_256)), 256.0)
}
