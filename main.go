package main

import (
	"fmt"
	"strings"
)

/*
Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

Examples

"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"
*/

func main() {
	str := ""

	if string(len(str)) == "" {
		return
	}

	word := strings.Split(str, "")

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	st := strings.Join(word, "")
	fmt.Println(st)
}
