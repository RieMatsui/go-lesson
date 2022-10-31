package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl2 :=  sl
	sl2[0] = 1000
	fmt.Println(sl)

	var i int = 10
	i2 := i
	i2 = 100
    fmt.Println(i, i2)

	slice := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5, 10)
	n := copy(slice2, slice)

	fmt.Println(n, slice2)
}