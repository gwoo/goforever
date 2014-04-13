# Goforever [![Build Status](https://travis-ci.org/gwoo/goforever.png)](https://travis-ci.org/gwoo/goforever)

Config based process manager. Goforever could be used in place of supervisor, runit, node-forever, etc.
Goforever will start an http server on the specified port.

	Usage of ./goforever:
	  -conf="goforever.toml": Path to config file.

## Running
Help.

	./goforever -h

Daemonize main process.

	./goforever start

Run main process and output to current session.

	./goforever

## CLI
	list				List processes.
	show [process]	    Show a main proccess or named process.
	start [process]		Start a main proccess or named process.
	stop [process]		Stop a main proccess or named process.
	restart [process]	Restart a main proccess or named process.

## HTTP API

Return a list of managed processes

	GET host:port/

Start the process

	POST host:port/:name

Restart the process

	PUT host:port/:name

Stop the process

	DELETE host:port/:name
