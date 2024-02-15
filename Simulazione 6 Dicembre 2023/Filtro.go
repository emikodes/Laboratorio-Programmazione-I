package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	stringa := scanner.Text()

	for i := 0; i < len(stringa)-1; i++ {
		if stringa[i] < unicode.MaxASCII && unicode.IsLetter(rune(stringa[i])) && unicode.IsDigit(rune(stringa[i+1])) {
			fmt.Println(string(stringa[i]))
		}
	}
}
