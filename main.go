package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	Id      int    `json:"id" binding :"required"`
	Name    string `json:"fullname" binding :"required"`
	Age     string `json:"age" binding :"required"`
	Address string `json:"address" binding :"required"`
	Phone   string `json:"phone" binding :"required"`
}

func main() {
	var DB *gorm.DB
	
	DB, err := gorm.Open("mysql", "root:rafi123@tcp(127.0.0.1:3306)/dmc_production")
	fmt.Println(DB)
	if err != nil {
		panic("failed to connect database")
	}

	if err == nil {

		fmt.Println("success to connect database")
	}

	defer DB.Close()
	fmt.Println(&Student{})
	// Migration()
}

func Migration() {
	var db *gorm.DB
	model := &Student{}
	db.AutoMigrate(model)
}
