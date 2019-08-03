package cmdmsg

type MessageHead struct {
	Sender string														`json:"sender"`
	Receiver string													`json:"receiver"`
	MsgType string													`json:"msg_type"`
	Sequence int32													`json:"sequence"`
}

