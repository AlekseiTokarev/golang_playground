package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	i := sum(4, 7)
	fmt.Println(i)

	all, err := sumAll([]int{5, 7, -78, 10})
	fmt.Println(all)

	all, err = sumAll([]int{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(all)

	dto := new(Dto)
	dto.Id = 4
	dto.Name = "super"

	bolB, _ := json.Marshal(dto)
	fmt.Println(string(bolB))

}

func sumAll(a []int) (n int, err error) {
	var total int
	if len(a) == 0 {
		return 0, errors.New("array shouldn't be empty")
	}
	for _, v := range a {
		if v < 0 {
			continue
		}
		total += v
	}
	return total, nil
}

func sum(a int, b int) int {
	i := a + b
	return i
}

type Dto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
