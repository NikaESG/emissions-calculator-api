package DAO

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"trino.com/trino-connectors/util/log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sqlx.DB

var schema = `CREATE SCHEMA IF NOT EXISTS db_trino_connectors DEFAULT CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS db_trino_connectors.User( 
											 UserId INT NOT NULL,
                                             Username VARCHAR(125) NOT NULL,
                                             Password VARCHAR(125) NULL,
                                             CompanyName VARCHAR(125) NULL,
                                             Status VARCHAR(45) NOT NULL,
                                             CreateTime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             UpdateTime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                             LatestLoginTime DATETIME NOT NULL,
                                             Etx TEXT NULL,
                                             PRIMARY KEY (id),
                                             UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
                                             UNIQUE INDEX UserId_UNIQUE (UserId ASC) VISIBLE,
                                             UNIQUE INDEX Username_UNIQUE (Username ASC) VISIBLE);

INSERT IGNORE INTO db_trino_connectors.User (UserId, Username, CompanyName, Status, LatestLoginTime) VALUES ('22222', 'Test2', 'TestCompany', 'SUCCESS', now(3));

CREATE TABLE IF NOT EXISTS db_trino_connectors.NameSpace (
                                                  NsId INT NOT NULL AUTO_INCREMENT,
                                                  NsName VARCHAR(125) NOT NULL,
                                                  NsDescription VARCHAR(125) NULL,
                                                  Status VARCHAR(45) NOT NULL,
                                                  CreateTime DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
                                                  UpdateTime DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                  PRIMARY KEY (nsId),
                                                  UNIQUE INDEX idNs_UNIQUE (nsId ASC) VISIBLE,
                                                  UNIQUE INDEX nsName_UNIQUE (nsName ASC) VISIBLE);

INSERT IGNORE INTO db_trino_connectors.NameSpace (NsName, NsDescription, Status) VALUES ('default', 'Init', 'SUCCESS');

CREATE TABLE IF NOT EXISTS db_trino_connectors.UserNameSpace (
                                                      UserId INT NOT NULL,
                                                      NsId INT NOT NULL,
                                                      PRIMARY KEY (UserId, NsId),
                                                      INDEX NsId_idx (NsId ASC) VISIBLE,
                                                      CONSTRAINT UserId
                                                          FOREIGN KEY (UserId)
                                                              REFERENCES db_trino_connectors.User (UserId)
                                                              ON DELETE NO ACTION
                                                              ON UPDATE NO ACTION,
                                                      CONSTRAINT NsId
                                                          FOREIGN KEY (NsId)
                                                              REFERENCES db_trino_connectors.NameSpace (NsId)
                                                              ON DELETE NO ACTION
                                                              ON UPDATE NO ACTION);

INSERT IGNORE INTO db_trino_connectors.UserNameSpace (UserId, NsId) VALUES ('22222', '1');

CREATE TABLE IF NOT EXISTS db_trino_connectors.Catalog (
                                                Id VARCHAR(125) NOT NULL,
                                                NsId INT NOT NULL DEFAULT 1,
                                                Status VARCHAR(45) NULL,
                                                CreateTime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                UpdateTime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                PRIMARY KEY (Id));

CREATE TABLE IF NOT EXISTS db_trino_connectors.UserCatalog (
                                                    UserId INT NOT NULL,
                                                    CatalogId VARCHAR(125) NOT NULL,
                                                    PRIMARY KEY (UserId, CatalogId),
                                                    INDEX CatelogId_idx (CatalogId ASC) VISIBLE);

INSERT IGNORE INTO db_trino_connectors.UserCatalog (UserId, CatalogId) VALUES ('22222', 'sheets');
`

func InitMysqlConfig() error {
	database, err := sqlx.Open("mysql", "root:hello@tcp(127.0.0.1:3306)/db_trino_connectors?parseTime=true&loc=Asia%2FSingapore&charset=utf8")
	// database, err := sqlx.Open("database type", "user_name:password@tcp(ip_address:port_number)/schema_name")
	if err != nil {
		fmt.Println("[InitConfig] open mysql failed,", err)
		return err
	}

	Db = database
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)

	// TODO: resolve panic: Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'CREATE TABLE IF NOT EXISTS db_trino_connectors.User( UserId INT NOT' at line 3
	Db.MustExec(schema)

	return nil
}

func InsertData(sql string, args ...interface{}) error {

	stmt, err := Db.Prepare(sql)
	if err != nil {
		log.Logger().Errorf("[InsertData] [Db.Prepare] error: %s", err)
		return err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		log.Logger().Errorf("[InsertData] [stmt.Exec] error: %s", err)
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		log.Logger().Errorf("[InsertData] [res.LastInsertId] error: %s", err)
		return err
	}

	return nil
}

func UpdateData(sql string, args ...interface{}) error {
	result, err := Db.Exec(sql, args...)
	if err != nil {
		log.Logger().Errorf("[UpdateData] [Db.Exec] error: %s", err)
		return err
	}
	log.Logger().Debugf("[UpdateData] [Db.Exec] result: %v", result)
	return nil
}

func DeleteData(sql string, args ...interface{}) error {
	result, err := Db.Exec(sql, args...)
	if err != nil {
		log.Logger().Errorf("[DeleteData] [Db.Exec] error: %s", err)
		return err
	}
	log.Logger().Debugf("[DeleteData] [Db.Exec] result: %v", result)
	return nil
}

func SelectData(dest interface{}, query string, args ...interface{}) (interface{}, error) {
	err := Db.Select(dest, query, args...)
	if err != nil {
		log.Logger().Errorf("[DeleteData] [Db.Exec] error: %s", err)
		return dest, err
	}
	log.Logger().Debugf("[DeleteData] [Db.Exec] result: %v", dest)
	return dest, nil
}
