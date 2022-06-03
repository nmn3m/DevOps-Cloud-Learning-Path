package main

import "fmt"

// Const should be first letter capital to be accessable from this file in those folders .
const GitLogin string = "fjewhfewgfiuea;fh"

func main() {

	// String Type
	var userName string = "Zero"
	fmt.Println(userName)
	fmt.Printf("The variable of type: %T \n", userName)

	// Boolean Type
	var isLoggedIn bool = true // can take true or false.
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable of type: %T \n", isLoggedIn)
	/*
		- Take a look at documentaion -> ("golang specific numeric type").
		- Take care of alises.
	*/

	// Int Type
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The variable of type: %T \n", smallVal)

	// Float Type
	var smallFloat float64 = 255.484897529587483288
	fmt.Println(smallFloat)
	fmt.Printf("The variable of type: %T \n", smallFloat)

	// Implicit Type -> you don't say the type of the variable "Lexer" do for you.
	var website = "golang.org"
	fmt.Println(website)
	fmt.Printf("The variable of type: %T \n", website)

	// no var style
	isNumber := false // i need to know the mean of := ?.
	fmt.Println(isNumber)

	// const style.
	fmt.Println(GitLogin)
	fmt.Printf("The type of variable is : %T \n", GitLogin)
}
