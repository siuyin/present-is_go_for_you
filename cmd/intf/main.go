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

	// Change 3 to -4 or 0
	q = myInt(3) // make 123 a myInt, eg woman("John") makes "John" a woman
	fmt.Println(q.Quantify())

	// change abc to four
	q = myStr("abc") // make "abc" be of type myStr
	fmt.Println(q.Quantify())
}

// 50 OMIT
