package main

import "fmt"

func main() {
	var numberOfStrings, numberOfQueries int
	fmt.Scan(&numberOfStrings)
	var stringsMap = make(map[string]int)
	var done = make(chan bool, 1)

	for i := 0; i < numberOfStrings; i++ {
		var inputString string
		fmt.Scan(&inputString)
		i, _ := stringsMap[inputString]
		stringsMap[inputString] = i + 1
	}

	fmt.Scan(&numberOfQueries)
	var queryResult = make([]int, numberOfQueries)

	for i := 0; i < numberOfQueries; i++ {
		var query string
		fmt.Scan(&query)

		//Simple ways is to use the logic inside the go routine within the loop itself
		//but used go routine just to get some hang of it
		go func(index int) {
			value, _ := stringsMap[query]
			queryResult[index] = queryResult[index] + value

			done <- true
		}(i)
	}

	for i := 0; i < numberOfQueries; i++ {
		<-done
	}

	for i := 0; i < numberOfQueries; i++ {
		fmt.Println(queryResult[i])
	}

}
