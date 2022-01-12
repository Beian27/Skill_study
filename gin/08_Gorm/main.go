package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type EXPersonn struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:Zhy5224146.+@/gin_stu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&EXPersonn{})
	//db.Create(&EXPersonn{
	//	Name: "Zhanghy",
	//	Sex:  true,
	//	Age:  1,
	//})
	//
	//db.Create(&EXPersonn{
	//	Name: "Zhanghy1",
	//	Sex:  true,
	//	Age:  2,
	//})

	var p []EXPersonn
	//db.Where("id > ?", 1).Find(&p).Update("age", 25)
	//db.Where("id > ?", 1).Find(&p).Updates(map[string]interface{}{
	//	"Name": "zhy",
	//	"Sex":  true,
	//	"Age":  24,
	//})
	//db.Delete(&EXPersonn{}, "id = ?", 1)
	db.Unscoped().Delete(&EXPersonn{}, "id = ?", 1)

	db.Where("age > ? ", 0).Find(&p)
	for i, person := range p {
		fmt.Println(i, " - ", person)
	}
}
