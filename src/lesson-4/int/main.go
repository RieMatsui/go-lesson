package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i)

	// intとint64は別の型
	var i2 int64 = 100
	fmt.Println(i2)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}

