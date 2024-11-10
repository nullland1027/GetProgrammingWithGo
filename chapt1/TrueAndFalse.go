package chapt1

import (
	"fmt"
	"strings"
)

func isLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func trueAndFalse() {
	var longSentence = "Fuck SB yangyang, 大傻逼"
	var command = "yang"
	var isContain = strings.Contains(longSentence, command)
	fmt.Println(isContain)

	var (
		apple  = "apple"
		banana = "banana"
	)
	if apple > banana {
		fmt.Println("apple is bigger")
	} else {
		fmt.Println("banana is bigger")
	}

	for year := 2100; year < 2130; year++ {
		fmt.Printf("year %v %v\n", year, isLeapYear(year))
	}

}
