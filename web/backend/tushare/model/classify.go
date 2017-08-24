package model

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE classify (
   id int(16) unsigned NOT NULL AUTO_INCREMENT,
   type tinyint(4) unsigned NOT NULL,
   tag char(128) NOT NULL,
   name char(128) NOT NULL,
   code char(32) NOT NULL,
   PRIMARY KEY (id)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

type Classify struct {
	ID      int    `db:"id"`
	Type    int    `db:"type"`
	Tag     string `db:"tag"`
	TagName string `db:"tag_name"`
	Name    string `db:"name"`
	Code    string `db:"code"`
}

type SqlAPI struct {
	Model *sqlx.DB
}

const (
	TypeIndustry = 1
	TypeConcept  = 2
)

func NewClassifyModel(DB *sqlx.DB) *SqlAPI {
	return &SqlAPI{
		Model: DB,
	}
}

func (api SqlAPI) Insert(param *Classify) int64 {
	stmt, err := api.Model.Prepare("INSERT classify SET type=?,tag=?,name=?,code=?,tag_name=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(param.Type, param.Tag, param.Name, param.Code, param.TagName)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func (api SqlAPI) Update(param *Classify) int64 {
	stmt, err := api.Model.Prepare("update classify set tag=?,name=?,code=?,tag_name=? where id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(param.Tag, param.Name, param.Code, param.TagName, param.ID)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return affect
}

func (api SqlAPI) Delete(param *Classify) int64 {
	stmt, err := api.Model.Prepare("delete from classify where id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(param.ID)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return affect
}

func (api SqlAPI) GetNumByTag(tag string, limit int) map[string]Classify {
	ret := []Classify{}
	limitStr := strconv.Itoa(limit)
	err := api.Model.Select(&ret, "select * from classify where tag = '"+tag+"' limit "+limitStr)
	if err != nil {
		panic(err)
	}
	result := map[string]Classify{}
	for _, v := range ret {
		result[v.Code] = v
	}
	return result
}
