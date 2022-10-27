package main

import "fmt"

func main() {
	// define map
	var x = make(map[string]int)
	x["k1"] = 22
	x["k2"] = 23

	fmt.Println("map: ", x)
}
