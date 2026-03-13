/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:21
	Notes       :


*/

package models

import "github.com/gofrs/uuid/v5"

type CourseType string

const (
	CourseTypeDepartment         CourseType = "DepartmentCourse"
	CourseTypeVocationalElective CourseType = "VocationalElective"
	CourseTypeSocialElective     CourseType = "SocialElective"
	CourseTypeFieldElective      CourseType = "FieldElective"
)

type Course struct {
	ID            uuid.UUID  `gorm:"primary_key"`
	CourseCode    string     `gorm:"size:50;not null"`
	CourseName    string     `gorm:"size:255;not null"`
	CourseType    CourseType `gorm:"size:50;not null"`
	Year          int        `gorm:"not null"`
	Group         int        `gorm:"not null;default:0"`
	Required      bool       `gorm:"not null;default:false"`
	TheoryHours   int        `gorm:"not null;default:0"`
	PracticeHours int        `gorm:"not null;default:0"`
	Credit        int        `gorm:"not null;default:0"`
	ECTS          int        `gorm:"not null;default:0"`
	Quota         int        `gorm:"not null;default:0"`

	BaseRecordFields
}
