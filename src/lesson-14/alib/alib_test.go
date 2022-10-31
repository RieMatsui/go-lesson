package alib

import "testing"

var Debug bool = false

func TestAvarage(t *testing.T) {
	if Debug {
		t.Skip("スキップします")
	}
	v := Avarage([]int{1, 2, 3 , 4, 5})
	if v != 3 {
		t.Errorf("%v vs %v", v, 3)
	}
}