package state

import (
	"testing"
)

func TestState(t *testing.T) {
	ctx := NewContext(true)
	t.Log(ctx.GetState())

	ctx.On()
	t.Log(ctx.GetState())

	ctx.Off()
	t.Log(ctx.GetState())
}
