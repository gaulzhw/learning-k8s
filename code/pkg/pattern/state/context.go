package state

type Context struct {
	state State
}

func NewContext(on bool) *Context {
	ctx := &Context{}
	if on {
		ctx.setState(&StartState{ctx: ctx})
	} else {
		ctx.setState(&StopState{ctx: ctx})
	}
	return ctx
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) GetState() State {
	return c.state
}

func (c *Context) On() {
	c.state.Action(true)
}

func (c *Context) Off() {
	c.state.Action(false)
}
