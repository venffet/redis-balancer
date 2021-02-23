package main

import (
	"log"
	"time"

	"github.com/venffet/redis-balancer"
	"gopkg.in/redis.v5"
)

func main() {
	clients := balancer.New(
		[]balancer.Options{
			{Options: redis.Options{Network: "tcp", Addr: "127.0.0.1:6376"}, CheckInterval: 600 * time.Millisecond},
			{Options: redis.Options{Network: "tcp", Addr: "127.0.0.1:6377"}, CheckInterval: 800 * time.Millisecond},
			{Options: redis.Options{Network: "tcp", Addr: "127.0.0.1:6378"}, CheckInterval: 800 * time.Millisecond},
		},
		balancer.ModeRoundRobin,
	)
	defer clients.Close()

	client := clients.Next()
	res, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
