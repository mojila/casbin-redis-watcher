package main

import (
	"fmt"
	"time"

	rediswatcher "github.com/mojila/casbin-redis-watcher"
)

func main() {
	w, err := rediswatcher.NewWatcher("localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	w.SetUpdateCallback(func(string) { fmt.Println("Updated") })

	time.Sleep(10 * time.Minute)
}
