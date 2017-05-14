package control

func Location(city string) (region, continent string) {
	switch city {
	case "NewYork", "Main":
		region = "NewYork"
		continent = "North America"
	case "LA", "Irvine":
		region = "California"
		continent = "North America"
	default:
		region = "unknown"
		continent = "unknown"
	}
	return
}

// func main() {
// 	region, continent := location("LA")
// 	fmt.Printf("Im living in %s, %s \n", region, continent)
// }
