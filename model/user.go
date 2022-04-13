package model

import "time"

type User struct {
	//gorm.Model
	ID        uint `xorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `xorm:"type:varchar(20);not null"`
	Telephone string `xorm:"varchar(110);not null;primary key"`
	Password  string `xorm:"varchar(110);not null"`
}

//type Model struct {
//
//	//DeletedAt DeletedAt `gorm:"index"`
//}
