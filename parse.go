package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func getNames(res string) string {
	res_json := Response{}
	_ = json.Unmarshal([]byte(res), &res_json)

	names := []string{}
	for i := 0; i < len(res_json.Data.Projects.Nodes); i++ {
		names = append(names, res_json.Data.Projects.Nodes[i].Name)
	}

	return strings.Join(names, ", ")
}

func getSumOfAllForks(res string) string {
	res_json := Response{}
	_ = json.Unmarshal([]byte(res), &res_json)

	sumOfAllForks := 0
	for i := 0; i < len(res_json.Data.Projects.Nodes); i++ {
		sumOfAllForks += res_json.Data.Projects.Nodes[i].ForksCount
	}

	return fmt.Sprint(sumOfAllForks)
}
