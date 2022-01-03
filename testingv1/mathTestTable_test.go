package testingv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
a				b
negative		negative
negative		positive
positive		negative
positive		positive

*/

func TestAdd_WithTestTable(t *testing.T) {
	testCase := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative & negative",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{
			name:     "negative & positive",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{
			name:     "positve & positive",
			a:        1,
			b:        1,
			expected: 2,
		},

		{
			name:     "positve & negative",
			a:        1,
			b:        -1,
			expected: 0,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			c := add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})

	}
}
