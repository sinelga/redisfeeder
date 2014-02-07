package yqljsonparser

import (
	"domains"
	"encoding/json"
	"io/ioutil"
	"log/syslog"
	"net/http"
	"time"
)

func Parser(golog syslog.Writer, feedLinksarr []domains.FeedLinks) []domains.RedisidItems{

	client := &http.Client{}
	
	var redisiditemsarr []domains.RedisidItems

	for _, feedLink := range feedLinksarr {
	
		var redisiditems domains.RedisidItems
		redisiditems.RedisID = feedLink.RedisID

		req, err := http.NewRequest("GET", feedLink.YQLlink, nil)
		if err != nil {

			golog.Err(err.Error())
		}
		req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.0; WOW64; rv:24.0) Gecko/20100101 Firefox/24.0") 

		resp, err := client.Do(req)
		if err != nil {

			golog.Err(err.Error())
		}

		defer resp.Body.Close()
		bodybytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			golog.Err(err.Error())

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
		
		redisiditems.Items = itemsarr		
		redisiditemsarr = append(redisiditemsarr,redisiditems)

	}
	
	return redisiditemsarr

}
