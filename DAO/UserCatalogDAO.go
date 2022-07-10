package DAO

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"trino.com/trino-connectors/data"
	"trino.com/trino-connectors/util/log"
)

const UserCatalogDB = "UserCatalog"

func GetDataByCatalogIds(catelogIds []string) ([]data.UserCatalogData, error) {
	d := []data.UserCatalogData{{
		UserId:    0,
		CatalogId: "",
	}}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %v WHERE CatalogId in (?)", UserCatalogDB), catelogIds)
	if err != nil {
		fmt.Printf("[GetDataByCatalogIds] [sqlx.In] error: %s", err)
		log.Logger().Errorf("[GetDataByCatalogIds] [sqlx.In] error: %s", err)
		return d, err
	}
	_, err = SelectData(&d, query, args...)
	if err != nil {
		fmt.Printf("[GetDataByCatalogIds] [SelectData] error: %s", err)
		log.Logger().Errorf("[GetDataByCatalogIds] [SelectData] error: %s", err)
		return d, err
	}

	return d, nil
}

func GetDataByUserIds(userIds []int64) ([]data.UserCatalogData, error) {
	d := []data.UserCatalogData{{
		UserId:    0,
		CatalogId: "",
	}}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %v WHERE UserId in (?)", UserCatalogDB), userIds)
	if err != nil {
		fmt.Printf("[GetDataByCatalogIds] [sqlx.In] error: %s", err)
		log.Logger().Errorf("[GetDataByCatalogIds] [sqlx.In] error: %s", err)
		return d, err
	}
	_, err = SelectData(&d, query, args...)
	if err != nil {
		fmt.Printf("[GetDataByCatalogIds] [SelectData] error: %s", err)
		log.Logger().Errorf("[GetDataByCatalogIds] [SelectData] error: %s", err)
		return d, err
	}

	return d, nil
}

func InsertUserCatalog(userIds int64, catalogId string) error {

	err := InsertData(fmt.Sprintf("INSERT INTO %v (UserId, CatalogId) VALUES (?, ?)", UserCatalogDB), userIds, catalogId)
	if err != nil {
		fmt.Printf("[InsertUserCatalog] [InsertData] error: %s", err)
		log.Logger().Errorf("[InsertUserCatalog] [InsertData] error: %s", err)
		return err
	}

	return nil
}

func UpdateUserCatalogStatusByCatalogId(catalogId string, status data.Status) error {

	err := UpdateData(fmt.Sprintf("update %v set Status = ? where Id = ?", UserCatalogDB), status, catalogId)
	if err != nil {
		fmt.Printf("[UpdateUserCatalogStatusByCatalogId] [UpdateData] error: %s", err)
		log.Logger().Errorf("[UpdateUserCatalogStatusByCatalogId] [UpdateData] error: %s", err)
		return err
	}

	return nil
}
