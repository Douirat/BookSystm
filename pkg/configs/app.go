package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to database:
var DataBaseConncetion = func() (*gorm.DB, error) {
	dsn := "root:zone2023@tcp(127.0.0.1:3306)/BookSystem?charset=utf8mb4&pasrseTime=True&loc=Local"
	dataBase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dataBase, nil
}