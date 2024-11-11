package utils

func CorrectMod(x, y int) int {
	if x >= 0 {
		return x % y
	} else {
		return (x%y + y) % y
	}
}
