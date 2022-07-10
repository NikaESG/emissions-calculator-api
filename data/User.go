package data

import (
	"database/sql"
	"time"
)

type Status string

const (
	SUCCESS    Status = "SUCCESS"
	FAIL       Status = "FAIL"
	PROCESSING Status = "PROCESSING"
	DELETED    Status = "DELETED"
	BLOCKED    Status = "BLOCKED"
	PENDING    Status = "PENDING"
)

type UserData struct {
	Id              int64          `db:"Id"`
	UserId          string         `db:"UserId"`
	Username        string         `db:"Username"`
	Password        sql.NullString `db:"Password,omitempty"`
	CompanyName     sql.NullString `db:"CompanyName,omitempty"`
	Status          Status         `db:"Status,omitempty"`
	CreateTime      time.Time      `db:"CreateTime"`
	UpdateTime      time.Time      `db:"UpdateTime"`
	LatestLoginTime time.Time      `db:"LatestLoginTime,omitempty"`
	Etx             sql.NullString `db:"Etx,omitempty"`
}

type CatalogData struct {
	Id         string    `db:"Id"`
	NsId       int64     `db:"NsId"`
	Status     Status    `db:"Status,omitempty"`
	CreateTime time.Time `db:"CreateTime"`
	UpdateTime time.Time `db:"UpdateTime"`
}

type NamespaceData struct {
	NsId          int64     `db:"NsId"`
	NsName        string    `db:"NsName"`
	NsDescription string    `db:"NsDescription"`
	Status        Status    `db:"Status,omitempty"`
	CreateTime    time.Time `db:"CreateTime"`
	UpdateTime    time.Time `db:"UpdateTime"`
}

type UserCatalogData struct {
	UserId    int64  `db:"UserId"`
	CatalogId string `db:"CatalogId"`
}
