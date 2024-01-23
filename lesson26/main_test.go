package main

import "testing"

var tests []struct {
	name string
	divident float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data": 100.0, 10.0, 10.0, false}
}

func TestDevide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}

}

func TestBadDevide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Dig not get an error when we should not have")
	}

}
