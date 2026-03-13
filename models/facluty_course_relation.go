/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:33
	Notes       :


*/

package models

import "github.com/gofrs/uuid/v5"

type DepartmentCourse struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	DepartmentID uuid.UUID `gorm:"not null;uniqueIndex:idx_department_course"`
	CourseID     uuid.UUID `gorm:"not null;uniqueIndex:idx_department_course"`
	InstructorID uuid.UUID `gorm:"not null"`

	Department Departments `gorm:"foreignKey:DepartmentID"`
	Course     Course      `gorm:"foreignKey:CourseID"`
	Instructor User        `gorm:"foreignKey:InstructorID"`
	BaseRecordFields
}
