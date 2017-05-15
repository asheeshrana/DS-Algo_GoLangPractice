package main

import "fmt"

func main() {
	var arraySize, numberOfLeftRotations int
	fmt.Scan(&arraySize)
	fmt.Scan(&numberOfLeftRotations)

	var inputArray = make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		fmt.Scan(&inputArray[i])
	}

	for i := 0; i < arraySize; i++ {
		var indexToPrint = (i + numberOfLeftRotations) % arraySize
		if i > 0 {
			fmt.Printf(" %d", inputArray[indexToPrint])
		} else {
			fmt.Printf("%d", inputArray[indexToPrint])
		}
	}

}
