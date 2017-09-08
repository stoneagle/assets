package model

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

var classifyDetailSchema = `
CREATE TABLE classify (
   id int(16) unsigned NOT NULL AUTO_INCREMENT,
   type tinyint(4) unsigned NOT NULL,
   tag char(128) NOT NULL,
   name char(128) NOT NULL,
   code char(32) NOT NULL,
   PRIMARY KEY (id)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8
`

type ClassifyDetail struct {
	ID         int    `db:"id"`
	ClassifyID int    `db:"classify_id"`
	Name       string `db:"name"`
	Code       string `db:"code"`
}

type ClassifyDetailAPI struct {
	Model *sqlx.DB
}

const (
	TableClassifyDetail = "classify_detail"
)

func NewClassifyDetailModel(DB *sqlx.DB) *ClassifyDetailAPI {
	return &ClassifyDetailAPI{
		Model: DB,
	}
}

func (api ClassifyDetailAPI) Insert(param *ClassifyDetail) int64 {
	stmt, err := api.Model.Prepare("INSERT " + TableClassifyDetail + " SET classify_id=?,name=?,code=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(param.ClassifyID, param.Name, param.Code)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func (api ClassifyDetailAPI) Update(param *ClassifyDetail) int64 {
	stmt, err := api.Model.Prepare("update " + TableClassifyDetail + " set classify_id=?,name=?,code=? where id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(param.ClassifyID, param.Name, param.Code, param.ID)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return affect
}

func (api ClassifyDetailAPI) Delete(param *ClassifyDetail) int64 {
	stmt, err := api.Model.Prepare("delete from " + TableClassifyDetail + " where id=?")
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

func (api ClassifyDetailAPI) GetNumByClassify(cid, limit int) map[string]ClassifyDetail {
	ret := []ClassifyDetail{}
	limitStr := strconv.Itoa(limit)
	cidStr := strconv.Itoa(cid)
	err := api.Model.Select(&ret, "select * from "+TableClassifyDetail+" where classify_id = "+cidStr+" limit "+limitStr)
	if err != nil {
		panic(err)
	}
	result := map[string]ClassifyDetail{}
	for _, v := range ret {
		result[v.Code] = v
	}
	return result
}
