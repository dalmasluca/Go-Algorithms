package main

import (
	"bufio"
	"fmt"
	"os"
)

const LEN_ALFABETO = 26

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int){
	if contaMinu == nil {return}
	for _, runa := range s {
		if runa >= 'a' && runa <= 'z' {
			(*contaMinu)[int(runa - 'a')] ++
	
		}
	}
}

func main(){
	var contaMinu [LEN_ALFABETO]int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	for scanner.Scan() {
		aggiornaConteggi(scanner.Text(),&contaMinu)
	}

	var runa = 'a'

	for pos := range contaMinu {
		fmt.Println(string(runa + rune(pos)), contaMinu[pos])
	}
}
