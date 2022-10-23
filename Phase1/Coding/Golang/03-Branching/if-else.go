package main

import "fmt"

func main() {
	if 8%2 == 0 {
		fmt.Println(" num is even ")
	} else {
		fmt.Println(" num is odd ")
	}

	if x := -5; x < 0 {
		fmt.Println(" num is negative ")
	} else if x > 0 {
		fmt.Println(" num is positive ")
	} else {
		fmt.Println(" num equal zero")
	}
}
