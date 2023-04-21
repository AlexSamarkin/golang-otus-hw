package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	const HelloString string = "Hello, OTUS!"
	fmt.Printf("%s", stringutil.Reverse(HelloString))
}
