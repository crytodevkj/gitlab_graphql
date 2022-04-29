/// [date] 2022-04-28

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	/// [note] 2 routines are running, one for service layer, other one for http server
	wg.Add(2)

	go func() { // Call to service layer
		Service()
		wg.Done()
	}()

	go func() { // Call to http server
		var url string = "http://localhost:8080/graphql?query={user(num:\"1\"){names}}"
		method := "GET"
		client := &http.Client{}
		for {
			// This block is optional, for user-friendly logging
			first_req, _ := http.NewRequest(method, url, strings.NewReader(""))
			first_res, _ := client.Do(first_req)
			if first_res == nil {
				continue
			}
			// Input prompt
			color.Set(color.FgMagenta, color.Bold)
			fmt.Println("")
			fmt.Println("Your Request:")
			color.Unset()
			// Input your request
			color.Set(color.FgYellow, color.Underline)
			fmt.Scanln(&url)
			color.Unset()
			// if request is 'q', app will be exited
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
			// output prompt
			fmt.Println("")
			color.Set(color.FgMagenta, color.Bold)
			fmt.Println("Your Response:")
			color.Unset()
			// json-formatted output
			color.Set(color.FgWhite)
			var prettyJSON bytes.Buffer
			er := json.Indent(&prettyJSON, body, "", "  ")
			if er != nil {
				log.Println("JSON parse error: ", er)
				return
			}
			log.Println(string(prettyJSON.Bytes()))
			color.Unset()
		}
	}()

	wg.Wait()
	color.Unset()
}
