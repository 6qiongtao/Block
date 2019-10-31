package main

import (
	"fmt"
	//"vtoken_digiccy_go/route/config"

	"github.com/garyburd/redigo/redis"
	log "github.com/micro/go-micro/util/log"
)

func main() {
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	c, err := redis.Dial("tcp", "192.168.0.129:6379")
	if err != nil {
		log.Info("redis connect err:", err)
	}
	defer c.Close()

	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

}