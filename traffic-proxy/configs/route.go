package configs

import (
	"fmt"
	"strings"
)

type Address string

func (r Address) parts() []string {
	var address = string(r)
	var cleaned = strings.ReplaceAll(address, "//", "")
	return strings.Split(cleaned, ":")
}

func (r Address) Port() string {
	return r.parts()[2]
}

func (r Address) Host() string {
	return r.parts()[1]
}

func (r Address) HostPort() string {
	return fmt.Sprintf("%s:%s", r.parts()[1], r.parts()[2])
}

func (r Address) Scheme() string {
	return r.parts()[0]
}
