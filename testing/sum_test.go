package testing

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"one", []int{1}, 1},
		{"two", []int{1, -2}, -1},
		{"three", []int{1, 4, 0}, 5},
	}

	// range over the tests and run them as subtests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// pass the input as variadic using the ... stuff
			got := Sum(tc.input...)
			if got != tc.want {
				t.Errorf("ERROR: got %v want %v", got, tc.want)
			}
		})
	}
}
