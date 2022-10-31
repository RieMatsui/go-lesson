package main

import (
	"testing"
)

var Debug bool = false

func TestIsOne(t *testing.T){
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}

func TestIsNotOne_0(t *testing.T){
    i := 0
	if Debug {
		t.Skip("スキップさせる")
	}
	v := IsOne(i)

	if v {
		t.Errorf("%v != %v", i, 1)
	}
}

func TestIsNotOne_2(t *testing.T){
    i := 2
	if Debug {
		t.Skip("スキップさせる")
	}
	v := IsOne(i)

	if v {
		t.Errorf("%v != %v", i, 1)
	}
}

func ExampleMain(){
	main()
	// Output:
	// true
	// false
	// 3
}