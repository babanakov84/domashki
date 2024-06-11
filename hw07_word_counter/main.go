package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func countWords(s string) map[string]int {
	mapWord := make(map[string]int)
	s = strings.ToLower(s)
	words := strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	})
	for _, word := range words {
		mapWord[word]++
	}
	return mapWord
}

func main() {
	fmt.Print("Введите текст: ")
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(countWords(text))
}

// hey, teacher
// We do not need no education, We do not need no thought control.
// teachers, leave those kids alone
//
