package services

import (
	"log"
	"traffic-proxy/configs"
)

func NewService(backend configs.ServiceConfig, storage configs.StorageConfig) *Service {
	log.Printf("starting service: %s\n", backend)
	log.Printf("    with storage: %s\n", storage)
	return nil
}

