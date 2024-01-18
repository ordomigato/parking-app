package utils

import (
	"testing"
)

func TestTrimStringEnd(t *testing.T) {
	var tests = []struct {
		inp1 string
		inp2 int
		ans  string
	}{
		{"ABCDEF", 2, "ABCD"},
		{"1234567", 4, "123"},
		{"AB", 2, ""},
		{"A", 2, ""},
	}

	for _, tt := range tests {
		result := TrimStringEnd(tt.inp1, tt.inp2)
		if result != tt.ans {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, tt.ans)
		}
	}
}
