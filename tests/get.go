package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {

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
	fmt.Println(resp.Status())
	fmt.Println(resp.Time())
	fmt.Println(resp.StatusCode())
	fmt.Println(resp.String())
	fmt.Println(resp.Header())

}
