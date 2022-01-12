package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	gorm.Model
	Studentname string
	ClassID     uint
	IDCard      IDCard
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Class       []Class `gorm:"many2many:Class_Teacher;"`
}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student

	//多对多
	Teachers []Teacher `gorm:"many2many:Class_Teacher;"`
}

func main() {
	db, err := gorm.Open("mysql", "root:Zhy5224146.+@/gin_stu?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.AutoMigrate(&Student{}, &IDCard{}, &Teacher{}, &Class{})

	id := IDCard{
		Num: 11111,
	}
	id1 := IDCard{
		Num: 11112,
	}
	id2 := IDCard{
		Num: 11113,
	}

	s := Student{
		Studentname: "zhy",
		IDCard:      id,
	}
	s1 := Student{
		Studentname: "zhy1",
		IDCard:      id1,
	}
	s2 := Student{
		Studentname: "zhy2",
		IDCard:      id2,
	}

	t := Teacher{
		TeacherName: "beian",
	}
	t1 := Teacher{
		TeacherName: "beian1",
	}
	t2 := Teacher{
		TeacherName: "beian2",
	}

	_ = Class{
		ClassName: "class",
		Students:  []Student{s, s1, s2},
		Teachers:  []Teacher{t, t1, t2},
	}

	//_ = db.Create(&c).Error

	r := gin.Default()

	r.POST("/student", func(c *gin.Context) {
		var stu Student
		_ = c.BindJSON(&stu)
		db.Create(&stu)
		c.JSON(200, stu)
	})

	r.GET("/student/:id", func(c *gin.Context) {
		var stu Student
		id := c.Param("id")
		db.Preload("IDCard").First(&stu, "id = ?", id)
		c.JSON(200, gin.H{
			"message": "成功",
			"data":    stu,
		})
	})

	r.GET("/class/:id", func(c *gin.Context) {
		var class Class
		id := c.Param("id")
		db.Preload("Students").Preload("Students.IDCard").Preload("Teachers").First(&class, "id = ?", id)
		c.JSON(200, gin.H{
			"message": "成功",
			"data":    class,
		})
	})

	r.Run(":8989")

}
