package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	if result != 3 {
		t.Errorf("%s", "error")
	}
}
