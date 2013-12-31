// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	"testing"
)

func TestPidfile(t *testing.T) {
	c := &Config{"", "",
		[]*Process{&Process{
			Name:    "test",
			Pidfile: "test.pid",
		}},
	}
	p := c.Get("test")
	err := p.Pidfile.write(100)
	if err != nil {
		t.Errorf("Error: %s.", err)
		return
	}
	ex := 100
	r := p.Pidfile.read()
	if ex != r {
		t.Errorf("Expected %#v. Result %#v\n", ex, r)
	}

	s := p.Pidfile.delete()
	if s != true {
		t.Error("Failed to remove pidfile.")
		return
	}
}

func TestProcessStart(t *testing.T) {
	c := &Config{
		"", "",
		[]*Process{&Process{
			Name:    "example",
			Command: "example/example",
			Pidfile: "example/example.pid",
			Logfile: "example/logs/example.debug.log",
			Errfile: "example/logs/example.errors.log",
			Respawn: 3,
		}},
	}
	p := c.Get("example")
	p.start("example")
	ex := 0
	r := p.x.Pid
	if ex >= r {
		t.Errorf("Expected %#v < %#v\n", ex, r)
	}
	p.stop()
}
