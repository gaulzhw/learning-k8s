package state

type State interface {
	Action(on bool)
	String() string
}
