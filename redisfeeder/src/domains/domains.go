package domains

import (
"time"
)

type Item struct {
	PubDate time.Time
	Link	string
	Title   string
	Cont    string
	ImgLink string
}

type RedisidItems struct {
	
	RedisID string
	Items []Item

}

type FeedLinks struct {

	RedisID string
	YQLlink string

}


