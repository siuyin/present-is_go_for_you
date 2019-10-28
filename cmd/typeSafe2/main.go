package main

import "fmt"

type myNum int

func main() {
	fmt.Println("go is type-safe (2)")
	fmt.Println("and very strongly typed")
	var i myNum = 2
	var j int = 3
	i = j
	fmt.Println(i)
}
