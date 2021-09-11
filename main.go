package main

import (
	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/server"
)

func main() {
	config.Init()
	// We do not need to connect to db right now
	// db.Init()
	server.Init()
}
