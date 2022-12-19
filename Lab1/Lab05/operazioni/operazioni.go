package main

import "fmt"

func operazioni(n1 int, n2 int) (int, int, int) {
	return n1 + n2, n1 * n2, n1 - n2
}

func main(){
  var n1, n2 int
  fmt.Scan(&n1,&n2)
  i,e,u := operazioni(n1,n2)
  fmt.Println(i,e,u)
}
