package reply

var (
	pongBytes           = []byte("+PONG\r\n")
	okBytes             = []byte("+OK\r\n")
	nullBulkBytes       = []byte("$-1\r\n")
	emptyMultiBulkBytes = []byte("*0\r\n")
	noBytes             = []byte("")
)

type PongReply struct{}

func NewPongReply() *PongReply {
	return &PongReply{}
}

func (r *PongReply) ToBytes() []byte {
	return pongBytes
}

type OKReply struct{}

func NewOKReply() *OKReply {
	return &OKReply{}
}

func (r *OKReply) ToBytes() []byte {
	return okBytes
}

type NullBulkReply struct{}

func NewNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

func (r *NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

type EmptyMultiBulkReply struct{}

func NewEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return &EmptyMultiBulkReply{}
}

func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

type NoReply struct{}

func NewNoReply() *NoReply {
	return &NoReply{}
}

func (r *NoReply) ToBytes() []byte {
	return noBytes
}
