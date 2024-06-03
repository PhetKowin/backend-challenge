package main

import (
	"encoding/json"
	"fmt"
	"net/http" // import package net-http เข้ามา
	"strings"
)

type request struct {
	Beef string
}

type responseData struct {
	Beef map[string]int
}

func beefSummary(w http.ResponseWriter, r *http.Request) {
	var requestBody request
	var responseData responseData
	myMap := make(map[string]int)
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		panic(err)
	}
	beef := strings.Fields(requestBody.Beef)
	replacer := strings.NewReplacer(",", "", ".", "")
	for i := 0; i < len(beef); i++ {
		beef[i] = replacer.Replace(beef[i])
		beef[i] = strings.ToLower(beef[i])
		value, exists := myMap[beef[i]]
		if exists {
			myMap[beef[i]] = value + 1
		} else {
			myMap[beef[i]] = 1
		}
	}
	responseData.Beef = myMap
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func handleRequest() {
	http.HandleFunc("/beef/summary", beefSummary)
	fmt.Println("Port 8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
