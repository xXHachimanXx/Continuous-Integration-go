package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(20, 20)

	if total != 40 {
		t.Errorf("Invalid sum result -> Result: %d Expected: %d", total, 40)
	}
}
