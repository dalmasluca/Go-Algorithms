package main
import "fmt"


func main(){
  var s string
  fmt.Println("Inserire mano di gioco : ")
  fmt.Scan(&s)
  if(len(s)==3){
    fmt.Println("mano",s,":",punti(s),"punti")
  }
}

func punti(s string) int{
  var punti int
  dim := len(s)
  for i:=0; i<dim; i++{
    switch s[i]{
      case 'A':
        punti += 11
      case '3':
        punti += 10
      case 'K':
        punti += 4
      case 'Q':
        punti += 3
      case 'J':
        punti += 2
      case '7','6','5','4','2':
        
      default:
        return -1
    }
  }
  return punti
}
