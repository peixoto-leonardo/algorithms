package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly_Push(t *testing.T) {
	cases := []struct {
		description    string
		singly         *Singly[int]
		value          int
		lengthExpected int
	}{
		{
			description:    "add a value to end of empty list",
			singly:         New[int](),
			value:          1,
			lengthExpected: 1,
		},
		{
			description:    "add a value to end of non-empty list",
			singly:         New(5, 2, 3),
			value:          1,
			lengthExpected: 4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			tt.singly.Push(tt.value)

			assert.Equal(t, tt.value, tt.singly.GetLast().Value)
			assert.Equal(t, tt.lengthExpected, tt.singly.length)
		})
	}
}
