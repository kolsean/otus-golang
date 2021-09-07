package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const helloMsg = "Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(helloMsg))
}
