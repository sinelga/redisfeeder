package main

import (
//	"domains"
//	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"lastdump"
	"log"
	"log/syslog"
//	"os"
	"time"
	"dumpkeyinfile"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	lastdumpdate := lastdump.GetLastDate(*golog)
	println(lastdumpdate.Format(time.RFC1123))
	
	
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Err(err.Error())
	}
	defer c.Close()
	
	dumpkeyinfile.DumpInFile(*golog,c,"it_IT:news:Home",lastdumpdate)
//
//	bitem, _ := redis.Strings(c.Do("ZREVRANGE", "it_IT:news:Home", "0", "-1"))
//
//	var itemsarr []domains.Item
//
//	for _, item := range bitem {
//
//		var itemobj domains.Item
//
//		b := []byte(item)
//		err := json.Unmarshal(b, &itemobj)
//		if err != nil {
//			log.Fatal(err)
//		}
//		itemsarr = append(itemsarr, itemobj)
//
//	}
//
//
//	f, err := os.OpenFile("dumpredis.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	for _, item := range itemsarr {
//
//		pubDate := item.PubDate
//
//		if pubDate.After(lastdumpdate) {
//			log.Println(pubDate)
//
//			if _, err = f.WriteString(item.Title + ".\n" + item.Cont + "\n"); err != nil {
//				panic(err)
//			}
//
//		}
//
//	}

}
