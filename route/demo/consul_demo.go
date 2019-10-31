package main

import (
	"fmt"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/util/log"
	"vtoken_digiccy_go/route/config"
)

var ConsulKV *consulApi.KV

func consulApiInit() *consulApi.KV {
	conf := consulApi.DefaultConfig()
	conf.Address = config.ConsulCfg.ConsulAddr
	client, err := consulApi.NewClient(conf)
	if err != nil {
		log.Info("consulApi NewClient err:", err)
	}

	ConsulKV = client.KV()
	return ConsulKV
}

func main()  {
	//pair, _, err := ConsulKV.Get("test", nil)
	//if err != nil {
	//	log.Info("client.KV().Get err:", err)
	//}
	//fmt.Printf("%+v\n", string(pair.Value))

	//创建key、velue
	p := &consulApi.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err := ConsulKV.Put(p, nil)
	if err != nil {
		panic(err)
	}

	pair, _, err := ConsulKV.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	_, err = ConsulKV.Delete("REDIS_MAXCLIENTS", nil)
	if err != nil {
		log.Info("ConsulKV.Delete err:", err)

	}

	pair, _, err = ConsulKV.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		log.Info("client.KV().Get err:", err)
	}
	if pair != nil {
		fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	}



}
