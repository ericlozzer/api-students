package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent() {
	db := Init()

	student := Student{
		Name:   "Eric",
		CPF:    12345678900,
		Email:  "ericlozzer@yahoo.com.br",
		Age:    33,
		Active: true,
	}

	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Create student!")
}
