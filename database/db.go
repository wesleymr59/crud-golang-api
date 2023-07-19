package database

import (
	logger "crud-golang-api/log"
	"crud-golang-api/models"
	"crud-golang-api/services"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnDataBase() {

	env_data := services.GetConfig()
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", env_data.Username, env_data.Password, env_data.Host, env_data.Port, env_data.Name)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Printf("Erro ao conectar com banco de dados")
	}
	conn.AutoMigrate(&models.Client{}, &models.Order{})

	DB = conn
}
