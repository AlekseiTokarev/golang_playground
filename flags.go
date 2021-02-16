package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(flag.Arg(0))

	arg1 := flag.String("lastname", "noone", "as last name")
	arg2 := flag.Bool("exist", true, "as bool")
	flag.Parse()
	fmt.Printf("last name: %v\n", *arg1)
	fmt.Printf("bool? %#v\n", *arg2)
}
