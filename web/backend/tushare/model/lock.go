package model

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

const (
	TableLock     = "lock"
	LockStatusOn  = 1
	LockStatusOff = 0

	LockTypeClassify = 1
)

type Lock struct {
	ID     int `db:"id"`
	Type   int `db:"type"`
	Status int `db:"status"`
}

type LockAPI struct {
	Model *sqlx.DB
}

func NewLockModel(DB *sqlx.DB) *LockAPI {
	return &LockAPI{
		Model: DB,
	}
}

func (api LockAPI) Insert(param *Lock) int64 {
	stmt, err := api.Model.Prepare("INSERT `" + TableLock + "` SET `type`=?,`status`=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(param.Type, param.Status)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func (api LockAPI) UpdateByType(t, status int) int64 {
	stmt, err := api.Model.Prepare("update `" + TableLock + "` set `status`=? where `type`=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(status, t)
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return affect
}

func (api LockAPI) Delete(param *Lock) int64 {
	stmt, err := api.Model.Prepare("delete from `" + TableLock + "` where id=?")
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

func (api LockAPI) CheckByType(t int) (ret bool) {
	ret = false
	target := Lock{}
	tStr := strconv.Itoa(t)
	err := api.Model.Get(&target, "select * from `"+TableLock+"` where `type` = '"+tStr+"'")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			newData := &Lock{
				Type:   t,
				Status: LockStatusOff,
			}
			api.Insert(newData)
		} else {
			panic(err)
		}
	} else {
		if target.Status == LockStatusOn {
			ret = true
		}
	}
	return ret
}
