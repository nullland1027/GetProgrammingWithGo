package chapt1

import (
	"fmt"
	"math/big"
)

func tryBig() {
	speed := new(big.Int)
	speed.SetString("100000000000000000000000000000000000000000000000", 10)
	fmt.Println("speed is ", speed)
	time := new(big.Int)
	time.SetInt64(20000000000000000)
	fmt.Println("time is ", time)
	distance := new(big.Int)
	distance.Mul(speed, time)
	fmt.Println("distance is ", distance)
}

func cains() {
	const distance = 236e15 // km
	const c = 299792.458    // km/s
	const secPerYear = 365 * 24 * 3600

	fmt.Println("light year", distance/c/secPerYear)
}
