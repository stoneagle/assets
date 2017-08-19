package resource

import "assets/web/backend/library"

type Area struct {
	Code string `json:"code"`
	Area string `json:"area"`
	Name string `json:"name"`
}

type Concept struct {
	Code  string `json:"code"`
	Cname string `json:"c_name"`
	Name  string `json:"name"`
}

type Industry struct {
	Code  string `json:"code"`
	Cname string `json:"c_name"`
	Name  string `json:"name"`
}

type SME struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type GEM struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ST struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type HS300 struct {
	Code   string       `json:"code"`
	Name   string       `json:"name"`
	Weight float32      `json:"weight"`
	Date   library.Date `json:"date"`
}

type SZ50 struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ZZ500 struct {
	Code   string       `json:"code"`
	Name   string       `json:"name"`
	Weight float32      `json:"weight"`
	Date   library.Date `json:"date"`
}

type Suspended struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	ODate library.Date `json:"oDate"`
	TDate library.Date `json:"tDate"`
}

type Terminated struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	ODate library.Date `json:"oDate"`
	TDate library.Date `json:"tDate"`
}
