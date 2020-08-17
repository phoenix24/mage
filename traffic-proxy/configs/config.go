package configs

type Mode string
type StorageURL string

const (
	PROXY  Mode = "PROXY"
	RECORD Mode = "RECORD"
	REPLAY Mode = "REPLAY"
)

type ServerConfig struct {
	Name    string     `mapstructure:"name"`
	Mode    Mode       `mapstructure:"mode"`
	Source  Address    `mapstructure:"source"`
	Backend Address    `mapstructure:"backend"`
	Storage StorageURL `mapstructure:"storage"`
}

type ProxyConfig struct {
	Name    string
	Port    int
	Servers []ServerConfig
}

//func (b ServerConfig) Scheme() services.Scheme {
//	var dst, src = b.Backend.Scheme(), b.Source.Scheme()
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
//	TCP       Scheme = iota + 1
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
