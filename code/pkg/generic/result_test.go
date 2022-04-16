package generic

import (
	"testing"
)

func TestUnwrap(t *testing.T) {
	(&Result[string]{Message: "test", Error: nil}).UnwrapOr("default")
	(&Result[int]{Message: 1, Error: nil}).UnwrapOr(-1)
}
