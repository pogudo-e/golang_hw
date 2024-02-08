package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.Trim(text, "\n")
	text = strings.ReplaceAll(text, " ", "")
	var loverText string
	for _, c := range text {
		loverText += string(unicode.ToLower(c))
	}
	charCunts := make(map[rune]int)
	for _, char := range loverText {
		charCunts[char]++
	}
	var x float64
	for char, count := range charCunts {
		x = 100 / float64(len(text)) * float64(count) / 100
		fmt.Printf("%c - %d %0.1f\n", char, count, x)
	}

}
