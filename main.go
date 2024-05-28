package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("MYSQL_CONN_STRING")
	//dsn := "root:18032002@tcp(127.0.0.1:3306)/socialheath?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("nooke")
	}
	println("oke")

	// Automatically migrate the schema
	db.AutoMigrate(&User{})

	// Use the database connection for CRUD operations
	db.Create(&User{Name: "John", Email: "john@example.com"})
}
