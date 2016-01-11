package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "Go fundamentals"
	author := "Gaurav"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}