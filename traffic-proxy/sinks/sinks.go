package sinks

import "fmt"

type NullSink struct {
	data chan []byte
}

func (sk *NullSink) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type ConsoleSink struct {
	data chan []byte
}

func (sk *ConsoleSink) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
