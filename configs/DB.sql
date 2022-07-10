CREATE SCHEMA IF NOT EXISTS `db_trino_connector` DEFAULT CHARACTER SET utf8 ;

CREATE TABLE IF NOT EXISTS `db_trino_connector`.`User` (
                                             `Id` INT NOT NULL AUTO_INCREMENT,
                                             `UserId` INT NOT NULL,
                                             `Username` VARCHAR(125) NOT NULL,
                                             `Password` VARCHAR(125) NULL,
                                             `CompanyName` VARCHAR(125) NULL,
                                             `Status` VARCHAR(45) NOT NULL,
                                             `CreateTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             `UpdateTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                             `LatestLoginTime` DATETIME NOT NULL,
                                             `Etx` TEXT NULL,
                                             PRIMARY KEY (`id`),
                                             UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
                                             UNIQUE INDEX `UserId_UNIQUE` (`UserId` ASC) VISIBLE,
                                             UNIQUE INDEX `Username_UNIQUE` (`Username` ASC) VISIBLE);


INSERT INTO `db_trino_connector`.`User` (`UserId`, `Username`, `CompanyName`, `Status`, `LatestLoginTime`) VALUES ('22222', 'Test2', 'TestCompany', 'SUCCESS', now(3));

CREATE TABLE IF NOT EXISTS `db_trino_connector`.`NameSpace` (
                                                  `NsId` INT NOT NULL AUTO_INCREMENT,
                                                  `NsName` VARCHAR(125) NOT NULL,
                                                  `NsDescription` VARCHAR(125) NULL,
                                                  `Status` VARCHAR(45) NOT NULL,
                                                  `CreateTime` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
                                                  `UpdateTime` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                  PRIMARY KEY (`nsId`),
                                                  UNIQUE INDEX `idNs_UNIQUE` (`nsId` ASC) VISIBLE,
                                                  UNIQUE INDEX `nsName_UNIQUE` (`nsName` ASC) VISIBLE);

INSERT INTO `db_trino_connector`.`NameSpace` (`NsName`, `NsDescription`, `Status`) VALUES ('default', 'Init', 'SUCCESS');

CREATE TABLE IF NOT EXISTS `db_trino_connector`.`UserNameSpace` (
                                                      `UserId` INT NOT NULL,
                                                      `NsId` INT NOT NULL,
                                                      PRIMARY KEY (`UserId`, `NsId`),
                                                      INDEX `NsId_idx` (`NsId` ASC) VISIBLE,
                                                      CONSTRAINT `UserId`
                                                          FOREIGN KEY (`UserId`)
                                                              REFERENCES `db_trino_connector`.`User` (`UserId`)
                                                              ON DELETE NO ACTION
                                                              ON UPDATE NO ACTION,
                                                      CONSTRAINT `NsId`
                                                          FOREIGN KEY (`NsId`)
                                                              REFERENCES `db_trino_connector`.`NameSpace` (`NsId`)
                                                              ON DELETE NO ACTION
                                                              ON UPDATE NO ACTION);

INSERT INTO `db_trino_connector`.`UserNameSpace` (`UserId`, `NsId`) VALUES ('22222', '1');

CREATE TABLE IF NOT EXISTS `db_trino_connector`.`Catalog` (
                                                `Id` INT NOT NULL AUTO_INCREMENT,
                                                `NsId` INT NOT NULL DEFAULT 1,
                                                `Status` VARCHAR(45) NULL,
                                                `CreateTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                `UpdateTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                PRIMARY KEY (`Id`));

CREATE TABLE IF NOT EXISTS `db_trino_connector`.`UserCatalog` (
                                                    `UserId` INT NOT NULL,
                                                    `CatalogId` INT NOT NULL,
                                                    PRIMARY KEY (`UserId`, `CatalogId`),
                                                    INDEX `CatelogId_idx` (`CatalogId` ASC) VISIBLE);

INSERT INTO `db_trino_connector`.`UserCatalog` (`UserId`, `CatalogId`) VALUES ('22222', '1');
