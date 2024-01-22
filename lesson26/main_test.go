package main

import "testing"

func TestDevide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}

}

func TestBadDevide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}

}