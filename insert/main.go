package main

import (
	"Learning-GORM/connect"
	"Learning-GORM/orm_struct"
	"gorm.io/gorm/clause"
	"time"
)

func main() {
	user := orm_struct.User{
		Name:     "大333哲哥",
		Age:      99,
		Birthday: time.Now(),
	}
	/*	create := connect.DB.Create(&user)
		fmt.Println(create.Error)
		fmt.Println(create.RowsAffected)
		fmt.Println(user.ID)*/

	//创建记录并更新给出的字段
	//connect.DB.Select("Name", "Age", "CreatedAt").Create(&user)

	//创建一个记录且一同忽略传递给略去的字段值
	//connect.DB.Omit("Name", "Age", "CreatedAt").Create(&user)

	/*var users = []orm_struct.User{{Name: "jinzhu1", Birthday: time.Now()},
		{Name: "jinzhu2", Birthday: time.Now()},
		{Name: "jinzhu3", Birthday: time.Now()}}

	connect.DB.CreateInBatches(users, 3)*/

	// 在`id`冲突时，将列更新为新值
	connect.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "jinzhu1"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&user)
}
