package sinks

import (
	"io"
	"log"
	"traffic-proxy/common"
	"traffic-proxy/configs"
)



func NewSink(config configs.SinkConfig) (*TrafficSink, error) {
	var writer io.WriteCloser

	switch config {
	case "null":
		writer = &NullSink{}

	case "console":
		writer = &ConsoleSink{}

	default:
		log.Fatalln("invalid sink config", config)
	}

	var sink = &TrafficSink{
		data: make(chan []byte),
		writer: writer,
	}

	go func() {
		for {
			select {
			case <- sink.quit:
				sink.writer.Close()

			case data := <- sink.data:
				sink.writer.Write(data)
			}
		}
	}()
	return sink, nil
}

func NewSinkFanout(channel chan *common.Packet, psinks []*TrafficSink) error {
	go func() {
		for {
			select {
			case packet := <- channel:
				for _, psink := range psinks {
					//psink.writer.Write(packet.Data)
					psink.data <- packet.Data
				}
			}
		}
	}()
	return nil
}
