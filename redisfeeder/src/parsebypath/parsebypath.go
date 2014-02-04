package parsebypath

import (
	"domains"
//	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"log/syslog"
	"net/http"
	"strings"
)

func Parse(golog syslog.Writer, itemsarr []domains.Item, xpath []string) []domains.Item{

	client := &http.Client{}
	
	var returnitemsarr []domains.Item
	
	for _, item := range itemsarr {

		newitem := item
		
		req, err := http.NewRequest("GET", item.Link, nil)
		if err != nil {

			golog.Err(err.Error())

		}

		resp, err := client.Do(req)
		if err != nil {
	
			golog.Err(err.Error())
		}

		defer resp.Body.Close()
		bodybytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			golog.Err(err.Error())
		}
		doc, _ := gokogiri.ParseHtml([]byte(bodybytes))
		res, _ := doc.Search(xpath[0])

		for _, itemdom := range res {

			if res2, err := itemdom.Search("img"); err != nil {

				golog.Err(err.Error())
			} else {

				imglink := res2[0].Attr("src")
				newitem.ImgLink = imglink 

			}

		}

		res, _ = doc.Search(xpath[1])

		var cont string
		for i, itemdom := range res {

			if i == 0 {
			cont = itemdom.Content()
			} else {
			
				cont = cont+" " +itemdom.Content()
			}

		}
		
		newitem.Cont = cont
		
		if strings.HasSuffix(newitem.ImgLink, ".JPG") {
//			fmt.Println(newitem.ImgLink[strings.Index(newitem.ImgLink, "images"):len(newitem.ImgLink)])

			imgLink := newitem.Link[0:strings.Index(newitem.Link, "-")] + "/" + newitem.ImgLink[strings.Index(newitem.ImgLink, "images"):len(newitem.ImgLink)]
//			fmt.Println(imgLink)
			newitem.ImgLink = imgLink 
			
			returnitemsarr =append(returnitemsarr,newitem)

		}

	}
	return returnitemsarr
}
