package main

import (
	"fmt"
	cache2 "main/cache"
	"time"
)

func main() {
	cache := cache2.New()

	cache.Set("name", "pavel", time.Second*3)
	time.Sleep(time.Second * 2)
	fmt.Println(cache.Get("name"))
}
