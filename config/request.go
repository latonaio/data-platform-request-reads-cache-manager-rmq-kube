package config

import (
	"fmt"
	"os"
)

func newREQUEST() *REQUEST {
	return &REQUEST{
		host: os.Getenv("NESTJS_DATA_CONNECTION_REQUEST_CONTROL_MANAGER_HOST"),
		port: os.Getenv("NESTJS_DATA_CONNECTION_REQUEST_CONTROL_MANAGER_PORT"),
	}
}

type REQUEST struct {
	host string
	port string
}

func (c *REQUEST) RequestURL() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}
