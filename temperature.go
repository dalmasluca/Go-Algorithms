package main

import (
	"fmt"
	"sort"
)

func main(){
    var num, somma  int
    var mediana float64
    sliceNum := new([]int)
    max3 := make([]int,3)
    min3 := make([]int,3)
    fmt.Scan(&num)   
    for num != 999 {
        *sliceNum = append(*sliceNum, num)
        somma += num
        fmt.Scan(&num)
    }
    
    media := float64(somma) / float64(len(*sliceNum))

    sort.Ints(*sliceNum)
    lunghezza := len(*sliceNum)
    meta := lunghezza / 2

    if lunghezza == 0 {
        mediana = float64( (*sliceNum)[meta] + (*sliceNum)[meta + 1] ) / 2.
    }else{
        mediana = float64((*sliceNum)[meta])
    }

    if lunghezza > 3 {
        copy(max3[:],(*sliceNum)[len((*sliceNum))-3:])
        copy( min3[:], (*sliceNum)[:3])
    }else{
        copy(max3[:],(*sliceNum)[:])
        copy(min3[:],(*sliceNum)[:])

    }


    fmt.Println("media: ",media)
    fmt.Println("mediana: ",mediana)
    fmt.Println("max3: ",max3)
    fmt.Println("min3: ",min3)

}
