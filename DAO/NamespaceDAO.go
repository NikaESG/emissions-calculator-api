package DAO

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"trino.com/trino-connectors/data"
	"trino.com/trino-connectors/util/log"
)

const NamespaceDB = "Catalog"

func GetNamespaces(nsId []int64) ([]data.NamespaceData, error) {
	d := []data.NamespaceData{{
		NsId:          0,
		NsName:        "",
		NsDescription: "",
		Status:        "",
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
	}}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %v WHERE NsId in (?)", NamespaceDB), nsId)
	if err != nil {
		fmt.Printf("[GetNamespaces] [sqlx.In] error: %s", err)
		log.Logger().Errorf("[GetNamespaces] [sqlx.In] error: %s", err)
		return d, err
	}
	_, err = SelectData(&d, query, args...)
	if err != nil {
		fmt.Printf("[GetNamespaces] [SelectData] error: %s", err)
		log.Logger().Errorf("[GetNamespaces] [SelectData] error: %s", err)
		return d, err
	}

	return d, nil
}

func GetNamespacesByName(nsName []string) ([]data.NamespaceData, error) {
	d := []data.NamespaceData{{
		NsId:          0,
		NsName:        "",
		NsDescription: "",
		Status:        "",
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
	}}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %v WHERE NsName in (?)", NamespaceDB), nsName)
	if err != nil {
		fmt.Printf("[GetNamespaces] [sqlx.In] error: %s", err)
		log.Logger().Errorf("[GetNamespaces] [sqlx.In] error: %s", err)
		return d, err
	}
	_, err = SelectData(&d, query, args...)
	if err != nil {
		fmt.Printf("[GetNamespaces] [SelectData] error: %s", err)
		log.Logger().Errorf("[GetNamespaces] [SelectData] error: %s", err)
		return d, err
	}

	return d, nil
}
