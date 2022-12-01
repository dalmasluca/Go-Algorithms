package main

import ("fmt")
import ("os")
import ("sort")
func main(){
    nomi:=new([]string)
    for i:=1;i<len(os.Args);i++{
        *nomi=append(*nomi,os.Args[i])}
    sort.Strings(*nomi)
    fmt.Println(*nomi)
}
