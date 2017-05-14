// Take in a regular array and return array with uniq element

package main

import "fmt"

func uniqArray(array []int) []int {
	var store = map[int]int{}
	var resultArr = []int{}

	for _, value := range array {
		freq, exist := store[value]
		if exist {
			store[value] = freq + 1
		} else {
			store[value] = 1
			resultArr = append(resultArr, value)
		}
	}

	return resultArr
}

func main() {
	var arr = []int{1, 3, 4, 7, 2, 2, 10, 7, 10}
	fmt.Println(uniqArray(arr))
}
