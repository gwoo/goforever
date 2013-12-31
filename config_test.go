// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	r, err := LoadConfig("goforever.toml")

	if err != nil {
		t.Errorf("Error creating config %s.", err)
		return
	}
	if r == nil {
		t.Errorf("Expected %#v. Result %#v\n", r, nil)
	}
}

func TestConfigGet(t *testing.T) {
	c, _ := LoadConfig("goforever.toml")
	ex := "example/example.pid"
	r := string(c.Get("example").Pidfile)
	if ex != r {
		t.Errorf("Expected %#v. Result %#v\n", ex, r)
	}
}

func TestConfigKeys(t *testing.T) {
	c, _ := LoadConfig("goforever.toml")
	ex := []string{"example", "example-panic"}
	r := c.Keys()
	if len(ex) != len(r) {
		t.Errorf("Expected %#v. Result %#v\n", ex, r)
	}
}
