// challenges/testing/begin/main_test.go
package linkedinTests

import "testing"

// write a test for letterCounter.count
func TestLetterCounter(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one letter", "12a45", 1},
		{"letters inside symbols", "a-b'c", 3},
		{"all letters", "tuk", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := letterCounter{identifier: "letter"}.count(tc.input)
			if got != tc.want {
				t.Errorf("ERROR: got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCounter(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one number", "a4cd", 1},
		{"numbers inside letters", "a1b4c", 2},
		{"all numbers", "123", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := numberCounter{designation: "number"}.count(tc.input)
			if got != tc.want {
				t.Errorf("ERROR: got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCounter(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one symbol", "≠", 1},
		{"symbols inside letters and numbers", "a¢∞§k", 3},
		{"all symbols", "¢∞•", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := symbolCounter{label: "symbol"}.count(tc.input)
			if got != tc.want {
				t.Errorf("ERROR: got %v want %v", got, tc.want)
			}
		})
	}
}
