package server

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "go-notify/config"
)

func Init() {
    config := config.GetConfig()
    router := NewRouter()

    srv := &http.Server{
        Addr:    "0.0.0.0:" + config.GetString("server.port"),
        Handler: router,
    }
    go Run(srv)
    GracefullyShutdown(srv)
}

// Run Start application and listen incoming request on target port
func Run(server *http.Server) {
    log.Print(fmt.Sprintf("Listening and serving HTTP on %s", server.Addr))
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Listen: %s\n", err)
    }
}

// GracefullyShutdown Gracefully shutdown server by release its resource and currently handling context.
func GracefullyShutdown(server *http.Server) {
    // Wait for interrupt signal to gracefully shut down the server with
    // a timeout of 5 seconds.
    quit := make(chan os.Signal)
    // kill (no param) default send syscall.SIGTERM
    // kill -2 is syscall.SIGINT
    // kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    // The context is used to inform the server it has 30 seconds to finish
    // the request it is currently handling
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exiting")
}
