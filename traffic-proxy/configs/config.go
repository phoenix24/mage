package configs

type ProxyConfig struct {
	Name     string
	Port     int
	Storage  StorageConfig
	Services []ServiceConfig
}

type StorageConfig struct {
	URL      string
	Username string
	Password string
}
