package redisin

import (
	"domains"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
//	"time"
)

//func InsertIn(itemarr []domains.Item) {
  func InsertIn(redisidItemsarr []domains.RedisidItems) {	

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	
	for _,redisidItems := range redisidItemsarr {

//	var queuename string		
	queuename := redisidItems.RedisID
	
	

	for _, item := range redisidItems.Items {

		bitem, err := json.Marshal(item)
		if err != nil {
			fmt.Println("error:", err)
		}
		//		os.Stdout.Write(b)
		fmt.Println(string(bitem))

		if pgq, err := c.Do("ZADD", queuename,item.PubDate.Unix(), bitem); err != nil {
			log.Fatal(err)

		} else {

			log.Println("in queue ", pgq)

		}

	}
}
}
