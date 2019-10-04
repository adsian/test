package main

import "fmt"

func main1() {
	fmt.Println("start...")
	request := NewRequest()
	cli := CLI{request}
	cli.Run()
}
