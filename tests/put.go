package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {

	client := resty.New()
	resp, err := client.R().
		SetHeader("content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(`{
			"name": "morpheus",
    		"job": "zion resident"
	}`).
		Put("https://reqres.in/api/users/2")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Status())
	//assertions

}
