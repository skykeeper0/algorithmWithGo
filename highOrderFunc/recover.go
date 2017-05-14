package main

import "fmt"

func main() {
	defer func() {
		function := recover()
		fmt.Println(function)
	}()
	panic("Panic")
}
