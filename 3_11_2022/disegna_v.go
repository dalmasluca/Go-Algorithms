package main
import "fmt"

func main(){
  var dimensione int
  fmt.Print("dimensione \\:\t")
  fmt.Scan(&dimensione)  
  for i:=0; i<dimensione; i++{
    for j:=0; j<dimensione*2; j++{
      if i-j==0 {
        fmt.Print("*")
      }else{
        fmt.Print(" ")
      }
    }
  fmt.Println()
  }
}
