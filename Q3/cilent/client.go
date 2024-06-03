package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type postBody struct {
	Beef string
}

func getBeefString() string {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(responseData)
}

func main() {

	postUrl := "http://localhost:8080/beef/summary"

	beefString := getBeefString()

	body := postBody{
		Beef: beefString,
	}

	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(bodyBytes)

	resp, err := http.Post(postUrl, "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		log.Println("Error response. Status Code: ", resp.StatusCode)
	}

	log.Println("Response:", string(responseBody))
}
