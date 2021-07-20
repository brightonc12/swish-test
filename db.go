package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type MySql struct {
	Client *gorm.DB
}

var (
	sqlOnce        sync.Once
	clientInstance *MySql
)

func GetMySql() *MySql {
	return clientInstance
}

func ConnectDB() (*MySql, error) {
	var err error

	sqlOnce.Do(func() {

		dsn := "root:root@tcp(127.0.0.1:3306)/swish?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}

		log.Print("MySql connected successfully")
		clientInstance = &MySql{Client: db}
	})

	return clientInstance, err
}
