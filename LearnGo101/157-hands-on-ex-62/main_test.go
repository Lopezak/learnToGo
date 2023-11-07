package main

import (
	"log"
	"testing"
)

func TestArea(t *testing.T) {
	s := Square{
		length: 3,
		width:  4,
	}
	got := s.Area()
	want := float64(12)
	if got != want {
		log.Fatalf("Got %v but wanted %v", got, want)
	}
}
