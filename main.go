package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/server"
)

func main() {
	environment := flag.String("e", "development", "Application runtime environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	// We do not need to connect to db right now
	// db.Init()
	server.Init()
}
