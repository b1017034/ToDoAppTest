package main

import (
	"b1017034/NodeGolangDockerTemplate/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	//fmt.Println("hello from docker")
	dsn := "user:password@tcp(database:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	migration(db)
	insert(db)
}

func migration(db *gorm.DB) {
	err := db.AutoMigrate(&model.ToDo{})
	if err != nil {
		fmt.Println(err)
	}
}

func insert(db *gorm.DB) {
	todo := &model.ToDo{Text: "test", CreatedAT: time.Now(), UpdatedAT: time.Now()}
	result := db.Create(todo).Error
	//fmt.Println(todo.ID)
	if result != nil {
		fmt.Println(result)
	}
}
