package main

import "fmt"

func main() {
	// Similar to Java Hashmaps

	// Declare
	// map[key-type]val-type
	fruitPricesMap := make(map[string]int)

	// Initialize
	fruitPricesMap["Apple"] = 5
	fruitPricesMap["Banana"] = 15
	fmt.Println("fruit prices map:", fruitPricesMap)

	// Get value from map
	priceOfApple := fruitPricesMap["Apple"]
	fmt.Println("Price of apple:", priceOfApple)

	// Get length (# of key-value pairs)
	fmt.Println("Length of the map:", len(fruitPricesMap))

	// Delete a pair from map
	fmt.Println("Delete a pair from map")
	delete(fruitPricesMap, "Banana") // delete if the key is present
	fmt.Println("Length of the map:", len(fruitPricesMap))
	fmt.Println("fruit prices map:", fruitPricesMap)

	// Check if key present
	_, isPresent := fruitPricesMap["Apple"]
	fmt.Println("is apple present:", isPresent)
	// _ is 'blank identifier', since we don't use it here we ignored it with _
	_, orangeIsPresent := fruitPricesMap["Orange"]
	fmt.Println("is orange present:", orangeIsPresent)

	// Declare & Initalize in one line
	movieCritics := map[string]int{"Lord of the Rings": 9, "Contratiempo": 10}
	fmt.Println("Movie critics:", movieCritics)
}
