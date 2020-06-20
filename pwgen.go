package go_pwentropy

import (
	"crypto/rand"
	"math/big"
)

type PwFormat int
type PwStrength int

const (
	symbolsEasy    = "abcdefghjkmnpqrstuvwxyzABCDEFHJKLMNRSTVW23456789"
	symbolsNumbers = "0123456789"
	symbolsComplex = "ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz0123456789!@$%^&*()[]{}?"

	FORMAT_EASY    PwFormat = iota
	FORMAT_NUMBERS PwFormat = iota
	FORMAT_COMPLEX PwFormat = iota

	STRENGTH_96  PwStrength = iota
	STRENGTH_128 PwStrength = iota
	STRENGTH_192 PwStrength = iota
	STRENGTH_256 PwStrength = iota
)

var pwLengthByStrengthAndFormat = map[PwStrength]map[PwFormat]int{
	STRENGTH_96:  {FORMAT_COMPLEX: 19, FORMAT_EASY: 19, FORMAT_NUMBERS: 22},
	STRENGTH_128: {FORMAT_COMPLEX: 24, FORMAT_EASY: 24, FORMAT_NUMBERS: 29},
	STRENGTH_192: {FORMAT_COMPLEX: 32, FORMAT_EASY: 33, FORMAT_NUMBERS: 42},
	STRENGTH_256: {FORMAT_COMPLEX: 41, FORMAT_EASY: 42, FORMAT_NUMBERS: 54},
}

func PwGen(pwGenFormat PwFormat, strength PwStrength) string {
	dash := 0
	var pool string
	switch pwGenFormat {
	case FORMAT_EASY:
		dash = 5
		pool = symbolsEasy
	case FORMAT_NUMBERS:
		dash = 4
		pool = symbolsNumbers
	case FORMAT_COMPLEX:
		dash = 5
		pool = symbolsComplex
	default:
		panic("I don't know that pwGenFormat")
	}

	length := pwLengthByStrengthAndFormat[strength][pwGenFormat]

	res := ""
	for l := 0; l < length; l++ {
		rb, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
		if err != nil {
			panic(err)
		}
		r := int(rb.Int64())

		res += pool[r : r+1]
		if dash > 0 && l%dash == (dash-1) && l+1 != length {
			res += "-"
		}
	}
	return res
}
