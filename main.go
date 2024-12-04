package main

import "fmt"

func main() {
	e := NewEngien()

	e.Set("foo", "bar")

	val, err := e.Get("foo")
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
