package main

import (
	"log"
	"testing"
)

func TestAddition(t *testing.T) {
	got := addition(7, 5)
	want := 12
	if got != want {
		log.Fatalf("error want %v, got %v", want, got)
	}
}
