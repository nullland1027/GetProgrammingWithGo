package typeConvertion

import (
	"GetProgrammingWithGo/unit2/stringAndRune"
	"fmt"
	"math"
	"strings"
)

func Decipher(cipherText, keyword string) string {
	repeatTimesFloat := float64(len(cipherText)) / float64(len(keyword))
	repeatTimesInt := int(math.Ceil(repeatTimesFloat))
	totalKeyword := strings.Repeat(keyword, repeatTimesInt)[:len(cipherText)]
	fmt.Println(len(cipherText), cipherText)
	fmt.Println(len(totalKeyword), totalKeyword)

	for i := 0; i < len(cipherText); i++ {
		decodedChar := stringAndRune.CaesarDecodeChar(rune(cipherText[i]), totalKeyword[i])
		fmt.Printf("%c ", decodedChar)
	}

	return ""
}
