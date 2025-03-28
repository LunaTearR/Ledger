package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbUser, dbName, dbPassword, dbHost string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=admin password=password dbname=ledger port=9920 sslmode=disable TimeZone=Asia/Bankok",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}