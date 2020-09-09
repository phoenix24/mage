package sinks

import (
	"log"
	"traffic-proxy/common"
	"traffic-proxy/configs"
	"traffic-proxy/sinks/message"
	"traffic-proxy/sinks/packet"
)

var sinkssofar = make(map[configs.SinkConfig]*message.MessageSink)

func NewMessageSink(config configs.SinkConfig) (*message.MessageSink, error) {
	if sink, ok := sinkssofar[config]; ok {
		return sink, nil
	}

	var sink message.MessageSink
	switch config {
	case "null":
		sink = &message.MessageSinkNull{}

	case "http":
		sink = &message.MessageSinkNull{}

	case "console":
		sink = &message.MessageSinkConsole{}

	default:
		log.Fatalln("invalid sink config", config)
	}
	sinkssofar[config] = &sink
	return &sink, nil
}

func NewPacketSink(sconfigs []configs.SinkConfig, chPackets chan *common.Packet, chCommand chan *common.Commands) (packet.PacketSink, error) {
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
	var psink = &packet.PacketSinkFanout{
		make(map[string]*common.Packet),
		sinksarr,
		chPackets,
		make(chan *common.MessagePair, 1),
		chCommand,
	}
	return psink, nil
}
