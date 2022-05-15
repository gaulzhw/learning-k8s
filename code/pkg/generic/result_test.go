package generic

import (
	"testing"
)

type Result[T any] struct {
	Message T
	Error   error
}

func (r *Result[T]) UnwrapOr(val T) T {
	if r.Error != nil {
		return val
	}
	return r.Message
}

func TestUnwrapString(t *testing.T) {
	tests := []struct {
		result Result[string]
		wrap   string
	}{
		{
			Result[string]{Message: "test", Error: nil},
			"default",
		},
	}

	for _, test := range tests {
		test.result.UnwrapOr(test.wrap)
	}
}

func TestUnwrapInt(t *testing.T) {
	tests := []struct {
		result Result[int]
		wrap   int
	}{
		{
			Result[int]{Message: 1, Error: nil},
			-1,
		},
	}

	for _, test := range tests {
		test.result.UnwrapOr(test.wrap)
	}
}
