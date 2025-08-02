package models

type Department struct {
	ID   int    `json:"id" gorm:"primaryKey;not null"`
	Name string `json:"name" gorm:"type:varchar(100);not null"`
}
