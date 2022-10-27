// range iterates over elements in variety of datastructure

package main

import "fmt"

func main() {

	x := []int{6, 7, 8}
	sum := 0
	for _, i := range x {
		sum += i
	}
	fmt.Println(sum)

	for i, v := range x {
		if v == 8 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "bnana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range of string iterate over unicode
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
