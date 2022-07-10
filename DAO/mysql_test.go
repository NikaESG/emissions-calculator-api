package DAO

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gotest.tools/v3/assert"
	"testing"
	"time"
	"trino.com/trino-connectors/data"
)

func TestInitMysqlConfig(t *testing.T) {
	err := InitMysqlConfig()
	assert.NilError(t, err)
}

func TestInsertData(t *testing.T) {
	err := InitMysqlConfig()
	err = InsertData("INSERT INTO User (UserId, Username, CompanyName, Status, LatestLoginTime) VALUES (?, ?, ?, ?, ?)", "222251", "Test351", "TestCompany", "SUCCESS", time.Now())
	assert.NilError(t, err)
}

func TestUpdateData(t *testing.T) {
	err := InitMysqlConfig()
	err = UpdateData("update User set Username = ? where UserId = ?", "Test38", "22225")
	assert.NilError(t, err)
}

func TestDeleteData(t *testing.T) {
	err := InitMysqlConfig()
	err = DeleteData("delete from User where UserId = ?", "22225")
	assert.NilError(t, err)
}

func TestSelectData(t *testing.T) {
	err := InitMysqlConfig()
	d := []data.UserData{{
		Id:              0,
		UserId:          "",
		Username:        "",
		Password:        sql.NullString{},
		CompanyName:     sql.NullString{},
		Status:          "",
		CreateTime:      time.Time{},
		UpdateTime:      time.Time{},
		LatestLoginTime: time.Time{},
		Etx:             sql.NullString{},
	}}
	e, err := SelectData(&d, "SELECT * FROM User WHERE UserId= ?", "2222")
	fmt.Println(e)
	assert.NilError(t, err)
}

func TestSelectInData(t *testing.T) {
	err := InitMysqlConfig()
	d := []data.UserData{{
		Id:              0,
		UserId:          "",
		Username:        "",
		Password:        sql.NullString{},
		CompanyName:     sql.NullString{},
		Status:          "",
		CreateTime:      time.Time{},
		UpdateTime:      time.Time{},
		LatestLoginTime: time.Time{},
		Etx:             sql.NullString{},
	}}

	query, args, err := sqlx.In("SELECT * FROM User WHERE UserId in (?)", []string{"2222"})

	e, err := SelectData(&d, query, args...)
	fmt.Println(e, d)
	assert.NilError(t, err)
}
