package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println(" weekend")
	default:
		fmt.Println("weekday")
	}

	// using switch to comapre types insted of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("intger")
		default:
			fmt.Printf("i don't know this type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(23)
	whatAmI("hello")
}
