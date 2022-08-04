package state

import (
	"fmt"
)

type StartState struct {
	ctx *Context
}

func (c *StartState) Action(on bool) {
	fmt.Println("in start state")
	if on {
		return
	}
	c.ctx.setState(&StopState{ctx: c.ctx})
}

func (c *StartState) String() string {
	return "Start State"
}
