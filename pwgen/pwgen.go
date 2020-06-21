package main

import (
	"flag"
	"fmt"
	pwe "github.com/kuking/go-pwentropy"
	"os"
)

var strength int
var format string
var doHelp bool

var pwFormat pwe.PwFormat
var pwStrength pwe.PwStrength

func main() {

	flag.IntVar(&strength, "s", 96, "Password strength, measured in entropy strength, valid options: 96, 128, 192, 256.")
	flag.StringVar(&format, "f", "easy", "Password format, options: easy, complex, numeric.")
	flag.BoolVar(&doHelp, "h", false, "Show usage")
	flag.Parse()
	if doHelp {
		fmt.Printf("Usage of %v:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	switch strength {
	case 96:
		pwStrength = pwe.Strength96
	case 128:
		pwStrength = pwe.Strength128
	case 192:
		pwStrength = pwe.Strength192
	case 256:
		pwStrength = pwe.Strength256
	default:
		fmt.Println("Strength requested not supported:", strength)
		os.Exit(1)
	}

	switch format {
	case "easy":
		pwFormat = pwe.FormatEasy
	case "numeric":
		pwFormat = pwe.FormatNumbers
	case "complex":
		pwFormat = pwe.FormatComplex
	default:
		fmt.Println("Format request not supported:", format)
		os.Exit(1)
	}

	fmt.Println(pwe.PwGen(pwFormat, pwStrength))
}
