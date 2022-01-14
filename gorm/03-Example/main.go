package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type User struct {
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
	gorm.Model
}

func main() {
	// 连接数据库
	url := "root:Zhy5224146.+@tcp(127.0.0.1:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 初始化 table schema
	// 可以取消表名的复数形式，使得表名和结构体名称一致
	db.SingularTable(true)
	if !db.HasTable("user") {
		db.AutoMigrate(&User{})
	}

	// 添加唯一索引
	db.Model(&User{}).AddUniqueIndex("name_email", "id", "name", "email")

	// 插入记录
	//db.Create(&User{Name: "张宏宇1", Age: 24, Email: "1229777944@qq.com"})
	//db.Create(&User{Name: "张宏宇2", Age: 24, Email: "1229777945@qq.com"})

	// 查看插入后的元素
	var user User
	var users []User
	db.Find(&users)
	fmt.Println("Users:", users)

	db.First(&user, "name = ?", "张小北")
	fmt.Println("First:", user)

	db.Model(&user).Update("name", "张小北")
	fmt.Println("更新后的记录:", user)

	db.Find(&users)
	fmt.Println("Find Users:", users)

	db.Take(&user)
	fmt.Println("Take User", user)

	db.Where("id > ?", 0).Find(&users)
	fmt.Println("Where Find:", users)

	db.Where("id <> ?", 3).Find(&users)
	fmt.Println("Where Find_1:", users)

	lastWeek := time.Now()
	db.Where("updated_at < ?", lastWeek).Find(&users)
	fmt.Println("Where Find_2:", users)

}
