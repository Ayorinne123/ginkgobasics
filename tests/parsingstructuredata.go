package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {

	type MyJsonName struct {
		Data []struct {
			Avatar    string `json:"avatar"`
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			ID        int64  `json:"id"`
			LastName  string `json:"last_name"`
		} `json:"data"`
		Page    int64 `json:"page"`
		PerPage int64 `json:"per_page"`
		Support struct {
			Text string `json:"text"`
			URL  string `json:"url"`
		} `json:"support"`
		Total      int64 `json:"total"`
		TotalPages int64 `json:"total_pages"`
	}
	client := resty.New()
	resp, err := client.R().
		SetQueryString("page=2").
		SetHeader("content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Get("http://reqres.in/api/users")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	var p1 MyJsonName
	json.Unmarshal(resp.Body(), &p1)

	fmt.Println(p1)
	fmt.Println(p1.TotalPages)
	fmt.Println(p1.Data[1].Avatar)

	for i := 0; i < len(p1.Data); i++ {
		fmt.Println(p1.Data[i].ID, p1.Data[i].Email)
	}
}
