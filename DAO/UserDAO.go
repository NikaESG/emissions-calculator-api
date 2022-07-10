package DAO

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"trino.com/trino-connectors/data"
	"trino.com/trino-connectors/util/log"
)

const UserDB = "User"

func GetUsers(userIds []int64) ([]data.UserData, error) {
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

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %v WHERE UserId in (?)", UserDB), userIds)
	if err != nil {
		fmt.Printf("[GetCatalogs] [sqlx.In] error: %s", err)
		log.Logger().Errorf("[GetCatalogs] [sqlx.In] error: %s", err)
		return d, err
	}
	_, err = SelectData(&d, query, args...)
	if err != nil {
		fmt.Printf("[GetCatalogs] [SelectData] error: %s", err)
		log.Logger().Errorf("[GetCatalogs] [SelectData] error: %s", err)
		return d, err
	}

	return d, nil
}
