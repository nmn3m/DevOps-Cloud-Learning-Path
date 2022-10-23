package main

import (
	"fmt"
	"math"
)

// we will talk about constants like character, string, boolean and  numeric

const a string = "constants"

func main() {
	fmt.Println("a is ", a)

	const n = 5000
	const d = 3e20 / n
	fmt.Println("d is ", d)
	// numeric constant has no type till you give one by explicit conversion
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
