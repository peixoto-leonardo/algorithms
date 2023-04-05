package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly_Pop(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description    string
		singly         *Singly[int]
		expected       int
		ok             bool
		lengthExpected int
	}{
		{
			description:    "remove at end of empty list",
			singly:         New[int](),
			lengthExpected: 0,
			expected:       0,
			ok:             false,
		},
		{
			description:    "remove at end of non-empty list",
			singly:         New(5, 2, 3),
			lengthExpected: 2,
			expected:       3,
			ok:             true,
		},
		{
			description:    "remove at end of non-empty list",
			singly:         New(5),
			lengthExpected: 0,
			expected:       5,
			ok:             true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			assert := assert.New(t)
			value, ok := tt.singly.Pop()

			assert.Equal(tt.ok, ok)
			assert.Equal(tt.expected, value)
			assert.Equal(tt.lengthExpected, tt.singly.length)
		})
	}
}
