package sinks

import (
	"io"
	"log"
	"traffic-proxy/common"
	"traffic-proxy/configs"
	"traffic-proxy/sinks/message"
	"traffic-proxy/sinks/packet"
)

func NewMessageSink(config configs.SinkConfig) (*message.MessageSink, error) {
	var writer io.WriteCloser

	switch config {
	case "null":
		writer = &message.MessageSinkNull{}

	case "console":
		writer = &message.MessageSinkConsole{}

	default:
		log.Fatalln("invalid sink config", config)
	}

	var sink = &message.MessageSink{
		Data: make(chan []byte),
		Writer: writer,
	}

	go func() {
		for {
			select {
			case <- sink.Quit:
				sink.Writer.Close()

			case data := <- sink.Data:
				sink.Writer.Write(data)
			}
		}
	}()
	return sink, nil
}

func NewPacketSink(sconfigs []configs.SinkConfig, channel chan *common.Packet) (packet.Sink, error) {
	var msinks []*message.MessageSink
	for _, config := range sconfigs {
		var sink, _ = NewMessageSink(config)
		msinks = append(msinks, sink)
	}
	var psink = &packet.SinkFanout{msinks, channel}
	return psink, nil
}
