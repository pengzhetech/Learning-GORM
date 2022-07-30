package main

import (
	"Learning-GORM/connect"
	"Learning-GORM/orm_struct"
	"fmt"
)

func main() {

	user := orm_struct.User{}
	connect.DB.First(&user)
	//fmt.Println(user)

	email := new(string)
	*email = "test"

	lastUser := orm_struct.User{}

	connect.DB.Last(&lastUser)
	fmt.Printf("lastUser--->%v", user.ID)
	/*
		type Result struct {
			ID   int
			Name string
			Age  int
		}
		var result Result
		connect.DB.Table("users").Raw("SELECT id, name, age FROM WHERE name = ?", "c").Scan(&result)
		fmt.Println(result)*/
}
