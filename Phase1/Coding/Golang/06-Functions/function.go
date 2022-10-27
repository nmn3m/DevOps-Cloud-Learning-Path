package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

func PlusPlus(x, y, z int) int {
	return x + y + z
}

func main() {

	fmt.Println(Plus(1, 3))
	fmt.Println(PlusPlus(1, 2, 3))
}
