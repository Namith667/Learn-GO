package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type University struct {
	Country string `json:"country"`
	Name    string `json:"name"`
}

func main() {
	resp, err := http.Get("http://universities.hipolabs.com/search?country=United+States")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	//fmt.Println("raw resp body", string(body))

	var response []University
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	count := len(response)
	num := 5
	for i, university := range response {
		if i < num {
			fmt.Println(i, university.Country, university.Name)

		}
	}

	for _, data := range response {
		fmt.Printf("University: %s, Country: %s\n", data.Name, data.Country)
	}

	fmt.Println("Got ", count, " Numbers")

}
