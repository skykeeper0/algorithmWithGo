package main

import "fmt"

func countRepeatedLetter(str string) int {
	var count int
	var repeatedLetterStore = make(map[string]int)

	for i := 0; i < len(str); i++ {
		value, exist := repeatedLetterStore[string(str[i])]
		if exist {
			repeatedLetterStore[string(str[i])] = value + 1
		} else {
			repeatedLetterStore[string(str[i])] = 1
		}
	}

	for _, value := range repeatedLetterStore {
		if value > 1 {
			count = count + 1
		}
	}

	return count
}

func main() {
	fmt.Println(countRepeatedLetter("mayoma"))
}
