package models

type Input struct {
	Username string `json:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
}
