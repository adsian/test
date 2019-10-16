package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := "hello"
	b := "world"

	c := bytes.Compare([]byte(a), []byte(b))
	fmt.Println(c)
}
