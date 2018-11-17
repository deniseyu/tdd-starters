package main

import "testing"

func DoSomething() string {
	return "yo"
}

func TestSomething(t *testing.T) {
	output := DoSomething()
	expected := "hello world"

	if output != expected {
		t.Errorf("No! Got: %v. Expected: %v", output, expected)
	}
}
