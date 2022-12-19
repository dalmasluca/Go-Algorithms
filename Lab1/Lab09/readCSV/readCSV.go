package main

import (
	"fmt"
	"strings"
)

func formatDate(s string) []string {
    sl := strings.Split(s, "/")
    return append(strings.Split(sl[0],"-"), append(strings.Split(sl[1], ":"))...)
}

func main () {
    var sanno, smese, sgiorno, sora, sminuto []int
    var stemp []float64
    var sumi []float64
    
    for {
        var anno, mese, giorno, ora, minuto int
        var temperatura, umidita float64
        _, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f\n", &anno, &mese, &giorno, &ora, &minuto, &temperatura, &umidita)
        if err != nil {
            break
        }
    
        sanno = append(sanno, anno)
        smese = append(smese, mese)
        sgiorno = append(sgiorno, giorno)
        sora = append(sora, ora)
        sminuto = append(sminuto, minuto)
        stemp = append(stemp, temperatura)
        sumi = append(sumi, umidita)

    }

    tempmin, tempmax := stemp[0], stemp[0]
    var imin, imax int
    for i, temp := range stemp {
        if temp > tempmax {
            tempmax = temp
            imax = i
        }
        if temp < tempmin {
            tempmin = temp
            imin = i
        }
    }
    umimax, umimin := sumi[0], sumi[0]
    var iminu, imaxu int
    for i, um := range sumi{
        if um > umimax {
            umimax = um
            imaxu = i
        }
        if um < umimin {
            umimin = um
            iminu = i
        }
    }
    fmt.Printf("minTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", tempmin, sora[imin], sminuto[imin], sgiorno[imin], smese[imin], sanno[imin])  
    fmt.Printf("maxTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", tempmax, sora[imax], sminuto[imax], sgiorno[imax], smese[imax], sanno[imax])
    fmt.Printf("minHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", umimin, sora[iminu], sminuto[iminu], sgiorno[iminu], smese[iminu], sanno[iminu])
    fmt.Printf("maxHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", umimax, sora[imaxu], sminuto[imaxu], sgiorno[imaxu], smese[imaxu], sanno[imaxu])

}
