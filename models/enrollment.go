/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:21
	Notes       :


*/

package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type EnrollmentType string

const (
	EnrollmentTypeRegular            EnrollmentType = "Regular"
	EnrollmentTypeTransfer           EnrollmentType = "Transfer"
	EnrollmentTypeVerticalTransfer   EnrollmentType = "VerticalTransfer"
	EnrollmentTypeHorizontalTransfer EnrollmentType = "HorizontalTransfer"
	EnrollmentTypeExchange           EnrollmentType = "Exchange"
)

type Enrollment struct {
	ID             uuid.UUID      `gorm:"primaryKey" json:"id"`
	UserID         uuid.UUID      `gorm:"not null" json:"user_id"`
	DepartmentID   uuid.UUID      `gorm:"not null" json:"department_id"`
	EnrollmentType EnrollmentType `gorm:"size:50;not null" json:"enrollment_type"`
	Year           int            `gorm:"not null" json:"year"`
	Semester       int            `gorm:"not null" json:"semester"`
	EnrollmentDate time.Time      `gorm:"not null" json:"enrollment_date"`
	Active         bool           `gorm:"not null;default:true" json:"active"`

	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Department Departments `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	BaseRecordFields
}
