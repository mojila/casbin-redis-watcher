package main

import (
	"fmt"
	"time"

	rediswatcher "github.com/billcobbler/casbin-redis-watcher/v2"
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
