package yqljsonparser

import (
	"domains"
	"fmt"
	"log"
	"log/syslog"
	"testing"
)

func TestParser(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	var feeslinksarr = []domains.FeedLinks{

		{RedisID: "it_IT:news:Home", YQLlink: "http://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Fquotidianohome.feedsportal.com%2Fc%2F33327%2Ff%2F565662%2Findex.rss%22&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"},
	}
	
//	Parser(*golog, feeslinksarr)

	testresult := Parser(*golog, feeslinksarr)
//	
	fmt.Println(testresult[0].Items[0].Title)
	fmt.Println(testresult[0].Items[0].Link)
//	

//	for _,item := range testresult {
//
//		fmt.Println(item)
//
//	}

}
