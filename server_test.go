package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var client = &http.Client{}
var url = "http://localhost:8081/query"
var method = "POST"

func helper(t *testing.T, method string, url string, payload io.Reader) {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		t.Fatal(err2)
	}
	res_json := Response{}
	_ = json.Unmarshal([]byte(body), &res_json)
	if len(res_json.Errors) > 0 {
		t.Fatal("error:", res_json.Errors)
	}
	t.Log(string(body))
}

func TestInit(t *testing.T) {

	payload := strings.NewReader("{\"query\":\"mutation Inits {  init {    num    names    sumOfAllForks  }}\",\"variables\":{}}")

	helper(t, method, url, payload)
}

func TestInitAppend(t *testing.T) {

	TestInit(t)
	payload := strings.NewReader("{\"query\":\"mutation Appends($n: NewRecord!) {  append (input: $n)}\",\"variables\":{  \"n\": {    \"num\": \"3\",    \"names\": \"a,b,c\",    \"sumOfAllForks\": \"11\"  }}}")

	helper(t, method, url, payload)
}

func TestInitAppendFetch(t *testing.T) {

	TestInitAppend(t)
	payload := strings.NewReader("{\"query\":\"mutation fetches($n: String!) {  fetch(input: $n) {    num    names    sumOfAllForks  }}\",\"variables\":{\"n\":\"3\"}}}")

	helper(t, method, url, payload)
}

func TestInitAppendGetrec(t *testing.T) {

	TestInitAppend(t)
	payload := strings.NewReader("{\"query\":\"query getRec($n: String!) {  record(input: $n) {    num    names    sumOfAllForks  }}\",\"variables\":{\"n\":\"3\"}}}")

	helper(t, method, url, payload)
}
