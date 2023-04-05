package linked_list

import (
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
			singly:         NewSingly[int](),
			value:          1,
			lengthExpected: 1,
		},
		{
			description:    "add a value to beginning of non-empty list",
			singly:         NewSingly(5, 2, 3),
			value:          1,
			lengthExpected: 4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			tt.singly.Unshift(tt.value)

			assert.Equal(t, tt.value, tt.singly.Head.Value)
			assert.Equal(t, tt.lengthExpected, tt.singly.length)
		})
	}
}

func BenchmarkSingly_Unshift(b *testing.B) {
	singly := NewSingly[int]()

	for i := 0; i < b.N; i++ {
		singly.Unshift(1)
	}
}

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
			singly:         NewSingly[int](),
			lengthExpected: 0,
			expected:       0,
			ok:             false,
		},
		{
			description:    "remove at beginning of non-empty list",
			singly:         NewSingly(5, 2, 3),
			lengthExpected: 2,
			expected:       5,
			ok:             true,
		},
		{
			description:    "remove at beginning of non-empty list",
			singly:         NewSingly(5),
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
	singly := NewSingly(1, 2, 3)

	for i := 0; i < b.N; i++ {
		singly.Shift()
	}
}

func TestSingly_Push(t *testing.T) {
	cases := []struct {
		description    string
		singly         *Singly[int]
		value          int
		lengthExpected int
	}{
		{
			description:    "add a value to end of empty list",
			singly:         NewSingly[int](),
			value:          1,
			lengthExpected: 1,
		},
		{
			description:    "add a value to end of non-empty list",
			singly:         NewSingly(5, 2, 3),
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
			singly:         NewSingly[int](),
			lengthExpected: 0,
			expected:       0,
			ok:             false,
		},
		{
			description:    "remove at end of non-empty list",
			singly:         NewSingly(5, 2, 3),
			lengthExpected: 2,
			expected:       3,
			ok:             true,
		},
		{
			description:    "remove at end of non-empty list",
			singly:         NewSingly(5),
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
