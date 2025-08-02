package models

type Account struct {
	Username   string `json:"username" gorm:"type:varchar(100)"`
	Password   string `json:"password" gorm:"type:varchar(100)"`
	EmployeeID int    `json:"employee_id" gorm:"primaryKey"`
}
