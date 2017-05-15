package main

import (
	"fmt"
)

func main() {
	var arraySize int
	fmt.Scan(&arraySize)

	var inputArray = make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		fmt.Scan(&inputArray[i])
	}

	for j := arraySize; j > 0; j-- {
		if j > 1 {
			fmt.Printf("%d ", inputArray[j-1])
		} else {
			fmt.Printf("%d", inputArray[j-1])
		}

	}

}
