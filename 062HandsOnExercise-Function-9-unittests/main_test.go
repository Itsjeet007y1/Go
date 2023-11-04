package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := paradise("Bail")
	want := "My Idea of paradise is: Bail"
	if got != want {
		log.Fatalf("error want %v, got %v", want, got)
	}
}
