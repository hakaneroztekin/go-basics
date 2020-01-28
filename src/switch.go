package main

import (
	"fmt"
	"time"
)

func main()  {
	orderCount := 2
	switch orderCount {
	case 1:
		fmt.Println("One order")
	case 2:
		fmt.Println("Two orders")
	case 3:
		fmt.Println("Three orders")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // Multiple expressions
		fmt.Println("Today is weekend")
	default: // optional
		fmt.Println("Today is weekday")
	}

	now := time.Now()
	switch { // w/o an expression, alternative to if/else
	case now.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("It's after noon")
	}

	// a type switch
	whatAmI := func(i interface{}) {
		switch t:= i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Type is neither bool nor int, it's %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

}