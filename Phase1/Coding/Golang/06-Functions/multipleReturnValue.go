package main

import "fmt"

func Vals() (int, int) {
	return 3, 7
}

func main() {

	x, y := Vals()
	fmt.Println(x)
	fmt.Println(y)

	_, z := Vals()
	fmt.Println(z)

}
