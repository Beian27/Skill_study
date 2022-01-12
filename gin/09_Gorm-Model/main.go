package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard
	Emails     []Emails

	BillingAddress   Address
	BillingAddressID sql.NullInt64

	ShippingAddress   Address
	ShippingAddressID int
	IgnoreMe          int        `gorm:"-"`
	Languages         []Language `gorm:"many2many:user_languages;"`
}

type Emails struct {
	ID         int
	UserID     int    `gorm:"index"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"`
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gor,:"index:idx_name_code"`
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

func (u User) TableName() string {
	return "dev_user"
}

func main() {
	db, err := gorm.Open("mysql", "root:Zhy5224146.+@/gin_stu?charset=utf8mb4  &parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.AutoMigrate(&User{}, &Emails{}, &Address{}, &Language{}, &CreditCard{})
	var user User
	var address Address
	db.Model(&user).Related(&address)

}
