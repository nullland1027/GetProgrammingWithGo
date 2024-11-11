package stringAndRune

import (
	"GetProgrammingWithGo/utils"
	"fmt"
)

func GetCharByIndex(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
}

func caesarEncodeChar(char rune, step byte) rune {
	var encodedChar rune
	if char >= 'A' && char <= 'Z' {
		encodedChar = 'A' + ((char-'A')+rune(step))%26
	} else if char >= 'a' && char <= 'z' {
		encodedChar = 'a' + ((char-'a')+rune(step))%26
	} else {
		encodedChar = char
	}
	return encodedChar
}

func caesarDecodeChar(encodedChar rune, step byte) rune {
	var decodedChar rune
	if encodedChar >= 'A' && encodedChar <= 'Z' {
		decodedChar = 'A' + rune(utils.CorrectMod(int((encodedChar-'A')-rune(step)), 26))
	} else if encodedChar >= 'a' && encodedChar <= 'z' {
		decodedChar = 'a' + rune(utils.CorrectMod(int((encodedChar-'a')-rune(step)), 26))
	} else {
		decodedChar = encodedChar
	}
	return decodedChar
}

func CaesarEncodeString(s string, step byte) string {
	encoded := ""
	for i := 0; i < len(s); i++ {
		encoded += string(caesarEncodeChar(rune(s[i]), step))
	}
	return encoded
}

func CaesarDecodeString(encoded string, step byte) string {
	decodedString := ""
	for i := 0; i < len(encoded); i++ {
		decodedChar := caesarDecodeChar(rune(encoded[i]), step)
		decodedString += string(decodedChar)
	}
	return decodedString
}
