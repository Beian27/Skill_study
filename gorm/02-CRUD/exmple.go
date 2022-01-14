package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Userinfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
	gorm.Model
}

func main() {
	url := "root:Zhy5224146.+@tcp(127.0.0.1:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	userinfo := &Userinfo{}
	db.AutoMigrate(userinfo)

	u1 := Userinfo{Name: "ZHY1", Gender: "男", Hobby: "Ball"}
	u2 := Userinfo{Name: "ZHY2", Gender: "男", Hobby: "Ball"}
	db.Create(&u1)
	db.Create(&u2)

	db.Find(userinfo)
	fmt.Println(userinfo)

	db.Find(userinfo, "name=?", "ZHY1")
	fmt.Println(userinfo)

	db.Model(userinfo).Where("name=?", "ZHY1").Update("hobby", "读书")
	db.Find(userinfo, "name=?", "ZHY1")
	fmt.Println(userinfo)

	take := db.Take(userinfo).Value
	last := db.Last(userinfo).Value
	find := db.Find(userinfo).Value
	first := db.First(userinfo, 10).Value
	fmt.Println("---------------------------")
	fmt.Println(take)
	fmt.Println("---------------------------")
	fmt.Println(last)
	fmt.Println("---------------------------")
	fmt.Println(find)
	fmt.Println("---------------------------")
	fmt.Println(first)
	//db.Delete(&Userinfo{})

	db.Unscoped().Where("ID > ?", 73).Delete(userinfo)
}
