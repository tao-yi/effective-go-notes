package main

import (
	"fmt"
)

var cache map[string]string = map[string]string{
	"a": "1",
	"b": "2",
	"c": "3",
}

func main() {
	for _, c := range "helloworld" {
		fmt.Printf("%q, %[1]T\n", c)
	}

	for key := range cache {
		fmt.Println(key)
	}

	for _, value := range cache {
		fmt.Println(value)
	}

	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}
