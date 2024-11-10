package chapt1

import (
	"fmt"
	"math/rand"
)

func guessNumber() {
	const targetNum = 81
	for true {
		var num = rand.Intn(100)
		if num == targetNum {
			fmt.Println(num, "You are at its target number.")
			break
		} else if num > targetNum {
			fmt.Println(num, "You are bigger than the target number.")
		} else {
			fmt.Println(num, "You are at smaller than target number.")
		}
	}

}
