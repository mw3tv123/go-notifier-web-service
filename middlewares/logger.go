package middlewares

import (
    "fmt"
    "io"
    "log"
    "os"
    "time"

    "github.com/gin-gonic/gin"
)

// ConfigLogger Setup Logger for Server
func ConfigLogger() {
    f, err := os.Create("gin.log")
    if err != nil {
        panic(err.Error())
    }
    // Force log's color
    gin.ForceConsoleColor()
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
    log.SetOutput(gin.DefaultWriter)
}

// LogFormatter Custom format log output
func LogFormatter(param gin.LogFormatterParams) string {
    return fmt.Sprintf("[%s] - %s %s %s %d %s\n",
        param.TimeStamp.Format(time.RFC1123),
        param.ClientIP,
        param.Method,
        param.Path,
        param.StatusCode,
        param.ErrorMessage,
    )
}
