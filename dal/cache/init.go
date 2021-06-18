package cache

import (
	"fmt"
	"github.com/JackTJC/backend/conf"
	"github.com/go-redis/redis/v8"
)
var cacheConn *cacheCli
type cacheCli struct {
	*redis.Client
}
func Init(){
	addr:=fmt.Sprintf("%s:%s",conf.Conf.Redis.Host,conf.Conf.Redis.Port)
	rdb:=redis.NewClient(&redis.Options{Addr: addr})
	cacheConn=&cacheCli{rdb}
}

func constructKey(format string,id int)string{
	return fmt.Sprintf(format,id)
}
