package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "math"
)

func main(){
  var ris int
  var valoreX int
  esponenti := []int{}
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)

  for scanner.Scan() {
    number, _ := strconv.Atoi(scanner.Text())
    esponenti = append(esponenti, number)
  }

  valoreX = len(esponenti) - 1

  for i:=0; i<valoreX; i++ {
    ris += esponenti[i] * int(math.Pow(float64(esponenti[valoreX]), float64(i)))
  }

  fmt.Println(ris)

}
