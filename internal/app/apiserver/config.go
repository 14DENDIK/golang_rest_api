package apiserver

import "github.com/14DENDIK/gopher_school/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
