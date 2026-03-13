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

type DayOfWeek string

const (
	DayMonday    DayOfWeek = "Monday"
	DayTuesday   DayOfWeek = "Tuesday"
	DayWednesday DayOfWeek = "Wednesday"
	DayThursday  DayOfWeek = "Thursday"
	DayFriday    DayOfWeek = "Friday"
	DaySaturday  DayOfWeek = "Saturday"
)

type AcademicSemester string

const (
	AcademicSemesterFall   AcademicSemester = "Fall"
	AcademicSemesterSpring AcademicSemester = "Spring"
)

type Timetable struct {
	ID                 uuid.UUID        `gorm:"primaryKey" json:"id"`
	FacultyID          uuid.UUID        `gorm:"not null" json:"faculty_id"`
	DepartmentCourseID uuid.UUID        `gorm:"not null" json:"department_course_id"`
	AcademicYear       int              `gorm:"not null" json:"academic_year"`
	Semester           AcademicSemester `gorm:"size:10;not null" json:"semester"`
	DayOfWeek          DayOfWeek        `gorm:"size:20;not null" json:"day_of_week"`
	StartTime          time.Time        `gorm:"not null" json:"start_time"`
	EndTime            time.Time        `gorm:"not null" json:"end_time"`
	Classroom          string           `gorm:"size:100;not null" json:"classroom"`

	Faculty          Faculty          `gorm:"foreignKey:FacultyID" json:"faculty,omitempty"`
	DepartmentCourse DepartmentCourse `gorm:"foreignKey:DepartmentCourseID" json:"department_course,omitempty"`
	BaseRecordFields
}
