package redisin

import (
	"domains"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func InsertIn(itemarr []domains.Item) {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	queuename := "it_IT:news"

	for _, item := range itemarr {

		bitem, err := json.Marshal(item)
		if err != nil {
			fmt.Println("error:", err)
		}
		//		os.Stdout.Write(b)
		fmt.Println(string(bitem))

		if pgq, err := c.Do("ZADD", queuename,time.Now().Unix(), bitem); err != nil {
			log.Fatal(err)

		} else {

			log.Println("in queue ", pgq)

		}

	}

}
