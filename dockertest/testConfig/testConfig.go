package testConfig

import "github.com/ory/dockertest/v3/docker"

type ConfigDocker struct {
	TestUser       string
	TestPassword   string
	TestHost       string
	TestDbname     string
	TestDbPort     string
	TestServerPort string
}

type RunOptions struct {
	Repository   string
	Tag          string
	Env          []string
	ExposedPorts string
	PortBindings map[docker.Port][]docker.PortBinding
}
