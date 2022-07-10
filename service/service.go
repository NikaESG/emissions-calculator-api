package service

import (
	"errors"
	"fmt"
	"strings"
	"trino.com/trino-connectors/DAO"
	"trino.com/trino-connectors/util"

	"trino.com/trino-connectors/data"
)

func ListCatalogs() (catalogs map[string]data.ConfigStruct, err error) {

	// TODO: hide our default connections like tpch and tpcds
	// solution 1: store these information in k8s secret,
	// but make sure that secret is not modifiable.
	// solution 2: store these information in a database.

	trinoCatalogDataMap, err := ListTrinoCatalogData()
	if err != nil {
		return nil, err
	}

	configMapData := make(map[string]data.ConfigMapData)
	for i, v := range trinoCatalogDataMap {
		detail := data.ConfigMapData{}
		d := strings.Split(v, "\n")
		for _, dv := range d {
			index := strings.Index(dv, "=")
			detail[dv[:index]] = dv[index+1:]
		}
		configMapData[i] = detail
	}

	catalogs = make(map[string]data.ConfigStruct)

	for trinoCatalogName, trinoCatalogData := range configMapData {
		catalogs[trinoCatalogName] = data.ConfigStruct{
			Params: trinoCatalogData,
		}
	}

	return catalogs, nil
}

func GetData(table string, field string) ([]string, error) {
	t, e := DAO.GetDatabaseTablePostgres(table)
	if e != nil {
		return []string{}, e
	}
	tId := t.Id
	f, e := DAO.GetDatabaseFieldPostgres(tId)

	if e != nil {
		return []string{}, e
	}

	name := ""
	var fId int64
	for _, v := range f {
		if v.Name == field {
			name = v.Name
			fId = v.Id
			break
		}
	}
	if name == "" {
		return []string{}, nil
	}

	return DAO.GetDatabaseDataPostgres(tId, fId)
}

func AddCatalog(catalogId string, data map[string]string) error {
	// verify if catalogId is distinct
	trinoCatalogDataMap, err := ListTrinoCatalogData()
	if err != nil {
		return err
	}
	for k, _ := range trinoCatalogDataMap {
		if strings.Split(k, ",")[0] == catalogId {
			return errors.New("catalogId '" + catalogId + "' already exists.")
		}
	}

	// TODO: validate all fields are present for each of these connector types

	var config string
	for k, v := range data {
		switch k {
		case "credentials-path":
			config += k + "=/etc/trino/secret/google-sheets-credentials.json\n"
			// TODO: add google sheets credentials as new k8s secret and mount it to trino pods
			// if err := AddGoogleSheetCredentialsSecret(v); err != nil {
			// 	return err
			// }
		default:
			config += k + "=" + v + "\n"
		}
	}
	fmt.Println("Add new catalog: " + catalogId)
	fmt.Println(config)

	catalogId = fmt.Sprintf("%v.properties", catalogId)
	dataMap, err := UpdateTrinoCatalogData(catalogId, data, trinoCatalogDataMap)
	if err != nil {
		return err
	}

	// verify if catalog is successfully added via k8s Patch API
	isPatched := false
	for k, _ := range dataMap[TrinoConfigMapName] {
		if k == catalogId {
			isPatched = true
			break
		}
	}
	if !isPatched {
		return errors.New("catalog '" + catalogId + "' is NOT added succesfully.")
	}

	return nil
}

func UpdateCatalog(catalogId string, data map[string]string) error {
	// verify if catalogId is distinct
	trinoCatalogDataMap, err := ListTrinoCatalogData()
	if err != nil {
		return err
	}

	if _, ok := util.Exits(trinoCatalogDataMap, catalogId, ".", 0); !ok {
		return errors.New("catalogId '" + catalogId + "' do not exist. ")
	}

	// TODO: validate all fields are present for each of these connector types

	var config string
	for k, v := range data {
		switch k {
		case "credentials-path":
			config += k + "=/etc/trino/secret/google-sheets-credentials.json\n"
			// TODO: add google sheets credentials as new k8s secret and mount it to trino pods
			// if err := AddGoogleSheetCredentialsSecret(v); err != nil {
			// 	return err
			// }
		default:
			config += k + "=" + v + "\n"
		}
	}

	catalogId = fmt.Sprintf("%v.properties", catalogId)
	_, err = UpdateTrinoCatalogData(catalogId, data, trinoCatalogDataMap)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCatalog(catalogId string) error {
	// verify if catalogId is distinct
	trinoCatalogDataMap, err := ListTrinoCatalogData()
	if err != nil {
		return err
	}

	if _, ok := util.Exits(trinoCatalogDataMap, catalogId, ".", 0); !ok {
		return errors.New("catalogId '" + catalogId + "' do not exist. ")
	}

	catalogId = fmt.Sprintf("%v.properties", catalogId)
	_, err = DeleteTrinoCatalogData(catalogId, trinoCatalogDataMap)
	if err != nil {
		return err
	}

	return nil
}

func TestConnection(catalogId string) error {
	// verify if catalogId is distinct
	//trinoCatalogDataMap, err := ListTrinoCatalogData()
	//if err != nil {
	//	return err
	//}
	//
	//if _, ok := util.Exits(trinoCatalogDataMap, catalogId, ".", 0); !ok {
	//	return errors.New("catalogId '" + catalogId + "' do not exist. ")
	//}
	_, err := ExecPod(TrinoCoordinator, fmt.Sprintf("trino --execute=\"SHOW SCHEMAS FROM %v\"", catalogId))
	if err != nil {
		return err
	}

	return nil
}

// func GetCatalogIdsByUserId(userId int64) (d []string, err error) {
// 	userCatalog, err := DAO.GetDataByUserIds([]int64{userId})
// 	if err != nil {
// 		fmt.Printf("[GetCatalogIdsByUserId] [GetDataByCatalogIds] error: %s", err)
// 		log.Logger().Errorf("[GetCatalogIdsByUserId] [GetDataByCatalogIds] error: %s", err)
// 		return d, err
// 	}
// 	for _, v := range userCatalog {
// 		d = append(d, v.CatalogId)
// 	}
// 	return d, nil
// }

// func GetCatalog(catalogId string) (d data.ConfigStruct, err error) {
// 	userCatalog, err := DAO.GetDataByCatalogIds([]string{catalogId})
// 	if err != nil {
// 		fmt.Printf("[GetCatalog] [GetDataByCatalogIds] error: %s", err)
// 		log.Logger().Errorf("[GetCatalog] [GetDataByCatalogIds] error: %s", err)
// 		return d, err
// 	}

// 	if len(userCatalog) == 0 {
// 		return d, err
// 	}

// 	d.UserId = userCatalog[0].UserId

// 	userData, err := DAO.GetUsers([]int64{d.UserId})
// 	if err != nil {
// 		fmt.Printf("[GetCatalog] [GetUsers] error: %s", err)
// 		log.Logger().Errorf("[GetCatalog] [GetUsers] error: %s", err)
// 		return d, err
// 	}

// 	if len(userData) == 0 {
// 		return d, err
// 	}

// 	d.UserName = userData[0].Username

// 	catelogData, err := DAO.GetCatalogs([]string{catalogId})
// 	if err != nil {
// 		fmt.Printf("[GetCatalog] [GetCatalogs] error: %s", err)
// 		log.Logger().Errorf("[GetCatalog] [GetCatalogs] error: %s", err)
// 		return d, err
// 	}

// 	if len(catelogData) == 0 {
// 		return d, err
// 	}

// 	nsId := catelogData[0].NsId

// 	namespaceData, err := DAO.GetNamespaces([]int64{nsId})
// 	if err != nil {
// 		fmt.Printf("[GetCatalog] [GetNamespaces] error: %s", err)
// 		log.Logger().Errorf("[GetCatalog] [GetNamespaces] error: %s", err)
// 		return d, err
// 	}

// 	if len(namespaceData) == 0 {
// 		return d, err
// 	}

// 	nsName := namespaceData[0].NsName

// 	// TODO: Get Catalog

// 	config, err := GetConfigmap(nsName, catalogId)
// 	if err != nil {
// 		fmt.Printf("[GetCatalog] [GetCatalogmap] error: %s", err)
// 		log.Logger().Errorf("[GetCatalog] [GetCatalogmap] error: %s", err)
// 		return d, err
// 	}
// 	d.Config = config

// 	return d, nil
// }

// func AddCatalog(userId int64, catalogId string, namespace string, d map[string]string) error {

// 	nsData, err := DAO.GetNamespacesByName([]string{namespace})
// 	if err != nil {
// 		fmt.Printf("[AddCatalog] [GetNamespacesByName] error: %s", err)
// 		log.Logger().Errorf("[AddCatalog] [GetNamespacesByName] error: %s", err)
// 		return err
// 	}

// 	err = DAO.InsertCatalogStatus(catalogId, nsData[0].NsId, data.PROCESSING)
// 	if err != nil {
// 		fmt.Printf("[AddCatalog] [InsertCatalogStatus PROCESSING] error: %s", err)
// 		log.Logger().Errorf("[AddCatalog] [InsertCatalogStatus PROCESSING] error: %s", err)
// 		return err
// 	}

// 	// TODO: Add catalog

// 	path := "./testFiles/config.yaml"
// 	err = util.WriteYaml(path, d)
// 	if err != nil {
// 		fmt.Printf("[AddCatalog] [util.WriteYaml] error: %s", err)
// 		log.Logger().Errorf("[AddCatalog] [util.WriteYaml] error: %s", err)
// 		return err
// 	}

// 	configName := fmt.Sprintf("%v-%v-config.yaml", userId, catalogId)

// 	err = AddConfigmap(namespace, configName)
// 	if err != nil {
// 		fmt.Printf("[AddCatalog] [AddConfigmap] error: %s", err)
// 		log.Logger().Errorf("[AddCatalog] [AddConfigmap] error: %s", err)
// 		return err
// 	}

// 	// Update DB Result

// 	err = DAO.UpdateCatelogStatus(catalogId, data.SUCCESS)
// 	if err != nil {
// 		fmt.Printf("[AddCatalog] [UpdateCatelogStatus SUCCESS] error: %s", err)
// 		log.Logger().Errorf("[AddCatalog] [UpdateCatelogStatus SUCCESS] error: %s", err)
// 		return err
// 	}

// 	return nil
// }

//func UpdateCatalog(catalogId string, d map[string]string) error {
//	err := DAO.UpdateCatelogStatus(catalogId, data.PROCESSING)
//
//	if err != nil {
//		fmt.Printf("[UpdateCatalog] [UpdateCatelogStatus PROCESSING] error: %s", err)
//		log.Logger().Errorf("[UpdateCatalog] [UpdateCatelogStatus PROCESSING] error: %s", err)
//		return err
//	}
//
//	// TODO: Update catalog
//
//	err = DAO.UpdateCatelogStatus(catalogId, data.SUCCESS)
//
//	if err != nil {
//		fmt.Printf("[UpdateCatalog] [UpdateCatelogStatus SUCCESS] error: %s", err)
//		log.Logger().Errorf("[UpdateCatalog] [UpdateCatelogStatus SUCCESS] error: %s", err)
//		return err
//	}
//	return nil
//}

//func DeleteCatalog(catalogId string) error {
//	err := DAO.UpdateCatelogStatus(catalogId, data.PROCESSING)
//
//	if err != nil {
//		fmt.Printf("[DeleteCatalog] [UpdateCatelogStatus PROCESSING] error: %s", err)
//		log.Logger().Errorf("[DeleteCatalog] [UpdateCatelogStatus PROCESSING] error: %s", err)
//		return err
//	}
//
//	// TODO: Delete catalog
//
//	// TODO: transaction should be used in the future
//
//	err = DAO.UpdateCatelogStatus(catalogId, data.DELETED)
//
//	if err != nil {
//		fmt.Printf("[DeleteCatalog] [UpdateCatelogStatus DELETED] error: %s", err)
//		log.Logger().Errorf("[DeleteCatalog] [UpdateCatelogStatus DELETED] error: %s", err)
//		return err
//	}
//
//	err = DAO.UpdateUserCatalogStatusByCatalogId(catalogId, data.DELETED)
//	if err != nil {
//		fmt.Printf("[DeleteCatalog] [UpdateUserCatalogStatusByCatalogId DELETED] error: %s", err)
//		log.Logger().Errorf("[DeleteCatalog] [UpdateUserCatalogStatusByCatalogId DELETED] error: %s", err)
//		return err
//	}
//
//	return nil
//}
