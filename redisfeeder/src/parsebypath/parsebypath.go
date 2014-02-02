package parsebypath

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"log/syslog"
	"net/http"
)

func Parse(golog syslog.Writer, link string, xpath string) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {

		golog.Err(err.Error())

	}

	resp, err := client.Do(req)
	if err != nil {
		//		log.Fatalln(err)
		golog.Err(err.Error())
	}

	defer resp.Body.Close()
	bodybytes, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		golog.Err(err.Error())
	}
	doc, _ := gokogiri.ParseHtml([]byte(bodybytes))
	res, _ := doc.Search(xpath)

	for _, itemdom := range res {

		fmt.Println(itemdom)
		if res2, err := itemdom.Search("img"); err != nil {
		
			golog.Err(err.Error())
		} else {
		
			fmt.Println(res2)
			
			imglink := res2[0].Attr("src")
			fmt.Println(imglink)
		
		}
		
		

	}

}
