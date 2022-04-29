/// [date] 2022-04-28

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var Url string = "https://gitlab.com/api/graphql"
var Req string = `
		query last_projects($n: Int = DISPLAY_NUM) {
			projects(last:$n) {
				nodes {
					name
					description
					forksCount
				}
			}
		}
	`

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		Service()
		wg.Done()
	}()

	go func() {
		var url string = "http://localhost:8080/graphql?query={user(num:\"1\"){names}}"
		for {
			// Taking input from user
			fmt.Println("Input your request:")
			fmt.Scanln(&url)
			if url == "q" {
				break
			}

			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, strings.NewReader(""))

			if err != nil {
				fmt.Println(err)
				return
			}
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
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Println(string(body))
			fmt.Println("")
		}
		wg.Done()
	}()

	wg.Wait()
	// Call to service layer

}
