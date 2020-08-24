package sinks

import (
	"fmt"
	"io"
	"traffic-proxy/configs"
)

type TrafficSink struct {
	writer  io.Writer
	channel chan []byte
}

func (sk *TrafficSink) Write(p []byte) (n int, err error) {
	sk.channel <- p
	return len(p), nil
}

type NullSink struct {
}

func (sk *NullSink) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type ConsoleSink struct {
}

func (sk *ConsoleSink) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func NewSink(config configs.Sink) (*TrafficSink, error) {
	var writer io.Writer
	switch config {
	case "null":
		writer = &NullSink{}
	case "console":
		writer = &ConsoleSink{}
	}

	var sink = &TrafficSink{
		writer:  writer,
		channel: make(chan []byte),
	}
	go func() {
		for {
			select {
			case data := <- sink.channel:
				sink.writer.Write(data)
			}
		}
	}()
	return sink, nil
}

func NewSinks(config configs.ServerConfig) ([]TrafficSink, error) {
	return nil, nil
}
