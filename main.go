package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var cluster = []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"}
var single = []string{":5555"}

func getRedisClient(hosts []string) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{Addrs: hosts})
}

func doPipeline(name string, hosts []string) {
	defer timeTrack(time.Now(), name+" Pipeline")
	client := getRedisClient(hosts)
	pipeline := client.Pipeline()
	for i := 0; i < 500000; i++ {
		pipeline.IncrBy(fmt.Sprint("{slot1}key:", i), 1)
	}

	if _, err := pipeline.Exec(); err != nil {
		fmt.Println("Ocorreu um erro :", err.Error())
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("%s took %s", name, elapsed))
}

func main() {
	doPipeline("Redis Single Node", single)
	doPipeline("Redis Cluster", cluster)
}
