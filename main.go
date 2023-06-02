package main

import "fmt"

type a struct {
	b int
	c int
}

func main() {
	type aa struct {
		b int
		c int
	}
	var aaa interface{} = aa{
		b: 1,
		c: 1,
	}

	switch v := aaa.(type) {
	case a:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}
}
