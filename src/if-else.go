package main

import "fmt"

func main() {
	myNumber := 7
	text := "My number is "
	if myNumber % 2 == 0 {
		text += "even"
	} else {
		text += "odd"
	}
	fmt.Println(text)

	// A statement can precede conditionals
	if myFavoriteNumber := 10; myFavoriteNumber >= 100 {
		fmt.Println("Favorite number is greater than 100")
	} else if myFavoriteNumber <= 50 {
		fmt.Println("Favorite number is greater than 5")
	} else {
		fmt.Println("Favorite number is between 50 and 100")
	}

	// There's no ternary if (with ?) in Go, so we need to use full if condition
}