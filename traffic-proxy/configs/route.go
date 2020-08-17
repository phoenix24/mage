package configs

import (
	"strings"
)

type Route string

func (r Route) parts() []string {
	var route = string(r)
	var cleaned = strings.ReplaceAll(route, "//", "")
	return strings.Split(cleaned, ":")
}

func (r Route) Host() string {
	return r.parts()[1]
}

func (r Route) Port() string {
	return r.parts()[2]
}

func (r Route) Service() string {
	return r.parts()[0]
}
