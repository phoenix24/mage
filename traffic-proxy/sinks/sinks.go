package sinks

import (
	"fmt"
	"io"
)

type TrafficSink struct {
	data chan []byte
	quit chan []byte
	writer  io.WriteCloser
}

func (sk *TrafficSink) Write(payload []byte) (n int, err error) {
	sk.data <- payload
	return len(payload), nil
}

func (sk *TrafficSink) Close() error {
	return nil
}


//null sink.
type NullSink struct {
}

func (sk *NullSink) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (sk *NullSink) Close() error {
	return nil
}


//console sink
type ConsoleSink struct {
}

func (sk *ConsoleSink) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func (sk *ConsoleSink) Close() error {
	return nil
}
