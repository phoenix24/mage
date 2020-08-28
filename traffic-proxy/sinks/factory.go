package sinks

import (
	"io"
	"log"
	"traffic-proxy/common"
	"traffic-proxy/configs"
	"traffic-proxy/sinks/message"
	"traffic-proxy/sinks/packet"
)

var sinkssofar = make(map[configs.SinkConfig]*message.MessageSink)

func NewMessageSink(config configs.SinkConfig) (*message.MessageSink, error) {
	var writer io.WriteCloser

	if sink, ok := sinkssofar[config]; ok {
		return sink, nil
	}

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
	sinkssofar[config] = sink

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

func NewPacketSink(sconfigs []configs.SinkConfig, channel chan *common.Packet) (packet.PacketSink, error) {
	//todo: move to utils.
	var sinksmap = make(map[configs.SinkConfig]*message.MessageSink)
	for _, config := range sconfigs {
		var sink, _ = NewMessageSink(config)
		sinksmap[config] = sink
	}

	var sinksarr []*message.MessageSink
	for _, sink := range sinksmap {
		sinksarr = append(sinksarr, sink)
	}
	var psink = &packet.PacketSinkFanout{sinksarr, channel}
	return psink, nil
}
