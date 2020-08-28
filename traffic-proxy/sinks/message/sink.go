package message

import (
	"fmt"
	"io"
)

type MessageSink struct {
	Data chan []byte
	//Bata chan *common.MessagePair
	Quit chan []byte
	Writer  io.WriteCloser
}

func (sk *MessageSink) Write(payload []byte) (n int, err error) {
	sk.Data <- payload
	return len(payload), nil
}

func (sk *MessageSink) Close() error {
	return nil
}


//null sink.
type MessageSinkNull struct {
}

func (sk *MessageSinkNull) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (sk *MessageSinkNull) Close() error {
	return nil
}


//console sink
type MessageSinkConsole struct {
}

func (sk *MessageSinkConsole) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func (sk *MessageSinkConsole) Close() error {
	return nil
}
