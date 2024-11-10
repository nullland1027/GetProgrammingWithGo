package chapt1

import (
	"fmt"
	"math/rand"
)

func floatNumberFormat() {
	fmt.Printf("%f\n", 2.71)
	fmt.Printf("%10.2f\n", 2.7)
}

func tenDime() {
	var dollar float64
	for i := 0; i < 11; i++ {
		dollar += 0.1
	}
	fmt.Println(dollar)
}

func randomMoney(n int) int {
	var cents int
	switch n {
	case 0:
		cents = 5
	case 1:
		cents = 10
	case 2:
		cents = 25
	}
	return cents
}

func piggySaveMoney() {
	var currentCent int
	for true {
		oneTimeMoney := randomMoney(rand.Intn(3))

		currentCent += oneTimeMoney
		fmt.Printf("You save %.2f $ this time. Now in total %.2f $.\n",
			float64(oneTimeMoney)/100,
			float64(currentCent)/100)
		if currentCent >= 2000 {
			break
		}
	}
}
