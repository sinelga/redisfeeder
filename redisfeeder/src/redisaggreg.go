package main

import (
	"domains"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"parsebypath"
	"redisin"
//	"strings"
	"time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

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

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Fquotidianohome.feedsportal.com%2Fc%2F33327%2Ff%2F565662%2Findex.rss%22&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var jItems map[string]map[string]interface{}

	err = json.Unmarshal(bodybytes, &jItems)
	if err != nil {
		golog.Err(err.Error())

	}

	items := jItems["query"]["results"].(map[string]interface{})

	itemsii := items["item"].([]interface{})

	var itemsarr []domains.Item

	for _, ii := range itemsii {

		var item domains.Item

		iii := ii.(map[string]interface{})

		pubDatestr := iii["pubDate"].(string)

		pubDate, err := time.Parse(time.RFC1123, pubDatestr)
		if err != nil {
			golog.Err(err.Error())

		}

		item.PubDate = pubDate
		title := iii["title"].(string)
		item.Title = title
		iiii := iii["guid"].(map[string]interface{})
		link := iiii["content"].(string)
		item.Link = link

		itemsarr = append(itemsarr, item)
	}

	xpaths := []string{"//div[@class='photo']", "//div[@class='description lp_links']/p"}

	itemsarrfull := parsebypath.Parse(*golog, itemsarr, xpaths)

//	for _, item := range itemsarrfull {
//
//		fmt.Println(item.Link)
//		fmt.Println(item.ImgLink)
//
//	}
	
	redisin.InsertIn(itemsarrfull) 

}
