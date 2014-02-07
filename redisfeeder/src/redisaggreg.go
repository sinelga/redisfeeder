package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"	
	"parsebypath"
	"redisin"
	"feedlinks"
	"yqljsonparser"

)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

//news
//select * from rss where url = "http://quotidianohome.feedsportal.com/c/33327/f/565662/index.rss"
//http://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Fquotidianohome.feedsportal.com%2Fc%2F33327%2Ff%2F565662%2Findex.rss%22&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys
//news:cronica
//select * from rss where url = "http://quotidianohome.feedsportal.com/c/33327/f/565663/index.rss"
//http://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Fquotidianohome.feedsportal.com%2Fc%2F33327%2Ff%2F565663%2Findex.rss%22&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys

// select * from  html where url='http://qn.quotidiano.net/esteri/2014/01/30/1018048-maro-india-governo-morte.shtml' AND xpath="//div[@class='photo']"
//
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

	feedlinksarr := feedlinks.GetLinks(*golog, "feedlinks.csv")

	redisidItemsarr := yqljsonparser.Parser(*golog, feedlinksarr)

	xpaths := []string{"//div[@class='photo']", "//div[@class='description lp_links']/p"}

	redisidItemsarrfull := parsebypath.Parse(*golog, redisidItemsarr, xpaths)

	for _, redisidItems := range redisidItemsarrfull {

		golog.Info(redisidItems.RedisID)

	}

	redisin.InsertIn(redisidItemsarrfull)

}
