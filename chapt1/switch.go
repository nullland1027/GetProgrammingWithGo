package chapt1

import (
	"fmt"
	"math/rand"
	"time"
)

func launch() {
	fmt.Println("Launching...")
	var count = 100
	for count > 0 {
		time.Sleep(time.Second)
		fmt.Println(count)
		if rand.Intn(10) == 0 {
			fmt.Println("Meet an error! Stop launch")
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("All good! and launched")
	}

}

func switchTest() {
	var room = "car"

	switch room {
	case "lake":
		fmt.Println("lake room")
	case "car", "motor":
		fmt.Println("car or motor room")
		fallthrough
	default:
		fmt.Println("undefined room")
	}
	launch()
}
