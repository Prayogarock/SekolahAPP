package database

import (
	"fmt"
	"sekolahApp/app/config"
	rolesData "sekolahApp/features/roles/data"
	usersData "sekolahApp/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InittialMigration(db *gorm.DB) {
	// Tambahkan migrasi jika diperlukan
	db.AutoMigrate(&usersData.User{})
	db.AutoMigrate(&rolesData.Role{})
}
