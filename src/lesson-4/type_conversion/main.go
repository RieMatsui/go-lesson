package main

import (
	"fmt"
	"strconv"
)

func main () {

	convertInt()
	convertStringIntoInt()
	convertStringIntoInt()
	convertStringIntoByte()
}

func convertInt() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n",fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)
}

func convertStringIntoInt(){

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)
}

func convertIntIntoString(){
	var i1 int = 200
    fmt.Printf("i1 = %T\n", i1)

	s1 := strconv.Itoa(i1)
	fmt.Println(s1)
	fmt.Printf("s1 = %T\n", s1)
}

func convertStringIntoByte(){
	var h string = "Hello, world!"
	b := []byte(h)
	fmt.Println(b)
	fmt.Printf("b = %T\n", b)

	h2 := string(b)
    fmt.Println(h2)
	fmt.Printf("h2 = %T\n", h2)
}