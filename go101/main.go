package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords(",olleH !dlrow") == "Hello, world!")
}

func SpinWords(s string) string {
	sb := new(strings.Builder)
	dlim := ""
	for _, s := range strings.Split(s, " ") {
		sb.WriteString(dlim)
		dlim = " "
		if len(s) > 4 {
			sb.WriteString(Reverse(s))
		} else {
			sb.WriteString(s)
		}
	}
	return sb.String()
}
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*
Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed:
*/
