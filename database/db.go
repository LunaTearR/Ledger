package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/LunaTearR/Ledgar/models"
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

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Ledger{},
		&models.AdminLog{},
	)
}