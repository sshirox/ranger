package ranger

import (
	"reflect"
	"testing"
)

func TestSequence_Filter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		filterFn func(int) bool
		expected []int
	}{
		{
			name:     "even numbers",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			filterFn: func(n int) bool { return n%2 == 0 },
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "none pass",
			input:    []int{1, 3, 5, 7},
			filterFn: func(n int) bool { return n%2 == 0 },
			expected: []int{},
		},
		{
			name:     "all pass",
			input:    []int{2, 4, 6},
			filterFn: func(n int) bool { return n%2 == 0 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "empty input",
			input:    []int{},
			filterFn: func(n int) bool { return n > 0 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := NewSequence(tt.input)
			got := seq.Filter(tt.filterFn)

			if !reflect.DeepEqual(got.Values(), tt.expected) {
				t.Errorf("Filter() = %v, want %v", got.Values(), tt.expected)
			}
		})
	}
}
