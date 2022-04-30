package main

import (
	"bufio"
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
	"github.com/haulerkonj/gqlgen_service/service"
)

// This Reader is important to aviod EOF in stdin.
var reader = bufio.NewReader(os.Stdin)
var url = "http://localhost:8081/query"
var method = "POST"
var payload_str = ""
var payload = strings.NewReader("{\"query\":\"mutation Inits {  init {    num    names    sumOfAllForks  }}\",\"variables\":{}}")
var client = &http.Client{}

/// This block is optional for good logging.
func pre_loop() {
	for {
		var pre_payload = strings.NewReader("{\"query\":\"mutation Inits {  init {    num    names    sumOfAllForks  }}\",\"variables\":{}}")
		req, err := http.NewRequest(method, url, pre_payload)
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if body != nil {
			break
		}
	}
}

/// [NOTE] USEFUL COMMANDS
///
/// 1. Init first and only once
///      {\"query\":\"mutation Inits {  init {    num    names    sumOfAllForks  }}\",\"variables\":{}}
/// 2. Query according to num
///      {"query":"mutation fetches($n: String!) {  fetch(input: $n) {    num    names    sumOfAllForks  }}","variables":{"n":"3"}}
/// 3. Query all records in raw text
///      {"query":"query allRec {  records }","variables":{}}
/// 4. Query one record that have been queried before
///      {"query":"query getRec($n: String!) {  record(input: $n) {    num    names    sumOfAllForks  }}", "variables":{"n":"3"}}
/// 5. Append record
///      {"query":"mutation Appends($n: NewRecord!) {  append (input: $n)}", "variables":{  "n": {    "num": "3",    "names": "a,b,c",    "sumOfAllForks": "11"  }}}

func my_loop() {
	for {
		// Input prompt
		color.Set(color.FgMagenta, color.Bold)
		fmt.Println("\nYour Request:")
		color.Unset()
		// Input your request
		color.Set(color.FgYellow, color.Underline)
		payload_str, _ = reader.ReadString(byte('\n'))
		color.Unset()

		// ----------------------------- I N P U T -------------------------------------------
		// if request is 'exit', app will be exited
		if payload_str == "exit\r\n" {
			os.Exit(1)
		}
		payload = strings.NewReader(payload_str)
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// ----------------------------- O U T P U T -------------------------------------------
		// Output prompt
		color.Set(color.FgMagenta, color.Bold)
		fmt.Println("Query Response:")
		color.Unset()
		// json-formatted output
		color.Set(color.FgWhite)
		var prettyJSON bytes.Buffer
		er := json.Indent(&prettyJSON, body, "", "  ")
		if er != nil {
			log.Println("JSON parse error: ", er)
			return
		}
		// print the result
		log.Println(string(prettyJSON.Bytes()))
		color.Unset()
	}
}

func main() {
	var wg sync.WaitGroup
	/// [note] 2 routines are running, one for service layer, other one for graphql endpoint
	wg.Add(2)

	go func() { // Call to service layer
		service.Server()
		wg.Done()
	}()

	go func() { // Call to graphql endpoint
		pre_loop()
		my_loop()
		wg.Done()
	}()

	wg.Wait()
	color.Unset()
}
