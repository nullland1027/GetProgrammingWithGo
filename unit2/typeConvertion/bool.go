package typeConvertion

import "fmt"

func Stob(s string) bool {
	var res bool
	switch s {
	case "true", "yes", "1":
		res = true
	case "false", "no", "0":
		res = false
	default:
		fmt.Println("ERROR")

	}

	return res
}
