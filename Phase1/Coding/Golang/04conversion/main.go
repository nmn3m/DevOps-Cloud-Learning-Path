package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5 ")

	reader := bufio.NewReader(os.Stdin) // we use reader as a refrence to use it you should use method for it .
	input, _ := reader.ReadString('\n') // single comma I don't know why for now.
	fmt.Println("Thanks for rating, ", input)

	/*
		- strconv package converts from string to what you want.
		- strings package uses to handle the space "\n" in the end of what bufio reading from os by using trim method.
	*/

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// this block of code will handle the err if it found.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}

}
