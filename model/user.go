package model

type User struct {
	//gorm.Model
	Name      string `xorm:"type:varchar(20);not null"`
	Telephone string `xorm:"varchar(110);not null;primary key"`
	Password  string `xorm:"varchar(110);not null"`
}
