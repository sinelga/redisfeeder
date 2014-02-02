package domains

import ()

type Item struct {
	Title   string
	Cont    string
	ImgLink string
}

type Ysqlquery struct {

	query Ysqlresults
}

type Ysqlresults struct {
	results Ysql
}

type Ysql struct {
	item []Ysqlitem
}

type Ysqlitem struct {
	title       string
	link        string
	description string
	pubDate     string
	author      string
	guid        []Ysqlguid
}

type Ysqlguid struct {
	isPermaLink string
	content     string
}
