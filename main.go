package main

import (
    "flag"
    "fmt"
    "os"

    "go-notify/config"
    "go-notify/db"
    "go-notify/server"
)

func main() {
    environment := flag.String("e", "development", "Application runtime environment")
    flag.Usage = func() {
        fmt.Println("Usage: server -e {mode}")
        os.Exit(1)
    }
    flag.Parse()
    config.Init(*environment)
    db.Init()
    server.Init()
}
