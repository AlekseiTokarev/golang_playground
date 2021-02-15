package main

import (
	"encoding/json"
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(sum(
		4,
		7,
	))
	fmt.Println(sumAll([]int{5, 7, -78, 10}))
	//
	dto := new(Dto)
	dto.Id = 4
	dto.Name = "super"

	bolB, _ := json.Marshal(dto)
	fmt.Println(string(bolB))

}

func sumAll(a []int) int {
	var total int
	for _, v := range a {
		if v < 0 {
			continue
		}
		total += v
	}
	return total
}

func sum(a int, b int) int {
	i := a + b
	return i
}

type Dto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
