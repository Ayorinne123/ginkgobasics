package main

import (
	"fmt"
	"reflect"

	"encoding/json"

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
	var data map[string]interface{}
	json.Unmarshal(resp.Body(), &data)

	data1 := data["data"]
	data2 := data1.([]interface{})
	fmt.Println(data2)
	fmt.Println(len(data2))

	for i := 0; i < len(data2); i++ {
		temp := data2[i].(map[string]interface{})
		fmt.Println(temp["id"])
		fmt.Println(reflect.TypeOf(temp["id"]))
		if temp["id"] == 10 {
			fmt.Println(temp["id"])
			fmt.Println(temp["first_name"])
			break

		}
	}
}
