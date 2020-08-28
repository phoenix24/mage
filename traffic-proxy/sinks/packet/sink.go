package packet

import (
	"traffic-proxy/common"
	"traffic-proxy/sinks/message"
)

type PacketSink interface {
	Consume() error
}

type PacketSinkFanout struct {
	MessageSinks   []*message.MessageSink
	ChannelPackets chan *common.Packet
}

func (p *PacketSinkFanout) Consume() error {
	go func() {
		for {
			select {
			case packet := <- p.ChannelPackets:
				//todo:
				// 1. packet - hold it, by connid (later by tcp).
				// 2. packet - if complete, parse
				// 3. packet - inspect request or response
				// 4. packet - parser (ideally stateless)
				// 5. packet - if request/ response complete, then flush to sink.
				for _, psink := range p.MessageSinks {
					psink.Data <- packet.Data
				}
			}
		}
	}()
	return nil
}
