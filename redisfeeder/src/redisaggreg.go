package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	//    "domains"
	"parsebypath"
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

	//    golog.Info(string(bodybytes))

	//    var jItems domains.Ysqlquery
	var jItems map[string]map[string]interface{}
	//	var jItems  map[string]interface{}

	err = json.Unmarshal(bodybytes, &jItems)
	if err != nil {
		golog.Err(err.Error())

	}

	items := jItems["query"]["results"].(map[string]interface{})

	itemsii := items["item"].([]interface{})

	var linksarr []string

	for _, ii := range itemsii {

		//    	fmt.Println(ii)

		iii := ii.(map[string]interface{})

		//    	fmt.Println(iii["guid"])

		iiii := iii["guid"].(map[string]interface{})

		//    	fmt.Println(iiii["content"])

		linksarr = append(linksarr, iiii["content"].(string))

	}

	for _, link := range linksarr {

		fmt.Println(link)

	}

	parsebypath.Parse(*golog, "http://qn.quotidiano.net/curiosita/2014/01/30/1018121-elisabetta-rolls-royce-bentley.shtml", "//div[@class='photo']")

	//    for i := range jItems["query"]{
	//
	//    	item := jItems["query"][i]
	//
	//
	//    	if i == "results" {
	//    		fmt.Println(i)
	//    		fmt.Println(item["item"])
	//    		itemii := item.(map[string])
	//
	//    	}
	//
	//
	//
	//    }

	//    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//
	//    var dat map[string]interface{}
	//
	//    if err := json.Unmarshal(byt, &dat); err != nil {
	//        panic(err)
	//    }
	//    fmt.Println(dat)

	//    jItems.

	//    fmt.Printf(jItems.query.results.item[0].title )

}
