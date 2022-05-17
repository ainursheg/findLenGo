package main

import "fmt"

func main() {
	fmt.Println("Testing algorithm...")
	findSubStringLength("abcdef")
	findSubStringLength("aaaaa")
	findSubStringLength("abaccab")
	
	
	LongestSubstring("abcdef")
	LongestSubstring("aaaaa")
	LongestSubstring("abaccab")
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

func LongestSubstring(s string) int {
	left := 0
	result := 0
	counter := make(map[byte]byte)
	
	for right := range s {
		counter[s[right]] += 1
		
		for len(counter) > 2 {
			leftChar := s[left]
			counter[leftChar] -= 1
			if counter[leftChar] == 0 {
				delete(counter, leftChar)
			}
			left += 1
		}
		result = max(result, right - left + 1)
	}
	fmt.Println("result", result)
	return result
}



func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
