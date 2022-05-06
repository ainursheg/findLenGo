package main

import "fmt"

func main() {
	fmt.Println("Testing algorithm...")
	findSubStringLength("abcdef")
	findSubStringLength("aaaaa")
	findSubStringLength("abaccab")
}

// Find longest substring with at most two distinct characters
// Input: String
// Output: int
func findSubStringLength(someString string) int {
	if (len(someString) == 0) {
		return 0
	}
	counter := 0
	firstChar := -1
	secondChar := -1
	for i := 0; i < len(someString); i++ {

		if (firstChar == -1) {
			firstChar = i
		}

		if (secondChar == -1 && 
			someString[i] != someString[firstChar]) {
			secondChar = i
		}

		if (someString[i] != someString[firstChar] &&
			 someString[i] != someString[secondChar]) {
			counter = max(counter, i - firstChar)
			firstChar = secondChar
			secondChar = -1
			i = firstChar
		}
	}
	fmt.Printf("Most longest substring of %s is %d\n", someString, max(counter, len(someString) - firstChar))
	return max(counter, len(someString) - firstChar)
}


func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}