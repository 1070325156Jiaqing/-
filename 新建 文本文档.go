package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a text to reverse.")
		return
	}

	text := os.Args[1]
	reversed := reverseText(text)

	fmt.Printf("Reversed text: %s", reversed)
}

func reverseText(text string) string {
	runes := []rune(text)
	reversed := make([]rune, len(runes))

	for i, j := len(runes)-1, 0; i >= 0; i-- {
		reversed[j] = runes[i]
		j++
	}

	return string(reversed)
}

// Программа принимает один аргумент командной строки - текст, и выводит его в обратном порядке.