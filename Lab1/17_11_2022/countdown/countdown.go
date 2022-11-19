package main

import (
	"fmt"
	"time"
)

type Clock struct {
	hour, min, sec int
}

func countdown(clock *Clock) bool {
	if (*clock).sec == 0 {
		return updateMin(clock)
	} else {
		(*clock).sec--
		return true
	}
}

func updateMin(clock *Clock) bool {
	if (*clock).min == 0 {
		return updateHour(clock)
	} else {
		(*clock).min--
		(*clock).sec = 59
		return true
	}
}

func updateHour(clock *Clock) bool {
	if (*clock).hour == 0 {
		return false
	} else {
		(*clock).hour--
		(*clock).min = 59
		return true
	}
}

func main() {
	var h, m, s int
	fmt.Print("time (hh mm ss): ")
	fmt.Scan(&h, &m, &s)
	clock := &Clock{h, m, s}

	for countdown(clock) {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("{", (*clock).hour, (*clock).min, (*clock).sec, "}")
	}

}
