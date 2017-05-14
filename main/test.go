package main

import (
	"fmt"

	"github.com/algorithmGo/control"
)

func main() {
	region, continent := control.Location("LA")
	fmt.Printf("Im living in %s, %s \n", region, continent)
}
