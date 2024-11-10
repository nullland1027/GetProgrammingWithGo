package chapt1

import (
	"fmt"
	"math/rand"
)

func randomDates() {
	for i := 0; i < 10; i++ {
		year := rand.Intn(2000) + 1000
		month := rand.Intn(12) + 1
		var days int
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			days = 31
		case 4, 6, 9, 11:
			days = 30
		case 2:
			if isLeapYear(year) {
				days = 29
			} else {
				days = 28
			}
		}
		day := rand.Intn(days)
		fmt.Printf("%d-%02d-%02d\n", year, month, day)
	}
}
