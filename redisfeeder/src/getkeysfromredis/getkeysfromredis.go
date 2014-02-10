package getkeysfromredis

import (
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func GetKeys(golog syslog.Writer, c redis.Conn, patern string) []string {

	var keysarr []string

	keys, _ := redis.Strings(c.Do("KEYS", patern))

	for _, key := range keys {

		//		golog.Info(key)
		keysarr = append(keysarr, key)

	}

	return keysarr

}
