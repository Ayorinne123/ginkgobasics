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
		"email": "eve.holt@reqres.in
		"password": "pitstol"
	}`).
		Post("http://reqres.in/api/register")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Status())
	//assertions

}
