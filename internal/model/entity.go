package model

type UserEntity struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name"`
	Phone    string `gorm:"column:phone_number"`
	Email    string `gorm:"column:email_id"`
	IsActive bool   `gorm:"column:is_active"`
}
