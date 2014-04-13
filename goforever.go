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

var conf = flag.String("conf", "goforever.toml", "Path to config file.")
var config *Config
var daemon *Process

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	usage := `
Commands
  list              List processes.
  show [name]       Show main proccess or named process.
  start [name]      Start main proccess or named process.
  stop [name]       Stop main proccess or named process.
  restart [name]    Restart main proccess or named process.
`
	fmt.Fprintln(os.Stderr, usage)
}

func init() {
	flag.Usage = Usage
	flag.Parse()
	setConfig()
	daemon = &Process{
		Name:    "goforever",
		Args:    []string{},
		Command: "goforever",
		Pidfile: config.Pidfile,
		Logfile: config.Logfile,
		Errfile: config.Errfile,
		Respawn: 1,
	}
}

func main() {
	if len(flag.Args()) > 0 {
		fmt.Printf("%s", Cli())
		return
	}
	if len(flag.Args()) == 0 {
		RunDaemon()
		HttpServer()
		return
	}
}

func Cli() string {
	var o []byte
	var err error
	sub := flag.Arg(0)
	name := flag.Arg(1)
	req := greq.New(host(), true)
	if sub == "list" {
		o, _, err = req.Get("/")
	}
	if name == "" {
		if sub == "start" {
			daemon.Args = append(daemon.Args, os.Args[2:]...)
			return daemon.start(daemon.Name)
		}
		_, _, err = daemon.find()
		if err != nil {
			return fmt.Sprintf("Error: %s.\n", err)
		}
		if sub == "show" {
			return fmt.Sprintf("%s.\n", daemon.String())
		}
		if sub == "stop" {
			message := daemon.stop()
			return message
		}
		if sub == "restart" {
			ch, message := daemon.restart()
			fmt.Print(message)
			return fmt.Sprintf("%s\n", <-ch)
		}
	}
	if name != "" {
		path := fmt.Sprintf("/%s", name)
		switch sub {
		case "show":
			o, _, err = req.Get(path)
		case "start":
			o, _, err = req.Post(path, nil)
		case "stop":
			o, _, err = req.Delete(path)
		case "restart":
			o, _, err = req.Put(path, nil)
		}
	}
	if err != nil {
		fmt.Printf("Process error: %s", err)
	}
	return fmt.Sprintf("%s\n", o)
}

func RunDaemon() {
	daemon.children = make(map[string]*Process, 0)
	for _, name := range config.Keys() {
		daemon.children[name] = config.Get(name)
	}
	daemon.run()
}

func setConfig() {
	var err error
	config, err = LoadConfig(*conf)
	if err != nil {
		log.Fatalf("%s", err)
		return
	}
	if config.Username == "" {
		log.Fatalf("Config error: %s", "Please provide a username.")
		return
	}
	if config.Password == "" {
		log.Fatalf("Config error: %s", "Please provide a password.")
		return
	}
	if config.Port == "" {
		config.Port = "2224"
	}
}

func host() string {
	scheme := "https"
	if isHttps() == false {
		scheme = "http"
	}
	return fmt.Sprintf("%s://%s:%s@0.0:%s",
		scheme, config.Username, config.Password, config.Port,
	)
}
