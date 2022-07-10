package DAO

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
	"strings"
	"trino.com/trino-connectors/data"
)

var Postgres *pg.DB

const (
	host     = "localhost"
	port     = 25432
	user     = "baserow"
	password = "baserow"
	dbname   = "baserow"
)

func InitialPostgres() error {

	// open database
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		User:     user,
		Password: password,
		Database: dbname,
	})

	orm.SetTableNameInflector(func(s string) string {
		return "database_table" + s
	})

	Postgres = db

	fmt.Println("Connected")
	return nil
}

func GetDatabaseTablePostgres(table string) (t data.Table, e error) {

	err := Postgres.Model(&t).
		Where("name = ?", table).
		Select()
	if err != nil {
		return t, err
	}

	return
}

func GetDatabaseFieldPostgres(tableId int64) (t []data.Field, e error) {

	e = Postgres.Model(&t).
		Where("table_id = ?", tableId).
		Select()
	if e != nil {
		return t, e
	}

	return
}

func GetDatabaseDataPostgres(tableId int64, fieldId int64) (t []string, e error) {

	table := fmt.Sprintf("database_table_%v", tableId)
	field := fmt.Sprintf("field_%v", fieldId)

	//type MyType struct {
	//	field string
	//}

	var data []map[string]interface{}

	_, err := Postgres.Query(&data, fmt.Sprintf("SELECT %v FROM %v", field, table))
	if err != nil {
		return t, err
	}

	for _, v := range data {
		d := fmt.Sprint(v[field])
		index := strings.Index(d, " ")
		if index == -1 {
			t = append(t, d)
			continue
		}
		t = append(t, d[index+1:len(d)-1])
	}

	return t, nil
}
