package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly_Shift(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description    string
		singly         *Singly[int]
		expected       int
		ok             bool
		lengthExpected int
	}{
		{
			description:    "remove at beginning of empty list",
			singly:         New[int](),
			lengthExpected: 0,
			expected:       0,
			ok:             false,
		},
		{
			description:    "remove at beginning of non-empty list",
			singly:         New(5, 2, 3),
			lengthExpected: 2,
			expected:       5,
			ok:             true,
		},
		{
			description:    "remove at beginning of non-empty list",
			singly:         New(5),
			lengthExpected: 0,
			expected:       5,
			ok:             true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			assert := assert.New(t)
			value, ok := tt.singly.Shift()

			assert.Equal(tt.ok, ok)
			assert.Equal(tt.expected, value)
			assert.Equal(tt.lengthExpected, tt.singly.length)
		})
	}
}

func BenchmarkSingly_Shift(b *testing.B) {
	singly := New(1, 2, 3)

	for i := 0; i < b.N; i++ {
		singly.Shift()
	}
}
