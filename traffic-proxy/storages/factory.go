package storages

import (
	"log"
	"traffic-proxy/configs"
)

type Storage interface {
}

func NewStorage(config configs.StorageURL) *Storage {
	log.Printf("starting storage: %s\n", config.URL)
	return nil
}
