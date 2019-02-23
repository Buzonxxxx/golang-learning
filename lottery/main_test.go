package main

import (
	"testing"
)

func TestGenerateNumbers(t *testing.T) {
	n := generateNumbers(49)
	if len(n) != 49 {
		t.Errorf("Expected number length of 49, bot got %v", len(n))
	}
}
