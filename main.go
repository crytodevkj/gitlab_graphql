/// [date] 2022-04-28

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
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
		method := "GET"
		client := &http.Client{}
		for {
			// Taking input from user
			first_req, _ := http.NewRequest(method, url, strings.NewReader(""))
			first_res, _ := client.Do(first_req)
			if first_res == nil {
				continue
			}
			color.Set(color.FgMagenta, color.Bold)
			fmt.Println("Your Request:")
			color.Unset()
			color.Set(color.FgYellow, color.Underline)
			fmt.Scanln(&url)
			color.Unset()
			if url == "q" {
				os.Exit(1)
			}

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
			color.Set(color.FgMagenta, color.Bold)
			fmt.Println("Your Response:")
			color.Unset()

			color.Set(color.FgWhite)
			fmt.Println(string(body))
			color.Unset()
		}
	}()

	wg.Wait()
	color.Unset()

	// Call to service layer

}
