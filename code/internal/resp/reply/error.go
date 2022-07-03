package reply

var (
	unknownErrBytes   = []byte("-Err unknown\r\n")
	syntaxErrBytes    = []byte("-Err syntax error\r\n")
	wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")
)

type UnknownErrReply struct{}

func NewUnknownErrReply() *UnknownErrReply {
	return &UnknownErrReply{}
}

func (r *UnknownErrReply) Error() string {
	return "Err unknown"
}

func (r *UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

type ArgNumErrReply struct {
	Cmd string
}

func NewArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}

func (r *ArgNumErrReply) Error() string {
	return "-ERR wrong number of arguments for '" + r.Cmd + "' command"
}

func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for '" + r.Cmd + "' command\r\n")
}

type SyntaxErrReply struct{}

func NewSyntaxErrReply() *SyntaxErrReply {
	return &SyntaxErrReply{}
}

func (r *SyntaxErrReply) Error() string {
	return "Err syntax error"
}

func (r *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

type WrongTypeErrReply struct{}

func NewWrongTypeErrReply() *WrongTypeErrReply {
	return &WrongTypeErrReply{}
}

func (r *WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

func (r *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

type ProtocolErrReply struct {
	Msg string
}

func NewProtocolErrReply(msg string) *ProtocolErrReply {
	return &ProtocolErrReply{
		Msg: msg,
	}
}

func (r *ProtocolErrReply) Error() string {
	return "ERR Protocol error: '" + r.Msg
}

func (r *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + r.Msg + "'\r\n")
}
