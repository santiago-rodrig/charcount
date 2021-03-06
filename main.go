package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cm := make(map[rune]uint)
	bc := [utf8.UTFMax + 1]uint{}
	invalidsCount := 0
	lettersCount := 0
	digitsCount := 0
	controlsCount := 0
	whiteSpacesCount := 0
	symbolsCount := 0

	for {
		r, n, err := reader.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalidsCount++

			continue
		}

		switch {
		case unicode.IsLetter(r):
			lettersCount++
		case unicode.IsDigit(r):
			digitsCount++
		case unicode.IsSpace(r):
			whiteSpacesCount++
		case unicode.IsControl(r):
			controlsCount++
		case unicode.IsSymbol(r):
			symbolsCount++
		}

		cm[r]++
		bc[n]++
	}

	fmt.Println("char\tcount")

	for r, c := range cm {
		fmt.Printf("%s\t%d\n", string(r), c)
	}

	fmt.Println()
	fmt.Println("bytes\tcount")

	for i, c := range bc[1:] {
		fmt.Printf("%d\t%d\n", i+1, c)
	}

	fmt.Printf("\ninvalid characters: %d\n", invalidsCount)
	fmt.Printf("letters: %d\n", lettersCount)
	fmt.Printf("digits: %d\n", digitsCount)
	fmt.Printf("control characters: %d\n", controlsCount)
	fmt.Printf("blank characters: %d\n", whiteSpacesCount)
	fmt.Printf("symbols: %d\n", symbolsCount)
}
