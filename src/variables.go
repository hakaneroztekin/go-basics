package main

import "fmt"

func main() {
	var goLooksAwesome = "Go looks awesome!"
	fmt.Println(goLooksAwesome)

	// Go infers the type of initialized variables
	var val1 int = 10
	var val2 = 10
	fmt.Println(val1)
	fmt.Println(val2)

	var myVal1, myVal2 = 20, 25
	fmt.Println(myVal1)
	fmt.Println(myVal2)

	var myBool = true
	fmt.Println(myBool)

	// Variables declared but not initialized will have default values
	var defaultInt int
	fmt.Println(defaultInt) // 0
	var defaultBool bool
	fmt.Println(defaultBool) // false
	var defaultString string
	fmt.Println(defaultString) // empty string
	var defFloat64 float64
	fmt.Println(defFloat64) // 0

	var apple1 = "apple 1"
	apple2 := "apple 2" // same with the above, shorthand for declaration & initialization
	fmt.Println(apple1)
	fmt.Println(apple2)
}
