package model

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

var classifyRetrySchema = `
CREATE TABLE classify_retry (
	id int(11) unsigned NOT NULL AUTO_INCREMENT,
	type tinyint(4) NOT NULL,
	status tinyint(4) NOT NULL,
	code text NOT NULL,
	PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

type ClassifyRetry struct {
	ID     int    `db:"id"`
	Type   int    `db:"type"`
	Status int    `db:"status"`
	Code   string `db:"code"`
}

type ClassifyRetryAPI struct {
	Model *sqlx.DB
}

const (
	TableClassifyRetry   = "classify_retry"
	ClassifyRetryWait    = 1
	ClassifyRetryProcess = 2
	ClassifyRetryEnd     = 3
)

func NewClassifyRetryModel(DB *sqlx.DB) *ClassifyRetryAPI {
	return &ClassifyRetryAPI{
		Model: DB,
	}
}

func (api ClassifyRetryAPI) Insert(param *ClassifyRetry) int64 {
	stmt, err := api.Model.Prepare("INSERT " + TableClassifyRetry + " SET status=?,code=?,type=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(param.Status, param.Code, param.Type)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func (api ClassifyRetryAPI) Update(param *ClassifyRetry) int64 {
	stmt, err := api.Model.Prepare("update " + TableClassifyRetry + " set status=?,code=?,type=? where id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(param.Status, param.Code, param.ID, param.Type)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return affect
}

func (api ClassifyRetryAPI) Delete(param *ClassifyRetry) int64 {
	stmt, err := api.Model.Prepare("delete from " + TableClassifyRetry + " where id=?")
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

func (api ClassifyRetryAPI) GetNumUnend(limit int, classifyType int) map[string]ClassifyRetry {
	ret := []ClassifyRetry{}
	limitStr := strconv.Itoa(limit)
	statusStr := strconv.Itoa(ClassifyRetryEnd)
	typeStr := strconv.Itoa(classifyType)
	err := api.Model.Select(&ret, "select * from "+TableClassifyRetry+" where status != "+statusStr+" and type = "+typeStr+" limit "+limitStr)
	if err != nil {
		panic(err)
	}
	result := map[string]ClassifyRetry{}
	for _, v := range ret {
		result[v.Code] = v
	}
	return result
}
