package internal

type User struct {
	Id       int    `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `json:"email" binding:"required" gorm:"unique; not null"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "users"
}
