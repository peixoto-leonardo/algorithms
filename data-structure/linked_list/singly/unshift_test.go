package singly

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly_Unshift(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description    string
		singly         *Singly[int]
		value          int
		lengthExpected int
	}{
		{
			description:    "add a value to beginning of empty list",
			singly:         New[int](),
			value:          1,
			lengthExpected: 1,
		},
		{
			description:    "add a value to beginning of non-empty list",
			singly:         New(5, 2, 3),
			value:          1,
			lengthExpected: 4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			tt.singly.Unshift(tt.value)

			fmt.Println(tt.singly)

			assert.Equal(t, tt.value, tt.singly.Head.Value)
			assert.Equal(t, tt.lengthExpected, tt.singly.length)
		})
	}
}

func BenchmarkSingly_Unshift(b *testing.B) {
	singly := New[int]()

	for i := 0; i < b.N; i++ {
		singly.Unshift(1)
	}
}
