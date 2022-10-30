package main

import "fmt"

const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775808

func main() {
	fmt.Println(Pi)

	fmt.Println(URL)
	fmt.Println(SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)
}
