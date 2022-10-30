package main

import "fmt"


/**
 * xとyを加算するメソッド
 * @param x
 * @param y
 * @return
 */
func Plus(x int, y int) int {
	return x + y
}

/**
 * xとyを除算し qを返す
 * xとyのあまり rを返す
 * @param x
 * @param y
 * @return
 */
func Div(x , y int) (int, int) {
	q := x / y
	r := x % y
    return q, r
}

func Double(price int) (result int) {
	result = price * 2
    return
}

func NoReturn() {
	fmt.Println("No Return")
	return
}

func main() {

	i := Plus(1, 2)
    fmt.Println(i)


	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i2_2, _ := Div(9, 5)
	fmt.Println(i2_2)

	i4 := Double(1000)
	fmt.Println(i4)

	NoReturn()
}

