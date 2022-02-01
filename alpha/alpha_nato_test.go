package main

import "testing"

func TestAlpha(t *testing.T) {
	for _, c := range testCases {
		if covert(c.input) == c.expected {
			t.Log("Pass for Alpha")
		} else {
			t.Fatal("Fail for Alpha")
		}
	}
}
