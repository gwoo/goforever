// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gwoo/greq"
)

var d = flag.Bool("d", false, "Daemonize goforever. Must be first flag")
var conf = flag.String("conf", "goforever.toml", "Path to config file.")
var port = flag.Int("port", 8080, "Port for the server.")
var username = flag.String("username", "demo", "Username for basic auth.")
var password = flag.String("password", "test", "Password for basic auth.")
var server string
var config *Config

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	usage := `
Process subcommands
  list			List processes.
  show <process>	Show a process.
  start <process>	Start a process.
  stop <process>	Stop a process.
  restart <process>	Restart a process.
`
	fmt.Fprintln(os.Stderr, usage)
}

func init() {
	setConfig()
	setHost()
	daemon = &Process{
		Name:    "goforever",
		Args:    []string{"./goforever"},
		Command: "goforever",
		Pidfile: "goforever.pid",
		Logfile: "goforever.debug.log",
		Errfile: "goforever.errors.log",
		Respawn: 1,
	}
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	daemon.Name = "goforever"
	if *d == true {
		daemon.Args = append(daemon.Args, os.Args[2:]...)
		daemon.start(daemon.Name)
		return
	}
	if len(flag.Args()) > 0 {
		fmt.Printf("%s\n", Cli())
		return
	}
	if len(flag.Args()) == 0 {
		RunDaemon()
		HttpServer()
		return
	}
}

func Cli() string {
	sub := flag.Arg(0)
	name := flag.Arg(1)
	var o []byte

	if sub == "list" {
		o, _ = greq.Get("/")
	}
	if name != "" {
		switch sub {
		case "show":
			o, _ = greq.Get("/" + name)
		case "start":
			o, _ = greq.Post("/"+name, nil)
		case "stop":
			o, _ = greq.Delete("/" + name)
		case "restart":
			o, _ = greq.Put("/"+name, nil)
		}
	}
	return string(o)
}

func RunDaemon() {
	fmt.Printf("Running %s.\n", daemon.Name)
	daemon.children = make(map[string]*Process, 0)
	for _, name := range config.Keys() {
		daemon.children[name] = config.Get(name)
	}
	daemon.run()
}

func setConfig() {
	c, err := LoadConfig(*conf)
	if err != nil {
		log.Fatalf("Config error: %s", err)
		return
	}
	config = c

	if config.Username != "" {
		username = &config.Username
	}
	if config.Password != "" {
		password = &config.Password
	}
}

func setHost() {
	scheme := "https"
	if isHttps() == false {
		scheme = "http"
	}
	greq.Host = fmt.Sprintf("%s://%s:%s@0.0:%d", scheme, *username, *password, *port)
}
