package DAO

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestGetDatabaseTablePostgres(t *testing.T) {
	InitialPostgres()
	tt, e := GetDatabaseTablePostgres("Projects")
	assert.Equal(t, nil, e)
	fmt.Println(tt)
	fmt.Println(tt)
	fmt.Println(tt.Id)
}

func TestGetDatabaseFieldPostgres(t *testing.T) {
	InitialPostgres()
	tt, e := GetDatabaseFieldPostgres(21)
	assert.Equal(t, nil, e)
	fmt.Println(tt)
	fmt.Println(tt)
}

func TestGetDatabaseDataPostgres(t *testing.T) {
	InitialPostgres()
	tt, e := GetDatabaseDataPostgres(21, 419)
	assert.Equal(t, nil, e)
	fmt.Println(tt)
	fmt.Println(tt)
}
