package main

import (
	"fmt"
)

func main() {
	var numberOfSequences int
	var numberOfQueries int

	fmt.Scan(&numberOfSequences)
	fmt.Scan(&numberOfQueries)

	var queriesList = make([][3]int, numberOfQueries)
	var seqList = make([][]int, numberOfSequences)

	for i := 0; i < numberOfQueries; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&queriesList[i][j])
		}
	}

	applyQuery(queriesList, numberOfQueries, &seqList, numberOfSequences)
}

func applyQuery(queriesList [][3]int, numberOfQueries int, seqList *[][]int, numberOfSequences int) int {
	var lastAnswer = 0
	for i := 0; i < numberOfQueries; i++ {
		if queriesList[i][0] == 1 {
			lastAnswer = applyQuery1(queriesList[i], seqList, lastAnswer, numberOfSequences)
		} else {
			lastAnswer = applyQuery2(queriesList[i], seqList, lastAnswer, numberOfSequences)
		}
	}

	return lastAnswer
}

func applyQuery1(queries [3]int, seqList *[][]int, lastAnswer int, numberOfSequences int) int {
	var index = (queries[1] ^ lastAnswer) % numberOfSequences
	(*seqList)[index] = append((*seqList)[index], queries[2])
	return lastAnswer
}

func applyQuery2(queries [3]int, seqList *[][]int, lastAnswer int, numberOfSequences int) int {
	var index = (queries[1] ^ lastAnswer) % numberOfSequences
	lastAnswer = (*seqList)[index][queries[2]%len((*seqList)[index])]
	fmt.Println(lastAnswer)
	return lastAnswer
}
