package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	resp, err := client.R().Get("http://reqres.in/api/users?pages=2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
