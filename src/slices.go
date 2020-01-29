package main

import "fmt"

func main() {
	// Slices are more common in Go, than arrays
	mySlice := make([]string, 3) // make slice of strings of length 3
	fmt.Println("slice:", mySlice)

	// We initialize slice just like arrays
	mySlice[0] = "Apple"
	mySlice[1] = "Banana"
	mySlice[2] = "Orange"
	fmt.Println("slice:", mySlice)
	fmt.Println("slice length:", len(mySlice))

	// append values (one of the features of slices)
	mySlice = append(mySlice, "Lemon")
	fmt.Println("slice:", mySlice)
	fmt.Println("slice length:", len(mySlice))

}
