package main

import (
	"bufio"
	"fmt"
	go_pwentropy "github.com/kuking/go-pwentropy"
	"os"
	"sort"
	"strings"
)

func remove(arr []string, i int) []string {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func removeWordsWithNonAlpha(words []string) (result []string) {
	result = make([]string, 0, len(words))
	for _, w := range words {
		if !containsNonAlpha(w) {
			result = append(result, w)
		}
	}
	return
}

func containsNonAlpha(w string) bool {
	for _, c := range w {
		if c < 'a' || c > 'z' {
			return true
		}
	}
	return false
}

func removeWithLessSymbols(words []string, symbols int) (result []string) {
	result = make([]string, 0, len(words))
	for _, w := range words {
		if go_pwentropy.UniqueSymbols(w) >= symbols {
			result = append(result, w)
		}
	}
	return
}

func removeShorterLongerThan(words []string, minSize, maxSize int) (result []string) {
	result = make([]string, 0, len(words))
	for _, w := range words {
		if len(w) >= minSize && len(w) <= maxSize {
			result = append(result, w)
		}
	}
	return
}

func unique(words []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range words {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func outputWord(outfile *os.File, bytes int, colSize int, word string) int {
	if (bytes + len(word) + 4) > colSize {
		outfile.WriteString("\n")
		bytes = 0
	}
	outfile.WriteString("\"" + word + "\", ")
	return bytes + len(word) + 4
}

func main() {
	file, err := os.Open("PwnedPasswordsTop100k.txt")
	if err != nil {
		fmt.Println("You probably want to do: $ wget https://www.ncsc.gov.uk/static-assets/documents/PwnedPasswordsTop100k.txt")
		panic(err)
	}
	defer file.Close()

	words := make([]string, 0, 100000)
	scanner := bufio.NewScanner(file)
	for i := 0; i < 6; i++ { // skips header, yes hardcoded
		scanner.Scan()
	}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)
		words = append(words, line)
	}

	fmt.Println("Read", len(words), "words... compacting...")

	words = removeShorterLongerThan(words, 5, 30)
	words = removeWordsWithNonAlpha(words)
	words = removeWithLessSymbols(words, 7)
	words = unique(words)
	sort.Strings(words)

	fmt.Println("Final word count", len(words))

	outfile, err := os.Create("common_charseqs.go")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	outfile.WriteString("package go_pwentropy\n\nvar COMMON_CHARSEQS = [...]string {")
	bytes := 36

	for _, word := range words {
		bytes = outputWord(outfile, bytes, 120, word)
	}
	outfile.WriteString("}")

}
