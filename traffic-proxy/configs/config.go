package configs

type Mode string
type SinkConfig string

const (
	PROXY  Mode = "PROXY"
	RECORD Mode = "RECORD"
	REPLAY Mode = "REPLAY"
)

type ServerConfig struct {
	Name   string  `mapstructure:"name"`
	Mode   Mode    `mapstructure:"mode"`
	Source Address `mapstructure:"source"`
	Remote Address `mapstructure:"remote"`
}

type HealthConfig struct {
	path string `mapstructure:"path"`
}

type DirectoryConfig struct {
	path string `mapstructure:"directory"`
}

type ProxyConfig struct {
	Name    string         `mapstructure:"name"`
	Port    int            `mapstructure:"port"`
	Sinks   []SinkConfig   `mapstructure:"sinks"`
	Servers []ServerConfig `mapstructure:"servers"`
}

//func (b ServerConfig) Protocol() services.Protocol {
//	var dst, src = b.backend.Protocol(), b.Source.Protocol()
//	if dst != src {
//		log.Fatalln("source and destination must be same service type.")
//	}
//	switch dst {
//	case "tcp":			return TCP
//	case "udp":			return UDP
//	case "dns":			return DNS
//	case "ntp":			return NTP
//	case "http":		return HTTP
//	case "https":		return HTTPS
//	case "websocket":	return WEBSOCKET
//	case "thrift":		return THRIFT
//	case "protobuf":	return PROTOBUF
//	case "mysql":		return MYSQL
//	case "pgsql":		return PGSQL
//	case "mongodb":		return MONGODB
//	case "cassandra":	return CASSANDRA
//	case "redis":		return REDIS
//	case "memcache":	return MEMCACHE
//	case "aerospike":	return AEROSPIKE
//	case "kafka":		return KAFKA
//	case "pulsar":		return PULSAR
//	case "zookeeper":	return ZOOKEEPER
//	default:			log.Fatalln("invalid or unsupported serviced :", dst)
//	}
//	return TCP
//}
//const (
//	TCP       Protocol = iota + 1
//	UDP
//	DNS
//	NTP
//	HTTP
//	HTTPS
//	WEBSOCKET
//	THRIFT
//	PROTOBUF
//	MYSQL
//	PGSQL
//	MONGODB
//	CASSANDRA
//	REDIS
//	MEMCACHE
//	AEROSPIKE
//	KAFKA
//	PULSAR
//	ZOOKEEPER
//)
