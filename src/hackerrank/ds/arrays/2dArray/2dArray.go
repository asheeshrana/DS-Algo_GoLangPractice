package main

import "fmt"

func main() {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	var inputMatrix [6][6]int

	var maxSum = MinInt
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&inputMatrix[i][j])
		}
	}

	for j := 0; j < 4; j++ {

		for k := 0; k < 4; k++ {
			var sum = inputMatrix[j][k] + inputMatrix[j][k+1] + inputMatrix[j][k+2] +
				inputMatrix[j+1][k+1] +
				inputMatrix[j+2][k] + inputMatrix[j+2][k+1] + inputMatrix[j+2][k+2]
			if sum >= maxSum {
				maxSum = sum
			}
		}
	}

	fmt.Print(maxSum)

}
