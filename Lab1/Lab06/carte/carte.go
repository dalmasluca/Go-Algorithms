package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Carta struct{
  valore,seme string
}

const nSemi = 4
const nValori = 14
const nCarteMazzo = 52

var VALORI = [] string {"A","1","2","3","4","5","6","7","8","9","10","J","Q","K"}

var SEMI = [] string {"\u2665", "\u2666", "\u2663", "\u2660"}

func carta(n int) (c Carta,t bool){
  if n>=0 || n<nCarteMazzo {
    t = true
    c.valore = VALORI[n%nValori]
    c.seme = SEMI[n/nValori]
  }else{
    t = false
  }
  return
}

func estraiCarta() Carta{
  c,_ := carta(rand.Intn(nCarteMazzo))
    return c
}

func dai4Carte() (arrayCarte [4] Carta) {
  for i := range arrayCarte{
    arrayCarte[i] = estraiCarta()
  }
  return arrayCarte
}

func main(){
  rand.Seed(time.Now().Unix())
  c := estraiCarta()
  fmt.Println(c)
  ac := dai4Carte()
  fmt.Println(ac)
}
