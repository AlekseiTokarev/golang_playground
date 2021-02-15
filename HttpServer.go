package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(processHello))
	http.Handle("/info", http.HandlerFunc(processInfo))

	http.ListenAndServe(":8080", nil)

	fmt.Printf("application started")
}

func processHello(writer http.ResponseWriter, request *http.Request) {
	var responseText = "processHello anonymous\n"

	param1 := request.URL.Query().Get("param1")
	if len(param1) > 0 {
		responseText = fmt.Sprintf("processHello %#v \n", param1)
	}
	fmt.Fprintf(writer, responseText)
}

const increment = 1

var infoResponse = struct {
	Name    string `json:"name"`
	Counter int
}{"go-playground", 0}

func processInfo(writer http.ResponseWriter, request *http.Request) {
	infoResponse.Counter += increment

	bytes, _ := json.Marshal(infoResponse)
	fmt.Fprintf(writer, string(bytes))
}
