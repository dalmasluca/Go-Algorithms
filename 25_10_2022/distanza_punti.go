package main
import "fmt"
import "math"

func main(){
  var x1,y1,x2,y2 float64;
  fmt.Print("x e y del primo punto:\t");
  fmt.Scan(&x1,&y1);
  fmt.Print("x e y del secondo punto:\t");
  fmt.Scan(&x2,&y2);
  fmt.Print("La disatanza tra i due punti è: ");
  fmt.Println(math.Sqrt(math.Pow((x2-x1),2)+math.Pow((y2-y1),2)));
}
