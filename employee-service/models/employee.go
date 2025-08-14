package models

type Employee struct {
	EmployeeID   int     `json:"employee_id" gorm:"primaryKey"`
	FirstName    string  `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName     string  `json:"last_name" gorm:"type:varchar(100);not null"`
	Email        string  `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone        string  `json:"phone" gorm:"type:varchar(10);unique;not null"`
	Gender       string  `json:"gender" gorm:"type:enum('Male','Female','Others');not null"`
	DepartmentID int     `json:"department_id"`
	Designation  string  `json:"designation" gorm:"type:varchar(100);not null;default:'New Joinee'"`
	Salary       float64 `json:"salary" gorm:"type:decimal(10,2);not null;default:40000.00"`
	HireDate     string  `json:"hire_date" gorm:"type:varchar(100)"`
}
