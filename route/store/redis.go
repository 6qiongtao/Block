// @Time : 2019/10/20 17:30
// @Author : Barry
package store

import (
	"time"
	"vtoken_digiccy_go/route/config"

	cache "vtoken_digiccy_go/route/tools/cache"

	log "github.com/micro/go-micro/util/log"
)


var (
	Re error
	Rs *cache.Cache

)

func init() {
	log.Info("Redis Engine Init")
	Rs, Re = cache.NewCache(config.RedisUrl)
	if Re != nil {
		log.Fatal(Re)
	}
}

func Get(prefix, key string) interface{} {
	return Rs.Get(prefix + key)
}

func RedisBytes(prefix, key string) (data []byte, err error) {
	return Rs.RedisBytes(prefix + key)
}

func RedisString(prefix, key string) (data string, err error) {
	return Rs.RedisString(prefix + key)
}

func RedisInt(prefix, key string) (data int, err error) {
	return Rs.RedisInt(prefix + key)
}

func Put(prefix, key string, val interface{}, timeout time.Duration) error {
	return Rs.Put(prefix+key, val, timeout)
}

func SetNX(prefix, key string, val interface{}, timeout time.Duration) bool {
	return Rs.SetNX(prefix+key, val, timeout)
}

func Delete(prefix, key string) error {
	return Rs.Delete(prefix + key)
}

func IsExist(merchant, key string) bool {
	return Rs.IsExist(merchant + key)
}

func LPush(prefix, key string, val interface{}) error {
	return Rs.LPush(prefix+key, val)
}

func LRem(prefix, key string, count int, val interface{}) error {
return Rs.LRem(prefix+key, count, val)
}

func Brpop(prefix, key string, callback func([]byte)) {
	Rs.Brpop(prefix+key, callback)
}

func GetRedisTTL(prefix, key string) time.Duration {
	return Rs.GetRedisTTL(prefix + key)
}

func Incrby(prefix, key string, num int) (interface{}, error) {
	return Rs.Incrby(prefix+key, num)
}
