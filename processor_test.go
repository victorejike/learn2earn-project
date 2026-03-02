package main

import "testing"

func TestCapitalizeMultiple(t *testing.T) {
	input := "it was the age of foolishness (cap, 6)"
	expected := "It Was The Age Of Foolishness"
	result := Process(input)
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestPunctuationFormatting(t *testing.T) {
	input := "Punctuation tests are ... kinda boring ,what do you think ?"
	expected := "Punctuation tests are... kinda boring, what do you think?"
	result := Process(input)
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

// Additional example tests
func TestHexBin(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin)"
	expected := "Simply add 66 and 2"
	result := Process(input)
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestArticles(t *testing.T) {
	input := "There is a apple and a banana"
	expected := "There is an apple and a banana"
	result := Process(input)
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}
