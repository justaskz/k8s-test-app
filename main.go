package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gomodule/redigo/redis"
)

func main() {
    http.HandleFunc("/", redisPingHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func redisPingHandler(w http.ResponseWriter, r *http.Request) {
    redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        http.Error(w, "REDIS_HOST environment variable not set", http.StatusInternalServerError)
        return
    }

    conn, err := redis.Dial("tcp", redisHost+":6379")
    if err != nil {
        http.Error(w, "Failed to connect to Redis", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    pong, err := redis.String(conn.Do("PING"))
    if err != nil {
        http.Error(w, "Redis PING failed", http.StatusInternalServerError)
        return
    }

    if pong == "PONG" {
        fmt.Fprintln(w, "Redis connection successful!")
    } else {
        http.Error(w, "Unexpected response from Redis", http.StatusInternalServerError)
    }
}
