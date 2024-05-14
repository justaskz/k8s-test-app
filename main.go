package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gomodule/redigo/redis"
)

func main() {
    http.HandleFunc("/", redisPingHandler)

    // Bind to all interfaces
    addr := ":8080"

    log.Printf("Listening on %s...\n", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}

func redisPingHandler(w http.ResponseWriter, r *http.Request) {
    redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        errMsg := "REDIS_HOST environment variable not set"
        http.Error(w, errMsg, http.StatusInternalServerError)
        log.Println(errMsg)
        return
    }

    conn, err := redis.Dial("tcp", redisHost+":6379")
    if err != nil {
        errMsg := fmt.Sprintf("Failed to connect to Redis: %v", err)
        http.Error(w, errMsg, http.StatusInternalServerError)
        log.Println(errMsg)
        return
    }
    defer conn.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    pong, err := redis.String(conn.Do("PING"))

    if err == redis.ErrNil {
        http.Error(w, "No response from Redis", http.StatusInternalServerError)
        log.Println("No response from Redis")
    } else if err != nil {
        errMsg := fmt.Sprintf("Redis PING failed: %v", err)
        http.Error(w, errMsg, http.StatusInternalServerError)
        log.Println(errMsg)
    } else if ctx.Err() == context.DeadlineExceeded {
        http.Error(w, "Redis PING timed out", http.StatusInternalServerError)
        log.Println("Redis PING timed out")
    } else if pong == "PONG" {
        fmt.Fprintln(w, "Redis connection successful!")
    } else {
        errMsg := fmt.Sprintf("Unexpected response from Redis: %s", pong)
        http.Error(w, errMsg, http.StatusInternalServerError)
        log.Println(errMsg)
    }
}
