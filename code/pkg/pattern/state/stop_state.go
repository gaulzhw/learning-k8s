package state

import (
	"fmt"
)

type StopState struct {
	ctx *Context
}

func (c *StopState) Action(on bool) {
	fmt.Println("in stop state")
	if !on {
		return
	}
	c.ctx.setState(&StartState{ctx: c.ctx})
}

func (c *StopState) String() string {
	return "Stop State"
}
