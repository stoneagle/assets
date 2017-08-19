package resource

import "assets/web/backend/library"

type LatestNew struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Classify string `json:"classify"`
	Time     string `json:"time"`
	Content  string `json:"content"`
}

type Notice struct {
	Title string       `json:"title"`
	Url   string       `json:"url"`
	Type  string       `json:"type"`
	Date  library.Date `json:"date"`
}

type SinaGuba struct {
	Rcounts int    `json:"rcounts"`
	Ptime   string `json:"ptime"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
