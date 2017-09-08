package model

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

var classifySchema = ``

const (
	TableClassify        = "classify"
	ClassifyTypeIndustry = 1
	ClassifyTypeConcept  = 2
)

type Classify struct {
	ID   int    `db:"id"`
	Type int    `db:"type"`
	Tag  string `db:"tag"`
	Name string `db:"name"`
}

type ClassifyAPI struct {
	Model *sqlx.DB
}

func NewClassifyModel(DB *sqlx.DB) *ClassifyAPI {
	return &ClassifyAPI{
		Model: DB,
	}
}

func (api ClassifyAPI) Insert(param *Classify) int64 {
	stmt, err := api.Model.Prepare("INSERT " + TableClassify + " SET type=?,tag=?,name=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(param.Type, param.Tag, param.Name)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func (api ClassifyAPI) Update(param *Classify) int64 {
	stmt, err := api.Model.Prepare("update " + TableClassify + " set type=?,name=?,tag=? where id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(param.Type, param.Name, param.Tag, param.ID)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return affect
}

func (api ClassifyAPI) Delete(param *Classify) int64 {
	stmt, err := api.Model.Prepare("delete from " + TableClassify + " where id=?")
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

func (api ClassifyAPI) GetNumByType(tid, limit int) map[string]Classify {
	ret := []Classify{}
	limitStr := strconv.Itoa(limit)
	tidStr := strconv.Itoa(tid)
	err := api.Model.Select(&ret, "select * from "+TableClassify+" where type = "+tidStr+" limit "+limitStr)
	if err != nil {
		panic(err)
	}
	result := map[string]Classify{}
	for _, v := range ret {
		result[v.Tag] = v
	}
	return result
}
