package main

import (
	"fmt"
	"time"
)

// 10 OMIT
func plus() {
	for {
		fmt.Print("+")
		time.Sleep(time.Second)
	}
}

// 20 OMIT
func dot() {
	for {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
}

// 30 OMIT
func main() {
	fmt.Println("go concurrency")
	// FIXME:
	// plus()
	// dot()
	fmt.Println("About to execute select{}")
	select {} // wait until Ctrl-C is pressed
}

// 40 OMIT
