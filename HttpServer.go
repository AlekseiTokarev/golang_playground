package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(hello))

	http.ListenAndServe(":8080", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	var responseText = "hello anonymous\n"

	param1 := request.URL.Query().Get("param1")
	if len(param1) > 0 {
		responseText = fmt.Sprintf("hello %#v \n", param1)
		//fmt.Printf(responseText, param1)
	}
	fmt.Fprintf(writer, responseText)
}
