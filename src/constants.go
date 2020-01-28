package main

import (
	"fmt"
	"math"
)

const macbook string = "Macbook"

func main() {
	fmt.Println(macbook)

	fmt.Println(3e2) // 3 * (10 ^ 2)

	// A numeric constant has no type until it’s given one
	const number1 = 3e20
	const number2 = 500000000
	const result = number1 / number2
	fmt.Println(int64(result)) // 600000000000

	fmt.Println(math.Sin(result))

}
