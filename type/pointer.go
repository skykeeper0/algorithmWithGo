package main

import "fmt"

// * is a swapper between regular address and value
func zero(xPtr *int) {
	*xPtr = 12
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(&x) // x is 0
}
