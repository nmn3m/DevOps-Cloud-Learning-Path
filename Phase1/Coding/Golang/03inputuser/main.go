package main

import (
	"bufio" // read about it
	"fmt"
	"os" // read about it
)

// If you want to search for any package, go to -> pkg.go.dev

func main() {
	welcome := "Welcome to input user"
	fmt.Println(welcome)
	/*
		- Reader will read from standard input for my os so we need to use bufio .
		- Notice pacakge in small letter and method in it use capial at first letter and camel case for other.
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza: ")

	/*
		- comma ok syntax || error err syntax
		- in golang we don't have try catch like python so we can't handle error so we use error ok syntax.
		- we use _ as err . if things go wrong you should do another thing .
		- we use "\n" to make him read until he use enter.
	*/

	input, _ := reader.ReadString('\n')

	fmt.Print("the rating of our pizza ,", input)
	fmt.Printf("the type of input is %T \n", input)
}
