// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	//"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	if daemon.Name != "goforever" {
		t.Error("Daemon name is not goforever")
	}
	daemon.Args = []string{"foo"}
	daemon.start(daemon.Name)
	if daemon.Args[0] != "foo" {
		t.Error("First arg not foo")
	}
	daemon.find()
	daemon.stop()
}
