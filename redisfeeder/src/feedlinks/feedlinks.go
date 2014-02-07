package feedlinks

import (
	"domains"
	"encoding/csv"
	"io"
	"log/syslog"
	"os"
)

func GetLinks(golog syslog.Writer, csvfile string) []domains.FeedLinks {

	var feedLinksarr []domains.FeedLinks
	
	csvFile, err := os.Open(csvfile)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	} else {

		csvReaderd := csv.NewReader(csvFile)
		csvReaderd.TrailingComma = true

		for {
			fields, err := csvReaderd.Read()
			var feedLinks domains.FeedLinks
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}

//			golog.Info(fields[0])
			feedLinks.RedisID = fields[0]
			feedLinks.YQLlink = fields[1]
			feedLinksarr = append(feedLinksarr,feedLinks)
		}

	}

	return feedLinksarr

}
