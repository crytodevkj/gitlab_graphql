package main

import (
	"fmt"
	"strings"
)

func Service(url string, req string) {
	res := Api(url, req)
	nodes := GetNodes(res)
	// fmt.Println(nodes)

	names := []string{}
	var sumOfAllForks float64 = 0
	for i := 0; i < len(nodes); i++ {
		names = append(names, GetName(nodes[i]))
		sumOfAllForks += GetForksCount(nodes[i])
	}

	fmt.Println("Names: ", strings.Join(names, ", "))
	fmt.Println("Sum of all forks: ", sumOfAllForks)
}
