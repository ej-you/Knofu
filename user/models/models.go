package models

import (
	"time"
)


// модель юзера
type User struct {
	ID			uint		`gorm:"primarykey" json:"id"`
	Email		string 		`gorm:"type:varchar(50); not null; uniqueIndex" json:"email"`
	FirstName	string 		`gorm:"type:varchar(100); not null" json:"firstName"`
	LastName	string		`gorm:"type:varchar(100); not null" json:"lastName"`
	Password	string 		`gorm:"type:varchar(50); not null; check: password > 8" json:"password"`
	CreatedAt	time.Time	`gorm:"not null; autoCreateTime" json:"createdAt"`
}
