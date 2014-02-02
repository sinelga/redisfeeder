package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	//	"github.com/moovweb/gokogiri/css"
	"io/ioutil"
	"log"
	"net/http"
	"domains"
	"redisin"
	"strings"

//	"os"
)

func main() {

	var itemarr []domains.Item

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://www.corriere.it/economia/corriereconomia/", nil)
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

	body := toUtf8(bodybytes)

	doc, _ := gokogiri.ParseHtml([]byte(body))
	res, _ := doc.Search("//div[@class='homearticle-box']")

	for _, itemdom := range res {

		if res2, err := itemdom.Search("a/img"); err != nil {

			fmt.Println(err.Error())

		} else {

			if len(res2) > 0 {

				var title string
				var cont string
				var imglink string

				title = res2[0].Attr("title")
				imglink = res2[0].Attr("src")

				if res3, err := itemdom.Search("p[2]"); err != nil {

					fmt.Println(err.Error())
				} else {

					//					fmt.Println(i, creanstr(res3[0].InnerHtml()))
					cont = creanstr(res3[0].InnerHtml())

				}

				if len(title) > 3 && len(imglink) > 10 && len(cont) > 5 {

					item := domains.Item{

						Title:   title,
						ImgLink: imglink,
						Cont:     cont,
					}

					itemarr = append(itemarr, item)

				}

			}

		}
			
	}
	redisin.InsertIn(itemarr)
	

}

func toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)

}

func creanstr(str string) string {

	var retstr string

	if len(str) > 10 {

		if strings.Index(str, "<") > 0 {

			retstr = str[0:strings.Index(str, "<")]

		} else {

			retstr = str

		}

	}
	retstr = strings.TrimLeft(retstr, "\n")

	return strings.TrimRight(retstr, " ") + "."

}
