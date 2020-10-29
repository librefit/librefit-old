package config

import (
	"sync"
)

var once sync.Once

// Config holds the app configuration (database, ports, etc)
type Config struct {
	Opts *Options
}

// NewConfig initialises a new configuration file
func NewConfig() *Config {
	o := newOptions()

	c := &Config{
		Opts: o,
	}

	return c
}
