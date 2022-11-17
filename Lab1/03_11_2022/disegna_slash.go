package main
import "fmt"

func main(){
  var dimensione int
  fmt.Print("dimensione \\:\t")
  fmt.Scan(&dimensione)  
  for i:=0; i<dimensione; i++{
    for j:=0; j<=i; j++{
      if i==j {
        fmt.Print("*")
      }else{
        fmt.Print(" ")
      }
    }
  fmt.Println()
  }
}
