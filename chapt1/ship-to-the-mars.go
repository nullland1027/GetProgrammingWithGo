package chapt1

import (
	"fmt"
	"math/rand"
)

func getSpaceShipCompany(n int) string {
	var company string
	switch n {
	case 0:
		company = "Space Adventures"
	case 1:
		company = "SpaceX"
	case 2:
		company = "Virgin Galactic"
	}
	return company
}

func oneTuneOrTwo(n int) string {
	var ticketType string
	switch n {
	case 0:
		ticketType = "单程"
	case 1:
		ticketType = "往返"
	}
	return ticketType
}

func shipToMars(n int) {
	const distance = 62100000 // km
	const secPerDay = 24 * 3600
	var (
		col1 = "太空航行公司"
		col2 = "飞行天数"
		col3 = "飞行类型"
		col4 = "价格（百万美元）"
	)
	fmt.Printf("%-10v %10v %10v %10v\n", col1, col2, col3, col4)
	for i := 0; i < n; i++ {
		company := getSpaceShipCompany(rand.Intn(3))
		speed := rand.Intn(15) + 16          // km/s
		days := distance / secPerDay / speed // days
		ticketType := oneTuneOrTwo(rand.Intn(2))
		oneTurnPrice := 66 - speed // million dollar
		var totalPrice int
		if ticketType == "往返" {
			totalPrice = 2 * oneTurnPrice
		} else if ticketType == "单程" {
			totalPrice = oneTurnPrice
		}
		fmt.Printf("%-20v %10v %10v %10v\n", company, days, ticketType, totalPrice)
	}

}
