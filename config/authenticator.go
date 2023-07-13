package config

import (
	"fmt"
	"os"
)

func newAuth() *AUTHENTICATOR {
	return &AUTHENTICATOR{
		host: os.Getenv("AUTHENTICATOR_HOST"),
		port: os.Getenv("AUTHENTICATOR_PORT"),
	}
}

type AUTHENTICATOR struct {
	host string
	port string
}

func (c *AUTHENTICATOR) RequestURL() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}
