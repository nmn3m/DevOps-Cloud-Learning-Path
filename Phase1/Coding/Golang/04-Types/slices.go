package main

import "fmt"

func main() {
	// define slice
	x := make([]string, 3)
	fmt.Println("emp: ", x)

	x[0] = "Hey"
	x[1] = "Hi"
	x[2] = "hello"

	fmt.Println("slice X is : ", x)

	fmt.Println("len", len(x))

	y := make([]string, len(x))
	copy(y, x)
	fmt.Println(y)
	y = append(y, "holla")
	fmt.Println("y", y)
}
