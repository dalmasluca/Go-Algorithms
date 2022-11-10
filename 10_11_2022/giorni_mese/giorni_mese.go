package main

import (
	"fmt"
	"strconv"
)

func giorniInMese(mese int) int {
	switch mese {
	case 2:
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func main() {
	var date string
	var datetemp string
	var ncount int

	fmt.Print("Insert a date with this format gg-mm-anno : ")
	fmt.Scan(&date)

	for _, n := range date {
		if n != '-' {
			datetemp += string(n)
		}
		ncount++
		i, _ := strconv.Atoi(datetemp)
		datetemp = ""
    if ncount == 1 {
			fmt.Println("il mese", i, "ha", giorniInMese(i), "giorni")
		}
	}
}
