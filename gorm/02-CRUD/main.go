package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type User struct {
	Name     string    `gorm:"name"`
	Age      int       `gorm:"age"`
	Birthday time.Time `gorm:"birthday"`
	gorm.Model
}

func connMysql() *gorm.DB {
	dsn := "root:Zhy5224146.+@tcp(127.0.0.1:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := connMysql()
	db.AutoMigrate(&User{})
	//db.Create()
	//user := User{
	//	Name:     "ZHY",
	//	Age:      24,
	//	Birthday: time.Now(),
	//}
	//create := user.Create(db)
	//fmt.Println(create)
	//
	//----------------------
	//var users = []User{
	//	{
	//		Name:     "ZYC",
	//		Age:      18,
	//		Birthday: time.Now(),
	//	},
	//	{
	//		Name:     "ZZM",
	//		Age:      19,
	//		Birthday: time.Now(),
	//	},
	//}
	//creates := db.Create(&users)
	//fmt.Println(creates)

	// -----------------------------
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "jinzhu_1", "Age": 18},
	//	{"Name": "jinzhu_2", "Age": 20},
	//})
	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "jinzhu", "Age": 18,
	//})

	//--------------------------------
	//var users = []User{{Name: "jinzhu1", Age: 11, Birthday: time.Now()}, {Name: "jinzhu2", Age: 12, Birthday: time.Now()}, {Name: "jinzhu3", Age: 13, Birthday: time.Now()}}
	//for _, user := range users {
	//	db.Create(&user)
	//}

	//--------------------------------
	//result := map[string]interface{}{}
	//value := db.Model(&User{}).First(&result)
	//fmt.Println(value)

	//--------------------------------
	value := db.First(&User{}, 10).Value
	fmt.Println(value)
}

func (u *User) Create(db *gorm.DB) *gorm.DB {
	create := db.Select("Name", "Age", "CreatedAt").Create(&u)
	return create
}
