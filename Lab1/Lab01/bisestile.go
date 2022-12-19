package main
import "fmt"

func main(){
	var g,m,a int
	fmt.Print("gg mm aaaa: ");fmt.Scan(&g,&m,&a);
	if(a<0){
		fmt.Print("bisestile non applicabile");	
	}else if(a<1582){
		if(a%4==0){
			fmt.Println("bisestile");
		}else{
			fmt.Println("non bisestile");
		}
	}else{
		if(a%400==0){
			fmt.Println("bisestile");
		}else if(a%4==0 && a%100!=0){
			fmt.Println("bisestile");
		}else{
			fmt.Println("non bisestile");
		}
	}
}
