package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// 2020-06-20 10:16:19.000000001 +0900 JST
	// 2020
	// 172
	// June
	// Saturday
	// 20
	// 10
	// time.Duration
	// 16
	// 19
	// 1
	// JST 32400
	// 1h0m0s
	// 1m0s
	// 1s
	// 1ms
	// 1µs
	// 1ns
	// 2h30m0s
	// 2020-06-20 10:16:19.000000001 +0900 JST
}

func TestChangeTime(*testing.T) {
	changeTime()
	// Output:
	// 2h30m0s
}

func TestConpairTime(*testing.T) {
	conpairTime()
	// Output:
	// 20737h6m1.000000002s
	// false
	// true
	// true
	// false
}

func TestSleep(*testing.T) {
	sleep()
	// Output:
	// 5秒停止後表示
}
