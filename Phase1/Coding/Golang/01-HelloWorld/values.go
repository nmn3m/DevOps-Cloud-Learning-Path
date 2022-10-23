package main

import "fmt"

// we will talk about values like Booleans, integers, float, string and so on.

func main() {
	fmt.Println("What" + " are" + " you" + " doing" + "?") // string can be concatenated by using +
	fmt.Println(" 1 + 1 = ", 1+1)                          // integers
	fmt.Println(" 7.0 / 2.0 = ", 7.0/3.0)                  // float
	fmt.Println(true && false)                             // booleans and
	fmt.Println(true || false)                             // booleans or
	fmt.Println(!false)                                    // booleans not

}
