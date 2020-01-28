package main

import "fmt"

func main() {

	// for is Go’s only looping construct
	i := 1
	for i <= 3 { // similar to the Java 'while' loop
		fmt.Println(i)
		i += 1
	}

	for j := 50; j < 55; j++ {
		fmt.Println(j)
	}

	for { // for w/o a condition loops repeatedly
		fmt.Println("Hello")
		break // until we 'break' or 'return'
	}

	for k := 80; k <= 90; k++ {
		if k % 2 == 0 {
			continue; // continue with the next iteration
		}
		fmt.Println(k) // print if the # is odd
	}

}