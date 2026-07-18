package main

import (
	"cache-ttl/cache"
	"fmt"
	"time"
)

func main() {
	ourCache := cache.NewCache(time.Second * 30)
	ourCache.Set("name", "Joaquim", time.Second*1)
	fmt.Println(ourCache.Get("name"))
	ourCache.Set("address", "King Avenue", time.Second*3)
	time.Sleep(time.Second * 10)
	fmt.Println(ourCache.Get("name"))
	fmt.Println(ourCache.Get("address"))
	ourCache.Stop()
}
