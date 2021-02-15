package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://dev.att.hylatest.com/info")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body is: ", string(body))

	res := Response{}

	if err := json.Unmarshal(body, &res); err != nil {
		panic(err)
	}
	fmt.Println(res.App.Name)
}

type Response struct {
	App App `json:"app"`
}

type App struct {
	Name string `json:"name"`
}
