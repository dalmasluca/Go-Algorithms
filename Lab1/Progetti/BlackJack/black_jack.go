package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Carta struct {
	valore, seme string
}

const nSemi = 4
const nValori = 13
const nCarteMazzo = 52

var VALORI = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

var SEMI = []string{"\u2665", "\u2666", "\u2663", "\u2660"}

func carta(n int) (c Carta, t bool) {
	if n >= 0 || n < nCarteMazzo {
		t = true
		c.valore = VALORI[n%nValori]
		c.seme = SEMI[n/nValori]
	} else {
		t = false
	}
	return
}

func estraiCarta() Carta {
	c, _ := carta(rand.Intn(nCarteMazzo))
	return c
}

func dai4Carte() (arrayCarte [4]Carta) {
	for i := range arrayCarte {
		arrayCarte[i] = estraiCarta()
	}
	return arrayCarte
}

func getValoreBJ(carta *Carta) int {
	switch (*carta).valore {
	case "J", "Q", "K":
		return 10
	case "A":
		return 11
	default:
		i, _ := strconv.Atoi((*carta).valore)
		return i
	}
}

func mazzoPoker() [nCarteMazzo]Carta {
	var mazzo [nCarteMazzo]Carta
	for i, seme := range SEMI {
		for j, valore := range VALORI {
			mazzo[i*nValori+j] = Carta{valore, seme}
		}
	}
	return mazzo
}

func mescola(mazzo *[nCarteMazzo]Carta) {
	for i := range *mazzo {
		r := rand.Intn(nCarteMazzo)
		(*mazzo)[i], (*mazzo)[r] = (*mazzo)[r], (*mazzo)[i]
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	c := estraiCarta()
	fmt.Println(c)
	ac := dai4Carte()
	fmt.Println(ac)
	fmt.Print("[ ")
	for i := range ac {
		fmt.Print(getValoreBJ(&ac[i]), " ")
	}
	fmt.Print("]\n")
	mazzo := mazzoPoker()
	fmt.Println(mazzo)
	mescola(&mazzo)
	fmt.Print(mazzo)
}
