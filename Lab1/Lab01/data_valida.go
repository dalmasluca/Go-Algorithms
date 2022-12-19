package main
import "fmt"

func main(){
  var g,m int;
  fmt.Print("gg mm: ");fmt.Scan(&g,&m);
  if(((m==1 || m==3 || m==5 || m==7 || m==8 || m==10 || m==12) && (g>0 && g<32)) || ((m==4 || m==6 || m==9 || m==11) && (g>0 && g<31)) || (m==2 && (g>0 && g<28)) ){
      fmt.Println("valida");
  }else{
    fmt.Println("non valida");
  }
}
