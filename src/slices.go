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

	// Some of the features of Slices
	// Append
	mySlice = append(mySlice, "Lemon")
	fmt.Println("slice:", mySlice)
	fmt.Println("slice length:", len(mySlice))

	// Copy
	copySlice := make([]string, 4)
	copy(copySlice, mySlice) // copy to copySlice from mySlice
	fmt.Println("Copy slice:", copySlice)

	// Slice
	fmt.Println("slice[1:3]:", mySlice[1:3])
	fmt.Println("slice[:3]:", mySlice[:3])
	fmt.Println("slice[2:]:", mySlice[2:])

	// Declare & Initialize in a single line
	mySlice2 := []string{"Book", "Mac", "Programming"}
	fmt.Println("mySlice2:", mySlice2)

	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoDSlice[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("twoDSlice:", twoDSlice)
}
