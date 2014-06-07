Project GGS
===========

This is a go on trying to implement a multiplayer game server in
[Go](http://golang.org/). It is desinged for concurrency and the ability to
easily implement more games on the server with as much ease as possible.

It currently has two games implemented:
[Tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe) and
[Achtung, die Kurve!](https://en.wikipedia.org/wiki/Achtung,_die_Kurve!).
A simple Java based client is available for both of those games.

Installation
------------

The server and client can be compiled and run using `make`.

Command          | Action
---------------- | ----------------------------------------
`make runserver` | Runs the server
`make runclient` | Runs the client
`make test`      | Runs test cases
`make docs`      | Starts documentation server on port 8050
`make clean`     | Cleans up generated binary files
