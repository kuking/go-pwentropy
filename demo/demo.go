package main

import (
	"fmt"
	pwe "github.com/kuking/go-pwentropy"
)

var formatToStr = map[pwe.PwFormat]string{pwe.FORMAT_COMPLEX: "Complex", pwe.FORMAT_EASY: "Easy", pwe.FORMAT_NUMBERS: "Numbers"}
var strengthToInt = map[pwe.PwStrength]int{pwe.STRENGTH_96: 96, pwe.STRENGTH_128: 128, pwe.STRENGTH_192: 192, pwe.STRENGTH_256: 256}

func printOne(format pwe.PwFormat, strength pwe.PwStrength) {
	pw := pwe.PwGen(format, strength)
	ent := pwe.FairEntropy(pw)
	fmt.Printf("%7v(min. %d, calc: %3.2f): %v\n", formatToStr[format], strengthToInt[strength], ent, pw)
}

func main() {
	fmt.Println("go-pwentropy demo")

	for _, st := range []pwe.PwStrength{pwe.STRENGTH_96, pwe.STRENGTH_128, pwe.STRENGTH_192, pwe.STRENGTH_256} {
		fmt.Println()
		fmt.Println("Minimum", strengthToInt[st], "bits of Entropy:")
		printOne(pwe.FORMAT_NUMBERS, st)
		printOne(pwe.FORMAT_EASY, st)
		printOne(pwe.FORMAT_COMPLEX, st)
	}

}
