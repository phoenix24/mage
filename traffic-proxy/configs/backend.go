package configs

import (
	"log"
	"traffic-proxy/services"
)

type Mode string

const (
	PROXY  Mode = "PROXY"
	RECORD Mode = "RECORD"
	REPLAY Mode = "REPLAY"
)

type ServiceConfig struct {
	Name string	`mapstructure:"name"`
	Mode Mode	`mapstructure:"mode"`
	Src  Route	`mapstructure:"route-src"`
	Dst  Route	`mapstructure:"route-dst"`
}

func (b ServiceConfig) Service() services.Service {
	var dst, src = b.Dst.Service(), b.Src.Service()
	if dst != src {
		log.Fatalln("source and destination must be same service type.")
	}
	switch dst {
	case "tcp":			return TCP
	case "udp":			return UDP
	case "dns":			return DNS
	case "ntp":			return NTP
	case "http":		return HTTP
	case "https":		return HTTPS
	case "websocket":	return WEBSOCKET
	case "thrift":		return THRIFT
	case "protobuf":	return PROTOBUF
	case "mysql":		return MYSQL
	case "pgsql":		return PGSQL
	case "mongodb":		return MONGODB
	case "cassandra":	return CASSANDRA
	case "redis":		return REDIS
	case "memcache":	return MEMCACHE
	case "aerospike":	return AEROSPIKE
	case "kafka":		return KAFKA
	case "pulsar":		return PULSAR
	case "zookeeper":	return ZOOKEEPER
	default:			log.Fatalln("invalid or unsupported serviced :", dst)
	}
	return TCP
}
