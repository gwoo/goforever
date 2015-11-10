// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	IP        string
	Port      string
	Username  string
	Password  string
	Daemonize bool
	Pidfile   Pidfile
	Logfile   string
	Errfile   string
	Processes []*Process `toml:"process"`
}

func (c Config) Keys() []string {
	keys := []string{}
	for _, p := range c.Processes {
		keys = append(keys, p.Name)
	}
	return keys
}

func (c Config) Get(key string) *Process {
	for _, p := range c.Processes {
		if p.Name == key {
			return p
		}
	}
	return nil
}

func LoadConfig(file string) (*Config, error) {
	if string(file[0]) != "/" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		file = filepath.Join(wd, file)
	}
	var c *Config
	if _, err := toml.DecodeFile(file, &c); err != nil {
		return nil, err
	}
	return c, nil
}
