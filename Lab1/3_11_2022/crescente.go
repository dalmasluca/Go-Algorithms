package main
import "fmt"

func main(){
  var parola string
  fmt.Scan(&parola)
  dim := len(parola)

  fmt.Print(string(parola[0]))
  for i:=1; i<dim; i++{
    if parola[i]<parola[i-1]{
      fmt.Print("-")
    }
    fmt.Print(string(parola[i]))
  }
}
