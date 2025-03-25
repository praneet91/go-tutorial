package internal

type Notes struct {
	Id          int    `gorm:"primary_key;auto_increment" json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
