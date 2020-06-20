package main

import (
	"fmt"
	pwe "github.com/kuking/go-pwentropy"
)

var formatToStr = map[pwe.PwFormat]string{pwe.FormatComplex: "Complex", pwe.FormatEasy: "Easy", pwe.FormatNumbers: "Numbers"}
var strengthToInt = map[pwe.PwStrength]int{pwe.Strength96: 96, pwe.Strength128: 128, pwe.Strength192: 192, pwe.Strength256: 256}

func printOne(format pwe.PwFormat, strength pwe.PwStrength) {
	pw := pwe.PwGen(format, strength)
	ent := pwe.FairEntropy(pw)
	fmt.Printf("%7v(min. %d, calc: %3.2f): %v\n", formatToStr[format], strengthToInt[strength], ent, pw)
}

func main() {
	fmt.Println("go-pwentropy demo")

	for st, bits := range strengthToInt {
		fmt.Println()
		fmt.Println("Minimum", bits, "bits of Entropy:")
		printOne(pwe.FormatNumbers, st)
		printOne(pwe.FormatEasy, st)
		printOne(pwe.FormatComplex, st)
	}

}
