package main

import "testing"

func TestGenerateTabulation(t *testing.T) {
	tabulation := generateTabulation(2)

	if tabulation != "  " {
		t.Errorf("expected 2-space tabulation, got _%v_", tabulation)
	}
}

func TestGetLongestSpliceNameLength(t *testing.T) {
	length := getLongestSpliceNameLength()

	if length != 10 {
		t.Errorf("expected 10, got _%v_", length)
	}
}
