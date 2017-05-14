package main

import "fmt"

func largestNumber(arr [6]int) (max int) {
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	return
}

func main() {
	var arr = [6]int{1, 4, 5, 67, 2, 2}
	fmt.Println(largestNumber(arr))
}
