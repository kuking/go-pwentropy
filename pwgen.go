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

	FormatEasy    PwFormat = iota
	FormatNumbers PwFormat = iota
	FormatComplex PwFormat = iota

	Strength96  PwStrength = iota
	Strength128 PwStrength = iota
	Strength192 PwStrength = iota
	Strength256 PwStrength = iota
)

var pwLengthByStrengthAndFormat = map[PwStrength]map[PwFormat]int{
	Strength96:  {FormatComplex: 19, FormatEasy: 19, FormatNumbers: 22},
	Strength128: {FormatComplex: 24, FormatEasy: 24, FormatNumbers: 29},
	Strength192: {FormatComplex: 32, FormatEasy: 33, FormatNumbers: 42},
	Strength256: {FormatComplex: 41, FormatEasy: 42, FormatNumbers: 54},
}

func PwGen(pwGenFormat PwFormat, strength PwStrength) string {
	dash := 0
	var pool string
	switch pwGenFormat {
	case FormatEasy:
		dash = 5
		pool = symbolsEasy
	case FormatNumbers:
		dash = 4
		pool = symbolsNumbers
	case FormatComplex:
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
