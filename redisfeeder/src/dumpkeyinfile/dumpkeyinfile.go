package dumpkeyinfile

import (
"log/syslog"
"github.com/garyburd/redigo/redis"
"domains"
"encoding/json"
"os"
"time"

)

func DumpInFile(golog syslog.Writer,c redis.Conn, key string,lastdumpdate time.Time) {

	bitem, _ := redis.Strings(c.Do("ZREVRANGE", key, "0", "-1"))

	var itemsarr []domains.Item

	for _, item := range bitem {

		var itemobj domains.Item

		b := []byte(item)
		err := json.Unmarshal(b, &itemobj)
		if err != nil {
		
			golog.Err(err.Error())
		}
		itemsarr = append(itemsarr, itemobj)

	}

//	log.Println(len(itemsarr))

	f, err := os.OpenFile("dumpredis.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, item := range itemsarr {

		pubDate := item.PubDate

		if pubDate.After(lastdumpdate) {
//			log.Println(pubDate)

			if _, err = f.WriteString(item.Title + ".\n" + item.Cont + "\n"); err != nil {
				panic(err)
			}

		}

	}


}

