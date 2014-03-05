// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	//"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	if daemon.Name != "goforever" {
		t.Error("Daemon name is not goforever")
	}
	os.Args = []string{"./goforever", "-d", "foo"}
	dize := true
	d = &dize
	main()
	if daemon.Args[0] != "foo" {
		t.Error("First arg not foo")
	}
	daemon.stop()
}
