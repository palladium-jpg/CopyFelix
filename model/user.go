package model

import "time"

type User struct {
	//gorm.Model

	ID        uint `xorm:"int  null pk autoincr"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `xorm:"varchar(20) null"`
	Telephone string `xorm:"varchar(110) null unique"`
	Password  string `xorm:"varchar(110) null"`
}

//type Model struct {
//
//	//DeletedAt DeletedAt `gorm:"index"`
//}
