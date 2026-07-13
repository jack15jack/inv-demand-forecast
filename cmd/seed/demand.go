package main

import (
	"math/rand"
	"time"
)

func baseDemand(itemNumber string, date time.Time) int {

	switch itemNumber {

	case "BOLT-001":
		return 20 + rand.Intn(10)

	case "LIGHT-001":
		return holidayDemand(date)

	case "FILTER-001":
		return 8 + rand.Intn(5)

	default:
		return 5
	}
}

func holidayDemand(date time.Time) int {

	month := date.Month()

	switch month {

	case time.November:
		return 35 + rand.Intn(15)

	case time.December:
		return 60 + rand.Intn(25)

	case time.January:
		return 5 + rand.Intn(3)

	default:
		return 2 + rand.Intn(2)
	}
}
