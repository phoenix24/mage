package main

import "fmt"

type SvcConf struct {
	Port int
	Name string
}

func (c SvcConf) HostPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
