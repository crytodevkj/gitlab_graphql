package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParseThenPrint(res string) {
	res_json := Response{}
	_ = json.Unmarshal([]byte(res), &res_json)

	names := []string{}
	sumOfAllForks := 0
	for i := 0; i < len(res_json.Data.Projects.Nodes); i++ {
		names = append(names, res_json.Data.Projects.Nodes[i].Name)
		sumOfAllForks += res_json.Data.Projects.Nodes[i].ForksCount
	}

	fmt.Println("Names: ", strings.Join(names, ", "))
	fmt.Println("Sum of all forks: ", sumOfAllForks)
}
