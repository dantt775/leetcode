package main

import "testing"

func TestMain(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"Test1": {
			input:    "abcabcbb",
			expected: 3,
		},
		"Test2": {
			input:    "bbbbb",
			expected: 1,
		},
		"Test4": {
			input:    "pwwkew",
			expected: 3,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})

	}
}
