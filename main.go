package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://gitlab.com/api/graphql"
	method := "POST"
	payload := strings.NewReader("{\"query\":\"query last_projects($n: Int = 5) {\\r\\n  projects(last:$n) {\\r\\n    nodes {\\r\\n      name\\r\\n      description\\r\\n      forksCount\\r\\n    }\\r\\n  }\\r\\n}\",\"variables\":{}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	ParseThenPrint(string(body))
}
