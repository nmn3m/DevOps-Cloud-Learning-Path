// for is only golang looping construct

package main

import "fmt"

func main() {

	x := 1
	for x < 3 {
		fmt.Println(x)
		x++
	}

	for y := 4; y <= 7; y++ {
		fmt.Println(y)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for z := 0; z <= 10; z++ {
		if z%2 == 0 {
			continue
		}
		fmt.Println(z)
	}
}
