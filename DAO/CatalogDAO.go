package DAO

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"trino.com/trino-connectors/data"
	"trino.com/trino-connectors/util/log"
)

const CatalogDB = "Catalog"

func GetCatalogs(catalogIds []string) ([]data.CatalogData, error) {
	d := []data.CatalogData{{
		Id:         "",
		NsId:       0,
		Status:     "",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %v WHERE Id in (?)", CatalogDB), catalogIds)
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

func UpdateCatelogStatus(catelogId string, status data.Status) error {

	err := UpdateData(fmt.Sprintf("update %v set Status = ? where Id = ?", CatalogDB), status, catelogId)
	if err != nil {
		fmt.Printf("[GetCatalogs] [SelectData] error: %s", err)
		log.Logger().Errorf("[GetCatalogs] [SelectData] error: %s", err)
		return err
	}

	return nil
}

func InsertCatalogStatus(catalogId string, nsId int64, status data.Status) error {

	err := InsertData(fmt.Sprintf("INSERT INTO %v  (Id, NsId, Status) VALUES (?, ?, ?)", CatalogDB), catalogId, nsId, status)
	if err != nil {
		fmt.Printf("[InsertCatalogStatus] [InsertData] error: %s", err)
		log.Logger().Errorf("[InsertCatalogStatus] [InsertData] error: %s", err)
		return err
	}
	return nil
}
