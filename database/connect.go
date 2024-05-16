package database

import (
	"fiber-project/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	var dsn = "root:senha@123@/bd_go?charset=utf8mb4&parseTime=True&loc=Local"
	var v = "Não conseguiu conectar ao banco de dados"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(v)
	}

	connection.AutoMigrate(&models.User{})
	fmt.Println("Conexão ok!") 
	
}
