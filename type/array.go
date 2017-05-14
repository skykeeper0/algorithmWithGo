package main

import "fmt"

func main() {
	// create array
	// var x = [5]int{1, 3, 4, 5, 7}

	// create slice
	// var y = []int{1, 2, 3}

	// create slice from array x
	// var y = x[1:5]

	// create slice inside array
	// y := make([]int, 5, 10)

	// var total int
	// for _, value := range y {
	// 	fmt.Println(value)
	// }

	// for i := 0; i < len(y); i++ {
	// 	fmt.Printf("%d\n", y[i])
	// }
	// fmt.Println(total / (len(x)))
	// fmt.Println(string(5))

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}

	fmt.Println(elements)

}
