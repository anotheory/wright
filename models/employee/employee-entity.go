package employee

import (
	"time"
)

type EmployeeEntity struct {
	Id        uint       `gorm:"primaryKey" json:"id"`
	PreName   *string    `gorm:"default:null" json:"preName"`
	FirstName string     `gorm:"not null" json:"firstName"`
	LastName  string     `gorm:"not null" json:"lastName"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func (EmployeeEntity) TableName() string {
	return "employee"
}

type EmployeeRequest struct {
	PreName   *string `json:"preName"`
	FirstName *string `json:"firstName" validate:"required"`
	LastName  *string `json:"lastName" validate:"required"`
}
