package main

import (
	"bytes"
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"strings"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	callback := req.Header.Get("X-CALLBACK")
	redisid := req.Header.Get("X-REDISID")

	feeder(*golog, resp, req, callback, redisid)

}

func main() {

	log.Println("Server Start")
	listener, err := net.Listen("tcp", "127.0.0.1:8010")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

func feeder(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, callback string, redisid string) {

	redisidconv := strings.Replace(redisid, "%3A", ":", -1)

	golog.Info("redisid-> " + redisidconv)
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	if redisid == "" {

		redisid = "it_IT:news:Home"

	}

	bitem, _ := redis.Strings(c.Do("ZREVRANGE", redisidconv, "0", "25"))

	var itemsarr []domains.Item

	for _, item := range bitem {

		var itemobj domains.Item

		b := []byte(item)
		err := json.Unmarshal(b, &itemobj)
		if err != nil {
			log.Fatal(err)
		}
		itemsarr = append(itemsarr, itemobj)

	}

	bout, err := json.Marshal(itemsarr)
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	buffer.WriteString(callback + "(")
	buffer.Write(bout)
	buffer.WriteString(");")

	resp.Write(buffer.Bytes())
}
