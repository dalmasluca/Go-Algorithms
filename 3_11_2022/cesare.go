package main
import "fmt"

func main(){
  var s1, s2 string
  var k int

  fmt.Scan(&s1)
  fmt.Scan(&k)

  for _,bt := range s1 {
    s2 += string( (((bt + rune(k) - 'a')  % ('z' - 'a' + 1))) + 'a' )
    fmt.Println(s2)
  }

  fmt.Println("chiave: ",k,"caratteri da cifrare:",s1,s2,"è il testo cifrato")
}
