package main

import "fmt"

func main() {
	// create a map
	x := make(map[string]int)
	x["key"] = 0
	x["new_key"] = 0
	fmt.Println(x)

	fmt.Printf("Need to delete last key\n")

	delete(x, "key")

	_, ok := x["key"]
	name2, ok2 := x["new_key"]
	// fmt.Println(x["no key"])
	fmt.Println(ok)
	fmt.Println(name2, ok2)
}
