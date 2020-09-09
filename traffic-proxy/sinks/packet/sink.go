package packet

import (
	"fmt"
	"github.com/google/uuid"
	"traffic-proxy/common"
	"traffic-proxy/sinks/message"
)

type PacketSink interface {
	Consume() error
}

type PacketSinkFanout struct {
	ConnPackets  map[string]*common.Packet //todo: change key to conn-string.
	MessageSinks []*message.MessageSink
	ChPackets    chan *common.Packet
	ChMessages   chan *common.MessagePair
	ChCommands   chan *common.Commands
}

func (p *PacketSinkFanout) Consume() error {
	go func() {
		for {
			select {
			case cmd := <- p.ChCommands:
				fmt.Println(cmd)

			case pair := <- p.ChMessages:
				for _, psink := range p.MessageSinks {
					(*psink).Write(pair)
				}

			case packet := <- p.ChPackets:
				var msgid = uuid.New().String()
				var connid = *packet.ConnID
				if request, v := p.ConnPackets[connid]; v {
					var pair = &common.MessagePair{
						ID:          msgid,
						RequestRaw:  request,
						ResponseRaw: packet,
						Protocol:    packet.Protocol,
					}
					delete(p.ConnPackets, connid)
					p.ChMessages <- pair

				} else {
					p.ConnPackets[connid] = packet
				}
				//todo:
				// 1. packet - hold it, by connid (for tcp, mysql etc).
				// 2. packet - if complete, parse (ideally stateless), else goto 1.
				// 3. packet - flush to sink.
			}
		}
	}()
	return nil
}
