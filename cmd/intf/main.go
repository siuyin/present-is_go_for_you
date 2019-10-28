package main

import "fmt"

// 10 OMIT
type Quantifier interface {
	Quantify() int
}

// 20 OMIT
type myStr string

func (s myStr) Quantify() int {
	return len(s)
}

// 30 OMIT
type myInt int

func (i myInt) Quantify() int {
	switch {
	case i == 0:
		return 0
	case i < 0:
		return -1
	case i > 0:
		return 1
	}
	return 0
}

// 40 OMIT
func main() {
	fmt.Println("go has implicitly satisfied interfaces")

	var q Quantifier

	q = myInt(3)
	fmt.Println(q.Quantify())

	q = myStr("abc")
	fmt.Println(q.Quantify())
}

// 50 OMIT
