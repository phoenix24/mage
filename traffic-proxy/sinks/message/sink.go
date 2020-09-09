package message

import (
	"fmt"
	"traffic-proxy/common"
)

type MessageSink interface {
	Write(msg *common.MessagePair) (n int, err error)
	Close() error
}


//null sink.
type MessageSinkNull struct {
}

func (sk *MessageSinkNull) Write(msg *common.MessagePair) (n int, err error) {
	return 0, nil
}

func (sk *MessageSinkNull) Close() error {
	return nil
}


//console sink
type MessageSinkConsole struct {
}

func (sk *MessageSinkConsole) Write(msg *common.MessagePair) (n int, err error) {
	var request = msg.RequestRaw.Data
	var response = msg.ResponseRaw.Data
	fmt.Printf(">>%s\n%s\n", msg.ID, string(request))
	fmt.Printf("<<%s\n%s\n", msg.ID, string(response))
	return len(request) + len(response), nil
}

func (sk *MessageSinkConsole) Close() error {
	return nil
}
