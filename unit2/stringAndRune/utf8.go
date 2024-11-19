package stringAndRune

import (
	"fmt"
	"unicode/utf8"
)

func ChineseInformation() {
	sentence := "你好。我爱学习中文。"
	fmt.Println(len(sentence), "byte.")

	fmt.Println(utf8.RuneCountInString(sentence), "runes.") // Chinese chars count

	c, size := utf8.DecodeRuneInString(sentence) // return the first char and its bytes
	fmt.Printf("First rune: %c %v", c, size)

	for i := 0; i < len(sentence); i++ {
		// fmt.Println(string(sentence[i])) // will not print the chinese char
		fmt.Printf("%c", sentence[i])
	}

	for idx, char := range sentence {
		fmt.Printf("%v - %c\n", idx, char)
	}
}
