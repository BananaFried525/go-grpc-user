package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint32    `gorm:"primaryKey;column:"`
	UserName     string    `gorm:"unique;column:user_name"`
	Email        string    `gorm:"unique;column:email"`
	Password     string    `gorm:"column:password"`
	ReferralCode string    `gorm:"unique;column:referral_code"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"index;column:updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUser(dbTxn *gorm.DB, id uint32) (*User, error) {
	user := &User{}
	dbTxn.Where("id = ?", id).Find(user)

	return user, nil
}
