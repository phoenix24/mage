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
	ID       string
	ConnID   *string //todo: temp?
	Protocol *string //todo: enum?
	Data     []byte
	Time     time.Time
	SrcIp    string
	SrcPort  string
	DstIp    string
	DstPort  string
}

type MessagePair struct {
	ID             string //todo: must be a uuid (type)?
	RequestRaw     *Packet
	RequestParsed  *Packet
	ResponseRaw    *Packet
	ResponseParsed *Packet
	Protocol       *string //MessageProtocol
}

type Commands struct {
}
