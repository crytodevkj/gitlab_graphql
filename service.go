/// [date] 2022-04-28

package main

import (
	"fmt"
	"strings"
)

func Service(url string, req string) {
	// call to repository layer
	res := Api(url, req)
	// parse into array of node
	nodes := GetNodes(res)

	names := []string{}
	var sumOfAllForks float64 = 0
	// loop through array to get names | sum of all forks
	for i := 0; i < len(nodes); i++ {
		names = append(names, GetName(nodes[i]))
		sumOfAllForks += GetForksCount(nodes[i])
	}

	// display results
	fmt.Println("Names: ", strings.Join(names, ", "))
	fmt.Println("Sum of all forks: ", sumOfAllForks)
}
