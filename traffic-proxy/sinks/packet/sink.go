package packet

import (
	"traffic-proxy/common"
	"traffic-proxy/sinks/message"
)

type Sink interface {
	Consume() error
}

type SinkFanout struct {
	MessageSinks  []*message.MessageSink
	PacketChannel chan *common.Packet
}

func (p *SinkFanout) Consume() error {
	go func() {
		for {
			select {
			case packet := <- p.PacketChannel:
				for _, psink := range p.MessageSinks {
					psink.Data <- packet.Data
				}
			}
		}
	}()
	return nil
}
