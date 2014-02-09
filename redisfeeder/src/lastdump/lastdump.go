package lastdump

import (
	"io/ioutil"
	"log/syslog"
	"os"
	"time"
	"strings"
)

func GetLastDate(golog syslog.Writer) time.Time {

	lasttime := time.Now()

	if _, err := os.Stat("lastdump.txt"); err != nil {

		if os.IsNotExist(err) {
			// file does not exist
			golog.Info(err.Error())
			f, err := os.OpenFile("lastdump.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			if _, err = f.WriteString(lasttime.Format(time.RFC1123)); err != nil {
				panic(err)
			}

		} else {
			// other error
		}

	} else {

		contents, _ := ioutil.ReadFile("lastdump.txt")
		
		contentsclean := strings.TrimSpace(string(contents)) 

//		println(string(contents))

		lastdumpdate, err := time.Parse(time.RFC1123, contentsclean)
		if err != nil {
			golog.Err(err.Error())
		}

//		println(lastdumpdate.Format(time.RFC1123))
		lasttime = lastdumpdate

		f, err := os.OpenFile("lastdump.txt", os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(time.Now().Format(time.RFC1123)); err != nil {
			panic(err)
		}

	}

	return lasttime
}
