package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	dsn := "root:Zhy5224146.+@tcp(127.0.0.1:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}
