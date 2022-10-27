package main

import "fmt"

func main() {
	// definae an array
	var x [4]int
	fmt.Println(" emp: ", x)

	x[3] = 5
	fmt.Println("index 3 in x is : ", x[3])

	y := [3]int{1, 3, 2}
	fmt.Println("array Y : ", y)

	var twoD [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
