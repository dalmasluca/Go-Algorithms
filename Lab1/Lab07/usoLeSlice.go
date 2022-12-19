package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	sliceCompleta := new([]string)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if len(os.Args) < 3 {
		fmt.Println("Nuemro di argomenti inserito non sufficiente")
		os.Exit(1)
	}

	*sliceCompleta = append(*sliceCompleta, os.Args[1:]...)

  fmt.Println(*sliceCompleta)

  for scanner.Scan() {
		*sliceCompleta = append(*sliceCompleta, scanner.Text())
	}

	fmt.Println(*sliceCompleta)
	sort.Strings(*sliceCompleta)
	fmt.Println(*sliceCompleta)

	*sliceCompleta = append(*new([]string), (*sliceCompleta)[:len(*sliceCompleta)-1]...)

	fmt.Println(*sliceCompleta)

	*sliceCompleta = append((*sliceCompleta)[:2], (*sliceCompleta)[4:]...)

	fmt.Println(*sliceCompleta)

  nuovaSlice := []string {"a","b","c"} 

	fmt.Println(nuovaSlice)
  
	*sliceCompleta = append((*sliceCompleta)[0:1], append(nuovaSlice,(*sliceCompleta)...)... )
	
  fmt.Println(*sliceCompleta)

	var parola string

	fmt.Scan(&parola)

	*sliceCompleta = append((*sliceCompleta)[:2] , append( [] string {parola}, *sliceCompleta...)...)

	fmt.Println(*sliceCompleta)

	fmt.Scan(&parola)

	*sliceCompleta = append(*sliceCompleta, parola)

	fmt.Println(*sliceCompleta)

	*sliceCompleta = append(*sliceCompleta, []string { "","","" }... )

	fmt.Println(*sliceCompleta)
	
	*sliceCompleta = append( (*sliceCompleta)[0:3] , append(make([]string,3),(*sliceCompleta)[3:]...)... )

	fmt.Println(*sliceCompleta)
	
	secondaSlcie := new([]string)
	
	*secondaSlcie = make([]string, len(*sliceCompleta))

	copy(*secondaSlcie,*sliceCompleta)

	fmt.Println(*secondaSlcie)
	
	*secondaSlcie = append((*secondaSlcie)[:len(*secondaSlcie)-1], )

	fmt.Println(*sliceCompleta)
	fmt.Println(*secondaSlcie)

}
