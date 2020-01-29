package main

import "fmt"

func main() {
	var myArray [5]int // all array int's are 0 by default
	fmt.Println("values:", myArray)

	myArray[3] = 100
	fmt.Println("values:", myArray)
	fmt.Println("array length is", len(myArray))

	myArray2 := [5]int{0, 5, 10, 15, 20} // Declare & Initialize an array in one line
	fmt.Println("values of array:", myArray2)

	var twoDArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDArray[i][j] = i + j
		}
	}
	fmt.Println(twoDArray)
}
