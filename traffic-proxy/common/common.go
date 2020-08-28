package common

import "time"

type MessageProtocol uint8

const (
	HTTP MessageProtocol = 0
	MYSQL
	PGSQL
	REDIS
)

type Packet struct {
	Data  []byte
	Time  time.Time
	//srcIp string
	//srcId SrcID
}

type MessagePair struct {
	ID       string //todo: must be a uuid (type)?
	Request  Packet
	Response Packet
	Protocol MessageProtocol
}
