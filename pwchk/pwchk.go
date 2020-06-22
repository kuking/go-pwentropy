package main

import (
	"bufio"
	"fmt"
	pwe "github.com/kuking/go-pwentropy"
	"os"
	"syscall"
)

func main() {
	fmt.Println("pwchk: enter a password to estimate its entropy")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Password: ")
	if setTermEcho(false) != nil {
		panic("Could not disable terminal echo")
	}
	pwd, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Printf("Entropy by Unique Symbols (lower-ish bound): %2.2f\n", pwe.EntropyByUniqueSymbols(pwd))
	fmt.Printf("Entropy by Unique Symbols (excluding common passwords): %2.2f\n", pwe.EntropyByUniqueExclCommonSeqs(pwd))
	fmt.Printf("Entropy by Classes involved (upper-ish bound): %2.2f\n", pwe.EntropyByClasses(pwd))
	fmt.Println()
	fmt.Printf("=> Fair Entropy: %2.2f\n", pwe.FairEntropy(pwd))
	fmt.Println()
	fmt.Println("On Password Entropy: Anything with less than 55 bits IS NOT good. Between 55 and 64 is poor.")
	fmt.Println("Aim to use passwords with 64 bits of entropy, or more. (128 or at least 96 would be good.)")

	if setTermEcho(true) != nil {
		panic("Could not re-enable the terminal echo")
	}
}

func setTermEcho(on bool) (err error) {
	attrs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil,
	}
	var echo string
	if on {
		echo = "echo"
	} else {
		echo = "-echo"
	}
	var ws syscall.WaitStatus
	pid, err := syscall.ForkExec("/bin/stty", []string{"stty", echo}, &attrs)
	if err != nil {
		return
	}
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	return
}
