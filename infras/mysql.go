package infras

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlSession() (*gorm.DB, error) {

	dbUsername := "root"
	dbPass := "12345678"
	dbName := "funfun"
	var addr = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUsername, dbPass, dbName)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
