package main

import "fmt"

func hasUpper(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func firstUpper(s string) int {
	for pos, c := range s {
		if c >= 'A' && c <= 'Z' {
			return pos
		}
	}
	return -1
}

func lastUpper(s string) int {
	for i := (len(s) - 1); i >= 0; i-- {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
}

func countDigitsLettersOthers(s string) (int, int, int) {
	var cC, cL, cA int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			cC++
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			cL++
		} else {
			cA++
		}
	}
	return cC, cL, cA
}

func isPalindrome(s string) bool {
	dim := len(s)
	for i := 0; i < dim/2; i++ {
		if s[i] != s[dim-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println("hasUpper :", hasUpper(s))
	fmt.Println("firstUpper : ", firstUpper(s))
	fmt.Println("lastUpper : ", lastUpper(s))
	i, e, f := countDigitsLettersOthers(s)
	fmt.Println("cDLO : ", i, e, f)
	fmt.Println("Palindrome :", isPalindrome(s))
}
