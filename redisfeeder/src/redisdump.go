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
	"dumpkeyinfile"
	"getkeysfromredis"
//	"time"
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
//	println(lastdumpdate.Format(time.RFC1123))

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Err(err.Error())
	}
	defer c.Close()

	keysarr := getkeysfromredis.GetKeys(*golog, c, "it_IT:news:*")

	for _, key := range keysarr {
		dumpkeyinfile.DumpInFile(*golog, c, key, lastdumpdate)

	}


}
