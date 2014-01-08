# Goforever [![Build Status](https://travis-ci.org/gwoo/goforever.png)](https://travis-ci.org/gwoo/goforever)


Config based process manager. Goforever could be used in place of supervisor, runit, node-forever, etc.
Goforever will start an http server on the specified port.


	Usage of ./goforever:
	  -conf="goforever.toml": Path to config file.
	  -d=false: Daemonize goforever. Must be first flag
	  -password="test": Password for basic auth.
	  -port=8080: Port for the server.
	  -username="demo": Username for basic auth.

## CLI
	list				List processes.
	show <process>		Show a process.
	start <process>		Start a process.
	stop <process>		Stop a process.
	restart <process>	Restart a process.


## HTTP API

Return a list of managed processes

	GET host:port/

Start the process

	POST host:port/:name

Restart the process

	PUT host:port/:name

Stop the process

	DELETE host:port/:name
