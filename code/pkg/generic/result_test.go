package generic

import (
	"testing"
)

func TestUnwrap(t *testing.T) {
	tests := struct[T any]{
		Message: T,
		Error: error,
	}{
		
	}
	(&Result[string]{Message: "test", Error: nil}).UnwrapOr("default")
	(&Result[int]{Message: 1, Error: nil}).UnwrapOr(-1)
}
